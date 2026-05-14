package tests

import (
	"math"
	"testing"

	"github.com/pholguinc/api-go-matrices/internal/services"
)

func TestComputeQR(t *testing.T) {
	a := [][]float64{
		{1, 0},
		{0, 1},
	}

	q, r := services.ComputeQR(a)

	if len(q) != 2 || len(r) != 2 {
		t.Errorf("Dimensiones incorrectas")
	}

	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			val := 0.0
			for k := 0; k < 2; k++ {
				val += q[i][k] * r[k][j]
			}
			if math.Abs(val-a[i][j]) > 1e-9 {
				t.Errorf("Fallo en A=QR: esperado %v, obtenido %v", a[i][j], val)
			}
		}
	}
}

func TestComputeQR_Rectangular(t *testing.T) {
	a := [][]float64{
		{12, -51},
		{6, 167},
		{-4, 24},
	}

	q, _ := services.ComputeQR(a)

	n := len(a[0])
	m := len(a)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			dot := 0.0
			for k := 0; k < m; k++ {
				dot += q[k][i] * q[k][j]
			}
			expected := 0.0
			if i == j {
				expected = 1.0
			}
			if math.Abs(dot-expected) > 1e-9 {
				t.Errorf("Q no es ortogonal")
			}
		}
	}
}
