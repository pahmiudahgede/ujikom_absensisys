package models

type SubjectAttendanceRecap struct {
	BaseModel
	StudentID     string `json:"student_id" gorm:"type:varchar(255);not null;uniqueIndex:idx_student_subject_month_year"`
	SubjectID     string `json:"subject_id" gorm:"type:varchar(255);not null;uniqueIndex:idx_student_subject_month_year"`
	Month         int    `json:"month" gorm:"not null;uniqueIndex:idx_student_subject_month_year;comment:'Bulan (1-12)'"`
	Year          int    `json:"year" gorm:"not null;uniqueIndex:idx_student_subject_month_year;comment:'Tahun'"`
	Present       int    `json:"present" gorm:"not null;default:0"`
	Alpha         int    `json:"alpha" gorm:"not null;default:0"`
	Sick          int    `json:"sick" gorm:"not null;default:0"`
	Permit        int    `json:"permit" gorm:"not null;default:0"`
	Late          int    `json:"late" gorm:"not null;default:0"`
	TotalSessions int    `json:"total_sessions" gorm:"not null;default:0;comment:'Total sesi pembelajaran'"`

	Student *Student `json:"student,omitempty" gorm:"foreignKey:StudentID"`
	Subject *Subject `json:"subject,omitempty" gorm:"foreignKey:SubjectID"`
}

func (SubjectAttendanceRecap) TableName() string {
	return "subject_attendance_recap"
}
