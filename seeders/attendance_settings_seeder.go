package seeders

import (
	"absensibe/models"
	"log"
	"time"

	"gorm.io/gorm"
)

func SeedAttendanceSettings(db *gorm.DB) error {
	log.Println("⚙️ Seeding attendance settings...")

	// Create time values properly for PostgreSQL TIME type
	checkInStart := parseTimeOnly("06:00")
	checkInEnd := parseTimeOnly("07:30")
	checkOutStart := parseTimeOnly("15:00")
	checkOutEnd := parseTimeOnly("17:00")

	settings := models.AttendanceSettings{
		ID:              "550e8400-e29b-41d4-a716-446655440060",
		SchoolID:        "550e8400-e29b-41d4-a716-446655440001",
		CheckInStart:    checkInStart,
		CheckInEnd:      checkInEnd,
		CheckOutStart:   checkOutStart,
		CheckOutEnd:     checkOutEnd,
		LateTolerance:   15,
		RequirePhoto:    true,
		RequireLocation: true,
		MaxDistance:     100,
	}

	if err := db.FirstOrCreate(&settings, models.AttendanceSettings{SchoolID: settings.SchoolID}).Error; err != nil {
		return err
	}

	log.Println("✅ Successfully seeded attendance settings")
	return nil
}

// Helper function untuk create time yang compatible dengan PostgreSQL TIME type
func parseTimeOnly(timeStr string) time.Time {
	// Parse time string dan set ke tanggal epoch (1970-01-01) untuk TIME type
	t, err := time.Parse("15:04", timeStr)
	if err != nil {
		// Fallback jika parsing gagal
		return time.Date(1970, 1, 1, 6, 0, 0, 0, time.UTC)
	}

	// Set ke tanggal epoch dengan jam yang diparsing
	return time.Date(1970, 1, 1, t.Hour(), t.Minute(), t.Second(), 0, time.UTC)
}
