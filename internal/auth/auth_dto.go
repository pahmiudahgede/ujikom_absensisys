package auth

import "time"

// LoginRequest represents student login request payload
type LoginRequest struct {
	Identifier string `json:"identifier" validate:"required" example:"2024000579"` // NISN or NIS
	Password   string `json:"password" validate:"required" example:"password123"`
	DeviceInfo string `json:"deviceinfo" example:"postmanstudent123x0=="`
} //	@name	LoginRequest

// LoginResponseData represents the actual login response data structure
type LoginResponseData struct {
	UserID       string    `json:"user_id" example:"003061ee-97ff-4d00-8155-f4bf15e319dd"`
	NISN         string    `json:"nisn" example:"2024000579"`
	NIS          string    `json:"nis" example:"20240001"`
	Name         string    `json:"name" example:"Eko Saputra"`
	SessionID    string    `json:"session_id" example:"b5e882d0eb21c5ba7752d2b8f216ccad"`
	AccessToken  string    `json:"access_token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
	CreatedAt    time.Time `json:"created_at" example:"0001-01-01T00:00:00Z"`
	LastActivity time.Time `json:"last_activity" example:"0001-01-01T00:00:00Z"`
	DeviceInfo   string    `json:"device_info" example:"postmanstudent123x0=="`
} //	@name	LoginResponseData

// LoginResponse represents complete login response structure
type LoginResponse struct {
	Meta struct {
		Status  int    `json:"status" example:"200"`
		Message string `json:"message" example:"login success"`
	} `json:"meta"`
	Data LoginResponseData `json:"data"`
} //	@name	LoginResponse

// LogoutResponse represents complete logout response structure
type LogoutResponse struct {
	Meta struct {
		Status  int    `json:"status" example:"200"`
		Message string `json:"message" example:"logout success"`
	} `json:"meta"`
} //	@name	LogoutResponse