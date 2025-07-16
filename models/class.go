package models

type Class struct {
	BaseModel
	Name              string  `json:"name" gorm:"type:varchar(100);not null;comment:'Nama kelas (contoh: X-RPL-1)'"`
	Grade             string  `json:"grade" gorm:"type:enum('X','XI','XII');not null;comment:'Tingkat kelas'"`
	JurusanID         string  `json:"jurusan_id" gorm:"type:varchar(255);not null;index"`
	SchoolID          string  `json:"school_id" gorm:"type:varchar(255);not null;index"`
	HomeroomTeacherID *string `json:"homeroom_teacher_id" gorm:"type:varchar(255);index;comment:'Wali kelas'"`
	MaxStudents       int     `json:"max_students" gorm:"not null;default:36;comment:'Maksimal siswa per kelas'"`
	AcademicYear      string  `json:"academic_year" gorm:"type:varchar(10);not null;index;comment:'Tahun ajaran (contoh: 2024/2025)'"`
	IsActive          bool    `json:"is_active" gorm:"not null;default:true;index"`

	Jurusan                   *Jurusan                    `json:"jurusan,omitempty" gorm:"foreignKey:JurusanID"`
	School                    *School                     `json:"school,omitempty" gorm:"foreignKey:SchoolID"`
	HomeroomTeacher           *Teacher                    `json:"homeroom_teacher,omitempty" gorm:"foreignKey:HomeroomTeacherID"`
	Students                  []Student                   `json:"students,omitempty" gorm:"foreignKey:ClassesID"`
	ClassSchedules            []ClassSchedule             `json:"class_schedules,omitempty" gorm:"foreignKey:ClassID"`
	SubjectAttendanceSettings []SubjectAttendanceSettings `json:"subject_attendance_settings,omitempty" gorm:"foreignKey:ClassID"`
}

func (Class) TableName() string {
	return "classes"
}
