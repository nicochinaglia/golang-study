package main

import (
	"fmt"
	"reflect"
)

// Pacote para retornar o tipo da variável

func main() {

	// Variável de texto.
	// Pode ser declarada sem string, exemplo: var nome = "Nicolas".
	var nome string = "Nicolas"

	// Pode-se usar o nome da variável com := valor para criar a variável e já atribuir um valor. Por padrão o Go já identifica o tipo da variável. Variável de número inteiro.
	idade := 28

	// Variável de verdadeiro ou falso.
	var estuda bool = true

	// Variável de número não inteiro.
	var versao float32 = 0.1

	var nomeCompleto = "Nicolas Roberto Chinaglia"

	fmt.Println("Olá, me chamo ", nome, "!\n Minha idade é: ", idade, ".\n Atualmente estou estudando? ", estuda, ".\n Qual a versão desse código? ", versao)

	//reflec.TypeOf(variável) - Retorna o tipo da variável selecionada.
	fmt.Println("Podemos usar a variável sem declarar seu tipo. Por exemplo: a variável nomeCompleto é do tipo:", reflect.TypeOf(nomeCompleto))
}
