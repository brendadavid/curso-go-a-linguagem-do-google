package main

import "fmt"

func main() {
	// O Go entende que a variável é do tipo string pois recebe uma String e a palavra chave var pode ser removida e substituida por :=
	nome := "Douglas"
	// No caso do float, como existem dois tipos, o Go vai inferir que o tipo da variável será float64
	versao := 1.1

	fmt.Println("Olá sr.", nome)
	fmt.Println("O projeto está na versão", versao)

	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")

	// O comando Scanf precisa receber o modificador com o tipo de variável e o endereço(ponteiro) da variável
	var comando int
	fmt.Scanf("%d", &comando)

	// If não usa parênteses e só entende expressões booleanas
	if comando == 1 {
		fmt.Println("Monitorando...")
	} else if comando == 2 {
		fmt.Println("Exibindo Logs...")
	} else if comando == 0 {
		fmt.Println("Saindo do Programa")
	} else {
		fmt.Println("Não conheço este comando")
	}
}
