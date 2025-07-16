package models

import "time"

type SubjectAttendance struct {
	BaseModel
	SessionID    string     `json:"session_id" gorm:"type:varchar(255);not null;uniqueIndex:idx_session_student"`
	StudentID    string     `json:"student_id" gorm:"type:varchar(255);not null;uniqueIndex:idx_session_student;index"`
	Status       string     `json:"status" gorm:"type:enum('hadir','alpha','sakit','izin','terlambat');not null;default:'alpha';index"`
	CheckInTime  *time.Time `json:"check_in_time" gorm:"type:timestamp;comment:'Waktu masuk kelas'"`
	CheckOutTime *time.Time `json:"check_out_time" gorm:"type:timestamp;comment:'Waktu keluar kelas'"`
	Latitude     *float64   `json:"latitude" gorm:"type:decimal(10,8)"`
	Longitude    *float64   `json:"longitude" gorm:"type:decimal(11,8)"`
	Photo        *string    `json:"photo" gorm:"type:text;comment:'Foto saat absen'"`
	Notes        *string    `json:"notes" gorm:"type:text;comment:'Catatan absensi'"`
	MarkedBy     *string    `json:"marked_by" gorm:"type:varchar(255);index;comment:'Dicatat oleh guru'"`
	MarkedAt     *time.Time `json:"marked_at" gorm:"type:timestamp;comment:'Waktu dicatat'"`

	Session *ClassSession `json:"session,omitempty" gorm:"foreignKey:SessionID"`
	Student *Student      `json:"student,omitempty" gorm:"foreignKey:StudentID"`
	Teacher *Teacher      `json:"teacher,omitempty" gorm:"foreignKey:MarkedBy"`
}

func (SubjectAttendance) TableName() string {
	return "subject_attendance"
}
