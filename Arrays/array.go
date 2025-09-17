package main

import (
	"fmt"        // INTERAGE COM O SISTEMA OPERACIONAL (ARQUIVOS, AMBIENTE, ETC)
)

func main() {
	/* ALGUMAS DECLARAÇÕES SOBRE O COMANDO
		. Estrutura fixa (não muda de tamanho).
		. Os elementos têm sempre o mesmo tipo.
		. Útil quando você sabe exatamente o tamanho.
		. Observação: arrays não são muito usados em Go, porque são limitados.
	*/
	var numeros [5] int // declarando o array de tamanho 5 que recebe valores inteiros
	numeros[0] = 10     // atribuindo valor 10 a primeira posição do array numeros
	numeros[1] = 20     // atribuindo valor 20 a segunada posição do array numeros

	fmt.Println("Array:", numeros)          // mostrando os valores em cada posição do array
	fmt.Println("Tamanho:", len(numeros))   // mostrando o tamanho/quantidade de posições no array
}