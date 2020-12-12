package entities

//PaginationRequest pagination request model
type FilterRequest struct {
	Search string
	Sort   string
	Filter string
	Page   string
	Limit  string
}

//FilterResponse response model
type FilterResponse struct {
	CurrentPage    int `json:"current_page"`
	PageCount      int `json:"page_count"`
	PageSize       int `json:"page_size"`
	RowCount       int `json:"total"`
	FirstRowOnPage int `json:"first_row_on_page"`
	LastRowOnPage  int `json:"last_row_on_page"`
}
