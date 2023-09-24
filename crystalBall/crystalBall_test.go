package crystalBall

import (
	"math"
	"math/rand"
	"testing"
)

func TestCrystalBall(t *testing.T) {
	idx := int(math.Floor(rand.Float64() * 10000))
	data := make([]bool, 10000)

	for i := idx; i < 10000; i++ {
		data[i] = true
	}

	if CrystalBall(data) != idx {
		t.Fail()
	}
	if CrystalBall(make([]bool, 821)) != -1 {
		t.Fail()
	}
}
