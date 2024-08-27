package entity

import "time"

type ProjectDetailsReq struct {
	Field        	string    `json:"field" db:"field"`
	Name         	string    `json:"name" db:"name"`
	Description  	string    `json:"description" db:"description"`
	Image        	string    `json:"image" db:"image"`
	IsRealProject   bool   	  `json:"isRealProject" db:"isRealProject"`
	URLProject   	string    `json:"urlproject" db:"urlProject"`
	UpdatedAt    	time.Time `json:"updatedat" db:"updatedat"`
	UploadedAt   	time.Time `json:"uploadedat" db:"uploadedat"`
	CompanyName  	string    `json:"companyname" db:"companyname"`
	About		 	string	  `json:"about" db:"about"`
}
