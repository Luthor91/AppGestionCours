package main

import (
	"log"
	"main_module/controller"
	"main_module/database"
	"main_module/database/migration"
	"main_module/database/sync"
	"main_module/server"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Erreur lors du chargement du fichier .env: %v", err)
	}

	rd := database.ConnectToRedis()

	database.CreatePostgresDatabase()
	db := database.ConnectToPostgres()

	migration.MigrateAllPostgresql(db)
	sync.SyncData(rd, db)
	controller.InsertTestDatas(db)

	server.DefineRoutes()
}
