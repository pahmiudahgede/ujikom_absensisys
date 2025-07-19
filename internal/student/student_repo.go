package student

import (
	"absensibe/config"
	"absensibe/models"
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type StudentRepository interface {
	GetStudentByID(ctx context.Context, studentID string, req *StudentInfoRequest) (*models.Student, error)
	GetAttendanceSummary(ctx context.Context, studentID string, month string) (*AttendanceSummary, error)
}

type studentRepository struct {
	db *gorm.DB
}

func NewStudentRepository() StudentRepository {
	return &studentRepository{
		db: config.DB,
	}
}

func (r *studentRepository) GetStudentByID(ctx context.Context, studentID string, req *StudentInfoRequest) (*models.Student, error) {
	var student models.Student

	query := r.db.WithContext(ctx)

	// Preload based on request
	if req.IncludeClass {
		query = query.Preload("Class").Preload("Class.School").Preload("Class.HomeroomTeacher")

		if req.IncludeSchedule {
			query = query.Preload("Class.ClassSchedules", "is_active = ?", true).
				Preload("Class.ClassSchedules.Subject").
				Preload("Class.ClassSchedules.Teacher")
		}
	}

	if req.IncludeAttendance {
		attendanceQuery := "1 = 1"
		args := []interface{}{}

		// Filter by month if specified
		if req.AttendanceMonth != "" {
			if monthTime, err := time.Parse("2006-01", req.AttendanceMonth); err == nil {
				startOfMonth := time.Date(monthTime.Year(), monthTime.Month(), 1, 0, 0, 0, 0, time.UTC)
				endOfMonth := startOfMonth.AddDate(0, 1, 0).Add(-time.Second)
				attendanceQuery += " AND date >= ? AND date <= ?"
				args = append(args, startOfMonth, endOfMonth)
			}
		}

		// Apply limit
		limit := req.AttendanceLimit
		if limit <= 0 || limit > 50 {
			limit = 10
		}

		query = query.Preload("Attendances", func(db *gorm.DB) *gorm.DB {
			return db.Where(attendanceQuery, args...).Order("date DESC").Limit(limit)
		})
	}

	err := query.First(&student, "id = ? AND status = ?", studentID, "aktif").Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("student not found")
		}
		return nil, fmt.Errorf("failed to get student: %w", err)
	}

	return &student, nil
}

func (r *studentRepository) GetAttendanceSummary(ctx context.Context, studentID string, month string) (*AttendanceSummary, error) {
	var summary AttendanceSummary

	baseQuery := r.db.WithContext(ctx).Model(&models.Attendance{}).Where("student_id = ?", studentID)

	// Apply month filter if specified
	if month != "" {
		if monthTime, err := time.Parse("2006-01", month); err == nil {
			startOfMonth := time.Date(monthTime.Year(), monthTime.Month(), 1, 0, 0, 0, 0, time.UTC)
			endOfMonth := startOfMonth.AddDate(0, 1, 0).Add(-time.Second)
			baseQuery = baseQuery.Where("date >= ? AND date <= ?", startOfMonth, endOfMonth)
			summary.CurrentMonth = month
		}
	} else {
		// Default to current month
		now := time.Now()
		startOfMonth := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.UTC)
		endOfMonth := startOfMonth.AddDate(0, 1, 0).Add(-time.Second)
		baseQuery = baseQuery.Where("date >= ? AND date <= ?", startOfMonth, endOfMonth)
		summary.CurrentMonth = now.Format("2006-01")
	}

	// Get total days
	if err := baseQuery.Count(&summary.TotalDays).Error; err != nil {
		return nil, fmt.Errorf("failed to count total days: %w", err)
	}

	// Count present days (hadir + terlambat)
	if err := baseQuery.Where("check_in_status IN (?)", []string{"hadir", "terlambat"}).Count(&summary.PresentDays).Error; err != nil {
		return nil, fmt.Errorf("failed to count present days: %w", err)
	}

	// Count late days
	if err := baseQuery.Where("check_in_status = ?", "terlambat").Count(&summary.LateDays).Error; err != nil {
		return nil, fmt.Errorf("failed to count late days: %w", err)
	}

	// Count absent days
	if err := baseQuery.Where("check_in_status = ?", "alpha").Count(&summary.AbsentDays).Error; err != nil {
		return nil, fmt.Errorf("failed to count absent days: %w", err)
	}

	// Calculate attendance rate
	if summary.TotalDays > 0 {
		summary.AttendanceRate = float64(summary.PresentDays) / float64(summary.TotalDays) * 100
	}

	return &summary, nil
}
