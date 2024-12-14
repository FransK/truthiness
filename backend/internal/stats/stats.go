package stats

import (
	"fmt"
	"log"
	"math"
)

type LinearLeastSquaresData struct {
	M  float64 `json:"m"`
	B  float64 `json:"b"`
	R2 float64 `json:"r2"`
}

func LinearLeastSquares(x, y []float64) (LinearLeastSquaresData, error) {
	if len(x) != len(y) {
		return LinearLeastSquaresData{0, 0, 0}, fmt.Errorf("x and y must have the same length")
	}

	n := float64(len(x))
	if n == 0 {
		return LinearLeastSquaresData{0, 0, 0}, fmt.Errorf("x and y must not be empty")
	}

	// Calculate means of x and y
	sumX, sumY := 0.0, 0.0
	for i := range x {
		sumX += x[i]
		sumY += y[i]
	}
	meanX := sumX / n
	meanY := sumY / n

	// Calculate slope (m) and intercept (b)
	sumXY, sumXX := 0.0, 0.0
	for i := range x {
		dx := x[i] - meanX
		dy := y[i] - meanY
		sumXY += dx * dy
		sumXX += dx * dx
	}

	if sumXX == 0 {
		return LinearLeastSquaresData{0, 0, 0}, fmt.Errorf("cannot compute a slope (vertical line)")
	}

	m := sumXY / sumXX
	b := meanY - m*meanX

	// Calculate R^2
	sst, ssr := 0.0, 0.0
	for i := range y {
		residual := y[i] - (m*x[i] + b)
		sst += math.Pow(y[i]-meanY, 2)
		ssr += math.Pow(residual, 2)
	}

	r2 := 1 - (ssr / sst)
	r := math.Sqrt(r2)

	log.Printf("Linear regression coefficients: m %v - b %v - r %v - r2 %v", m, b, r, r2)

	return LinearLeastSquaresData{m, b, r2}, nil
}

// A regression model that utilizes iterative refinement
func NonLinearLeastSquares() {}
