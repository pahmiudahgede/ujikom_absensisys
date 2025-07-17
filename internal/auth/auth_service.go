package auth

import (
	"absensibe/middleware"
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type StudentService interface {
	Login(ctx context.Context, req LoginRequest, c *fiber.Ctx) (*middleware.SessionData, error)
	Logout(ctx context.Context, sessionID string) error
}

type studentService struct {
	repo StudentRepository
}

func NewStudentService(repo StudentRepository) StudentService {
	return &studentService{
		repo: repo,
	}
}

func (s *studentService) Login(ctx context.Context, req LoginRequest, c *fiber.Ctx) (*middleware.SessionData, error) {

	student, err := s.repo.GetByNISN(ctx, req.NISN)
	if err != nil {
		return nil, fmt.Errorf("invalid NISN or password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(student.Password), []byte(req.Password))
	if err != nil {
		return nil, fmt.Errorf("invalid NISN or password")
	}

	sessionData, err := middleware.CreateSession(
		student.ID,
		student.NISN,
		student.Fullname,
		"student",
		c,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create session: %w", err)
	}

	response := &middleware.SessionData{
		UserID:      student.ID,
		NISN:        student.NISN,
		Name:        student.Fullname,
		AccessToken: sessionData.AccessToken,
		SessionID:   sessionData.SessionID,
		DeviceInfo:  req.DeviceInfo,
	}

	return response, nil
}

func (s *studentService) Logout(ctx context.Context, sessionID string) error {
	err := middleware.DestroySession(sessionID)
	if err != nil {
		return fmt.Errorf("failed to logout: %w", err)
	}

	return nil
}
