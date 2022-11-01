package fuzz

import (
	"testing"
	"unicode/utf8"
)

func FuzzReverse(f *testing.F) {
	testecase := []string{"Teste", " ", "tesTeDois"}

	for _, c := range testecase {
		f.Add(c)
	}
	f.Fuzz(func(t *testing.T, orig string) {
		rev, err1 := Reverse(orig)
		rev2, err2 := Reverse(rev)

		if err1 != nil {
			return
		}
		if err2 != nil {
			return
		}

		if orig != rev2 {
			t.Errorf("Esperado : %q, obitido: %q", orig, rev2)
		}
		if utf8.ValidString(orig) && !utf8.ValidString(rev) {
			t.Errorf("A string reversa não é UTF-8 : %q", orig)
		}
	})
}
