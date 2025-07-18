package models

import "time"

type Parent struct {
	ID        string    `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();unique;not null" json:"id"`
	StudentID string    `gorm:"type:uuid;not null;index;constraint:OnDelete:CASCADE" json:"student_id"`
	Fullname  string    `gorm:"type:varchar(255);not null" json:"fullname"`
	Phone     string    `gorm:"type:varchar(20);not null" json:"phone"`
	Job       string    `gorm:"type:varchar(100);not null" json:"job"`
	Relation  string    `gorm:"type:varchar(10);not null;check:relation IN ('ayah','ibu','wali');comment:'Hubungan dengan siswa'" json:"relation"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`

	// Relations
	Student *Student `json:"student,omitempty" gorm:"foreignKey:StudentID"`
}

func (Parent) TableName() string {
	return "parents"
}
