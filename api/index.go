// package main

// import (
// 	"log"
// 	"os"

// 	// "webportfolio/config"
// 	"webportfolio/database"
// 	"webportfolio/database/migration"
// 	route "webportfolio/routers"

// 	"github.com/gofiber/fiber/v2"
// 	"github.com/gofiber/fiber/v2/middleware/logger"

// 	// "github.com/gofiber/template/html/v2"
// 	_ "github.com/lib/pq"
// )

// // func main() {
// // 	// Initialize the HTML template engine
// // 	engine := html.New("./public/views", ".html")

// // 	app := fiber.New(fiber.Config{
// // 		Views: engine,
// // 	})

// // 	// Middleware for logging requests
// // 	app.Use(logger.New(logger.Config{
// // 		Format: "${time} ${ip} ${method} ${path} ${status} ${latency}\n",
// // 	}))

// // 	// Serve static files from "public" directory
// // 	app.Static("/", config.StaticDir)

// // 	// Database and migration
// // 	database.ConnectDB()
// // 	migration.Migration()

// // 	// Routes
// // 	route.RouterApp(app)

// // 	// Start server
// // 	port := ":8080"
// // 	err := app.Listen(port)
// // 	if err != nil {
// // 		log.Fatal(err)
// // 	}
// // }

// func main() {
// 	app := fiber.New()

// 	// Middleware for logging requests
// 	app.Use(logger.New(logger.Config{
// 		Format: "${time} ${ip} ${method} ${path} ${status} ${latency}\n",
// 	}))

// 	// Database and migration
// 	database.ConnectDB()
// 	migration.Migration()

// 	// Routes
// 	route.RouterApp(app)

// 	// Get port from environment variable or use default
// 	port := os.Getenv("PORT")
// 	if port == "" {
// 		port = ":8080" // Default port if not set
// 	}

//		// Start server
//		app.Listen(port)
//		err := app.Listen(port)
//		if err != nil {
//			log.Fatal(err)
//		}
//	}
package handler

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"webportfolio/models/entity"
	"webportfolio/models/req"

	_ "github.com/lib/pq"
)

// Handler function for Vercel deployment
func Handler(w http.ResponseWriter, r *http.Request) {
	
	// Create a new ServeMux for routing
	mux := http.NewServeMux()	
	// Initialize the database and migration
	// database.ConnectDB()
	// migration.Migration()
	ConnectDB()
	Migration()

	// // Serve static files from "public" directory
	// staticDir := "./public" // Adjust the static directory path as needed
	// fileServer := http.FileServer(http.Dir(staticDir))
	// mux.Handle("/", http.StripPrefix("/", fileServer))

	// Set up routes
	RouterApp(mux) // Assuming you have a function to set up routes in `webportfolio/routers`
	// Tambahkan middleware CORS ke seluruh aplikasi

	// Atur header CORS
	w.Header().Set("Access-Control-Allow-Origin", "https://www.arshad.my.id, https://aboutme-pied-theta.vercel.app/, https://aboutme-pied-theta.vercel.app/api/get/projects, https://www.arshad.my.id/about, https://www.google.com/, https://www.muhammadfawwazarshadsaid.github.io/about, https://www.muhammadfawwazarshadsaid.github.io/")
	w.Header().Set("Access-Control-Allow-Methods", "OPTIONS, GET, POST, PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-CSRF-Token")




	// Lanjutkan ke handler berikutnya
	mux.ServeHTTP(w,r)

	// Get port from environment variable or use default
	// port := os.Getenv("PORT")
	// if port == "" {
	// 	port = "8080" // Default port if not set
	// }

	// // Start server
	// log.Printf("Server is starting on port %s", port)
	// err := http.ListenAndServe(":"+port, mux)
	// if err != nil {
	// 	log.Fatal(err)
	// }

}

func RouterApp(mux *http.ServeMux) {
	// Serve static files
	// staticDir := "./public" // Adjust the static directory path as needed
	// fileServer := http.FileServer(http.Dir(staticDir))
	// mux.Handle("/", http.StripPrefix("/", fileServer))

	// API routes
	mux.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request){
	// Atur header CORS
	// Atur header CORS
	w.Header().Set("Access-Control-Allow-Origin", "https://www.arshad.my.id, https://aboutme-pied-theta.vercel.app/, https://aboutme-pied-theta.vercel.app/api/get/projects, https://www.arshad.my.id/about, https://www.google.com/, https://www.muhammadfawwazarshadsaid.github.io/about, https://www.muhammadfawwazarshadsaid.github.io/")
	w.Header().Set("Access-Control-Allow-Methods", "OPTIONS, GET, POST, PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-CSRF-Token")


	mux.ServeHTTP(w,r)
	})

	mux.HandleFunc("/api/get/projects", GetProjectsHandler)
	mux.HandleFunc("/api/get/companies", GetCompaniesHandler)
	mux.HandleFunc("/api/get/projects/", GetAProjectHandler)
	mux.HandleFunc("/api/get/companies/", GetACompanyHandler)
	mux.HandleFunc("/api/get/projectdetails/", GetProjectDetailsHandler)

	mux.HandleFunc("/api/create/projects", SetProjectsHandler)
	mux.HandleFunc("/api/create/companies", SetCompaniesHandler)
}
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

	rows, err := db.Query("SELECT p.id, p.companyid, p.field, p.name, p.description, p.image, p.urlproject, p.updatedAt, p.uploadedAt, p.isRealProject, c.name AS companyName, c.about AS companyAbout FROM PROJECT p JOIN company c ON p.companyid = c.id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var projects []entity.Project
	for rows.Next() {
		var p entity.Project
		err := rows.Scan(&p.ID, &p.CompanyID, &p.Field, &p.Name, &p.Description, &p.Image, &p.URLProject, &p.UpdatedAt, &p.UploadedAt, &p.IsRealProject, &p.CompanyName, &p.CompanyAbout)
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
	query := `SELECT p.field, p.name, p.description, p.image, p.urlproject, p.updatedat, p.uploadedat, p.isRealProject, c.name AS company_name, c.about AS company_about
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
		&projectdetails.IsRealProject,
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
	query := `SELECT id, companyid, field, name, description, image, urlproject, updatedAt, uploadedAt, isRealProject FROM PROJECT WHERE id = $1`
	err := db.QueryRow(query, id).Scan(&project.ID, &project.CompanyID, &project.Field, &project.Name, &project.Description, &project.Image, &project.URLProject, &project.UpdatedAt, &project.UploadedAt, &project.IsRealProject)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // No project found
		}
		log.Println("Failed to get project by ID:", err)
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

func CreateProjects(c *entity.Project) error {
	db := ConnectDB()
	defer db.Close()

	query := `INSERT INTO PROJECT (companyid, field, name, description, image, urlproject, updatedAt, uploadedAt, isRealProject) 
			  VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
			  RETURNING id, companyid`
	err := db.QueryRow(query, c.CompanyID, c.Field, c.Name, c.Description, c.Image, c.URLProject, c.UpdatedAt, c.UploadedAt, c.IsRealProject).Scan(&c.ID, &c.CompanyID)
	if err != nil {
		return err
	}
	return nil
}


func Migration(){

	db := ConnectDB()
	defer db.Close()
}


func Index(w http.ResponseWriter, r *http.Request) {
	projects, err := GetProjects()
	if err != nil {
		http.Error(w, "Gagal mengambil data", http.StatusInternalServerError)
		return
	}

	// Implement rendering of HTML templates if needed, or respond with JSON
	// Example: Responding with JSON for simplicity
	response := map[string]interface{}{
		"Projects": projects,
		"Title":    "Welcome",
	}
	json.NewEncoder(w).Encode(response)
}

func GetProjectDetailsHandler(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Gagal menemukan id", http.StatusBadRequest)
		return
	}
	projectdetails, err := GetProjectDetails(id)
	if err != nil {
		http.Error(w, "Gagal mendapat project details", http.StatusBadGateway)
		return
	}
	if projectdetails == nil {
		http.Error(w, "Project tidak ditemukan", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(projectdetails)
}

func GetProjectsHandler(w http.ResponseWriter, r *http.Request) {
	projects, err := GetProjects()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(projects)
}

func GetCompaniesHandler(w http.ResponseWriter, r *http.Request) {
	companies, err := GetCompanies()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(companies)
}

func GetACompanyHandler(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Invalid company ID", http.StatusBadRequest)
		return
	}

	company, err := GetCompanyByID(id)
	if err != nil {
		http.Error(w, "Failed to get company", http.StatusInternalServerError)
		return
	}
	if company == nil {
		http.Error(w, "Company not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(company)
}

func GetAProjectHandler(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Invalid project ID", http.StatusBadRequest)
		return
	}

	project, err := GetProjectByID(id)
	if err != nil {
		http.Error(w, "Failed to get project", http.StatusInternalServerError)
		return
	}
	if project == nil {
		http.Error(w, "Project not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(project)
}

func SetProjectsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var projects req.ProjectReq
	if err := json.NewDecoder(r.Body).Decode(&projects); err != nil {
		http.Error(w, "Gagal parsing body", http.StatusInternalServerError)
		return
	}

	transferProjects := &entity.Project{
		Name:         projects.Name,
		Description:  projects.Description,
		Image:        projects.Image,
		Field:        projects.Field,
		URLProject:   projects.URLProject,
		UpdatedAt:    projects.UpdatedAt,
		UploadedAt:   projects.UploadedAt,
	}
	if err := CreateProjects(transferProjects); err != nil {
		http.Error(w, "Gagal buat data baru pada project", http.StatusInternalServerError)
		return
	}
	response := map[string]interface{}{
		"message": "Berhasil buat data baru di project",
		"data":    transferProjects,
	}
	json.NewEncoder(w).Encode(response)
}

func SetCompaniesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var companies req.CompanyReq
	if err := json.NewDecoder(r.Body).Decode(&companies); err != nil {
		http.Error(w, "Gagal parsing body", http.StatusInternalServerError)
		return
	}

	transferCompanies := &entity.Company{
		Name:  companies.Name,
		About: companies.About,
	}
	if err := CreateCompanies(transferCompanies); err != nil {
		http.Error(w, "Gagal buat data baru pada company", http.StatusInternalServerError)
		return
	}
	response := map[string]interface{}{
		"message": "Berhasil buat data baru di companies",
		"data":    transferCompanies,
	}
	json.NewEncoder(w).Encode(response)
}