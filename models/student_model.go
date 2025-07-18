package models

import "time"

type Student struct {
	ID           string    `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();unique;not null" json:"id"`
	NISN         string    `gorm:"type:varchar(20);not null;uniqueIndex;comment:'Nomor Induk Siswa Nasional'" json:"nisn"`
	NIS          string    `gorm:"type:varchar(20);not null;uniqueIndex;comment:'Nomor Induk Siswa'" json:"nis"`
	ClassID      string    `gorm:"type:uuid;not null;index;constraint:OnDelete:CASCADE" json:"class_id"`
	Password     string    `gorm:"type:varchar(255);not null;comment:'Hashed password'" json:"-"`
	Fullname     string    `gorm:"type:varchar(255);not null" json:"fullname"`
	PhotoProfile *string   `gorm:"type:text" json:"photo_profile"`
	Gender       string    `gorm:"type:varchar(1);not null;check:gender IN ('L','P');comment:'L=Laki-laki, P=Perempuan'" json:"gender"`
	BirthDate    time.Time `gorm:"type:date;not null" json:"birth_date"`
	PlaceOfBirth string    `gorm:"type:varchar(100);not null" json:"place_of_birth"`
	Phone        string    `gorm:"type:varchar(20);not null" json:"phone"`
	Status       string    `gorm:"type:varchar(20);not null;default:'aktif';check:status IN ('aktif','nonaktif','lulus','keluar','mutasi');index" json:"status"`
	EntryYear    int       `gorm:"not null;comment:'Tahun masuk'" json:"entry_year"`
	CreatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`

	// Relations
	Class                   *Class                   `json:"class,omitempty" gorm:"foreignKey:ClassID"`
	Parents                 []Parent                 `json:"parents,omitempty" gorm:"foreignKey:StudentID"`
	StudentAddress          *StudentAddress          `json:"student_address,omitempty" gorm:"foreignKey:StudentID"`
	Attendances             []Attendance             `json:"attendances,omitempty" gorm:"foreignKey:StudentID"`
	SubjectAttendances      []SubjectAttendance      `json:"subject_attendances,omitempty" gorm:"foreignKey:StudentID"`
	AttendanceRecaps        []AttendanceRecap        `json:"attendance_recaps,omitempty" gorm:"foreignKey:StudentID"`
	SubjectAttendanceRecaps []SubjectAttendanceRecap `json:"subject_attendance_recaps,omitempty" gorm:"foreignKey:StudentID"`
	Permits                 []Permit                 `json:"permits,omitempty" gorm:"foreignKey:StudentID"`
	SubjectPermits          []SubjectPermit          `json:"subject_permits,omitempty" gorm:"foreignKey:StudentID"`
}

func (Student) TableName() string {
	return "students"
}
