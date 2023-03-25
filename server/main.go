package main

import (
	"github.com/rohan184/server/pkg/database"
	"github.com/rohan184/server/pkg/router"
)

func main() {
	database.DBConnection()

	router.Router().Run()
}
