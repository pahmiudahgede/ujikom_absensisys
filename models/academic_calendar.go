package models

import "time"

type AcademicCalendar struct {
	BaseModel
	SchoolID     string    `json:"school_id" gorm:"type:varchar(255);not null;uniqueIndex:idx_school_year_semester"`
	AcademicYear string    `json:"academic_year" gorm:"type:varchar(10);not null;uniqueIndex:idx_school_year_semester"`
	Semester     string    `json:"semester" gorm:"type:enum('ganjil','genap');not null;uniqueIndex:idx_school_year_semester"`
	StartDate    time.Time `json:"start_date" gorm:"type:date;not null"`
	EndDate      time.Time `json:"end_date" gorm:"type:date;not null"`
	Description  *string   `json:"description" gorm:"type:text"`
	IsActive     bool      `json:"is_active" gorm:"not null;default:true"`

	School *School `json:"school,omitempty" gorm:"foreignKey:SchoolID"`
}

func (AcademicCalendar) TableName() string {
	return "academic_calendar"
}
