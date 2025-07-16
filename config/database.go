package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDatabase() {
	var err error

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local&sql_mode=''",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	config := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NowFunc: func() time.Time {
			return time.Now().Local()
		},
		DisableForeignKeyConstraintWhenMigrating: true,
	}

	DB, err = gorm.Open(mysql.Open(dsn), config)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatal("Failed to configure database connection pool:", err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	DB.Exec("SET sql_mode = ''")
	DB.Exec("SET time_zone = '+00:00'")

	log.Println("✅ Database connected successfully!")
}
