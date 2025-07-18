package models

import "time"

type Permit struct {
	ID             string    `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();unique;not null" json:"id"`
	StudentID      string    `gorm:"type:uuid;not null;index:idx_student_date;constraint:OnDelete:CASCADE" json:"student_id"`
	Type           string    `gorm:"type:varchar(10);not null;check:type IN ('sakit','izin')" json:"type"`
	Reason         string    `gorm:"type:text;not null" json:"reason"`
	StartDate      time.Time `gorm:"type:date;not null;index:idx_student_date" json:"start_date"`
	EndDate        time.Time `gorm:"type:date;not null;index:idx_student_date" json:"end_date"`
	DocumentProof  *string   `gorm:"type:text;comment:'Dokumen pendukung (surat sakit, dll)'" json:"document_proof"`
	ApprovedBy     *string   `gorm:"type:uuid;index" json:"approved_by"`
	ApprovalStatus string    `gorm:"type:varchar(20);not null;default:'pending';check:approval_status IN ('pending','approved','rejected');index" json:"approval_status"`
	ApprovalNotes  *string   `gorm:"type:text" json:"approval_notes"`
	CreatedAt      time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt      time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`

	// Relations
	Student           *Student `json:"student,omitempty" gorm:"foreignKey:StudentID"`
	ApprovedByTeacher *Teacher `json:"approved_by_teacher,omitempty" gorm:"foreignKey:ApprovedBy"`
}

func (Permit) TableName() string {
	return "permits"
}

type SubjectPermit struct {
	ID             string    `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();unique;not null" json:"id"`
	StudentID      string    `gorm:"type:uuid;not null;index:idx_student_date;constraint:OnDelete:CASCADE" json:"student_id"`
	SessionID      *string   `gorm:"type:uuid;index" json:"session_id"`
	SubjectID      string    `gorm:"type:uuid;not null;index;constraint:OnDelete:CASCADE" json:"subject_id"`
	Type           string    `gorm:"type:varchar(10);not null;check:type IN ('sakit','izin')" json:"type"`
	Reason         string    `gorm:"type:text;not null" json:"reason"`
	StartDate      time.Time `gorm:"type:date;not null;index:idx_student_date" json:"start_date"`
	EndDate        time.Time `gorm:"type:date;not null;index:idx_student_date" json:"end_date"`
	DocumentProof  *string   `gorm:"type:text;comment:'Dokumen pendukung (surat sakit, dll)'" json:"document_proof"`
	ApprovedBy     *string   `gorm:"type:uuid;index" json:"approved_by"`
	ApprovalStatus string    `gorm:"type:varchar(20);not null;default:'pending';check:approval_status IN ('pending','approved','rejected');index" json:"approval_status"`
	ApprovalNotes  *string   `gorm:"type:text" json:"approval_notes"`
	CreatedAt      time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt      time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`

	// Relations
	Student           *Student      `json:"student,omitempty" gorm:"foreignKey:StudentID"`
	Session           *ClassSession `json:"session,omitempty" gorm:"foreignKey:SessionID"`
	Subject           *Subject      `json:"subject,omitempty" gorm:"foreignKey:SubjectID"`
	ApprovedByTeacher *Teacher      `json:"approved_by_teacher,omitempty" gorm:"foreignKey:ApprovedBy"`
}

func (SubjectPermit) TableName() string {
	return "subject_permits"
}
