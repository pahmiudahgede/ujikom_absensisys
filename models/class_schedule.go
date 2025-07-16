package models

import (
	"fmt"
	"time"
)

type ClassSchedule struct {
	BaseModel
	ClassID      string   `json:"class_id" gorm:"type:varchar(255);not null;uniqueIndex:idx_class_schedule"`
	SubjectID    string   `json:"subject_id" gorm:"type:varchar(255);not null;index"`
	TeacherID    string   `json:"teacher_id" gorm:"type:varchar(255);not null;index"`
	DayOfWeek    string   `json:"day_of_week" gorm:"type:enum('senin','selasa','rabu','kamis','jumat','sabtu');not null;uniqueIndex:idx_class_schedule"`
	StartTime    TimeOnly `json:"start_time" gorm:"uniqueIndex:idx_class_schedule;comment:'Jam mulai pelajaran'"`
	EndTime      TimeOnly `json:"end_time" gorm:"comment:'Jam selesai pelajaran'"`
	Room         *string  `json:"room" gorm:"type:varchar(50);comment:'Ruang kelas'"`
	AcademicYear string   `json:"academic_year" gorm:"type:varchar(10);not null;index"`
	Semester     string   `json:"semester" gorm:"type:enum('ganjil','genap');not null;index"`
	IsActive     bool     `json:"is_active" gorm:"not null;default:true;index"`

	Class         *Class         `json:"class,omitempty" gorm:"foreignKey:ClassID"`
	Subject       *Subject       `json:"subject,omitempty" gorm:"foreignKey:SubjectID"`
	Teacher       *Teacher       `json:"teacher,omitempty" gorm:"foreignKey:TeacherID"`
	ClassSessions []ClassSession `json:"class_sessions,omitempty" gorm:"foreignKey:ScheduleID"`
}

func (ClassSchedule) TableName() string {
	return "class_schedules"
}

func (cs *ClassSchedule) GetStartTimeString() string {
	return cs.StartTime.String()
}

func (cs *ClassSchedule) GetEndTimeString() string {
	return cs.EndTime.String()
}

func (cs *ClassSchedule) SetStartTimeFromString(timeStr string) error {
	timeOnly, err := NewTimeOnlyFromString(timeStr)
	if err != nil {
		return err
	}
	cs.StartTime = timeOnly
	return nil
}

func (cs *ClassSchedule) SetEndTimeFromString(timeStr string) error {
	timeOnly, err := NewTimeOnlyFromString(timeStr)
	if err != nil {
		return err
	}
	cs.EndTime = timeOnly
	return nil
}

func (cs *ClassSchedule) ValidateTimeRange() error {
	if cs.EndTime.Before(cs.StartTime) || cs.EndTime.Equal(cs.StartTime) {
		return fmt.Errorf("end time must be after start time")
	}
	return nil
}

func (cs *ClassSchedule) GetDuration() time.Duration {
	return cs.EndTime.Time.Sub(cs.StartTime.Time)
}
