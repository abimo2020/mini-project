package main

import (
	"mini-project/config"
	route "mini-project/routes"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func init() {
	config.InitDB()
	config.InitialMigration()
}

func main() {

	e := route.New()
	e.Logger.Fatal(e.Start(":8080"))
}
