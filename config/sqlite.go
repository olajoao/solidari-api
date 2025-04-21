package config

import (
	"os"

	"github.com/olajoao/solidari-api/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")
	db_path := "./db/donation.db"
	_, err := os.Stat(db_path)
	if os.IsNotExist(err) {
		logger.Info("database file not found, creating...")
		err = os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			return nil, err
		}
		file, err := os.Create(db_path)
		if err != nil {
			return nil, err
		}
		file.Close()
	}
	db, err := gorm.Open(sqlite.Open(db_path), &gorm.Config{})
	if err != nil {
		logger.Errorf("InitializeSQLite %v", err)
		return nil, err
	}

	err = db.AutoMigrate(&schemas.Donation{})
	if err != nil {
		logger.Errorf("AutoMigrate error %v", err)
		return nil, err
	}

	return db, nil
}
