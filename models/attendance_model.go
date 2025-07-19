package models

import "time"

type Attendance struct {
	ID             string     `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();unique;not null" json:"id"`
	StudentID      string     `gorm:"type:uuid;not null;uniqueIndex:idx_student_date;constraint:OnDelete:CASCADE" json:"student_id"`
	CheckInTime    *time.Time `gorm:"type:timestamp;comment:'Waktu masuk'" json:"check_in_time"`
	CheckOutTime   *time.Time `gorm:"type:timestamp;comment:'Waktu pulang'" json:"check_out_time"`
	CheckOutStatus string     `gorm:"type:varchar(20);not null;default:'alpha';check:check_out_status IN ('hadir','alpha','sakit','izin');index" json:"check_out_status"`
	Date           time.Time  `gorm:"type:date;not null;uniqueIndex:idx_student_date;index:idx_date_status;comment:'Tanggal absensi'" json:"date"`
	CheckInStatus  string     `gorm:"type:varchar(20);not null;default:'alpha';index:idx_date_status;check:check_in_status IN ('hadir','terlambat','alpha','sakit','izin')" json:"check_in_status"`
	// Check-in data
	CheckInLatitude  *float64 `gorm:"type:decimal(10,8)" json:"check_in_latitude"`
	CheckInLongitude *float64 `gorm:"type:decimal(11,8)" json:"check_in_longitude"`
	CheckInPhoto     *string  `gorm:"type:text;comment:'Foto saat masuk'" json:"check_in_photo"`
	CheckInNote      *string  `gorm:"type:text;comment:'Catatan masuk'" json:"check_in_note"`

	// Check-out data
	CheckOutLatitude  *float64 `gorm:"type:decimal(10,8)" json:"check_out_latitude"`
	CheckOutLongitude *float64 `gorm:"type:decimal(11,8)" json:"check_out_longitude"`
	CheckOutPhoto     *string  `gorm:"type:text;comment:'Foto saat pulang'" json:"check_out_photo"`
	CheckOutNote      *string  `gorm:"type:text;comment:'Catatan pulang'" json:"check_out_note"`

	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`

	// Relations
	Student *Student `json:"student,omitempty" gorm:"foreignKey:StudentID"`
}

func (Attendance) TableName() string {
	return "attendances"
}
