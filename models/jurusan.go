package models

type Jurusan struct {
	BaseModel
	Code        string  `json:"code" gorm:"type:varchar(10);not null;uniqueIndex;comment:'Kode jurusan (contoh: RPL, TKJ)'"`
	Name        string  `json:"name" gorm:"type:varchar(255);not null;comment:'Nama lengkap jurusan'"`
	SchoolID    string  `json:"school_id" gorm:"type:varchar(255);not null;index"`
	Description *string `json:"description" gorm:"type:text"`
	IsActive    bool    `json:"is_active" gorm:"not null;default:true;index"`

	School  *School `json:"school,omitempty" gorm:"foreignKey:SchoolID"`
	Classes []Class `json:"classes,omitempty" gorm:"foreignKey:JurusanID"`
}

func (Jurusan) TableName() string {
	return "jurusan"
}
