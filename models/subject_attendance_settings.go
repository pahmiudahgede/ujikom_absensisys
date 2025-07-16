package models

type SubjectAttendanceSettings struct {
	BaseModel
	SubjectID             string `json:"subject_id" gorm:"type:varchar(255);not null;uniqueIndex:idx_subject_class"`
	ClassID               string `json:"class_id" gorm:"type:varchar(255);not null;uniqueIndex:idx_subject_class"`
	CheckInTolerance      int    `json:"check_in_tolerance" gorm:"not null;default:15;comment:'Toleransi keterlambatan dalam menit'"`
	AutoCheckout          bool   `json:"auto_checkout" gorm:"not null;default:true;comment:'Otomatis checkout di akhir jam'"`
	RequirePhoto          bool   `json:"require_photo" gorm:"not null;default:false"`
	RequireLocation       bool   `json:"require_location" gorm:"not null;default:false"`
	MinAttendanceDuration int    `json:"min_attendance_duration" gorm:"not null;default:30;comment:'Durasi minimum kehadiran dalam menit'"`

	Subject *Subject `json:"subject,omitempty" gorm:"foreignKey:SubjectID"`
	Class   *Class   `json:"class,omitempty" gorm:"foreignKey:ClassID"`
}

func (SubjectAttendanceSettings) TableName() string {
	return "subject_attendance_settings"
}
