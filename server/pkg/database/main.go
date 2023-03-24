package database

import (
	"fmt"
	"os"

	"github.com/rohan184/server/pkg/model"
	"github.com/rohan184/server/pkg/resources"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func DBConnection() {
	var err error
	db, err = gorm.Open(sqlite.Open("crawler.db"), &gorm.Config{})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	db.AutoMigrate(model.Crawler{})
}

func Insert(resp *resources.Insight) error {
	if err := db.Table("crawlers").Create(&model.Crawler{
		Url:        resp.URL,
		WordCount:  resp.WordCount,
		WebLinks:   fmt.Sprintf("%v", resp.WebLinks),
		MediaLinks: fmt.Sprintf("%v", resp.MediaLinks),
		Fav:        false,
	}).Error; err != nil {
		return err
	}
	return nil
}

func Query() (res []model.Crawler, err error) {
	if err = db.Find(&res).Error; err != nil {
		return res, err
	}

	return res, nil
}

func DeleteInsight(id int) error {
	return db.Delete(&model.Crawler{}, id).Error
}

func MarkInsightFav(id int) error {
	return db.Table("crawlers").Where("id=?", id).Update("Fav", true).Error
}
