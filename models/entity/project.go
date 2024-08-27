package entity

import "time"

type Project struct {
	ID         		int      	`json:"id" db:"id"`               // ID of the project
	CompanyID  		int      	`json:"companyId" db:"companyid"` // Foreign key to the company
	Field      		string    	`json:"field" db:"field"`         // Field of the project
	Name      		string  	`json:"name" db:"name"`           // Name of the project
	Description		string   	`json:"description" db:"description"` // Description of the project
	Image      		string   	`json:"image" db:"image"`         // URL of the project image
	IsRealProject 	bool    	`json:"isRealProject" db:"isRealProject"` // URL of the project
	URLProject 		string    	`json:"urlProject" db:"urlproject"` // isRealProject
	UpdatedAt  		time.Time 	`json:"updatedAt" db:"updatedAt"` // Last update timestamp
	UploadedAt 		time.Time 	`json:"uploadedAt" db:"uploadedAt"` // Creation timestamp
	CompanyName		string 		`json:"companyName" db:"companyName"`
	CompanyAbout 	string 		`json:"companyAbout" db:"companyAbout"`
}

