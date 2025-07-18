package models

import "time"

type Teacher struct {
	ID           string    `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();unique;not null" json:"id"`
	NIP          string    `gorm:"type:varchar(20);not null;uniqueIndex;comment:'Nomor Induk Pegawai'" json:"nip"`
	SchoolID     string    `gorm:"type:uuid;not null;index;constraint:OnDelete:CASCADE" json:"school_id"`
	Fullname     string    `gorm:"type:varchar(255);not null" json:"fullname"`
	Email        string    `gorm:"type:varchar(255);not null;uniqueIndex" json:"email"`
	Phone        string    `gorm:"type:varchar(20);not null" json:"phone"`
	Password     string    `gorm:"type:varchar(255);not null;comment:'Hashed password'" json:"-"`
	PhotoProfile *string   `gorm:"type:text" json:"photo_profile"`
	Gender       string    `gorm:"type:varchar(1);not null;check:gender IN ('L','P');comment:'L=Laki-laki, P=Perempuan'" json:"gender"`
	BirthDate    time.Time `gorm:"type:date;not null" json:"birth_date"`
	PlaceOfBirth string    `gorm:"type:varchar(100);not null" json:"place_of_birth"`
	EmployeeType string    `gorm:"type:varchar(10);not null;check:employee_type IN ('PNS','PPPK','GTT','GTY');comment:'Jenis kepegawaian'" json:"employee_type"`
	Status       string    `gorm:"type:varchar(20);not null;default:'aktif';check:status IN ('aktif','nonaktif','pensiun','mutasi');index" json:"status"`
	CreatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`

	// Relations
	School                 *School             `json:"school,omitempty" gorm:"foreignKey:SchoolID"`
	ClassesAsWali          []Class             `json:"classes_as_wali,omitempty" gorm:"foreignKey:HomeroomTeacherID"`
	ClassSchedules         []ClassSchedule     `json:"class_schedules,omitempty" gorm:"foreignKey:TeacherID"`
	ClassSessions          []ClassSession      `json:"class_sessions,omitempty" gorm:"foreignKey:CreatedBy"`
	SubjectAttendance      []SubjectAttendance `json:"subject_attendance,omitempty" gorm:"foreignKey:MarkedBy"`
	PermitsApproved        []Permit            `json:"permits_approved,omitempty" gorm:"foreignKey:ApprovedBy"`
	SubjectPermitsApproved []SubjectPermit     `json:"subject_permits_approved,omitempty" gorm:"foreignKey:ApprovedBy"`
	TeacherAddress         *TeacherAddress     `json:"teacher_address,omitempty" gorm:"foreignKey:TeacherID"`
}

func (Teacher) TableName() string {
	return "teachers"
}
