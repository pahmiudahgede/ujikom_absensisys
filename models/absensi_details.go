package models

type AbsensiDetails struct {
	BaseModel
	AbsensiID         string   `json:"absensi_id" gorm:"type:varchar(255);not null;uniqueIndex"`
	CheckinLatitude   *float64 `json:"checkin_latitude" gorm:"type:decimal(10,8)"`
	CheckinLongitude  *float64 `json:"checkin_longitude" gorm:"type:decimal(11,8)"`
	CheckoutLatitude  *float64 `json:"checkout_latitude" gorm:"type:decimal(10,8)"`
	CheckoutLongitude *float64 `json:"checkout_longitude" gorm:"type:decimal(11,8)"`
	CheckinPhoto      *string  `json:"checkin_photo" gorm:"type:text;comment:'Foto saat check-in'"`
	CheckoutPhoto     *string  `json:"checkout_photo" gorm:"type:text;comment:'Foto saat check-out'"`
	CheckinNote       *string  `json:"checkin_note" gorm:"type:text;comment:'Catatan check-in'"`
	CheckoutNote      *string  `json:"checkout_note" gorm:"type:text;comment:'Catatan check-out'"`

	Absensi *Absensi `json:"absensi,omitempty" gorm:"foreignKey:AbsensiID"`
}

func (AbsensiDetails) TableName() string {
	return "absensi_details"
}
