package models

import (
	"time"
)

type School struct {
	ID            string     `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();unique;not null" json:"id"`
	NPSN          string     `gorm:"type:varchar(20);not null;uniqueIndex;comment:'Nomor Pokok Sekolah Nasional'" json:"npsn"`
	Name          string     `gorm:"type:varchar(255);not null" json:"name"`
	Email         string     `gorm:"type:varchar(255);not null;uniqueIndex" json:"email"`
	Phone         string     `gorm:"type:varchar(20);not null" json:"phone"`
	Fax           *string    `gorm:"type:varchar(20)" json:"fax"`
	Website       *string    `gorm:"type:varchar(255)" json:"website"`
	StandFrom     *time.Time `gorm:"type:date;comment:'Tanggal berdiri'" json:"stand_from"`
	Akreditasi    string     `gorm:"type:varchar(20);not null;default:'belum_terakreditasi';check:akreditasi IN ('A','B','C','belum_terakreditasi')" json:"akreditasi"`
	KepalaSekolah *string    `gorm:"type:varchar(255);comment:'Nama kepala sekolah'" json:"kepala_sekolah"`
	IsActive      bool       `gorm:"not null;default:true;index" json:"is_active"`
	CreatedAt     time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt     time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`

	// Relations
	Majors             []Major             `json:"majors,omitempty" gorm:"foreignKey:SchoolID"`
	Subjects           []Subject           `json:"subjects,omitempty" gorm:"foreignKey:SchoolID"`
	Teachers           []Teacher           `json:"teachers,omitempty" gorm:"foreignKey:SchoolID"`
	Classes            []Class             `json:"classes,omitempty" gorm:"foreignKey:SchoolID"`
	SafeAreas          []SafeArea          `json:"safe_areas,omitempty" gorm:"foreignKey:SchoolID"`
	AcademicCalendars  []AcademicCalendar  `json:"academic_calendars,omitempty" gorm:"foreignKey:SchoolID"`
	AttendanceSettings *AttendanceSettings `json:"attendance_settings,omitempty" gorm:"foreignKey:SchoolID"`
	SchoolAddress      *SchoolAddress      `json:"school_address,omitempty" gorm:"foreignKey:SchoolID"`
}

func (School) TableName() string {
	return "schools"
}
