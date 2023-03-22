package database

import (
	"fmt"
	"os"

	"github.com/rohan184/server/pkg/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func DBConnection() {
	// dsn := "postgres://user:user@localhost:5432/crawler"
	var err error
	db, err = gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	db.AutoMigrate(model.Crawler{})
}

func Insert(url string, wordCount int) error {
	err := db.Create(&model.Crawler{
		Url:       url,
		WordCount: fmt.Sprint(wordCount),
		Fav:       false,
	}).Error
	if err != nil {
		return err
	}
	return nil
}

func Query() (res []model.Crawler, err error) {
	err = db.Find(&res).Error
	if err != nil {
		return res, err
	}

	return res, nil
}
