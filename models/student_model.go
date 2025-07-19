package models

import "time"

type Student struct {
	ID           string    `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();unique;not null" json:"id"`
	NIS          string    `gorm:"type:varchar(20);not null;uniqueIndex;comment:'Nomor Induk Siswa'" json:"nis"`
	NISN         string    `gorm:"type:varchar(20);not null;uniqueIndex;comment:'Nomor Induk Siswa Nasional'" json:"nisn"`
	ClassID      string    `gorm:"type:uuid;not null;index;constraint:OnDelete:CASCADE" json:"class_id"`
	Password     string    `gorm:"type:varchar(255);not null;comment:'Hashed password'" json:"-"`
	Fullname     string    `gorm:"type:varchar(255);not null" json:"fullname"`
	PhotoProfile *string   `gorm:"type:text" json:"photo_profile"`
	Gender       string    `gorm:"type:varchar(1);not null;check:gender IN ('L','P');comment:'L=Laki-laki, P=Perempuan'" json:"gender"`
	BirthDate    time.Time `gorm:"type:date;not null" json:"birth_date"`
	PlaceOfBirth string    `gorm:"type:varchar(100);not null" json:"place_of_birth"`
	Phone        string    `gorm:"type:varchar(20);not null" json:"phone"`
	Address      string    `gorm:"type:text;not null" json:"address"`
	Religion     string    `gorm:"type:varchar(20);not null;check:religion IN ('islam','kristen','katolik','hindu','budha','konghucu')" json:"religion"`

	// Parent Information
	FatherName string `gorm:"type:varchar(255);not null;comment:'Nama Ayah'" json:"father_name"`
	MotherName string `gorm:"type:varchar(255);not null;comment:'Nama Ibu'" json:"mother_name"`

	Status    string    `gorm:"type:varchar(20);not null;default:'aktif';check:status IN ('aktif','nonaktif','lulus','keluar');index" json:"status"`
	EntryYear int       `gorm:"not null;comment:'Tahun masuk'" json:"entry_year"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`

	// Relations
	Class       *Class       `json:"class,omitempty" gorm:"foreignKey:ClassID"`
	Attendances []Attendance `json:"attendances,omitempty" gorm:"foreignKey:StudentID"`
}

func (Student) TableName() string {
	return "students"
}
