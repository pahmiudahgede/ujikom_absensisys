package models

import "time"

type Holiday struct {
	ID          string    `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();unique;not null" json:"id"`
	SchoolID    string    `gorm:"type:uuid;not null;index;constraint:OnDelete:CASCADE" json:"school_id"`
	Title       string    `gorm:"type:varchar(255);not null;comment:'Nama libur/event'" json:"title"`
	Description *string   `gorm:"type:text" json:"description"`
	StartDate   time.Time `gorm:"type:date;not null;index" json:"start_date"`
	EndDate     time.Time `gorm:"type:date;not null;index" json:"end_date"`
	Type        string    `gorm:"type:varchar(20);not null;check:type IN ('libur_nasional','libur_sekolah','event','ujian');index" json:"type"`
	IsActive    bool      `gorm:"not null;default:true;index" json:"is_active"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`

	// Relations
	School *School `json:"school,omitempty" gorm:"foreignKey:SchoolID"`
}

func (Holiday) TableName() string {
	return "holidays"
}
