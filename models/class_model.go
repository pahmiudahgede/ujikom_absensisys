package models

import "time"

type Class struct {
	ID                string    `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();unique;not null" json:"id"`
	Name              string    `gorm:"type:varchar(100);not null;comment:'Nama kelas (contoh: X-RPL-1)'" json:"name"`
	Grade             string    `gorm:"type:varchar(5);not null;check:grade IN ('X','XI','XII');comment:'Tingkat kelas'" json:"grade"`
	MajorID           string    `gorm:"type:uuid;not null;index;constraint:OnDelete:CASCADE" json:"major_id"`
	SchoolID          string    `gorm:"type:uuid;not null;index;constraint:OnDelete:CASCADE" json:"school_id"`
	HomeroomTeacherID *string   `gorm:"type:uuid;index;comment:'Wali kelas'" json:"homeroom_teacher_id"`
	MaxStudents       int       `gorm:"not null;default:36;comment:'Maksimal siswa per kelas'" json:"max_students"`
	AcademicYear      string    `gorm:"type:varchar(10);not null;index;comment:'Tahun ajaran (contoh: 2024/2025)'" json:"academic_year"`
	IsActive          bool      `gorm:"not null;default:true;index" json:"is_active"`
	CreatedAt         time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt         time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`

	// Relations
	Major                     *Major                      `json:"major,omitempty" gorm:"foreignKey:MajorID"`
	School                    *School                     `json:"school,omitempty" gorm:"foreignKey:SchoolID"`
	HomeroomTeacher           *Teacher                    `json:"homeroom_teacher,omitempty" gorm:"foreignKey:HomeroomTeacherID"`
	Students                  []Student                   `json:"students,omitempty" gorm:"foreignKey:ClassID"`
	ClassSchedules            []ClassSchedule             `json:"class_schedules,omitempty" gorm:"foreignKey:ClassID"`
	SubjectAttendanceSettings []SubjectAttendanceSettings `json:"subject_attendance_settings,omitempty" gorm:"foreignKey:ClassID"`
}

func (Class) TableName() string {
	return "classes"
}
