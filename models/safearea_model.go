package models

import "time"

type SafeArea struct {
	ID          string    `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();unique;not null" json:"id"`
	SchoolID    string    `gorm:"type:uuid;not null;index;constraint:OnDelete:CASCADE" json:"school_id"`
	Name        string    `gorm:"type:varchar(255);not null;comment:'Nama area aman'" json:"name"`
	Latitude    float64   `gorm:"type:decimal(10,8);not null" json:"latitude"`
	Longitude   float64   `gorm:"type:decimal(11,8);not null" json:"longitude"`
	Radius      float64   `gorm:"type:decimal(8,2);not null;comment:'Radius dalam meter'" json:"radius"`
	Description *string   `gorm:"type:text" json:"description"`
	IsActive    bool      `gorm:"not null;default:true;index" json:"is_active"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`

	// Relations
	School *School `json:"school,omitempty" gorm:"foreignKey:SchoolID"`
}

func (SafeArea) TableName() string {
	return "safe_areas"
}
