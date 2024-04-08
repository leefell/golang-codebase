package enderecos

import "testing"

type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

func TestTipoDeEndereco(t *testing.T) {
	// O arquivo de teste deve terminar com _test.go
	// E a funcao de Teste tem que ter 'Test' no inicio
	// Para fazer todos os testes de um diretório usa 'go test ./...'
	// Para verificar se um teste está testando a função inteira usa 'go test --cover'
	// Para criar um arquivo de texto com o resultado do go test --cover usa 'go test --coverprofile nomedoarquivo.txt'
	// e para ler esse arquivo usa 'go tool cover --func=nomedoarquivo.txt'
	// existe também o 'go tool cover --html=nomedoarquivo.txt'

	t.Parallel() // esses testes podem rodar em paralelo

	cenariosDeTeste := []cenarioDeTeste{
		{"Rua ABC", "Rua"},
		{"Avenida XYZ", "Avenida"},
		{"Estrada 123", "Estrada"},
		{"Rodovia ABC-DEF", "Rodovia"},
		{"Rua da Amizade", "Rua"},
		{"Praças das Rosas", "Tipo Inválido"},
		{"RUA DOS TONTO", "Rua"},
		{"Avenida Principal", "Avenida"},
		{"Estrada da Montanha", "Estrada"},
		{"RODOVIA dos Sonhos", "Rodovia"},
		{"Rua da Escola", "Rua"},
		{"", "Tipo Inválido"},
	}

	for _, cenario := range cenariosDeTeste {
		retornoRecebido := TipoDeEndereco(cenario.enderecoInserido)
		if retornoRecebido != cenario.retornoEsperado {
			t.Errorf("O tipo recebido %s é diferente do esperado %s", retornoRecebido, cenario.retornoEsperado)
		}
	}
}

func TestQualquer(t *testing.T) {

	t.Parallel()

	if 1 > 2 {
		t.Error("Teste Quebrou!")
	}
}
