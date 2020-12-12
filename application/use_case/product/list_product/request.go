package list_product

import domain "altastore-api/domain/entities"

//ProductFilterRequest .
type ProductFilterRequest struct {
	Search string `json:"search"`
	Page   string `json:"page"`
	Limit  string `json:"limit"`
}

//RequestMapper mapper request
func RequestMapper(req ProductFilterRequest) domain.FilterRequest {
	return domain.FilterRequest{
		Search: req.Search,
		Page:   req.Page,
		Limit:  req.Limit,
	}
}
