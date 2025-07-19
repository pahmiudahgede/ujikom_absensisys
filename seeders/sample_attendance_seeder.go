package seeders

import (
	"absensibe/models"
	"fmt"
	"log"
	"math/rand"
	"time"

	"gorm.io/gorm"
)

func SeedSampleAttendances(db *gorm.DB) error {
	log.Println("📋 Seeding sample attendances...")

	var students []models.Student
	if err := db.Find(&students).Error; err != nil {
		return err
	}

	if len(students) == 0 {
		log.Println("No students found, skipping attendance seeding")
		return nil
	}

	now := time.Now()
	var attendances []models.Attendance
	attendanceCounter := 1000

	for i := 0; i < 30; i++ {
		date := now.AddDate(0, 0, -i)

		if date.Weekday() == time.Saturday || date.Weekday() == time.Sunday {
			continue
		}

		for _, student := range students {
			attendanceID := fmt.Sprintf("550e8400-e29b-41d4-a716-44665544%04d", attendanceCounter)
			attendanceCounter++

			checkInStatus, checkOutStatus := generateAttendanceStatus()

			var checkInTime, checkOutTime *time.Time
			var checkInLat, checkInLng, checkOutLat, checkOutLng *float64
			var checkInPhoto, checkOutPhoto, checkInNote, checkOutNote *string

			if checkInStatus == "hadir" || checkInStatus == "terlambat" {
				checkIn := generateCheckInTime(date, checkInStatus == "terlambat")
				checkInTime = &checkIn

				lat := -6.9174 + (rand.Float64()-0.5)*0.001
				lng := 107.6191 + (rand.Float64()-0.5)*0.001
				checkInLat = &lat
				checkInLng = &lng

				photoPath := fmt.Sprintf("uploads/attendance/%s_checkin.jpg", attendanceID)
				checkInPhoto = &photoPath

				if checkInStatus == "terlambat" {
					note := "Terlambat karena macet di jalan"
					checkInNote = &note
				} else {
					note := "Hadir tepat waktu"
					checkInNote = &note
				}
			} else if checkInStatus == "sakit" {
				note := "Sakit demam dan flu"
				checkInNote = &note
			} else if checkInStatus == "izin" {
				note := "Izin ada keperluan keluarga"
				checkInNote = &note
			}

			if checkOutStatus == "hadir" {
				checkOut := generateCheckOutTime(date)
				checkOutTime = &checkOut

				lat := -6.9174 + (rand.Float64()-0.5)*0.001
				lng := 107.6191 + (rand.Float64()-0.5)*0.001
				checkOutLat = &lat
				checkOutLng = &lng

				photoPath := fmt.Sprintf("uploads/attendance/%s_checkout.jpg", attendanceID)
				checkOutPhoto = &photoPath

				note := "Pulang sesuai jadwal"
				checkOutNote = &note
			}

			attendance := models.Attendance{
				ID:                attendanceID,
				StudentID:         student.ID,
				Date:              date,
				CheckInTime:       checkInTime,
				CheckOutTime:      checkOutTime,
				CheckInStatus:     checkInStatus,
				CheckOutStatus:    checkOutStatus,
				CheckInLatitude:   checkInLat,
				CheckInLongitude:  checkInLng,
				CheckOutLatitude:  checkOutLat,
				CheckOutLongitude: checkOutLng,
				CheckInPhoto:      checkInPhoto,
				CheckOutPhoto:     checkOutPhoto,
				CheckInNote:       checkInNote,
				CheckOutNote:      checkOutNote,
			}

			attendances = append(attendances, attendance)
		}
	}

	successCount := 0
	for _, attendance := range attendances {
		if err := db.FirstOrCreate(&attendance, models.Attendance{
			StudentID: attendance.StudentID,
			Date:      attendance.Date,
		}).Error; err != nil {
			log.Printf("Error creating attendance for student %s on %s: %v",
				attendance.StudentID, attendance.Date.Format("2006-01-02"), err)
			continue
		}
		successCount++
	}

	log.Printf("✅ Successfully seeded %d sample attendances", successCount)
	return nil
}

func generateAttendanceStatus() (checkInStatus, checkOutStatus string) {

	// rand.Seed(time.Now().UnixNano() + int64(rand.Intn(1000)))
	rand.New(rand.NewSource(time.Now().UnixNano() + int64(rand.Intn(1000))))

	r := rand.Intn(100) + 1

	switch {
	case r <= 85:
		return "hadir", "hadir"
	case r <= 93:
		return "terlambat", "hadir"
	case r <= 96:
		return "sakit", "alpha"
	case r <= 98:
		return "izin", "alpha"
	default:
		return "alpha", "alpha"
	}
}

func generateCheckInTime(date time.Time, isLate bool) time.Time {

	rand.Seed(time.Now().UnixNano() + date.Unix())

	if isLate {

		baseMinutes := 31 + rand.Intn(60)
		hour := 7
		minutes := baseMinutes
		if minutes >= 60 {
			hour = 8
			minutes = minutes - 60
		}
		return time.Date(date.Year(), date.Month(), date.Day(), hour, minutes, rand.Intn(60), 0, date.Location())
	} else {

		baseMinutes := 30 + rand.Intn(60)
		hour := 6
		minutes := baseMinutes
		if minutes >= 60 {
			hour = 7
			minutes = minutes - 60
		}
		return time.Date(date.Year(), date.Month(), date.Day(), hour, minutes, rand.Intn(60), 0, date.Location())
	}
}

func generateCheckOutTime(date time.Time) time.Time {

	rand.Seed(time.Now().UnixNano() + date.Unix() + 3600)

	hour := 15 + rand.Intn(2)
	minute := rand.Intn(60)

	if hour == 16 && minute > 30 {
		minute = 30
	}

	return time.Date(date.Year(), date.Month(), date.Day(), hour, minute, rand.Intn(60), 0, date.Location())
}

func SeedTodayAttendances(db *gorm.DB) error {
	log.Println("📋 Seeding today's attendances for testing...")

	var students []models.Student
	if err := db.Find(&students).Error; err != nil {
		return err
	}

	if len(students) == 0 {
		log.Println("No students found, skipping today's attendance seeding")
		return nil
	}

	today := time.Now()

	if today.Weekday() == time.Saturday || today.Weekday() == time.Sunday {
		log.Println("Today is weekend, skipping attendance seeding")
		return nil
	}

	attendanceCounter := 2000
	successCount := 0

	for _, student := range students {
		attendanceID := fmt.Sprintf("550e8400-e29b-41d4-a716-44665544%04d", attendanceCounter)
		attendanceCounter++

		attendance := models.Attendance{
			ID:             attendanceID,
			StudentID:      student.ID,
			Date:           today,
			CheckInStatus:  "alpha",
			CheckOutStatus: "alpha",
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
		}

		var existingAttendance models.Attendance
		err := db.Where("student_id = ? AND date = ?", student.ID, today.Format("2006-01-02")).First(&existingAttendance).Error

		if err == gorm.ErrRecordNotFound {
			if err := db.Create(&attendance).Error; err != nil {
				log.Printf("Error creating today's attendance for student %s: %v", student.Fullname, err)
				continue
			}
			successCount++
		}
	}

	log.Printf("✅ Successfully seeded %d today's attendance records", successCount)
	return nil
}

func SeedDemoAttendances(db *gorm.DB) error {
	log.Println("🎭 Seeding demo attendances for presentation...")

	var students []models.Student
	if err := db.Limit(3).Find(&students).Error; err != nil {
		return err
	}

	if len(students) == 0 {
		log.Println("No students found, skipping demo attendance seeding")
		return nil
	}

	today := time.Now()
	yesterday := today.AddDate(0, 0, -1)

	attendanceCounter := 3000
	var demoAttendances []models.Attendance

	for i, student := range students {
		attendanceID := fmt.Sprintf("550e8400-e29b-41d4-a716-44665544%04d", attendanceCounter)
		attendanceCounter++

		var checkInStatus, checkOutStatus string
		var checkInTime, checkOutTime *time.Time
		var checkInNote, checkOutNote *string

		switch i {
		case 0:
			checkInStatus = "hadir"
			checkOutStatus = "hadir"
			checkIn := time.Date(yesterday.Year(), yesterday.Month(), yesterday.Day(), 7, 15, 0, 0, yesterday.Location())
			checkOut := time.Date(yesterday.Year(), yesterday.Month(), yesterday.Day(), 15, 30, 0, 0, yesterday.Location())
			checkInTime = &checkIn
			checkOutTime = &checkOut
			note1 := "Hadir tepat waktu"
			note2 := "Pulang sesuai jadwal"
			checkInNote = &note1
			checkOutNote = &note2

		case 1:
			checkInStatus = "terlambat"
			checkOutStatus = "hadir"
			checkIn := time.Date(yesterday.Year(), yesterday.Month(), yesterday.Day(), 8, 5, 0, 0, yesterday.Location())
			checkOut := time.Date(yesterday.Year(), yesterday.Month(), yesterday.Day(), 15, 45, 0, 0, yesterday.Location())
			checkInTime = &checkIn
			checkOutTime = &checkOut
			note1 := "Terlambat karena macet"
			note2 := "Pulang sesuai jadwal"
			checkInNote = &note1
			checkOutNote = &note2

		case 2:
			checkInStatus = "sakit"
			checkOutStatus = "sakit"
			note1 := "Sakit demam tinggi"
			checkInNote = &note1
		}

		lat := -6.9174
		lng := 107.6191
		photoPath1 := fmt.Sprintf("uploads/demo/%s_checkin.jpg", attendanceID)
		photoPath2 := fmt.Sprintf("uploads/demo/%s_checkout.jpg", attendanceID)

		attendance := models.Attendance{
			ID:                attendanceID,
			StudentID:         student.ID,
			Date:              yesterday,
			CheckInTime:       checkInTime,
			CheckOutTime:      checkOutTime,
			CheckInStatus:     checkInStatus,
			CheckOutStatus:    checkOutStatus,
			CheckInLatitude:   &lat,
			CheckInLongitude:  &lng,
			CheckOutLatitude:  &lat,
			CheckOutLongitude: &lng,
			CheckInPhoto:      &photoPath1,
			CheckOutPhoto:     &photoPath2,
			CheckInNote:       checkInNote,
			CheckOutNote:      checkOutNote,
		}

		demoAttendances = append(demoAttendances, attendance)
	}

	successCount := 0
	for _, attendance := range demoAttendances {
		if err := db.FirstOrCreate(&attendance, models.Attendance{
			StudentID: attendance.StudentID,
			Date:      attendance.Date,
		}).Error; err != nil {
			log.Printf("Error creating demo attendance: %v", err)
			continue
		}
		successCount++
	}

	log.Printf("✅ Successfully seeded %d demo attendances", successCount)
	return nil
}
