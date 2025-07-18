package models

import "time"

type AddressBase struct {
	Province      string   `gorm:"type:varchar(100);not null" json:"province"`
	District      string   `gorm:"type:varchar(100);not null" json:"district"`
	Subdistrict   string   `gorm:"type:varchar(100);not null" json:"subdistrict"`
	Hamlet        string   `gorm:"type:varchar(100);not null" json:"hamlet"`
	Village       string   `gorm:"type:varchar(100);not null" json:"village"`
	Neighbourhood string   `gorm:"type:varchar(100);not null" json:"neighbourhood"`
	PostalCode    string   `gorm:"type:varchar(10);not null" json:"postal_code"`
	Detail        *string  `gorm:"type:text" json:"detail"`
	Religion      string   `gorm:"type:varchar(20);not null;check:religion IN ('islam','kristen','katholik','hindu','budha','konghucu')" json:"religion"`
	Latitude      *float64 `gorm:"type:decimal(10,8)" json:"latitude"`
	Longitude     *float64 `gorm:"type:decimal(11,8)" json:"longitude"`
}

type StudentAddress struct {
	ID        string `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();unique;not null" json:"id"`
	StudentID string `gorm:"type:uuid;not null;uniqueIndex;constraint:OnDelete:CASCADE" json:"student_id"`
	AddressBase
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`

	Student *Student `json:"student,omitempty" gorm:"foreignKey:StudentID"`
}

func (StudentAddress) TableName() string {
	return "student_addresses"
}

type TeacherAddress struct {
	ID        string `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();unique;not null" json:"id"`
	TeacherID string `gorm:"type:uuid;not null;uniqueIndex;constraint:OnDelete:CASCADE" json:"teacher_id"`
	AddressBase
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`

	Teacher *Teacher `json:"teacher,omitempty" gorm:"foreignKey:TeacherID"`
}

func (TeacherAddress) TableName() string {
	return "teacher_addresses"
}

type SchoolAddress struct {
	ID       string `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();unique;not null" json:"id"`
	SchoolID string `gorm:"type:uuid;not null;uniqueIndex;constraint:OnDelete:CASCADE" json:"school_id"`
	AddressBase
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`

	School *School `json:"school,omitempty" gorm:"foreignKey:SchoolID"`
}

func (SchoolAddress) TableName() string {
	return "school_addresses"
}
