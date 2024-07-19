package route

import (
	"webportfolio/config"
	"webportfolio/controller"

	"github.com/gofiber/fiber/v2"
)

func RouterApp(app *fiber.App) {
	app.Get("/", controller.UserControllerShow)
	app.Static("/public", config.StaticDir)

	json 		:= app.Group("/api")
	get			:= json.Group("/get")
	create		:= json.Group("/create")

	get.Get("/projects", controller.GetProjectsHandler)
	get.Get("/companies", controller.GetCompaniesHandler)
	get.Get("/projects/:id", controller.GetAProjectHandler)
	get.Get("/companies/:id", controller.GetACompanyHandler)

	create.Post("/projects", controller.SetProjectsHandler)
	create.Post("/companies", controller.SetCompaniesHandler)
	
}