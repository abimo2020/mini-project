package main

import (
	"mini-project/config"
	"mini-project/lib/seeder"
	route "mini-project/routes"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func init() {
	config.InitDB()
	config.InitialMigration()
	seeder.DBSeed(config.DB)
}

func main() {

	e := route.New()
	e.Logger.Fatal(e.Start(":8080"))
}
