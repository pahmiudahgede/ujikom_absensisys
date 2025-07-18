package models

import "time"

type AcademicCalendar struct {
	ID           string    `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();unique;not null" json:"id"`
	SchoolID     string    `gorm:"type:uuid;not null;uniqueIndex:idx_school_year_semester;constraint:OnDelete:CASCADE" json:"school_id"`
	AcademicYear string    `gorm:"type:varchar(10);not null;uniqueIndex:idx_school_year_semester" json:"academic_year"`
	Semester     string    `gorm:"type:varchar(10);not null;uniqueIndex:idx_school_year_semester;check:semester IN ('ganjil','genap')" json:"semester"`
	StartDate    time.Time `gorm:"type:date;not null" json:"start_date"`
	EndDate      time.Time `gorm:"type:date;not null" json:"end_date"`
	Description  *string   `gorm:"type:text" json:"description"`
	IsActive     bool      `gorm:"not null;default:true" json:"is_active"`
	CreatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`

	// Relations
	School *School `json:"school,omitempty" gorm:"foreignKey:SchoolID"`
}

func (AcademicCalendar) TableName() string {
	return "academic_calendars"
}
