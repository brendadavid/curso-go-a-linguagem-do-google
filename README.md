# Anotações sobre o código


 O Go entende que a variável é do tipo que ela recebe quando é inicializada, por exemplo: nome =: "Douglas", por isso a palavra chave var pode ser removida e substituida por :=

No caso do float, como existem dois tipos (float32 e float64), o Go vai inferir que o tipo da variável será float64

O pacote Reflect tem uma função para saber os tipos das variáveis

    `fmt.Println("Variável nome é do tipo:", reflect.TypeOf(nome), "Variável versao é to tipo:", reflect.TypeOf(versao))
    Saída: Variável nome é do tipo: string Variável versao é to tipo: float64`


O comando Scanf precisa receber o modificador (ex: "%d") com o tipo de variável e o endereço(ponteiro) da variável que é indicado pelo "&" antes do nome da variável

O comando Scan não precisa receber o modificador com o tipo de variável, mas se o tipo for diferente do declarado, a variável recebe 0

Não precisa escrever break após cada case do switch, ele executa apenas um case e encerra mas caso seja escrito, compila normalmente

Quando é retornado duas variáveis e uma só for utilizada, utilizar o operador de identificador em branco _ (underline) no lugar da variável a ser "descartada"

Não existe *while* e *foreach* no Golang, apenas *for*

A instrução *for*, se for passada vazia, sem argumentos depois da palavra chave "for {intrução}", corresponde a um *while true*