package main

import (
	"fmt"
)

func main() {
	fmt.Println("Bem vindo ao programa de estudo, por favor selecione uma opção: ")
	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Exibir logs")
	fmt.Println("0 - Sair")

	var comando int

	// fmt.Scanf é responsável por receber o input do usuário para a variável comando.
	// Precisa utilizar o & na variável para sinalizar que é o endereço da variável comando previamente criada.
	fmt.Scanf("%d", &comando)

	//Também é possível utilizar a função fmt.Scan(&comando), não sendo necessário inserir o modificador de variável.

	fmt.Println("Opção escolhida:", comando)

	// If sempre precisa retornar um boolean.
	/*
		if comando == 1 {
			fmt.Println("Iniciando monitoramento")
		} else if comando == 2 {
			fmt.Println("Exibindo logs")
		} else if comando == 0 {
			fmt.Println("saindo do programa")
		} else {
			fmt.Println("Você não selecionou uma opção válida")
		}
	*/

	//Switch
	switch comando {
	case 1:
		fmt.Println("Iniciando monitoramento")
	case 2:
		fmt.Println("Exibindo logs")
	case 3:
		fmt.Println("saindo do programa")
	default:
		fmt.Println("Você não selecionou uma opção válida")
	}
}
