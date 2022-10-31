package formas

import (
	"math"
	"testing"
)

func TestArea(t *testing.T) {

	t.Run("Retangulo", func(t *testing.T) {
		t.Parallel()
		ret := Retangulo{10, 12}
		areaEsperada := float64(120)
		areRecebida := ret.Area()
		if areaEsperada != areRecebida {
			t.Errorf("A área recebida %f é diferente da área esperada %f", areRecebida, areaEsperada)
		}
	})

	t.Run("Circulo", func(t *testing.T) {
		t.Parallel()
		circ := Circulo{10}
		areaEsperada := float64(math.Pi * 100)
		areaRecebida := circ.Area()

		if areaEsperada != areaRecebida {
			t.Errorf("A área recebida %f é diferente da área esperada %f", areaRecebida, areaEsperada)
		}
	})
}
