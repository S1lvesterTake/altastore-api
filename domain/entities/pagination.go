package entities

//PaginationRequest pagination request model
type PaginationRequest struct {
	CategoryName string
	PerPage      string
	Limit        string
}

//PaginationResponse pagination response model
type PaginationResponse struct {
	Page      uint `json:"page"`
	Limit     uint `json:"limit"`
	TotalPage uint `json:"total_page"`
}
