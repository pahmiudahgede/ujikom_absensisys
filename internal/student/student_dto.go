package student

import (
	"absensibe/models"
	"time"
)

// StudentInfoRequest represents query parameters for student info
type StudentInfoRequest struct {
	IncludeClass      bool   `query:"class"`
	IncludeAttendance bool   `query:"absen"`
	IncludeSchedule   bool   `query:"schedule"`
	AttendanceMonth   string `query:"month"` // format: 2024-01
	AttendanceLimit   int    `query:"limit"` // default: 10
} //	@name	StudentInfoRequest

// StudentInfoResponse represents student information response
type StudentInfoResponse struct {
	// Basic student info
	ID           string    `json:"id"`
	NIS          string    `json:"nis"`
	NISN         string    `json:"nisn"`
	Fullname     string    `json:"fullname"`
	PhotoProfile *string   `json:"photo_profile"`
	Gender       string    `json:"gender"`
	BirthDate    time.Time `json:"birth_date"`
	PlaceOfBirth string    `json:"place_of_birth"`
	Phone        string    `json:"phone"`
	Address      string    `json:"address"`
	Religion     string    `json:"religion"`
	FatherName   string    `json:"father_name"`
	MotherName   string    `json:"mother_name"`
	Status       string    `json:"status"`
	EntryYear    int       `json:"entry_year"`

	// Optional relations
	Class       *models.Class       `json:"class,omitempty"`
	Attendances []models.Attendance `json:"attendances,omitempty"`

	// Meta info
	AttendanceSummary *AttendanceSummary `json:"attendance_summary,omitempty"`
} //	@name	StudentInfoResponse

// AttendanceSummary represents attendance statistics
type AttendanceSummary struct {
	TotalDays      int64   `json:"total_days"`
	PresentDays    int64   `json:"present_days"`
	LateDays       int64   `json:"late_days"`
	AbsentDays     int64   `json:"absent_days"`
	AttendanceRate float64 `json:"attendance_rate"`
	CurrentMonth   string  `json:"current_month"`
} //	@name	AttendanceSummary

// Convert models.Student to StudentInfoResponse
func ToStudentResponse(student *models.Student) *StudentInfoResponse {
	return &StudentInfoResponse{
		ID:           student.ID,
		NIS:          student.NIS,
		NISN:         student.NISN,
		Fullname:     student.Fullname,
		PhotoProfile: student.PhotoProfile,
		Gender:       student.Gender,
		BirthDate:    student.BirthDate,
		PlaceOfBirth: student.PlaceOfBirth,
		Phone:        student.Phone,
		Address:      student.Address,
		Religion:     student.Religion,
		FatherName:   student.FatherName,
		MotherName:   student.MotherName,
		Status:       student.Status,
		EntryYear:    student.EntryYear,
		Class:        student.Class,
		Attendances:  student.Attendances,
	}
}
