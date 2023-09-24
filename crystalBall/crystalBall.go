package crystalBall

import "math"

func CrystalBall(breaks []bool) int {
	jumpAmount := int(math.Floor(math.Sqrt(float64(len(breaks)))))
	i := jumpAmount

	for ; i < len(breaks); i += jumpAmount {
		if breaks[i] {
			break
		}
	}

	i -= jumpAmount

	for j := 0; j <= jumpAmount && i < len(breaks); i, j = i+1, j+1 {
		if breaks[i] {
			return i
		}
	}

	return -1
}
