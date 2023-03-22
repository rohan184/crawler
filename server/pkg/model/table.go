package model

import "gorm.io/gorm"

type Crawler struct {
	gorm.Model
	Url       string
	WordCount string
	Fav       bool
}
