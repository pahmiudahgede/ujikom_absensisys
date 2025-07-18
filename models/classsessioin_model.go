package models

import "time"

type ClassSession struct {
	ID              string     `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();unique;not null" json:"id"`
	ScheduleID      string     `gorm:"type:uuid;not null;uniqueIndex:idx_schedule_date;constraint:OnDelete:CASCADE" json:"schedule_id"`
	Date            time.Time  `gorm:"type:date;not null;uniqueIndex:idx_schedule_date;index;comment:'Tanggal pelaksanaan'" json:"date"`
	ActualStartTime *time.Time `gorm:"type:timestamp;comment:'Waktu mulai sebenarnya'" json:"actual_start_time"`
	ActualEndTime   *time.Time `gorm:"type:timestamp;comment:'Waktu selesai sebenarnya'" json:"actual_end_time"`
	Topic           *string    `gorm:"type:varchar(255);comment:'Topik pembelajaran'" json:"topic"`
	Material        *string    `gorm:"type:text;comment:'Materi yang diajarkan'" json:"material"`
	Status          string     `gorm:"type:varchar(20);not null;default:'scheduled';check:status IN ('scheduled','ongoing','completed','cancelled');index" json:"status"`
	Notes           *string    `gorm:"type:text;comment:'Catatan guru'" json:"notes"`
	CreatedBy       string     `gorm:"type:uuid;not null;index;constraint:OnDelete:CASCADE" json:"created_by"`
	CreatedAt       time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt       time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`

	// Relations
	Schedule           *ClassSchedule      `json:"schedule,omitempty" gorm:"foreignKey:ScheduleID"`
	CreatedByTeacher   *Teacher            `json:"created_by_teacher,omitempty" gorm:"foreignKey:CreatedBy"`
	SubjectAttendances []SubjectAttendance `json:"subject_attendances,omitempty" gorm:"foreignKey:SessionID"`
	SubjectPermits     []SubjectPermit     `json:"subject_permits,omitempty" gorm:"foreignKey:SessionID"`
}

func (ClassSession) TableName() string {
	return "class_sessions"
}
