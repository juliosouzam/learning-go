package matematica

import (
	"fmt"
	"strconv"
)

// Media é responsável por calcular o que você já sabe...
func Media(n ...float64) float64 {
	total := 0.0

	for _, nu := range n {
		total += nu
	}

	media := total / float64(len(n))
	mediaArredondada, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", media), 64)

	return mediaArredondada
}
