package stats

import (
	"fmt"
	"math"
)

func LinearLeastSquares(x, y []float64) (float64, float64, float64, error) {
	if len(x) != len(y) {
		return 0, 0, 0, fmt.Errorf("x and y must have the same length")
	}

	n := float64(len(x))
	if n == 0 {
		return 0, 0, 0, fmt.Errorf("x and y must not be empty")
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
		return 0, 0, 0, fmt.Errorf("cannot compute a slope (vertical line)")
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

	return m, b, r2, nil
}

// A regression model that utilizes iterative refinement
func NonLinearLeastSquares() {}
