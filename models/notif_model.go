package models

import "time"

type Notification struct {
	ID            string     `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();unique;not null" json:"id"`
	RecipientID   string     `gorm:"type:uuid;not null;index:idx_recipient" json:"recipient_id"`
	RecipientType string     `gorm:"type:varchar(20);not null;index:idx_recipient;check:recipient_type IN ('student','teacher','parent');comment:'ID penerima (siswa/guru/orang tua)'" json:"recipient_type"`
	Title         string     `gorm:"type:varchar(255);not null" json:"title"`
	Message       string     `gorm:"type:text;not null" json:"message"`
	Type          string     `gorm:"type:varchar(20);not null;check:type IN ('attendance','permit','announcement','reminder');index" json:"type"`
	ReferenceID   *string    `gorm:"type:uuid;comment:'ID referensi terkait'" json:"reference_id"`
	ReferenceType *string    `gorm:"type:varchar(50);comment:'Jenis referensi'" json:"reference_type"`
	IsRead        bool       `gorm:"not null;default:false;index" json:"is_read"`
	SentAt        time.Time  `gorm:"not null;index;default:CURRENT_TIMESTAMP" json:"sent_at"`
	ReadAt        *time.Time `gorm:"type:timestamp" json:"read_at"`
	CreatedAt     time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt     time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (Notification) TableName() string {
	return "notifications"
}

type AttendanceSettings struct {
	ID              string    `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();unique;not null" json:"id"`
	SchoolID        string    `gorm:"type:uuid;not null;uniqueIndex;constraint:OnDelete:CASCADE" json:"school_id"`
	CheckinStart    time.Time `gorm:"type:time;not null;default:'06:00:00';comment:'Format: HH:mm:ss'" json:"checkin_start"`
	CheckinEnd      time.Time `gorm:"type:time;not null;default:'07:30:00';comment:'Format: HH:mm:ss'" json:"checkin_end"`
	CheckoutStart   time.Time `gorm:"type:time;not null;default:'15:00:00';comment:'Format: HH:mm:ss'" json:"checkout_start"`
	CheckoutEnd     time.Time `gorm:"type:time;not null;default:'17:00:00';comment:'Format: HH:mm:ss'" json:"checkout_end"`
	LateTolerance   int       `gorm:"not null;default:15;comment:'Toleransi terlambat dalam menit'" json:"late_tolerance"`
	RequirePhoto    bool      `gorm:"not null;default:true" json:"require_photo"`
	RequireLocation bool      `gorm:"not null;default:true" json:"require_location"`
	MaxDistance     int       `gorm:"not null;default:100;comment:'Jarak maksimal dalam meter'" json:"max_distance"`
	CreatedAt       time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt       time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`

	// Relations
	School *School `json:"school,omitempty" gorm:"foreignKey:SchoolID"`
}

func (AttendanceSettings) TableName() string {
	return "attendance_settings"
}
