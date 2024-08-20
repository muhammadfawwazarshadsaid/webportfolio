package route

import (
	// "webportfolio/config"
	"webportfolio/controller"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func RouterApp(app *fiber.App) {
	app.Use(logger.New())
	app.Get("/", controller.Index)
	app.Get("/projectdetails/:id", controller.GetProjectDetails)
	// app.Static("/public", config.StaticDir)
    // app.Static("/css", "./public/css")
    // app.Static("/js", "./public/js")

	json := app.Group("/api")
	get := json.Group("/get")
	create := json.Group("/create")

	get.Get("/projects", controller.GetProjectsHandler)
	get.Get("/companies", controller.GetCompaniesHandler)
	get.Get("/projects/:id", controller.GetAProjectHandler)
	get.Get("/companies/:id", controller.GetACompanyHandler)
	get.Get("/projectdetails/:id", controller.GetProjectDetailsHandler)

	create.Post("/projects", controller.SetProjectsHandler)
	create.Post("/companies", controller.SetCompaniesHandler)
}
