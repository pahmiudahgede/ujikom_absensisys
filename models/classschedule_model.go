package models

import "time"

type ClassSchedule struct {
	ID           string    `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();unique;not null" json:"id"`
	ClassID      string    `gorm:"type:uuid;not null;uniqueIndex:idx_class_schedule;constraint:OnDelete:CASCADE" json:"class_id"`
	SubjectID    string    `gorm:"type:uuid;not null;index;constraint:OnDelete:CASCADE" json:"subject_id"`
	TeacherID    string    `gorm:"type:uuid;not null;index;constraint:OnDelete:CASCADE" json:"teacher_id"`
	DayOfWeek    string    `gorm:"type:varchar(10);not null;uniqueIndex:idx_class_schedule;check:day_of_week IN ('senin','selasa','rabu','kamis','jumat','sabtu')" json:"day_of_week"`
	StartTime    time.Time `gorm:"type:time;not null;uniqueIndex:idx_class_schedule;comment:'Jam mulai pelajaran'" json:"start_time"`
	EndTime      time.Time `gorm:"type:time;not null;comment:'Jam selesai pelajaran'" json:"end_time"`
	Room         *string   `gorm:"type:varchar(50);comment:'Ruang kelas'" json:"room"`
	AcademicYear string    `gorm:"type:varchar(10);not null;index" json:"academic_year"`
	Semester     string    `gorm:"type:varchar(10);not null;check:semester IN ('ganjil','genap');index" json:"semester"`
	IsActive     bool      `gorm:"not null;default:true;index" json:"is_active"`
	CreatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`

	// Relations
	Class         *Class         `json:"class,omitempty" gorm:"foreignKey:ClassID"`
	Subject       *Subject       `json:"subject,omitempty" gorm:"foreignKey:SubjectID"`
	Teacher       *Teacher       `json:"teacher,omitempty" gorm:"foreignKey:TeacherID"`
	ClassSessions []ClassSession `json:"class_sessions,omitempty" gorm:"foreignKey:ScheduleID"`
}

func (ClassSchedule) TableName() string {
	return "class_schedules"
}
