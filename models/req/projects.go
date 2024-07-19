package req

import "time"

type ProjectReq struct {
	Field      	string    `json:"field" db:"field"`         // Field of the project
	Name      	string    `json:"name" db:"name"`           // Name of the project
	Description string   `json:"description" db:"description"` // Description of the project
	Image      	string    `json:"image" db:"image"`         // URL of the project image
	URLProject 	string    `json:"urlProject" db:"urlproject"` // URL of the project
	UpdatedAt  	time.Time `json:"updatedAt" db:"updatedAt"` // Last update timestamp
	UploadedAt 	time.Time `json:"uploadedAt" db:"uploadedAt"` // Creation timestamp
}

