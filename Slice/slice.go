package main

import "fmt"

func main(){
	/* ALGUMAS DECLARAÇÕES SOBRE O COMANDO
		. Mais usado que array no Go.
		. É um "ponteiro" para um array interno, com tamanho dinâmico.
		. Você pode adicionar/remover itens facilmente.
		. Slices usam arrays internos, mas você nunca precisa se preocupar em gerenciar isso manualmente.
		. len(slice) → número de elementos.
		. cap(slice) → capacidade máxima antes de precisar realocar.
	*/
	// criação slice diretamente
	frutas := []string{"maçã", "banana"}
	fmt.Println("Frutas:", frutas)   // print da slice antes da adição

	// adicionando elementos dentro do slice
	frutas = append(frutas, "laranja", "uva")
	fmt.Println("Frutas:", frutas)      // print da slice depois da adição

	// fatiando um slice (igual python)
	fmt.Println("Primeiras frutas;", frutas[0:2])   // selecionando os valores da posição 0 à 2 (posição dois ignorada)

	// criando um slice numérico
	numeros := make([]int, 3, 5)  // slice de 3 elementos, capacidade 5
	fmt.Println("Slice:", numeros, "Tamanho:", len(numeros), "Capacidade:", cap(numeros))

	// atribuindo valores
	numeros = append(numeros, 10, 20, 30)
	fmt.Println("Após append:", numeros, "Tamanho:", len(numeros), "Capacidade:", cap(numeros))

}