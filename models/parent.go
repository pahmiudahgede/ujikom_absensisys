package models

type Parent struct {
	BaseModel
	StudentID string `json:"student_id" gorm:"type:varchar(255);not null;index"`
	Fullname  string `json:"fullname" gorm:"type:varchar(255);not null"`
	Phone     string `json:"phone" gorm:"type:varchar(20);not null"`
	Job       string `json:"job" gorm:"type:varchar(100);not null"`
	Sebagai   string `json:"sebagai" gorm:"type:enum('ayah','ibu','wali');not null;comment:'Hubungan dengan siswa'"`

	Student *Student `json:"student,omitempty" gorm:"foreignKey:StudentID"`
}

func (Parent) TableName() string {
	return "parents"
}
