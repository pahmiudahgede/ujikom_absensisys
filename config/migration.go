package config

import (
	"absensibe/models"
	"absensibe/seeders"
	"fmt"
	"log"
	"strings"
)

func RunMigrations() {
	log.Println("🔄 Starting PostgreSQL database migrations...")

	if DB == nil {
		log.Fatal("❌ Database connection is not initialized")
	}

	if err := DB.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"").Error; err != nil {
		log.Printf("Warning: Could not create uuid-ossp extension: %v", err)
	}

	DB.Exec("SET TIME ZONE 'UTC'")

	modelsToMigrate := []interface{}{
		&models.School{},
		&models.SchoolAddress{},
		&models.Major{},
		&models.Subject{},
		&models.Teacher{},
		&models.TeacherAddress{},
		&models.Class{},
		&models.Student{},
		&models.StudentAddress{},
		&models.Parent{},
		&models.SafeArea{},
		&models.ClassSchedule{},
		&models.ClassSession{},
		&models.Attendance{},
		&models.AttendanceDetails{},
		&models.SubjectAttendance{},
		&models.AttendanceRecap{},
		&models.SubjectAttendanceRecap{},
		&models.AcademicCalendar{},
		&models.AttendanceSettings{},
		&models.SubjectAttendanceSettings{},
		&models.Permit{},
		&models.SubjectPermit{},
		&models.Notification{},
	}

	for i, model := range modelsToMigrate {
		modelName := fmt.Sprintf("%T", model)

		modelName = strings.TrimPrefix(modelName, "*models.")

		log.Printf("🔄 [%d/%d] Migrating %s...", i+1, len(modelsToMigrate), modelName)

		if err := DB.AutoMigrate(model); err != nil {
			log.Printf("❌ Failed to migrate %s: %v", modelName, err)
			continue
		}

		log.Printf("✅ [%d/%d] Migrated %s successfully", i+1, len(modelsToMigrate), modelName)
	}

	CreateIndexes()

	log.Println("🎉 PostgreSQL database migrations completed successfully!")
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
	log.Println("⚠️  Dropping all PostgreSQL tables...")

	DB.Exec("SET session_replication_role = replica")

	modelsToDrop := []interface{}{
		&models.Notification{},
		&models.SubjectPermit{},
		&models.Permit{},
		&models.SubjectAttendanceSettings{},
		&models.AttendanceSettings{},
		&models.AcademicCalendar{},
		&models.SubjectAttendanceRecap{},
		&models.AttendanceRecap{},
		&models.SubjectAttendance{},
		&models.AttendanceDetails{},
		&models.Attendance{},
		&models.ClassSession{},
		&models.ClassSchedule{},
		&models.SafeArea{},
		&models.Parent{},
		&models.StudentAddress{},
		&models.Student{},
		&models.Class{},
		&models.TeacherAddress{},
		&models.Teacher{},
		&models.Subject{},
		&models.Major{},
		&models.SchoolAddress{},
		&models.School{},
	}

	for _, model := range modelsToDrop {
		modelName := fmt.Sprintf("%T", model)
		modelName = strings.TrimPrefix(modelName, "*models.")

		if err := DB.Migrator().DropTable(model); err != nil {
			log.Printf("⚠️  Failed to drop %s: %v", modelName, err)
		} else {
			log.Printf("🗑️  Dropped %s successfully", modelName)
		}
	}

	DB.Exec("SET session_replication_role = DEFAULT")

	log.Println("🎉 All PostgreSQL tables dropped successfully!")
}

func CreateIndexes() {
	log.Println("🔧 Creating custom PostgreSQL indexes...")

	indexes := []struct {
		Name  string
		Query string
	}{
		{
			Name:  "idx_student_addresses_student",
			Query: "CREATE INDEX IF NOT EXISTS idx_student_addresses_student ON student_addresses(student_id)",
		},
		{
			Name:  "idx_teacher_addresses_teacher",
			Query: "CREATE INDEX IF NOT EXISTS idx_teacher_addresses_teacher ON teacher_addresses(teacher_id)",
		},
		{
			Name:  "idx_school_addresses_school",
			Query: "CREATE INDEX IF NOT EXISTS idx_school_addresses_school ON school_addresses(school_id)",
		},
		{
			Name:  "idx_class_schedules_academic",
			Query: "CREATE INDEX IF NOT EXISTS idx_class_schedules_academic ON class_schedules(academic_year, semester, is_active)",
		},
		{
			Name:  "idx_notifications_recipient_full",
			Query: "CREATE INDEX IF NOT EXISTS idx_notifications_recipient_full ON notifications(recipient_id, recipient_type, is_read)",
		},
		{
			Name:  "idx_attendances_date_status",
			Query: "CREATE INDEX IF NOT EXISTS idx_attendances_date_status ON attendances(date, status)",
		},
		{
			Name:  "idx_subject_attendances_session",
			Query: "CREATE INDEX IF NOT EXISTS idx_subject_attendances_session ON subject_attendances(session_id, status)",
		},
		{
			Name:  "idx_students_class_status",
			Query: "CREATE INDEX IF NOT EXISTS idx_students_class_status ON students(class_id, status)",
		},
		{
			Name:  "idx_teachers_school_status",
			Query: "CREATE INDEX IF NOT EXISTS idx_teachers_school_status ON teachers(school_id, status)",
		},
		{
			Name:  "idx_class_sessions_date_status",
			Query: "CREATE INDEX IF NOT EXISTS idx_class_sessions_date_status ON class_sessions(date, status)",
		},
		{
			Name:  "idx_permits_dates",
			Query: "CREATE INDEX IF NOT EXISTS idx_permits_dates ON permits(start_date, end_date, approval_status)",
		},
		{
			Name:  "idx_subject_permits_dates",
			Query: "CREATE INDEX IF NOT EXISTS idx_subject_permits_dates ON subject_permits(start_date, end_date, approval_status)",
		},
		{
			Name:  "idx_class_schedules_time",
			Query: "CREATE INDEX IF NOT EXISTS idx_class_schedules_time ON class_schedules(day_of_week, start_time, end_time)",
		},
		{
			Name:  "idx_students_entry_year",
			Query: "CREATE INDEX IF NOT EXISTS idx_students_entry_year ON students(entry_year)",
		},
		{
			Name:  "idx_attendances_student_month",
			Query: "CREATE INDEX IF NOT EXISTS idx_attendances_student_month ON attendances(student_id, EXTRACT(MONTH FROM date), EXTRACT(YEAR FROM date))",
		},
	}

	successCount := 0
	for _, idx := range indexes {

		var exists bool
		checkQuery := `
			SELECT EXISTS (
				SELECT 1 
				FROM pg_indexes 
				WHERE indexname = $1
			)
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
		"schools", "school_addresses", "majors", "subjects", "teachers", "teacher_addresses",
		"classes", "students", "student_addresses", "parents", "safe_areas",
		"class_schedules", "class_sessions", "attendances", "attendance_details",
		"subject_attendances", "attendance_recaps", "subject_attendance_recaps",
		"academic_calendars", "attendance_settings", "subject_attendance_settings",
		"permits", "subject_permits", "notifications",
	}

	for _, table := range requiredTables {
		if !DB.Migrator().HasTable(table) {
			log.Printf("❌ Table '%s' does not exist", table)
			return false
		}
	}

	log.Println("✅ All required PostgreSQL tables exist")
	return true
}

func ResetDatabase() {
	log.Println("🔄 Resetting PostgreSQL database...")

	DropAllTables()
	RunMigrationsWithSeed()

	log.Println("🎉 PostgreSQL database reset completed successfully!")
}

func CreateIndexIfNotExists(indexName, query string) error {
	var exists bool
	checkQuery := `
		SELECT EXISTS (
			SELECT 1 
			FROM pg_indexes 
			WHERE indexname = $1
		)
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
