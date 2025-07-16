package models

import "time"

type SubjectPermit struct {
	BaseModel
	StudentID      string    `json:"student_id" gorm:"type:varchar(255);not null;index:idx_student_date"`
	SessionID      *string   `json:"session_id" gorm:"type:varchar(255);index"`
	SubjectID      string    `json:"subject_id" gorm:"type:varchar(255);not null;index"`
	Type           string    `json:"type" gorm:"type:enum('sakit','izin');not null"`
	Reason         string    `json:"reason" gorm:"type:text;not null"`
	StartDate      time.Time `json:"start_date" gorm:"type:date;not null;index:idx_student_date"`
	EndDate        time.Time `json:"end_date" gorm:"type:date;not null;index:idx_student_date"`
	DocumentProof  *string   `json:"document_proof" gorm:"type:text;comment:'Dokumen pendukung (surat sakit, dll)'"`
	ApprovedBy     *string   `json:"approved_by" gorm:"type:varchar(255);index"`
	ApprovalStatus string    `json:"approval_status" gorm:"type:enum('pending','approved','rejected');not null;default:'pending';index"`
	ApprovalNotes  *string   `json:"approval_notes" gorm:"type:text"`

	Student           *Student      `json:"student,omitempty" gorm:"foreignKey:StudentID"`
	Session           *ClassSession `json:"session,omitempty" gorm:"foreignKey:SessionID"`
	Subject           *Subject      `json:"subject,omitempty" gorm:"foreignKey:SubjectID"`
	ApprovedByTeacher *Teacher      `json:"approved_by_teacher,omitempty" gorm:"foreignKey:ApprovedBy"`
}

func (SubjectPermit) TableName() string {
	return "subject_permits"
}
