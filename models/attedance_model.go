package models

import "time"

type Attendance struct {
	ID           string     `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();unique;not null" json:"id"`
	StudentID    string     `gorm:"type:uuid;not null;uniqueIndex:idx_student_date;constraint:OnDelete:CASCADE" json:"student_id"`
	Date         time.Time  `gorm:"type:date;not null;uniqueIndex:idx_student_date;index;comment:'Tanggal absensi'" json:"date"`
	CheckInTime  *time.Time `gorm:"type:timestamp;comment:'Waktu check-in'" json:"check_in_time"`
	CheckOutTime *time.Time `gorm:"type:timestamp;comment:'Waktu check-out'" json:"check_out_time"`
	Status       string     `gorm:"type:varchar(20);not null;default:'alpha';check:status IN ('hadir','alpha','sakit','izin','terlambat');index" json:"status"`
	CreatedAt    time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt    time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`

	// Relations
	Student           *Student           `json:"student,omitempty" gorm:"foreignKey:StudentID"`
	AttendanceDetails *AttendanceDetails `json:"attendance_details,omitempty" gorm:"foreignKey:AttendanceID"`
}

func (Attendance) TableName() string {
	return "attendances"
}

type AttendanceDetails struct {
	ID                string    `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();unique;not null" json:"id"`
	AttendanceID      string    `gorm:"type:uuid;not null;uniqueIndex;constraint:OnDelete:CASCADE" json:"attendance_id"`
	CheckinLatitude   *float64  `gorm:"type:decimal(10,8)" json:"checkin_latitude"`
	CheckinLongitude  *float64  `gorm:"type:decimal(11,8)" json:"checkin_longitude"`
	CheckoutLatitude  *float64  `gorm:"type:decimal(10,8)" json:"checkout_latitude"`
	CheckoutLongitude *float64  `gorm:"type:decimal(11,8)" json:"checkout_longitude"`
	CheckinPhoto      *string   `gorm:"type:text;comment:'Foto saat check-in'" json:"checkin_photo"`
	CheckoutPhoto     *string   `gorm:"type:text;comment:'Foto saat check-out'" json:"checkout_photo"`
	CheckinNote       *string   `gorm:"type:text;comment:'Catatan check-in'" json:"checkin_note"`
	CheckoutNote      *string   `gorm:"type:text;comment:'Catatan check-out'" json:"checkout_note"`
	CreatedAt         time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt         time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`

	// Relations
	Attendance *Attendance `json:"attendance,omitempty" gorm:"foreignKey:AttendanceID"`
}

func (AttendanceDetails) TableName() string {
	return "attendance_details"
}

type AttendanceRecap struct {
	ID        string    `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();unique;not null" json:"id"`
	StudentID string    `gorm:"type:uuid;not null;uniqueIndex:idx_student_month_year;constraint:OnDelete:CASCADE" json:"student_id"`
	Month     int       `gorm:"not null;uniqueIndex:idx_student_month_year;comment:'Bulan (1-12)'" json:"month"`
	Year      int       `gorm:"not null;uniqueIndex:idx_student_month_year;comment:'Tahun'" json:"year"`
	Present   int       `gorm:"not null;default:0" json:"present"`
	Alpha     int       `gorm:"not null;default:0" json:"alpha"`
	Sick      int       `gorm:"not null;default:0" json:"sick"`
	Permit    int       `gorm:"not null;default:0" json:"permit"`
	Late      int       `gorm:"not null;default:0" json:"late"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`

	// Relations
	Student *Student `json:"student,omitempty" gorm:"foreignKey:StudentID"`
}

func (AttendanceRecap) TableName() string {
	return "attendance_recaps"
}

type SubjectAttendance struct {
	ID          string     `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();unique;not null" json:"id"`
	SessionID   string     `gorm:"type:uuid;not null;uniqueIndex:idx_session_student;constraint:OnDelete:CASCADE" json:"session_id"`
	StudentID   string     `gorm:"type:uuid;not null;uniqueIndex:idx_session_student;index;constraint:OnDelete:CASCADE" json:"student_id"`
	Status      string     `gorm:"type:varchar(20);not null;default:'alpha';check:status IN ('hadir','alpha','sakit','izin','terlambat');index" json:"status"`
	CheckInTime *time.Time `gorm:"type:timestamp;comment:'Waktu di kelas'" json:"check_in_time"`
	Latitude    *float64   `gorm:"type:decimal(10,8)" json:"latitude"`
	Longitude   *float64   `gorm:"type:decimal(11,8)" json:"longitude"`
	Photo       *string    `gorm:"type:text;comment:'Foto saat absen'" json:"photo"`
	Notes       *string    `gorm:"type:text;comment:'Catatan absensi'" json:"notes"`
	MarkedBy    *string    `gorm:"type:uuid;index;comment:'Dicatat oleh guru'" json:"marked_by"`
	MarkedAt    *time.Time `gorm:"type:timestamp;comment:'Waktu dicatat'" json:"marked_at"`
	CreatedAt   time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`

	// Relations
	Session *ClassSession `json:"session,omitempty" gorm:"foreignKey:SessionID"`
	Student *Student      `json:"student,omitempty" gorm:"foreignKey:StudentID"`
	Teacher *Teacher      `json:"teacher,omitempty" gorm:"foreignKey:MarkedBy"`
}

func (SubjectAttendance) TableName() string {
	return "subject_attendances"
}

type SubjectAttendanceRecap struct {
	ID            string    `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();unique;not null" json:"id"`
	StudentID     string    `gorm:"type:uuid;not null;uniqueIndex:idx_student_subject_month_year;constraint:OnDelete:CASCADE" json:"student_id"`
	SubjectID     string    `gorm:"type:uuid;not null;uniqueIndex:idx_student_subject_month_year;constraint:OnDelete:CASCADE" json:"subject_id"`
	Month         int       `gorm:"not null;uniqueIndex:idx_student_subject_month_year;comment:'Bulan (1-12)'" json:"month"`
	Year          int       `gorm:"not null;uniqueIndex:idx_student_subject_month_year;comment:'Tahun'" json:"year"`
	Present       int       `gorm:"not null;default:0" json:"present"`
	Alpha         int       `gorm:"not null;default:0" json:"alpha"`
	Sick          int       `gorm:"not null;default:0" json:"sick"`
	Permit        int       `gorm:"not null;default:0" json:"permit"`
	Late          int       `gorm:"not null;default:0" json:"late"`
	TotalSessions int       `gorm:"not null;default:0;comment:'Total sesi pembelajaran'" json:"total_sessions"`
	CreatedAt     time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt     time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`

	// Relations
	Student *Student `json:"student,omitempty" gorm:"foreignKey:StudentID"`
	Subject *Subject `json:"subject,omitempty" gorm:"foreignKey:SubjectID"`
}

func (SubjectAttendanceRecap) TableName() string {
	return "subject_attendance_recaps"
}

type SubjectAttendanceSettings struct {
	ID                    string    `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();unique;not null" json:"id"`
	SubjectID             string    `gorm:"type:uuid;not null;uniqueIndex:idx_subject_class;constraint:OnDelete:CASCADE" json:"subject_id"`
	ClassID               string    `gorm:"type:uuid;not null;uniqueIndex:idx_subject_class;constraint:OnDelete:CASCADE" json:"class_id"`
	CheckInTolerance      int       `gorm:"not null;default:15;comment:'Toleransi keterlambatan dalam menit'" json:"check_in_tolerance"`
	AutoCheckout          bool      `gorm:"not null;default:true;comment:'Otomatis checkout di akhir jam'" json:"auto_checkout"`
	RequirePhoto          bool      `gorm:"not null;default:false" json:"require_photo"`
	RequireLocation       bool      `gorm:"not null;default:false" json:"require_location"`
	MinAttendanceDuration int       `gorm:"not null;default:30;comment:'Durasi minimum kehadiran dalam menit'" json:"min_attendance_duration"`
	CreatedAt             time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt             time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`

	// Relations
	Subject *Subject `json:"subject,omitempty" gorm:"foreignKey:SubjectID"`
	Class   *Class   `json:"class,omitempty" gorm:"foreignKey:ClassID"`
}

func (SubjectAttendanceSettings) TableName() string {
	return "subject_attendance_settings"
}
