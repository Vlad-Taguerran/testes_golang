package enderecos

import "strings"

// TipoDeEndereco verifica se um endereço é valido pela primeira palavra
func TipoDeEndereco(endereco string) string {
	tiposValidos := []string{"rua", "avenida", "estrada", "rodovia"}

	enderecoLetraMinuscula := strings.ToLower(endereco)
	primeiraPalavra := strings.Split(enderecoLetraMinuscula, " ")[0]

	enderecoTemTipoValido := false

	for _, tipo := range tiposValidos {
		if tipo == primeiraPalavra {
			enderecoTemTipoValido = true
		}
	}

	if enderecoTemTipoValido {
		return strings.Title(primeiraPalavra)
	}

	return "tipo invalido"
}
