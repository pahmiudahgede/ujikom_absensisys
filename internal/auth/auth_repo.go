package auth

import (
	"absensibe/models"
	"context"
	"fmt"

	"gorm.io/gorm"
)

type StudentRepository interface {
	GetByIdentifier(ctx context.Context, identifier string) (*models.Student, error)
}

type studentRepository struct {
	db *gorm.DB
}

func NewStudentRepository(db *gorm.DB) StudentRepository {
	return &studentRepository{
		db: db,
	}
}

func (r *studentRepository) GetByIdentifier(ctx context.Context, identifier string) (*models.Student, error) {
	var student models.Student
	
	// Try to find by NISN or NIS
	err := r.db.WithContext(ctx).
		Where("(nisn = ? OR nis = ?) AND status = ?", identifier, identifier, "aktif").
		First(&student).Error
		
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("student not found")
		}
		return nil, fmt.Errorf("failed to get student: %w", err)
	}

	return &student, nil
}