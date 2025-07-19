package models

import "time"

type School struct {
	ID        string    `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();unique;not null" json:"id"`
	NPSN      string    `gorm:"type:varchar(20);not null;uniqueIndex;comment:'Nomor Pokok Sekolah Nasional'" json:"npsn"`
	Name      string    `gorm:"type:varchar(255);not null" json:"name"`
	Email     string    `gorm:"type:varchar(255);not null;uniqueIndex" json:"email"`
	Phone     string    `gorm:"type:varchar(20);not null" json:"phone"`
	Website   *string   `gorm:"type:varchar(255)" json:"website"`
	Address   string    `gorm:"type:text;not null" json:"address"`
	Principal string    `gorm:"type:varchar(255);not null;comment:'Nama kepala sekolah'" json:"principal"`
	Latitude  float64   `gorm:"type:decimal(10,8);not null;comment:'Koordinat sekolah'" json:"latitude"`
	Longitude float64   `gorm:"type:decimal(11,8);not null;comment:'Koordinat sekolah'" json:"longitude"`
	IsActive  bool      `gorm:"not null;default:true;index" json:"is_active"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`

	// Relations
	Classes            []Class             `json:"classes,omitempty" gorm:"foreignKey:SchoolID"`
	Teachers           []Teacher           `json:"teachers,omitempty" gorm:"foreignKey:SchoolID"`
	Subjects           []Subject           `json:"subjects,omitempty" gorm:"foreignKey:SchoolID"`
	SafeAreas          []SafeArea          `json:"safe_areas,omitempty" gorm:"foreignKey:SchoolID"`
	Holidays           []Holiday           `json:"holidays,omitempty" gorm:"foreignKey:SchoolID"`
	AttendanceSettings *AttendanceSettings `json:"attendance_settings,omitempty" gorm:"foreignKey:SchoolID"`
}

func (School) TableName() string {
	return "schools"
}
