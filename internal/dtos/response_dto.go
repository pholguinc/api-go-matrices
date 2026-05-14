package dtos

type ApiResponse struct {
	// Status code of the response
	// @example 200
	Status int `json:"status"`
	// Message describing the result
	// @example "Operación exitosa"
	Message string `json:"message"`
	// Data returned by the operation
	Data interface{} `json:"data,omitempty"`
}

func NewSuccessResponse(status int, message string, data interface{}) ApiResponse {
	return ApiResponse{
		Status:  status,
		Message: message,
		Data:    data,
	}
}

func NewErrorResponse(status int, message string) ApiResponse {
	return ApiResponse{
		Status:  status,
		Message: message,
		Data:    nil,
	}
}
