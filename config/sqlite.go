package config

import (
	"os"

	"github.com/gbmoraes-dev/gopportunities/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// InitSQLite initialize the SQLite database.
func InitSQLite() (*gorm.DB, error) {
	// Initialize logger.
	logger := GetLogger("sqlite")

	// Define database path.
	dbPath := "./database/main.db"

	// Check if database file exists, if not, create it.
	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		logger.Infof("Database file does not exist, creating it...")
		err = os.MkdirAll("./database", os.ModePerm)
		if err != nil {
			return nil, err
		}
		file, err := os.Create(dbPath)
		if err != nil {
			return nil, err
		}
		file.Close()
	}

	// Initialize database and connect.
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		logger.Errorf("Error initializing database: %v", err)
		return nil, err
	}

	// Migrate the schema.
	err = db.AutoMigrate(&schemas.Opportunity{})
	if err != nil {
		logger.Errorf("Error migrating database: %v", err)
		return nil, err
	}

	// Return the database.
	return db, nil
}
