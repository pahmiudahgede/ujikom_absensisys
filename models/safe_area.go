package models

type SafeArea struct {
	BaseModel
	SchoolID    string  `json:"school_id" gorm:"type:varchar(255);not null;index"`
	Name        string  `json:"name" gorm:"type:varchar(255);not null;comment:'Nama area aman'"`
	Latitude    float64 `json:"latitude" gorm:"type:decimal(10,8);not null"`
	Longitude   float64 `json:"longitude" gorm:"type:decimal(11,8);not null"`
	Radius      float64 `json:"radius" gorm:"type:decimal(8,2);not null;comment:'Radius dalam meter'"`
	Description *string `json:"description" gorm:"type:text"`
	IsActive    bool    `json:"is_active" gorm:"not null;default:true;index"`

	School *School `json:"school,omitempty" gorm:"foreignKey:SchoolID"`
}

func (SafeArea) TableName() string {
	return "safe_area"
}
