package models

import "time"

type Major struct {
	ID          string    `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();unique;not null" json:"id"`
	Code        string    `gorm:"type:varchar(10);not null;uniqueIndex;comment:'Kode jurusan (contoh: RPL, TKJ)'" json:"code"`
	Name        string    `gorm:"type:varchar(255);not null;comment:'Nama lengkap jurusan'" json:"name"`
	SchoolID    string    `gorm:"type:uuid;not null;index;constraint:OnDelete:CASCADE" json:"school_id"`
	Description *string   `gorm:"type:text" json:"description"`
	IsActive    bool      `gorm:"not null;default:true;index" json:"is_active"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`

	// Relations
	School  *School `json:"school,omitempty" gorm:"foreignKey:SchoolID"`
	Classes []Class `json:"classes,omitempty" gorm:"foreignKey:MajorID"`
}

func (Major) TableName() string {
	return "majors"
}
