package seeders

import (
	"crypto/rand"
	"fmt"
	"time"

	"gorm.io/gorm"
)

func GenerateUUID() string {
	b := make([]byte, 16)
	rand.Read(b)
	return fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}

func CreateClassScheduleRaw(db *gorm.DB, data ClassScheduleData) error {
	query := `
		INSERT INTO class_schedules 
		(id, created_at, updated_at, class_id, subject_id, teacher_id, day_of_week, start_time, end_time, room, academic_year, semester, is_active) 
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`

	return db.Exec(query,
		GenerateUUID(),
		time.Now(),
		time.Now(),
		data.ClassID,
		data.SubjectID,
		data.TeacherID,
		data.DayOfWeek,
		data.StartTime,
		data.EndTime,
		data.Room,
		data.AcademicYear,
		data.Semester,
		data.IsActive,
	).Error
}

func CreateAttendanceSettingsRaw(db *gorm.DB, data AttendanceSettingsData) error {
	query := `
		INSERT INTO attendance_settings 
		(id, created_at, updated_at, school_id, checkin_start, checkin_end, checkout_start, checkout_end, late_tolerance, require_photo, require_location, max_distance) 
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`

	return db.Exec(query,
		GenerateUUID(),
		time.Now(),
		time.Now(),
		data.SchoolID,
		data.CheckinStart,
		data.CheckinEnd,
		data.CheckoutStart,
		data.CheckoutEnd,
		data.LateTolerance,
		data.RequirePhoto,
		data.RequireLocation,
		data.MaxDistance,
	).Error
}

type ClassScheduleData struct {
	ClassID      string
	SubjectID    string
	TeacherID    string
	DayOfWeek    string
	StartTime    string
	EndTime      string
	Room         string
	AcademicYear string
	Semester     string
	IsActive     bool
}

type AttendanceSettingsData struct {
	SchoolID        string
	CheckinStart    string
	CheckinEnd      string
	CheckoutStart   string
	CheckoutEnd     string
	LateTolerance   int
	RequirePhoto    bool
	RequireLocation bool
	MaxDistance     int
}

func ValidateTimeFormat(timeStr string) bool {
	_, err := time.Parse("15:04:05", timeStr)
	return err == nil
}

func GenerateTimeSlots() []struct {
	Start string
	End   string
} {
	return []struct {
		Start string
		End   string
	}{
		{"07:30:00", "08:15:00"},
		{"08:15:00", "09:00:00"},
		{"09:15:00", "10:00:00"},
		{"10:00:00", "10:45:00"},
		{"11:00:00", "11:45:00"},
		{"12:30:00", "13:15:00"},
		{"13:15:00", "14:00:00"},
		{"14:00:00", "14:45:00"},
	}
}

func FixTimeFieldsInDatabase(db *gorm.DB) error {

	queries := []string{
		"SET sql_mode = 'TRADITIONAL,ALLOW_INVALID_DATES'",
		"SET time_zone = '+07:00'",
	}

	for _, query := range queries {
		if err := db.Exec(query).Error; err != nil {
			return fmt.Errorf("failed to execute query %s: %v", query, err)
		}
	}

	return nil
}
