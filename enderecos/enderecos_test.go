package enderecos

import (
	"testing"
)

// cenarioDeTeste usada para passar varios tipos de teste
type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

func TestTipoDeEndereco(t *testing.T) {

	cenarioDeTeste := []cenarioDeTeste{
		{"Rua Abc", "Rua"},
		{"rua paulista", "Rua"},
		{"Rua da", "Rua"},
		{"RUA Abc", "Rua"},
		{"Avenida Abc", "Avenida"},
		{"AVENIDA Abc", "Avenida"},
		{"avenida Abc", "Avenida"},
		{"Estrada Abc", "Estrada"},
		{"EsTrAda Abc", "Estrada"},
		{"estrADA Abc", "Estrada"},
		{"Rodovia Abc", "Rodovia"},
		{"Rodovia Abc", "Rodovia"},
		{"Rodovia Abc", "Rodovia"},
	}
	for _, cenario := range cenarioDeTeste {
		tipoDeEnderecoRecebido := TipoDeEndereco(cenario.enderecoInserido)
		if tipoDeEnderecoRecebido != cenario.retornoEsperado {
			t.Errorf("O tipo de endereço recebido %s é diferente do endereço esperado %s", tipoDeEnderecoRecebido, cenario.retornoEsperado)
		}
	}
}
