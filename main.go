package main

import (
	"log"
	"webportfolio/database"
	"webportfolio/database/migration"
	route "webportfolio/routers"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)


func main(){
	database.ConnectDB()
	migration.Migration()
	app := fiber.New()
	port := ":8080"

	route.RouterApp(app)
	err := app.Listen(port)
	if err!= nil{
		log.Fatal()
	}
}