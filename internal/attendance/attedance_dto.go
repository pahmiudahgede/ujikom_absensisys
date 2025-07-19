package attendance

import (
	"absensibe/models"
	"time"
)

// AttendanceRulesResponse represents attendance rules and safe areas
type AttendanceRulesResponse struct {
	Settings    *AttendanceSettingsDTO `json:"settings"`
	SafeAreas   []SafeAreaDTO          `json:"safe_areas"`
	CurrentTime string                 `json:"current_time"`
	CanCheckIn  bool                   `json:"can_check_in"`
	CanCheckOut bool                   `json:"can_check_out"`
	Status      string                 `json:"status"` // "before_checkin", "checkin_time", "between", "checkout_time", "after_checkout"
	Message     string                 `json:"message"`
} //	@name	AttendanceRulesResponse

// AttendanceSettingsDTO represents attendance settings
type AttendanceSettingsDTO struct {
	CheckInStart    string `json:"check_in_start"`  // "07:00"
	CheckInEnd      string `json:"check_in_end"`    // "08:00"
	CheckOutStart   string `json:"check_out_start"` // "15:00"
	CheckOutEnd     string `json:"check_out_end"`   // "17:00"
	LateTolerance   int    `json:"late_tolerance"`  // minutes
	RequirePhoto    bool   `json:"require_photo"`
	RequireLocation bool   `json:"require_location"`
	MaxDistance     int    `json:"max_distance"` // meters
} //	@name	AttendanceSettingsDTO

// SafeAreaDTO represents safe area for attendance
type SafeAreaDTO struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	Radius      float64 `json:"radius"` // meters
	Description *string `json:"description"`
	IsActive    bool    `json:"is_active"`
} //	@name	SafeAreaDTO

// Convert models to DTOs
func ToAttendanceSettingsDTO(settings *models.AttendanceSettings) *AttendanceSettingsDTO {
	if settings == nil {
		return nil
	}

	return &AttendanceSettingsDTO{
		CheckInStart:    settings.CheckInStart.Format("15:04"),
		CheckInEnd:      settings.CheckInEnd.Format("15:04"),
		CheckOutStart:   settings.CheckOutStart.Format("15:04"),
		CheckOutEnd:     settings.CheckOutEnd.Format("15:04"),
		LateTolerance:   settings.LateTolerance,
		RequirePhoto:    settings.RequirePhoto,
		RequireLocation: settings.RequireLocation,
		MaxDistance:     settings.MaxDistance,
	}
}

func ToSafeAreaDTO(safeArea *models.SafeArea) SafeAreaDTO {
	return SafeAreaDTO{
		ID:          safeArea.ID,
		Name:        safeArea.Name,
		Latitude:    safeArea.Latitude,
		Longitude:   safeArea.Longitude,
		Radius:      safeArea.Radius,
		Description: safeArea.Description,
		IsActive:    safeArea.IsActive,
	}
}

func ToSafeAreaDTOs(safeAreas []models.SafeArea) []SafeAreaDTO {
	result := make([]SafeAreaDTO, len(safeAreas))
	for i, area := range safeAreas {
		result[i] = ToSafeAreaDTO(&area)
	}
	return result
}

// Helper function to determine attendance status and permissions
func DetermineAttendanceStatus(settings *models.AttendanceSettings) (bool, bool, string, string) {
	now := time.Now()
	currentTime := time.Date(0, 1, 1, now.Hour(), now.Minute(), 0, 0, time.UTC)

	canCheckIn := false
	canCheckOut := false
	status := ""
	message := ""

	// Check current time against attendance windows
	if currentTime.Before(settings.CheckInStart) {
		status = "before_checkin"
		message = "Belum waktunya absen masuk"
	} else if currentTime.After(settings.CheckInStart) && currentTime.Before(settings.CheckInEnd) {
		status = "checkin_time"
		canCheckIn = true
		if currentTime.After(addMinutes(settings.CheckInStart, settings.LateTolerance)) {
			message = "Waktu absen masuk (terlambat)"
		} else {
			message = "Waktu absen masuk"
		}
	} else if currentTime.After(settings.CheckInEnd) && currentTime.Before(settings.CheckOutStart) {
		status = "between"
		message = "Diluar waktu absen"
	} else if currentTime.After(settings.CheckOutStart) && currentTime.Before(settings.CheckOutEnd) {
		status = "checkout_time"
		canCheckOut = true
		message = "Waktu absen pulang"
	} else {
		status = "after_checkout"
		message = "Sudah lewat waktu absen"
	}

	return canCheckIn, canCheckOut, status, message
}

// Helper function to add minutes to time
func addMinutes(t time.Time, minutes int) time.Time {
	return t.Add(time.Duration(minutes) * time.Minute)
}
