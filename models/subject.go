package models

type Subject struct {
	BaseModel
	Code        string  `json:"code" gorm:"type:varchar(10);not null;uniqueIndex;comment:'Kode mata pelajaran (contoh: MTK, BIND)'"`
	Name        string  `json:"name" gorm:"type:varchar(255);not null;comment:'Nama mata pelajaran'"`
	SchoolID    string  `json:"school_id" gorm:"type:varchar(255);not null;index"`
	CreditHours int     `json:"credit_hours" gorm:"not null;default:2;comment:'Jumlah jam pelajaran per minggu'"`
	Description *string `json:"description" gorm:"type:text"`
	IsActive    bool    `json:"is_active" gorm:"not null;default:true;index"`

	School                    *School                     `json:"school,omitempty" gorm:"foreignKey:SchoolID"`
	ClassSchedules            []ClassSchedule             `json:"class_schedules,omitempty" gorm:"foreignKey:SubjectID"`
	SubjectAttendanceRecaps   []SubjectAttendanceRecap    `json:"subject_attendance_recaps,omitempty" gorm:"foreignKey:SubjectID"`
	SubjectPermits            []SubjectPermit             `json:"subject_permits,omitempty" gorm:"foreignKey:SubjectID"`
	SubjectAttendanceSettings []SubjectAttendanceSettings `json:"subject_attendance_settings,omitempty" gorm:"foreignKey:SubjectID"`
}

func (Subject) TableName() string {
	return "subjects"
}
