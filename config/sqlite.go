package config

import (
	"os"

	"github.com/gilbertouk/gopportunities-api/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")

	dbPath := "./db/main.db"

	// Check if the database file exists
	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		logger.Info("Database file does not exist, creating a new one.")

		// Create the database file and directory
		err = os.MkdirAll("./db", os.ModePerm) // Create the directory
		if err != nil {
			logger.Errorf("Error creating database directory: %v", err)
			return nil, err
		}

		file, err := os.Create(dbPath) // Create the database file
		if err != nil {
			logger.Errorf("Error creating database file: %v", err)
			return nil, err
		}
		file.Close()
	} else {
		logger.Info("Database file exists, opening it.")
	}

	// create a new SQLite database connection
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		logger.Errorf("Error initializing SQLite database: %v", err)
		return nil, err
	}

	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil {
		logger.Errorf("Error migrating database schema: %v", err)
		return nil, err
	}

	return db, nil
}
