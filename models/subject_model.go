package models

import "time"

type Subject struct {
	ID          string    `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();unique;not null" json:"id"`
	Code        string    `gorm:"type:varchar(50);not null;uniqueIndex;comment:'Kode mata pelajaran (contoh: MTK, BIND)'" json:"code"`
	Name        string    `gorm:"type:varchar(255);not null;comment:'Nama mata pelajaran'" json:"name"`
	SchoolID    string    `gorm:"type:uuid;not null;index;constraint:OnDelete:CASCADE" json:"school_id"`
	CreditHours int       `gorm:"not null;default:2;comment:'Jumlah jam pelajaran per minggu'" json:"credit_hours"`
	Description *string   `gorm:"type:text" json:"description"`
	IsActive    bool      `gorm:"not null;default:true;index" json:"is_active"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`

	// Relations
	School         *School         `json:"school,omitempty" gorm:"foreignKey:SchoolID"`
	ClassSchedules []ClassSchedule `json:"class_schedules,omitempty" gorm:"foreignKey:SubjectID"`
}

func (Subject) TableName() string {
	return "subjects"
}
