package main

import "fmt"

func main() {
	/* ALGUMAS DECLARAÇÕES SOBRE O COMANDO
		. Estrutura chave:valor.
		. Super rápido para buscas.
		. Chaves precisam ser comparáveis (ex.: string, int, bool).
	*/

	// declaração de uma map
	idade := map[string]int{
		"Ana": 20,
		"João": 25,
		"Maria": 30,
	} 

	fmt.Println("Mapa:", idade)

	// acessando valor
	fmt.Println("Idade de Ana:", idade["Ana"])

	// alterando valores
	idade["João"] = 22
	fmt.Println("Mapa:", idade)  // mostrar o valor após a alteração

	// verificando se existe
	val, existe := idade["Pedro"]
	fmt.Println("Pedro existe?", existe, "Valor:", val)

	// removendo valores
	delete(idade, "Maria")
	fmt.Println("Após delete:", idade)
}