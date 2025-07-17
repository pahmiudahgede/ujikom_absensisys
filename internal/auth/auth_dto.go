package auth

type LoginRequest struct {
	NISN       string `json:"nisn" validate:"required" example:"1234567890"`
	Password   string `json:"password" validate:"required" example:"password123"`
	DeviceInfo string `json:"deviceinfo" example:"Mobile App v1.0"`
}
