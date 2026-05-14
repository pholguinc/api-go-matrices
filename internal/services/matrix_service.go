package services

import (
	"math"
	"github.com/pholguinc/api-go-matrices/internal/dtos"
	"github.com/pholguinc/api-go-matrices/internal/models"
	"github.com/pholguinc/api-go-matrices/internal/repositories"
)

type MatrixService interface {
	FactorizeQR(userID string, matrix [][]float64) (dtos.QRResponseData, error)
}

type matrixService struct {
	repo repositories.MatrixRepository
}

func NewMatrixService(repo repositories.MatrixRepository) MatrixService {
	return &matrixService{repo: repo}
}

func (s *matrixService) FactorizeQR(userID string, matrix [][]float64) (dtos.QRResponseData, error) {
	q, r := ComputeQR(matrix)
	
	result := dtos.QRResponseData{
		Q: q,
		R: r,
	}

	record := models.MatrixRecord{
		UserID: userID,
		Input:  matrix,
		Q:      q,
		R:      r,
	}
	
	_ = s.repo.SaveRecord(&record)

	return result, nil
}

func ComputeQR(a [][]float64) ([][]float64, [][]float64) {
	m := len(a)
	if m == 0 {
		return nil, nil
	}
	n := len(a[0])

	qData := make([]float64, m*n)
	rData := make([]float64, n*n)
	
	q := make([][]float64, m)
	for i := range q {
		q[i] = qData[i*n : (i+1)*n]
	}
	r := make([][]float64, n)
	for i := range r {
		r[i] = rData[i*n : (i+1)*n]
	}

	v := make([]float64, m)

	for j := 0; j < n; j++ {
		for i := 0; i < m; i++ {
			v[i] = a[i][j]
		}

		for i := 0; i < j; i++ {
			dot := 0.0
			for k := 0; k < m; k++ {
				dot += q[k][i] * v[k]
			}
			
			r[i][j] = dot
			for k := 0; k < m; k++ {
				v[k] -= dot * q[k][i]
			}
		}

		normVal := 0.0
		for k := 0; k < m; k++ {
			normVal += v[k] * v[k]
		}
		
		normVal = math.Sqrt(normVal)

		r[j][j] = normVal
		if normVal > 1e-10 {
			for i := 0; i < m; i++ {
				q[i][j] = v[i] / normVal
			}
		}
	}

	return q, r
}
