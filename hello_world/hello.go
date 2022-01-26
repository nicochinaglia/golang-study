package main

// fmt é um pacote com funções para formatar texto, como imprimir no console.
import "fmt"

// Utilizar o comando go mod tidy para adicionar o pacote ao código.
import "rsc.io/quote"

func main() {
    fmt.Println("Oi, Mundo! =)")
	fmt.Println("Esse é meu primeiro código em Go!")
	
	// Chama a função da biblioteca "quote" adicionada ao projeto.
	fmt.Println(quote.Go())
}