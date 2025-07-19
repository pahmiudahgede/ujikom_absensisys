package attendance

import (
	"absensibe/utils"
	"context"
	"fmt"
	"time"
)

type AttendanceRulesService interface {
	GetAttendanceRules(ctx context.Context, studentID string) (*AttendanceRulesResponse, error)
	ClearAttendanceRulesCache(ctx context.Context, schoolID string) error
}

type attendanceRulesService struct {
	repo AttendanceRulesRepository
}

func NewAttendanceRulesService(repo AttendanceRulesRepository) AttendanceRulesService {
	return &attendanceRulesService{
		repo: repo,
	}
}

func (s *attendanceRulesService) GetAttendanceRules(ctx context.Context, studentID string) (*AttendanceRulesResponse, error) {
	// Get school ID first
	schoolID, err := s.repo.GetSchoolIDByStudentID(ctx, studentID)
	if err != nil {
		return nil, err
	}

	// Try to get from cache
	cacheKey := fmt.Sprintf("attendance_rules:%s", schoolID)
	var response AttendanceRulesResponse

	if err := utils.Redis.GetCache(cacheKey, &response); err == nil {
		// Update dynamic fields (time-dependent)
		s.updateTimeStatus(&response)
		return &response, nil
	}

	// Cache miss, get from database
	settings, err := s.repo.GetAttendanceSettings(ctx, schoolID)
	if err != nil {
		return nil, err
	}

	safeAreas, err := s.repo.GetSafeAreas(ctx, schoolID)
	if err != nil {
		return nil, err
	}

	// Build response
	response = AttendanceRulesResponse{
		Settings:  ToAttendanceSettingsDTO(settings),
		SafeAreas: ToSafeAreaDTOs(safeAreas),
	}

	// Add time status
	s.updateTimeStatus(&response)

	// Cache for 1 hour (settings don't change frequently)
	if err := utils.Redis.SetCache(cacheKey, response, utils.TTL_MEDIUM); err != nil {
		fmt.Printf("Warning: failed to cache attendance rules: %v\n", err)
	}

	return &response, nil
}

func (s *attendanceRulesService) updateTimeStatus(response *AttendanceRulesResponse) {
	now := time.Now()
	response.CurrentTime = now.Format("15:04")

	if response.Settings == nil {
		response.CanCheckIn = false
		response.CanCheckOut = false
		response.Status = "no_settings"
		response.Message = "Pengaturan absen belum dikonfigurasi"
		return
	}

	// Parse time strings back to time.Time for comparison
	checkInStart, _ := time.Parse("15:04", response.Settings.CheckInStart)
	checkInEnd, _ := time.Parse("15:04", response.Settings.CheckInEnd)
	checkOutStart, _ := time.Parse("15:04", response.Settings.CheckOutStart)
	checkOutEnd, _ := time.Parse("15:04", response.Settings.CheckOutEnd)

	currentTime := time.Date(0, 1, 1, now.Hour(), now.Minute(), 0, 0, time.UTC)

	// Determine status
	if currentTime.Before(checkInStart) {
		response.Status = "before_checkin"
		response.Message = fmt.Sprintf("Absen masuk dimulai pada %s", response.Settings.CheckInStart)
		response.CanCheckIn = false
		response.CanCheckOut = false
	} else if !currentTime.After(checkInEnd) {
		response.Status = "checkin_time"
		response.CanCheckIn = true
		response.CanCheckOut = false

		// Check if late
		lateThreshold := checkInStart.Add(time.Duration(response.Settings.LateTolerance) * time.Minute)
		if currentTime.After(lateThreshold) {
			response.Message = fmt.Sprintf("Waktu absen masuk (terlambat). Batas: %s", response.Settings.CheckInEnd)
		} else {
			response.Message = fmt.Sprintf("Waktu absen masuk. Batas: %s", response.Settings.CheckInEnd)
		}
	} else if currentTime.Before(checkOutStart) {
		response.Status = "between"
		response.Message = fmt.Sprintf("Absen pulang dimulai pada %s", response.Settings.CheckOutStart)
		response.CanCheckIn = false
		response.CanCheckOut = false
	} else if !currentTime.After(checkOutEnd) {
		response.Status = "checkout_time"
		response.CanCheckIn = false
		response.CanCheckOut = true
		response.Message = fmt.Sprintf("Waktu absen pulang. Batas: %s", response.Settings.CheckOutEnd)
	} else {
		response.Status = "after_checkout"
		response.Message = "Waktu absen sudah berakhir"
		response.CanCheckIn = false
		response.CanCheckOut = false
	}
}

func (s *attendanceRulesService) ClearAttendanceRulesCache(ctx context.Context, schoolID string) error {
	cacheKey := fmt.Sprintf("attendance_rules:%s", schoolID)
	return utils.Redis.Delete(cacheKey)
}
