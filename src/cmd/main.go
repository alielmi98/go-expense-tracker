package main

import (
	"log"

	"github.com/alielmi98/golang-expense-tracker-api/api"
	"github.com/alielmi98/golang-expense-tracker-api/config"
	"github.com/alielmi98/golang-expense-tracker-api/constants"
	"github.com/alielmi98/golang-expense-tracker-api/data/db"
	"github.com/alielmi98/golang-expense-tracker-api/data/db/migrations"
)

// @securityDefinitions.apikey AuthBearer
// @in header
// @name Authorization
func main() {
	cfg := config.GetConfig()

	err := db.InitDb(cfg)
	defer db.CloseDb()
	if err != nil {
		log.Fatalf("caller:%s  Level:%s  Msg:%s", constants.Postgres, constants.Startup, err.Error())
	}
	migrations.Up_1()
	api.InitServer(cfg)
}
