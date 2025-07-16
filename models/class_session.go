package models

import "time"

type ClassSession struct {
	BaseModel
	ScheduleID      string     `json:"schedule_id" gorm:"type:varchar(255);not null;uniqueIndex:idx_schedule_date"`
	Date            time.Time  `json:"date" gorm:"type:date;not null;uniqueIndex:idx_schedule_date;index;comment:'Tanggal pelaksanaan'"`
	ActualStartTime *time.Time `json:"actual_start_time" gorm:"type:timestamp;comment:'Waktu mulai sebenarnya'"`
	ActualEndTime   *time.Time `json:"actual_end_time" gorm:"type:timestamp;comment:'Waktu selesai sebenarnya'"`
	Topic           *string    `json:"topic" gorm:"type:varchar(255);comment:'Topik pembelajaran'"`
	Material        *string    `json:"material" gorm:"type:text;comment:'Materi yang diajarkan'"`
	Status          string     `json:"status" gorm:"type:enum('scheduled','ongoing','completed','cancelled');not null;default:'scheduled';index"`
	Notes           *string    `json:"notes" gorm:"type:text;comment:'Catatan guru'"`
	CreatedBy       string     `json:"created_by" gorm:"type:varchar(255);not null;index"`

	Schedule          *ClassSchedule      `json:"schedule,omitempty" gorm:"foreignKey:ScheduleID"`
	CreatedByTeacher  *Teacher            `json:"created_by_teacher,omitempty" gorm:"foreignKey:CreatedBy"`
	SubjectAttendance []SubjectAttendance `json:"subject_attendance,omitempty" gorm:"foreignKey:SessionID"`
	SubjectPermits    []SubjectPermit     `json:"subject_permits,omitempty" gorm:"foreignKey:SessionID"`
}

func (ClassSession) TableName() string {
	return "class_sessions"
}
