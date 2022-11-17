package main

import (
	"fmt"
	"reflect"
)

func main() {
	// O Go entende que a variável é do tipo string pois recebe uma String e a palavra chave var pode ser removida e substituida por :=
	nome := "Douglas"
	// No caso do float, como existem dois tipos, o Go vai inferir que o tipo da variável será float64
	versao := 1.1
	fmt.Println("Olá sr.", nome)
	fmt.Println("O projeto está na versão", versao)

	// O pacote Reflect tem uma função para saber os tipos das variáveis
	fmt.Println("Variável nome é do tipo:", reflect.TypeOf(nome), "Variável versao é to tipo:", reflect.TypeOf(versao))
	// Saída: Variável nome é do tipo: string Variável versao é to tipo: float64
}
