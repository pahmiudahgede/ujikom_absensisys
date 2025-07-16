package models

import (
	"time"

	"gorm.io/gorm"
)

type Notification struct {
	BaseModel
	RecipientID   string     `json:"recipient_id" gorm:"type:varchar(255);not null;index:idx_recipient"`
	RecipientType string     `json:"recipient_type" gorm:"type:enum('student','teacher','parent');not null;index:idx_recipient;comment:'ID penerima (siswa/guru/orang tua)'"`
	Title         string     `json:"title" gorm:"type:varchar(255);not null"`
	Message       string     `json:"message" gorm:"type:text;not null"`
	Type          string     `json:"type" gorm:"type:enum('attendance','permit','announcement','reminder');not null;index"`
	ReferenceID   *string    `json:"reference_id" gorm:"type:varchar(255);comment:'ID referensi terkait'"`
	ReferenceType *string    `json:"reference_type" gorm:"type:varchar(50);comment:'Jenis referensi'"`
	IsRead        bool       `json:"is_read" gorm:"not null;default:false;index"`
	SentAt        time.Time  `json:"sent_at" gorm:"not null;index"`
	ReadAt        *time.Time `json:"read_at" gorm:"type:timestamp"`
}

func (Notification) TableName() string {
	return "notifications"
}

func (n *Notification) BeforeCreate(tx *gorm.DB) error {

	if err := n.BaseModel.BeforeCreate(tx); err != nil {
		return err
	}

	if n.SentAt.IsZero() {
		n.SentAt = time.Now()
	}

	return nil
}
