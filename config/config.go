package config

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

// Init initialize the configs.
func Init() error {
	var err error

	db, err = InitSQLite()

	if err != nil {
		return fmt.Errorf("error initializing database: %v", err)
	}

	return nil
}

// GetDB returns the database.
func GetDB() *gorm.DB {
	return db
}

// GetLogger returns the logger.
func GetLogger(p string) *Logger {
	logger = NewLogger(p)
	return logger
}
