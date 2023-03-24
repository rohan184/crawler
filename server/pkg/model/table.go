package model

import (
	"gorm.io/gorm"
)

type Crawler struct {
	gorm.Model
	Url        string
	WordCount  int
	WebLinks   string
	MediaLinks string
	Fav        bool
}
