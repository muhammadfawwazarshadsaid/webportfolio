package req

// Company represents a company entity in the database
type CompanyReq struct {
	Name 	string `json:"name" db:"name"`    // Name of the company
	About 	string `json:"about" db:"about"` // About information of the company
}