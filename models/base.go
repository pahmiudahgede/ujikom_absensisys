package models

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BaseModel struct {
	ID        string    `json:"id" gorm:"primaryKey;type:varchar(255);not null"`
	CreatedAt time.Time `json:"created_at" gorm:"not null;autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"not null;autoUpdateTime"`
}

func (base *BaseModel) BeforeCreate(tx *gorm.DB) error {
	if base.ID == "" {
		base.ID = uuid.New().String()
	}
	return nil
}

func ParseTimeString(timeStr string) (time.Time, error) {
	if timeStr == "" {
		return time.Time{}, fmt.Errorf("time string is empty")
	}
	return time.Parse("15:04:05", timeStr)
}

func TimeToString(t time.Time) string {
	return t.Format("15:04:05")
}

func ParseDateString(dateStr string) (time.Time, error) {
	if dateStr == "" {
		return time.Time{}, fmt.Errorf("date string is empty")
	}
	return time.Parse("2006-01-02", dateStr)
}

func DateToString(t time.Time) string {
	return t.Format("2006-01-02")
}

func ValidateTimeString(timeStr string) error {
	_, err := ParseTimeString(timeStr)
	return err
}

func ValidateDateString(dateStr string) error {
	_, err := ParseDateString(dateStr)
	return err
}

func CompareTimeStrings(time1, time2 string) (int, error) {
	t1, err := ParseTimeString(time1)
	if err != nil {
		return 0, fmt.Errorf("invalid time1 format: %v", err)
	}

	t2, err := ParseTimeString(time2)
	if err != nil {
		return 0, fmt.Errorf("invalid time2 format: %v", err)
	}

	if t1.Before(t2) {
		return -1, nil
	} else if t1.After(t2) {
		return 1, nil
	}
	return 0, nil
}

func IsTimeInRange(timeStr, startStr, endStr string) (bool, error) {
	targetTime, err := ParseTimeString(timeStr)
	if err != nil {
		return false, fmt.Errorf("invalid target time format: %v", err)
	}

	startTime, err := ParseTimeString(startStr)
	if err != nil {
		return false, fmt.Errorf("invalid start time format: %v", err)
	}

	endTime, err := ParseTimeString(endStr)
	if err != nil {
		return false, fmt.Errorf("invalid end time format: %v", err)
	}

	return (targetTime.Equal(startTime) || targetTime.After(startTime)) &&
		(targetTime.Equal(endTime) || targetTime.Before(endTime)), nil
}

func GetCurrentTimeString() string {
	return TimeToString(time.Now())
}

func GetCurrentDateString() string {
	return DateToString(time.Now())
}
