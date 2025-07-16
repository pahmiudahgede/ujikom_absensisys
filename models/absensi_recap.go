package models

type AbsensiRecap struct {
	BaseModel
	StudentID string `json:"student_id" gorm:"type:varchar(255);not null;uniqueIndex:idx_student_month_year"`
	Month     int    `json:"month" gorm:"not null;uniqueIndex:idx_student_month_year;comment:'Bulan (1-12)'"`
	Year      int    `json:"year" gorm:"not null;uniqueIndex:idx_student_month_year;comment:'Tahun'"`
	Present   int    `json:"present" gorm:"not null;default:0"`
	Alpha     int    `json:"alpha" gorm:"not null;default:0"`
	Sick      int    `json:"sick" gorm:"not null;default:0"`
	Permit    int    `json:"permit" gorm:"not null;default:0"`
	Late      int    `json:"late" gorm:"not null;default:0"`

	Student *Student `json:"student,omitempty" gorm:"foreignKey:StudentID"`
}

func (AbsensiRecap) TableName() string {
	return "absensi_recap"
}
