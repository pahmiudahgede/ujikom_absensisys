package config

import (
	"absensibe/models"
	"absensibe/seeders"
	"fmt"
	"log"
)

func RunMigrations() {
	log.Println("🔄 Starting database migrations...")

	if DB == nil {
		log.Fatal("❌ Database connection is not initialized")
	}

	DB.Exec("SET sql_mode = 'TRADITIONAL'")
	DB.Exec("SET time_zone = '+07:00'")

	modelsToMigrate := []interface{}{

		&models.School{},
		&models.Jurusan{},
		&models.Subject{},
		&models.Teacher{},
		&models.Class{},
		&models.Student{},
		&models.Parent{},
		&models.Address{},

		&models.ClassSchedule{},
		&models.ClassSession{},

		&models.Absensi{},
		&models.AbsensiDetails{},
		&models.SubjectAttendance{},
		&models.AbsensiRecap{},
		&models.SubjectAttendanceRecap{},

		&models.SafeArea{},
		&models.AcademicCalendar{},
		&models.Holiday{},
		&models.AttendanceSettings{},
		&models.SubjectAttendanceSettings{},

		&models.Permit{},
		&models.SubjectPermit{},
		&models.Notification{},
	}

	for i, model := range modelsToMigrate {
		modelName := fmt.Sprintf("%T", model)

		log.Printf("🔄 [%d/%d] Migrating %s...", i+1, len(modelsToMigrate), modelName)

		if err := DB.AutoMigrate(model); err != nil {
			log.Printf("❌ Failed to migrate %s: %v", modelName, err)
			continue
		}

		log.Printf("✅ [%d/%d] Migrated %s successfully", i+1, len(modelsToMigrate), modelName)
	}

	CreateIndexes()

	log.Println("🎉 Database migrations completed successfully!")
}

func RunMigrationsWithSeed() {
	RunMigrations()
	RunSeeders()
}

func RunSeeders() {
	log.Println("🌱 Starting database seeding...")
	seeders.RunAllSeeders(DB)
}

func DropAllTables() {
	log.Println("⚠️  Dropping all tables...")

	DB.Exec("SET foreign_key_checks = 0")

	modelsToDrop := []interface{}{
		&models.Notification{},
		&models.SubjectPermit{},
		&models.Permit{},
		&models.SubjectAttendanceSettings{},
		&models.AttendanceSettings{},
		&models.Holiday{},
		&models.AcademicCalendar{},
		&models.SafeArea{},
		&models.SubjectAttendanceRecap{},
		&models.AbsensiRecap{},
		&models.SubjectAttendance{},
		&models.AbsensiDetails{},
		&models.Absensi{},
		&models.ClassSession{},
		&models.ClassSchedule{},
		&models.Address{},
		&models.Parent{},
		&models.Student{},
		&models.Class{},
		&models.Teacher{},
		&models.Subject{},
		&models.Jurusan{},
		&models.School{},
	}

	for _, model := range modelsToDrop {
		modelName := fmt.Sprintf("%T", model)
		if err := DB.Migrator().DropTable(model); err != nil {
			log.Printf("⚠️  Failed to drop %s: %v", modelName, err)
		} else {
			log.Printf("🗑️  Dropped %s successfully", modelName)
		}
	}

	DB.Exec("SET foreign_key_checks = 1")

	log.Println("🎉 All tables dropped successfully!")
}

func CreateIndexes() {
	log.Println("🔧 Creating custom indexes...")

	indexes := []struct {
		Name  string
		Query string
	}{
		{
			Name:  "idx_addresses_reference",
			Query: "CREATE INDEX idx_addresses_reference ON addresses(reference_id, reference_type)",
		},
		{
			Name:  "idx_class_schedules_academic",
			Query: "CREATE INDEX idx_class_schedules_academic ON class_schedules(academic_year, semester, is_active)",
		},
		{
			Name:  "idx_notifications_recipient_full",
			Query: "CREATE INDEX idx_notifications_recipient_full ON notifications(recipient_id, recipient_type, is_read)",
		},
		{
			Name:  "idx_absensi_date_status",
			Query: "CREATE INDEX idx_absensi_date_status ON absensi(date, status)",
		},
		{
			Name:  "idx_subject_attendance_session",
			Query: "CREATE INDEX idx_subject_attendance_session ON subject_attendance(session_id, status)",
		},
		{
			Name:  "idx_students_class_status",
			Query: "CREATE INDEX idx_students_class_status ON students(classes_id, status)",
		},
		{
			Name:  "idx_teachers_school_status",
			Query: "CREATE INDEX idx_teachers_school_status ON teachers(school_id, status)",
		},
		{
			Name:  "idx_class_sessions_date_status",
			Query: "CREATE INDEX idx_class_sessions_date_status ON class_sessions(date, status)",
		},
		{
			Name:  "idx_permits_dates",
			Query: "CREATE INDEX idx_permits_dates ON permits(start_date, end_date, approval_status)",
		},
		{
			Name:  "idx_subject_permits_dates",
			Query: "CREATE INDEX idx_subject_permits_dates ON subject_permits(start_date, end_date, approval_status)",
		},
	}

	successCount := 0
	for _, idx := range indexes {

		var exists bool
		checkQuery := `
			SELECT COUNT(*) > 0 
			FROM information_schema.statistics 
			WHERE table_schema = DATABASE() 
			AND index_name = ?
		`

		if err := DB.Raw(checkQuery, idx.Name).Scan(&exists).Error; err != nil {
			log.Printf("⚠️  Failed to check index %s: %v", idx.Name, err)
			continue
		}

		if exists {
			log.Printf("ℹ️  Index %s already exists, skipping...", idx.Name)
			successCount++
			continue
		}

		if err := DB.Exec(idx.Query).Error; err != nil {
			log.Printf("⚠️  Failed to create index %s: %v", idx.Name, err)
		} else {
			log.Printf("✅ Created index %s successfully", idx.Name)
			successCount++
		}
	}

	log.Printf("✅ Custom indexes processed successfully! (%d/%d)", successCount, len(indexes))
}

func CheckTablesExist() bool {
	requiredTables := []string{
		"schools", "jurusan", "subjects", "teachers", "classes", "students",
		"parents", "addresses", "class_schedules", "class_sessions",
		"absensi", "absensi_details", "subject_attendance",
		"absensi_recap", "subject_attendance_recap", "safe_area",
		"academic_calendar", "holidays", "attendance_settings",
		"subject_attendance_settings", "permits", "subject_permits",
		"notifications",
	}

	for _, table := range requiredTables {
		if !DB.Migrator().HasTable(table) {
			log.Printf("❌ Table '%s' does not exist", table)
			return false
		}
	}

	log.Println("✅ All required tables exist")
	return true
}

func ResetDatabase() {
	log.Println("🔄 Resetting database...")

	DropAllTables()
	RunMigrationsWithSeed()

	log.Println("🎉 Database reset completed successfully!")
}

func CreateIndexIfNotExists(indexName, query string) error {
	var exists bool
	checkQuery := `
		SELECT COUNT(*) > 0 
		FROM information_schema.statistics 
		WHERE table_schema = DATABASE() 
		AND index_name = ?
	`

	if err := DB.Raw(checkQuery, indexName).Scan(&exists).Error; err != nil {
		return fmt.Errorf("failed to check index existence: %v", err)
	}

	if exists {
		log.Printf("ℹ️  Index %s already exists", indexName)
		return nil
	}

	if err := DB.Exec(query).Error; err != nil {
		return fmt.Errorf("failed to create index: %v", err)
	}

	log.Printf("✅ Created index %s", indexName)
	return nil
}
