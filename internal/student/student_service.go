package student

import (
	"absensibe/utils"
	"context"
	"fmt"
)

type StudentService interface {
	GetStudentInfo(ctx context.Context, studentID string, req *StudentInfoRequest) (*StudentInfoResponse, error)
	ClearStudentCache(ctx context.Context, studentID string) error
}

type studentService struct {
	repo StudentRepository
}

func NewStudentService(repo StudentRepository) StudentService {
	return &studentService{
		repo: repo,
	}
}

func (s *studentService) GetStudentInfo(ctx context.Context, studentID string, req *StudentInfoRequest) (*StudentInfoResponse, error) {
	// Generate cache key
	cacheKey := fmt.Sprintf("student_info:%s:%v:%v:%v:%s:%d",
		studentID, req.IncludeClass, req.IncludeAttendance, req.IncludeSchedule, req.AttendanceMonth, req.AttendanceLimit)

	// Try get from cache first
	var response StudentInfoResponse
	if err := utils.Redis.GetCache(cacheKey, &response); err == nil {
		return &response, nil
	}

	// Get from database
	student, err := s.repo.GetStudentByID(ctx, studentID, req)
	if err != nil {
		return nil, err
	}

	// Convert to response
	response = *ToStudentResponse(student)

	// Get attendance summary if requested
	if req.IncludeAttendance {
		summary, err := s.repo.GetAttendanceSummary(ctx, studentID, req.AttendanceMonth)
		if err != nil {
			// Don't fail the whole request, just log the error
			fmt.Printf("Warning: failed to get attendance summary: %v\n", err)
		} else {
			response.AttendanceSummary = summary
		}
	}

	// Cache the response
	cacheTTL := utils.TTL_MEDIUM
	if req.IncludeAttendance {
		cacheTTL = utils.TTL_SHORT // Attendance data changes frequently
	}

	if err := utils.Redis.SetCache(cacheKey, response, cacheTTL); err != nil {
		// Log error but don't fail the request
		fmt.Printf("Warning: failed to cache student info: %v\n", err)
	}

	return &response, nil
}

func (s *studentService) ClearStudentCache(ctx context.Context, studentID string) error {
	pattern := fmt.Sprintf("student_info:%s:*", studentID)
	return utils.Redis.FlushPattern(pattern)
}
