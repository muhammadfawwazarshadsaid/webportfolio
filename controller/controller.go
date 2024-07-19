package controller

import (
	"strconv"
	"webportfolio/database"
	"webportfolio/models/entity"
	"webportfolio/models/req"

	"github.com/gofiber/fiber/v2"
)

func UserControllerShow(c *fiber.Ctx) error{
	return c.JSON(fiber.Map{
		"message" : "Hello World",
	})
}

func GetCompaniesHandler(c *fiber.Ctx) error {
	companies, err := database.GetCompanies()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.JSON(companies)
}

func GetACompanyHandler(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid company ID",
		})
	}

	company, err := database.GetCompanyByID(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to get company",
		})
	}
	if company == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Company not found",
		})
	}

	return c.JSON(company)
}

func GetProjectsHandler(c *fiber.Ctx) error {
	projects, err := database.GetProjects()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.JSON(projects)
}

func GetAProjectHandler(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid project ID",
		})
	}

	project, err := database.GetProjectByID(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to get project",
		})
	}
	if project == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "project not found",
		})
	}

	return c.JSON(project)
}

func SetProjectsHandler(fc *fiber.Ctx) error{
	projects := new(req.ProjectReq)
	if err := fc.BodyParser(projects); err!= nil{
		return fc.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal parsing body",
		})
	}

	transferProjects := &entity.Project{
		Name: projects.Name,
		Description : projects.Description,
		Image: projects.Image,
		Field: projects.Field,
		URLProject: projects.URLProject,
		UpdatedAt: projects.UpdatedAt,
		UploadedAt: projects.UploadedAt,
	}
	if err := database.CreateProjects(transferProjects); err!= nil{
		return fc.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message" : "Gagal buat data baru pada project",
		})
	}
	return fc.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message" 	: "Berhasil buat data baru di project",
		"data"		: transferProjects,
	})
}

func SetCompaniesHandler(fc *fiber.Ctx) error{
	companies := new(req.CompanyReq)
	if err := fc.BodyParser(companies); err!=nil{
		return fc.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message" : "Gagal parsing body",
		})
	}

	transferCompanies := &entity.Company{
		Name: companies.Name,
		About: companies.About,
	}
	if err:= database.CreateCompanies(transferCompanies); err!=nil{
		return fc.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message" : "Gagal buat data baru pada company",
		})
	}
	return fc.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message" 	: "Berhasil buat data baru di companies",
		"data" 		: transferCompanies,
	})
}