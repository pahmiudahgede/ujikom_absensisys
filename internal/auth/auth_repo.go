package auth

import (
	"absensibe/models"
	"context"
	"fmt"

	"gorm.io/gorm"
)

type StudentRepository interface {
	GetByNISN(ctx context.Context, nisn string) (*models.Student, error)
}

type studentRepository struct {
	db *gorm.DB
}

func NewStudentRepository(db *gorm.DB) StudentRepository {
	return &studentRepository{
		db: db,
	}
}

func (r *studentRepository) GetByNISN(ctx context.Context, nisn string) (*models.Student, error) {
	var student models.Student
	err := r.db.WithContext(ctx).Where("nisn = ? AND status = ?", nisn, "aktif").First(&student).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("student not found")
		}
		return nil, fmt.Errorf("failed to get student: %w", err)
	}

	return &student, nil
}
