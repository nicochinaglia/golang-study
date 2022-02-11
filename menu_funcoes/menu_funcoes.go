package main

import (
	"fmt"
	"os"
)

func main() {
	introducao()

	comando := selecionaOpcoes()

	switch comando {
	case 1:
		fmt.Println("Iniciando monitoramento")
	case 2:
		fmt.Println("Exibindo logs")
	case 3:
		fmt.Println("saindo do programa")
		os.Exit(0)
	default:
		fmt.Println("Você não selecionou uma opção válida")
		os.Exit(-1)
	}
}

func introducao() {
	fmt.Println("Olá, informe seu nome por favor:")

	var nome string
	fmt.Scanf("%s", &nome)

	fmt.Println("Olá,", nome, "seja bem vindo!")
}

func selecionaOpcoes() int {

	var comandoSelecionado int

	fmt.Println("Por favor selecione uma opção: ")
	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Exibir logs")
	fmt.Println("0 - Sair")

	fmt.Scan(&comandoSelecionado)

	return comandoSelecionado
}
