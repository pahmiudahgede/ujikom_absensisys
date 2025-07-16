package models

import "time"

// Student represents the students table
type Student struct {
	BaseModel
	NISN         string    `json:"nisn" gorm:"type:varchar(20);not null;uniqueIndex;comment:'Nomor Induk Siswa Nasional'"`
	NIS          string    `json:"nis" gorm:"type:varchar(20);not null;uniqueIndex;comment:'Nomor Induk Siswa'"`
	ClassesID    string    `json:"classes_id" gorm:"type:varchar(255);not null;index"`
	Password     string    `json:"-" gorm:"type:varchar(255);not null;comment:'Hashed password'"`
	Fullname     string    `json:"fullname" gorm:"type:varchar(255);not null"`
	PhotoProfile *string   `json:"photo_profile" gorm:"type:text"`
	Gender       string    `json:"gender" gorm:"type:enum('L','P');not null;comment:'L=Laki-laki, P=Perempuan'"`
	BirthDate    time.Time `json:"birth_date" gorm:"type:date;not null"`
	PlaceOfBirth string    `json:"place_of_birth" gorm:"type:varchar(100);not null"`
	Phone        string    `json:"phone" gorm:"type:varchar(20);not null"`
	Status       string    `json:"status" gorm:"type:enum('aktif','nonaktif','lulus','keluar','mutasi');not null;default:'aktif';index"`
	EntryYear    int       `json:"entry_year" gorm:"not null;comment:'Tahun masuk'"`

	// Relations
	Class                   *Class                    `json:"class,omitempty" gorm:"foreignKey:ClassesID"`
	Parents                 []Parent                  `json:"parents,omitempty" gorm:"foreignKey:StudentID"`
	Addresses               []Address                 `json:"addresses,omitempty" gorm:"foreignKey:ReferenceID;where:reference_type = 'student'"`
	Absensi                 []Absensi                 `json:"absensi,omitempty" gorm:"foreignKey:StudentID"`
	SubjectAttendance       []SubjectAttendance       `json:"subject_attendance,omitempty" gorm:"foreignKey:StudentID"`
	AbsensiRecaps           []AbsensiRecap            `json:"absensi_recaps,omitempty" gorm:"foreignKey:StudentID"`
	SubjectAttendanceRecaps []SubjectAttendanceRecap  `json:"subject_attendance_recaps,omitempty" gorm:"foreignKey:StudentID"`
	Permits                 []Permit                  `json:"permits,omitempty" gorm:"foreignKey:StudentID"`
	SubjectPermits          []SubjectPermit           `json:"subject_permits,omitempty" gorm:"foreignKey:StudentID"`
}

func (Student) TableName() string {
	return "students"
}