# Anotações sobre o código


 O Go entende que a variável é do tipo que ela recebe quando é inicializada, por exemplo: nome := "Douglas", por isso a palavra chave var pode ser removida e substituida por :=

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

Arrays em Go tem tamanho fixo e deve ser declarado o tipo de dado, ex: var teste [5]string

A melhor prática é usar Slice, que é uma abstração do Array e não precisa informar o tamanho (parecido com ArrayList do Java), e os valores podem ser passados na declaração. Ex: `testeSlice := []string{"teste1", "teste2", "teste3"}`. Tanto o Array quanto o slice, pode passar direto a variável para imprimir, não precisa fazer laço ou usar funções como no Java: `fmt.Println(testeSlice)`

Usar append para adicionar itens ao Array, ex: append(array, "propriedade"), usar "len" para ler a quantidade de itens e "cap" para saber a capacidade do Array. Quando o append é usado no slice, o Go dobra a capacidade do Array.

for pode ser usado com range ao invés do famoso `i:=0, i<len(array), i++`, ex: `for i, site := range sites`