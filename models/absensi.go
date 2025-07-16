package models

import "time"

type Absensi struct {
	BaseModel
	StudentID    string     `json:"student_id" gorm:"type:varchar(255);not null;uniqueIndex:idx_student_date"`
	Date         time.Time  `json:"date" gorm:"type:date;not null;uniqueIndex:idx_student_date;index;comment:'Tanggal absensi'"`
	CheckInTime  *time.Time `json:"check_in_time" gorm:"type:timestamp;comment:'Waktu check-in'"`
	CheckOutTime *time.Time `json:"check_out_time" gorm:"type:timestamp;comment:'Waktu check-out'"`
	Status       string     `json:"status" gorm:"type:enum('hadir','alpha','sakit','izin','terlambat');not null;default:'alpha';index"`

	Student        *Student        `json:"student,omitempty" gorm:"foreignKey:StudentID"`
	AbsensiDetails *AbsensiDetails `json:"absensi_details,omitempty" gorm:"foreignKey:AbsensiID"`
}

func (Absensi) TableName() string {
	return "absensi"
}
