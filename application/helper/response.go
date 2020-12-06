package helper

//BaseResponse base response for all response
type BaseResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
