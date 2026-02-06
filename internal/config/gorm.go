package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/rafmasloman/mis-modernisasi-backend/internal/infrastructure/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func GormApp() (*gorm.DB, error) {
	gormEnv := EnvGorm()

	// logger
	dbLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // slow SQL Threshold
			LogLevel:      logger.Info, // log level
			Colorful:      true,        // disable color
		},
	)

	// open db
	dsn := fmt.Sprintf("host=%s user=%s dbname=%s port=%s password=%s sslmode=disable TimeZone=Asia/Jakarta", gormEnv.DBHost, gormEnv.DBUser, gormEnv.DBName, gormEnv.DBPort, gormEnv.DBPass)

	var err error

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: dbLogger})

	// Migrate DB

	// MigrateGorm(db)

	if err != nil {
		panic("Cannot connect database")
	}

	return db, nil
}

func MigrateGorm(db *gorm.DB) {
	db.AutoMigrate(&model.ReportBuilderModel{},
		&model.FilterBuilderModel{})
}
