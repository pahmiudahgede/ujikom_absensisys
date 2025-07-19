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
		&models.Teacher{},
		&models.Class{},
		&models.Student{},
		&models.Subject{},
		&models.ClassSchedule{},
		&models.Attendance{},
		&models.AttendanceSettings{},
		&models.SafeArea{},
		&models.Holiday{},
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
		&models.School{},
		&models.Teacher{},
		&models.Class{},
		&models.Student{},
		&models.Subject{},
		&models.ClassSchedule{},
		&models.Attendance{},
		&models.AttendanceSettings{},
		&models.SafeArea{},
		&models.Holiday{},
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

func CheckTablesExist() bool {
	requiredTables := []string{
		"students", "teachers", "schools", "subjects", "safe_areas", "holidays", "class_schedules", "classes", "attendance_settings", "attendances",
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
