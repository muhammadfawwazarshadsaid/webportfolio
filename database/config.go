package database

import (
	"database/sql"
	"log"
	"webportfolio/models/entity"
)

func ConnectDB() *sql.DB{
	connectSQL := "postgresql://webportfolio_owner:yl2OuIc8Uqai@ep-round-feather-a1z58zox.ap-southeast-1.aws.neon.tech/webportfolio?sslmode=require"
	db, err := sql.Open("postgres", connectSQL)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func GetProjects() ([]entity.Project, error) {
	db := ConnectDB()
	defer db.Close()

	rows, err := db.Query("SELECT p.id, p.companyid, p.field, p.name, p.description, p.image, p.urlproject, p.updatedAt, p.uploadedAt, c.name AS companyName, c.about AS companyAbout FROM PROJECT p JOIN company c ON p.companyid = c.id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var projects []entity.Project
	for rows.Next() {
		var p entity.Project
		err := rows.Scan(&p.ID, &p.CompanyID, &p.Field, &p.Name, &p.Description, &p.Image, &p.URLProject, &p.UpdatedAt, &p.UploadedAt, &p.CompanyName, &p.CompanyAbout)
		if err != nil {
			return nil, err
		}
		projects = append(projects, p)
	}

	return projects, nil
}

func GetProjectDetails(id int) (*entity.ProjectDetailsReq, error) {
	db := ConnectDB()
	defer db.Close()

	var projectdetails entity.ProjectDetailsReq
	query := `SELECT p.field, p.name, p.description, p.image, p.urlproject, p.updatedat, p.uploadedat, c.name AS company_name, c.about AS company_about
			  FROM project p
			  JOIN company c ON p.companyid = c.id
			  WHERE p.id = $1`
	err := db.QueryRow(query, id).Scan(
		&projectdetails.Field,
		&projectdetails.Name,
		&projectdetails.Description,
		&projectdetails.Image,
		&projectdetails.URLProject,
		&projectdetails.UpdatedAt,
		&projectdetails.UploadedAt,
		&projectdetails.CompanyName,
		&projectdetails.About)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		log.Println("Failed to get project by ID:", err)
		return nil, err
	}
	return &projectdetails, nil
}


func GetProjectByID(id int) (*entity.Project, error) {
	db := ConnectDB()
	defer db.Close()

	var project entity.Project
	query := `SELECT id, companyid, field, name, description, image, urlproject, updatedAt, uploadedAt FROM PROJECT WHERE id = $1`
	err := db.QueryRow(query, id).Scan(&project.ID, &project.CompanyID, &project.Field, &project.Name, &project.Description, &project.Image, &project.URLProject, &project.UpdatedAt, &project.UploadedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // No company found
		}
		log.Println("Failed to get company by ID:", err)
		return nil, err
	}

	return &project, nil
}

func GetCompanies() ([]entity.Company, error) {
	db := ConnectDB()
	defer db.Close()

	rows, err := db.Query("SELECT id, name, about FROM COMPANY")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var companies []entity.Company
	for rows.Next() {
		var c entity.Company
		err := rows.Scan(&c.ID, &c.Name, &c.About)
		if err != nil {
			return nil, err
		}
		companies = append(companies, c)
	}

	return companies, nil
}

func GetCompanyByID(id int) (*entity.Company, error) {
	db := ConnectDB()
	defer db.Close()

	var company entity.Company
	query := `SELECT id, name, about FROM COMPANY WHERE id = $1`
	err := db.QueryRow(query, id).Scan(&company.ID, &company.Name, &company.About)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // No company found
		}
		log.Println("Failed to get company by ID:", err)
		return nil, err
	}

	return &company, nil
}

func CreateCompanies(c *entity.Company) error {
	db := ConnectDB()
	defer db.Close()

	query := `INSERT INTO COMPANY (name, about) 
			  VALUES ($1, $2)
			  RETURNING id`
	err := db.QueryRow(query, c.Name, c.About).Scan(&c.ID)
	if err!= nil{
		return err
	}
	return nil
}

func CreateProjects(c *entity.Project) error{
	db := ConnectDB()
	defer db.Close()

	query := `INSERT INTO PROJECT (companyid, field, name, description, image, urlproject, updatedAt, uploadedAt) 
			  VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
			  RETURNING id, companyid`
	err := db.QueryRow(query, c.Name, c.Image, c.Field, c.Description, c.URLProject).Scan(&c.ID,&c.CompanyID)
	if err!= nil {
		return err
	}
	return nil
}