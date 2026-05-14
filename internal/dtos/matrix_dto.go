package dtos

type QRRequest struct {
	// Matrix is a 2D array of numbers
	// @example [[12, -51], [6, 167]]
	Matrix [][]float64 `json:"matrix" validate:"required"`
}

type QRResponseData struct {
	Q [][]float64 `json:"q"`
	R [][]float64 `json:"r"`
}
