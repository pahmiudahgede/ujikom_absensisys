package models

import (
	"time"

	"gorm.io/gorm"
)

type AttendanceSettings struct {
	BaseModel
	SchoolID        string   `json:"school_id" gorm:"type:varchar(255);not null;uniqueIndex"`
	CheckinStart    TimeOnly `json:"checkin_start" gorm:"comment:'Format: HH:mm:ss'"`
	CheckinEnd      TimeOnly `json:"checkin_end" gorm:"comment:'Format: HH:mm:ss'"`
	CheckoutStart   TimeOnly `json:"checkout_start" gorm:"comment:'Format: HH:mm:ss'"`
	CheckoutEnd     TimeOnly `json:"checkout_end" gorm:"comment:'Format: HH:mm:ss'"`
	LateTolerance   int      `json:"late_tolerance" gorm:"not null;default:15;comment:'Toleransi terlambat dalam menit'"`
	RequirePhoto    bool     `json:"require_photo" gorm:"not null;default:true"`
	RequireLocation bool     `json:"require_location" gorm:"not null;default:true"`
	MaxDistance     int      `json:"max_distance" gorm:"not null;default:100;comment:'Jarak maksimal dalam meter'"`

	School *School `json:"school,omitempty" gorm:"foreignKey:SchoolID"`
}

func (AttendanceSettings) TableName() string {
	return "attendance_settings"
}

func (a *AttendanceSettings) BeforeCreate(tx *gorm.DB) error {

	if err := a.BaseModel.BeforeCreate(tx); err != nil {
		return err
	}

	if a.CheckinStart.IsZero() {
		a.CheckinStart = NewTimeOnly(6, 0, 0)
	}
	if a.CheckinEnd.IsZero() {
		a.CheckinEnd = NewTimeOnly(7, 30, 0)
	}
	if a.CheckoutStart.IsZero() {
		a.CheckoutStart = NewTimeOnly(15, 0, 0)
	}
	if a.CheckoutEnd.IsZero() {
		a.CheckoutEnd = NewTimeOnly(17, 0, 0)
	}

	return nil
}

func (a *AttendanceSettings) GetCheckinStartTime() time.Time {
	return a.CheckinStart.GetTime()
}

func (a *AttendanceSettings) GetCheckinEndTime() time.Time {
	return a.CheckinEnd.GetTime()
}

func (a *AttendanceSettings) GetCheckoutStartTime() time.Time {
	return a.CheckoutStart.GetTime()
}

func (a *AttendanceSettings) GetCheckoutEndTime() time.Time {
	return a.CheckoutEnd.GetTime()
}

func (a *AttendanceSettings) SetCheckinStartTime(hour, min, sec int) {
	a.CheckinStart = NewTimeOnly(hour, min, sec)
}

func (a *AttendanceSettings) SetCheckinEndTime(hour, min, sec int) {
	a.CheckinEnd = NewTimeOnly(hour, min, sec)
}

func (a *AttendanceSettings) SetCheckoutStartTime(hour, min, sec int) {
	a.CheckoutStart = NewTimeOnly(hour, min, sec)
}

func (a *AttendanceSettings) SetCheckoutEndTime(hour, min, sec int) {
	a.CheckoutEnd = NewTimeOnly(hour, min, sec)
}

func (a *AttendanceSettings) IsCheckinTime(currentTime time.Time) bool {
	current := NewTimeOnly(currentTime.Hour(), currentTime.Minute(), currentTime.Second())
	return !current.Before(a.CheckinStart) && !current.After(a.CheckinEnd)
}

func (a *AttendanceSettings) IsCheckoutTime(currentTime time.Time) bool {
	current := NewTimeOnly(currentTime.Hour(), currentTime.Minute(), currentTime.Second())
	return !current.Before(a.CheckoutStart) && !current.After(a.CheckoutEnd)
}

func (a *AttendanceSettings) IsLate(checkinTime time.Time) bool {
	current := NewTimeOnly(checkinTime.Hour(), checkinTime.Minute(), checkinTime.Second())
	toleranceEnd := NewTimeOnly(a.CheckinEnd.Hour(), a.CheckinEnd.Minute()+a.LateTolerance, a.CheckinEnd.Second())
	return current.After(toleranceEnd)
}
