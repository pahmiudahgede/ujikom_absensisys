package attendance

import (
	"absensibe/config"
	"absensibe/models"
	"context"
	"fmt"

	"gorm.io/gorm"
)

type AttendanceRulesRepository interface {
	GetSchoolIDByStudentID(ctx context.Context, studentID string) (string, error)
	GetAttendanceSettings(ctx context.Context, schoolID string) (*models.AttendanceSettings, error)
	GetSafeAreas(ctx context.Context, schoolID string) ([]models.SafeArea, error)
}

type attendanceRulesRepository struct {
	db *gorm.DB
}

func NewAttendanceRulesRepository() AttendanceRulesRepository {
	return &attendanceRulesRepository{
		db: config.DB,
	}
}

func (r *attendanceRulesRepository) GetSchoolIDByStudentID(ctx context.Context, studentID string) (string, error) {
	var student models.Student

	err := r.db.WithContext(ctx).
		Select("students.id, classes.school_id").
		Joins("JOIN classes ON students.class_id = classes.id").
		Where("students.id = ? AND students.status = ?", studentID, "aktif").
		First(&student).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return "", fmt.Errorf("student not found")
		}
		return "", fmt.Errorf("failed to get student school: %w", err)
	}

	// Get school_id from the join
	var schoolID string
	err = r.db.WithContext(ctx).
		Table("students").
		Select("classes.school_id").
		Joins("JOIN classes ON students.class_id = classes.id").
		Where("students.id = ?", studentID).
		Scan(&schoolID).Error

	if err != nil {
		return "", fmt.Errorf("failed to get school ID: %w", err)
	}

	return schoolID, nil
}

func (r *attendanceRulesRepository) GetAttendanceSettings(ctx context.Context, schoolID string) (*models.AttendanceSettings, error) {
	var settings models.AttendanceSettings

	err := r.db.WithContext(ctx).
		Where("school_id = ?", schoolID).
		First(&settings).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("attendance settings not found for school")
		}
		return nil, fmt.Errorf("failed to get attendance settings: %w", err)
	}

	return &settings, nil
}

func (r *attendanceRulesRepository) GetSafeAreas(ctx context.Context, schoolID string) ([]models.SafeArea, error) {
	var safeAreas []models.SafeArea

	err := r.db.WithContext(ctx).
		Where("school_id = ? AND is_active = ?", schoolID, true).
		Order("name ASC").
		Find(&safeAreas).Error

	if err != nil {
		return nil, fmt.Errorf("failed to get safe areas: %w", err)
	}

	return safeAreas, nil
}
