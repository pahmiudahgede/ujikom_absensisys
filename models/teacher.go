package models

import "time"

type Teacher struct {
	BaseModel
	NIP          string    `json:"nip" gorm:"type:varchar(20);not null;uniqueIndex;comment:'Nomor Induk Pegawai'"`
	SchoolID     string    `json:"school_id" gorm:"type:varchar(255);not null;index"`
	Fullname     string    `json:"fullname" gorm:"type:varchar(255);not null"`
	Email        string    `json:"email" gorm:"type:varchar(255);not null;uniqueIndex"`
	Phone        string    `json:"phone" gorm:"type:varchar(20);not null"`
	Password     string    `json:"-" gorm:"type:varchar(255);not null;comment:'Hashed password'"`
	PhotoProfile *string   `json:"photo_profile" gorm:"type:text"`
	Gender       string    `json:"gender" gorm:"type:enum('L','P');not null;comment:'L=Laki-laki, P=Perempuan'"`
	BirthDate    time.Time `json:"birth_date" gorm:"type:date;not null"`
	PlaceOfBirth string    `json:"place_of_birth" gorm:"type:varchar(100);not null"`
	EmployeeType string    `json:"employee_type" gorm:"type:enum('PNS','PPPK','GTT','GTY');not null;comment:'Jenis kepegawaian'"`
	Status       string    `json:"status" gorm:"type:enum('aktif','nonaktif','pensiun','mutasi');not null;default:'aktif';index"`

	School                 *School             `json:"school,omitempty" gorm:"foreignKey:SchoolID"`
	ClassesAsWali          []Class             `json:"classes_as_wali,omitempty" gorm:"foreignKey:HomeroomTeacherID"`
	ClassSchedules         []ClassSchedule     `json:"class_schedules,omitempty" gorm:"foreignKey:TeacherID"`
	ClassSessions          []ClassSession      `json:"class_sessions,omitempty" gorm:"foreignKey:CreatedBy"`
	SubjectAttendance      []SubjectAttendance `json:"subject_attendance,omitempty" gorm:"foreignKey:MarkedBy"`
	PermitsApproved        []Permit            `json:"permits_approved,omitempty" gorm:"foreignKey:ApprovedBy"`
	SubjectPermitsApproved []SubjectPermit     `json:"subject_permits_approved,omitempty" gorm:"foreignKey:ApprovedBy"`
	Addresses              []Address           `json:"addresses,omitempty" gorm:"foreignKey:ReferenceID;where:reference_type = 'teacher'"`
}

func (Teacher) TableName() string {
	return "teachers"
}
