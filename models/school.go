package models

import "time"

type School struct {
	BaseModel
	NPSN          string     `json:"npsn" gorm:"type:varchar(20);not null;uniqueIndex;comment:'Nomor Pokok Sekolah Nasional'"`
	Name          string     `json:"name" gorm:"type:varchar(255);not null"`
	Email         string     `json:"email" gorm:"type:varchar(255);not null;uniqueIndex"`
	Phone         string     `json:"phone" gorm:"type:varchar(20);not null"`
	Fax           *string    `json:"fax" gorm:"type:varchar(20)"`
	Website       *string    `json:"website" gorm:"type:varchar(255)"`
	StandFrom     *time.Time `json:"stand_from" gorm:"type:date;comment:'Tanggal berdiri'"`
	Akreditasi    string     `json:"akreditasi" gorm:"type:enum('A','B','C','belum_terakreditasi');not null;default:'belum_terakreditasi'"`
	KepalaSekolah *string    `json:"kepala_sekolah" gorm:"type:varchar(255);comment:'Nama kepala sekolah'"`
	IsActive      bool       `json:"is_active" gorm:"not null;default:true;index"`

	Jurusan            []Jurusan           `json:"jurusan,omitempty" gorm:"foreignKey:SchoolID"`
	Subjects           []Subject           `json:"subjects,omitempty" gorm:"foreignKey:SchoolID"`
	Teachers           []Teacher           `json:"teachers,omitempty" gorm:"foreignKey:SchoolID"`
	Classes            []Class             `json:"classes,omitempty" gorm:"foreignKey:SchoolID"`
	SafeAreas          []SafeArea          `json:"safe_areas,omitempty" gorm:"foreignKey:SchoolID"`
	AcademicCalendars  []AcademicCalendar  `json:"academic_calendars,omitempty" gorm:"foreignKey:SchoolID"`
	Holidays           []Holiday           `json:"holidays,omitempty" gorm:"foreignKey:SchoolID"`
	AttendanceSettings *AttendanceSettings `json:"attendance_settings,omitempty" gorm:"foreignKey:SchoolID"`
	Addresses          []Address           `json:"addresses,omitempty" gorm:"foreignKey:ReferenceID;where:reference_type = 'school'"`
}

func (School) TableName() string {
	return "schools"
}
