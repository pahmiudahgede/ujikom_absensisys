package models

import "time"

type AttendanceSettings struct {
	ID              string    `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();unique;not null" json:"id"`
	SchoolID        string    `gorm:"type:uuid;not null;uniqueIndex;constraint:OnDelete:CASCADE" json:"school_id"`
	CheckInStart    time.Time `gorm:"type:time;not null;comment:'Mulai bisa absen masuk'" json:"check_in_start"`
	CheckInEnd      time.Time `gorm:"type:time;not null;comment:'Batas akhir absen masuk'" json:"check_in_end"`
	CheckOutStart   time.Time `gorm:"type:time;not null;comment:'Mulai bisa absen pulang'" json:"check_out_start"`
	CheckOutEnd     time.Time `gorm:"type:time;not null;comment:'Batas akhir absen pulang'" json:"check_out_end"`
	LateTolerance   int       `gorm:"not null;default:15;comment:'Toleransi terlambat dalam menit'" json:"late_tolerance"`
	RequirePhoto    bool      `gorm:"not null;default:true" json:"require_photo"`
	RequireLocation bool      `gorm:"not null;default:true" json:"require_location"`
	MaxDistance     int       `gorm:"not null;default:100;comment:'Jarak maksimal dari safe area dalam meter'" json:"max_distance"`
	CreatedAt       time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt       time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`

	// Relations
	School *School `json:"school,omitempty" gorm:"foreignKey:SchoolID"`
}

func (AttendanceSettings) TableName() string {
	return "attendance_settings"
}
