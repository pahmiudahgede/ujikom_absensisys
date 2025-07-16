package models

type Address struct {
	BaseModel
	ReferenceID   string   `json:"reference_id" gorm:"type:varchar(255);not null;index:idx_reference"`
	ReferenceType string   `json:"reference_type" gorm:"type:enum('student','school','teacher');not null;index:idx_reference;comment:'Jenis referensi'"`
	Province      string   `json:"province" gorm:"type:varchar(100);not null"`
	District      string   `json:"district" gorm:"type:varchar(100);not null"`
	Subdistrict   string   `json:"subdistrict" gorm:"type:varchar(100);not null"`
	Hamlet        string   `json:"hamlet" gorm:"type:varchar(100);not null"`
	Village       string   `json:"village" gorm:"type:varchar(100);not null"`
	Neighbourhood string   `json:"neighbourhood" gorm:"type:varchar(100);not null"`
	PostalCode    string   `json:"postal_code" gorm:"type:varchar(10);not null"`
	Detail        *string  `json:"detail" gorm:"type:text"`
	Religion      string   `json:"religion" gorm:"type:enum('islam','kristen','katholik','hindu','budha','konghucu');not null"`
	Latitude      *float64 `json:"latitude" gorm:"type:decimal(10,8)"`
	Longitude     *float64 `json:"longitude" gorm:"type:decimal(11,8)"`
}

func (Address) TableName() string {
	return "addresses"
}
