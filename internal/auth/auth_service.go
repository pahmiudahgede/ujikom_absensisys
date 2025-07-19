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
	// Get student by identifier (NISN or NIS)
	student, err := s.repo.GetByIdentifier(ctx, req.Identifier)
	if err != nil {
		return nil, fmt.Errorf("invalid NISN/NIS or password")
	}

	// Check password
	err = bcrypt.CompareHashAndPassword([]byte(student.Password), []byte(req.Password))
	if err != nil {
		return nil, fmt.Errorf("invalid NISN/NIS or password")
	}

	// Create session
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

	// Build response with both NISN and NIS
	response := &middleware.SessionData{
		UserID:       student.ID,
		NISN:         student.NISN,
		Name:         student.Fullname,
		Role:         "student",
		SessionID:    sessionData.SessionID,
		AccessToken:  sessionData.AccessToken,
		CreatedAt:    sessionData.CreatedAt,
		LastActivity: sessionData.LastActivity,
		DeviceInfo:   req.DeviceInfo,
		IPAddress:    c.IP(),
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