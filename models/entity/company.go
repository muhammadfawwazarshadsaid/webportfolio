package entity

// Company represents a company entity in the database
type Company struct {
	ID   	int    `json:"id" db:"id"`        // ID of the company
	Name 	string `json:"name" db:"name"`    // Name of the company
	About 	string `json:"about" db:"about"` // About information of the company
}