package models

import "time"

type Holiday struct {
	BaseModel
	SchoolID    string    `json:"school_id" gorm:"type:varchar(255);not null;index:idx_school_date"`
	Name        string    `json:"name" gorm:"type:varchar(255);not null"`
	Date        time.Time `json:"date" gorm:"type:date;not null;index:idx_school_date"`
	Type        string    `json:"type" gorm:"type:enum('nasional','agama','sekolah');not null;index"`
	Description *string   `json:"description" gorm:"type:text"`

	School *School `json:"school,omitempty" gorm:"foreignKey:SchoolID"`
}

func (Holiday) TableName() string {
	return "holidays"
}
