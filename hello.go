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

	// O comando Scan não precisa receber o modificador com o tipo de variável, mas se o tipo for diferente de int a variável recebe 0
	var comando int
	fmt.Scan(&comando)

	// switch não precisa de break, executa apenas um case e encerra
	switch comando {
	case 1:
		fmt.Println("Monitorando...")
	case 2:
		fmt.Println("Exibindo Logs...")
	case 0:
		fmt.Println("Saindo do Programa")
	default:
		fmt.Println("Não conheço este comando")
	}
}
