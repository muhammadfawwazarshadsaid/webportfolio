package main

import (
	"log"
	"os"

	// "webportfolio/config"
	"webportfolio/database"
	"webportfolio/database/migration"
	route "webportfolio/route"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"

	// "github.com/gofiber/template/html/v2"
	_ "github.com/lib/pq"
)

// func main() {
// 	// Initialize the HTML template engine
// 	engine := html.New("./public/views", ".html")

// 	app := fiber.New(fiber.Config{
// 		Views: engine,
// 	})

// 	// Middleware for logging requests
// 	app.Use(logger.New(logger.Config{
// 		Format: "${time} ${ip} ${method} ${path} ${status} ${latency}\n",
// 	}))

// 	// Serve static files from "public" directory
// 	app.Static("/", config.StaticDir)

// 	// Database and migration
// 	database.ConnectDB()
// 	migration.Migration()

// 	// Routes
// 	route.RouterApp(app)

// 	// Start server
// 	port := ":8080"
// 	err := app.Listen(port)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }
func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = ":3000"
	} else {
		port = ":" + port
	}

	return port
}

func main() {
	app := fiber.New()

	// Middleware for logging requests
	app.Use(logger.New(logger.Config{
		Format: "${time} ${ip} ${method} ${path} ${status} ${latency}\n",
	}))

	// Database and migration
	database.ConnectDB()
	migration.Migration()

	// Routes
	route.RouterApp(app)

	err := app.Listen(getPort())
	if err != nil {
		log.Fatal(err)
	}
}


