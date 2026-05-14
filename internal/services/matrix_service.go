package services

import (
	"math"
	"github.com/pholguinc/api-go-matrices/internal/dtos"
)

type MatrixService interface {
	FactorizeQR(matrix [][]float64) (dtos.QRResponseData, error)
}

type matrixService struct{}

func NewMatrixService() MatrixService {
	return &matrixService{}
}

func (s *matrixService) FactorizeQR(matrix [][]float64) (dtos.QRResponseData, error) {
	q, r := computeQR(matrix)
	
	result := dtos.QRResponseData{
		Q: q,
		R: r,
	}

	return result, nil
}

func computeQR(a [][]float64) ([][]float64, [][]float64) {
	m := len(a)
	if m == 0 {
		return nil, nil
	}
	n := len(a[0])

	q := make([][]float64, m)
	for i := range q {
		q[i] = make([]float64, n)
	}
	r := make([][]float64, n)
	for i := range r {
		r[i] = make([]float64, n)
	}

	for j := 0; j < n; j++ {
		v := make([]float64, m)
		for i := 0; i < m; i++ {
			v[i] = a[i][j]
		}

		for i := 0; i < j; i++ {
			r[i][j] = dotProduct(getColumn(q, i), v)
			for k := 0; k < m; k++ {
				v[k] = v[k] - r[i][j]*q[k][i]
			}
		}

		r[j][j] = norm(v)
		if r[j][j] > 1e-10 {
			for i := 0; i < m; i++ {
				q[i][j] = v[i] / r[j][j]
			}
		}
	}

	return q, r
}

func dotProduct(v1, v2 []float64) float64 {
	sum := 0.0
	for i := range v1 {
		sum += v1[i] * v2[i]
	}
	return sum
}

func norm(v []float64) float64 {
	return math.Sqrt(dotProduct(v, v))
}

func getColumn(matrix [][]float64, col int) []float64 {
	column := make([]float64, len(matrix))
	for i := range matrix {
		column[i] = matrix[i][col]
	}
	return column
}
