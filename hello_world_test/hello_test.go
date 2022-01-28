package main

// Importar biblioteca de testes
import "testing"

//Como criar um teste?
// O arquivo precisa ter como nome xxx_test.go
// A função de teste precisa começar com a palavra "Teste"
// A função de teste utiliza apenas o argumento t *testing.T
// Para utilizar *testing.T, é preciso importar a biblioteca "testing"

func TestHello(t *testing.T) {

	// var := serve para criar e atribuir um valor a variável.
	got := Hello()
	want := "Hello, World"

	//Verifica se o valor da função Hello() é o mesmo valor da variável want
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
