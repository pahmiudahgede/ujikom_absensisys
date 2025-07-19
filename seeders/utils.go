package seeders

import (
	"time"
)

// Helper functions
func stringPtr(s string) *string {
	return &s
}

func parseTime(timeStr string) time.Time {
	t, _ := time.Parse("15:04", timeStr)
	return t
}
