package models

import "time"

type Teacher struct {
	ID           string    `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();unique;not null" json:"id"`
	TeacherCode  *string   `gorm:"type:varchar(10);uniqueIndex;comment:'Kode guru internal'" json:"teacher_code"`
	NUPTK        string    `gorm:"type:varchar(20);not null;uniqueIndex;comment:'Nomor Unik Pendidik dan Tenaga Kependidikan'" json:"nuptk"`
	SchoolID     string    `gorm:"type:uuid;not null;index;constraint:OnDelete:CASCADE" json:"school_id"`
	Fullname     string    `gorm:"type:varchar(255);not null" json:"fullname"`
	Gender       string    `gorm:"type:varchar(1);not null;constraint:chk_teacher_gender,check:gender IN ('L','P');comment:'L=Laki-laki, P=Perempuan'" json:"gender"`
	Email        string    `gorm:"type:varchar(255);not null;uniqueIndex" json:"email"`
	Phone        string    `gorm:"type:varchar(20);not null" json:"phone"`
	Password     string    `gorm:"type:varchar(255);not null;comment:'Hashed password'" json:"-"`
	PhotoProfile *string   `gorm:"type:text" json:"photo_profile"`
	BirthDate    time.Time `gorm:"type:date;not null" json:"birth_date"`
	PlaceOfBirth string    `gorm:"type:varchar(100);not null" json:"place_of_birth"`
	PositionType string    `gorm:"type:varchar(30);not null;constraint:chk_teacher_position,check:position_type IN ('kepala sekolah','Guru Mapel','Wali Kelas','Guru BK')" json:"position_type"`
	Status       string    `gorm:"type:varchar(20);not null;default:'aktif';constraint:chk_teacher_status,check:status IN ('aktif','nonaktif','pensiun');index" json:"status"`
	CreatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`

	// Relations
	School         *School         `json:"school,omitempty" gorm:"foreignKey:SchoolID"`
	ClassesAsWali  []Class         `json:"classes_as_wali,omitempty" gorm:"foreignKey:HomeroomTeacherID"`
	ClassSchedules []ClassSchedule `json:"class_schedules,omitempty" gorm:"foreignKey:TeacherID"`
}

func (Teacher) TableName() string {
	return "teachers"
}
