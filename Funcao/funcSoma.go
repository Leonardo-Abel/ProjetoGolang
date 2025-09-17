package main

import (
	"fmt"
)

func soma(a int, b int) int {
	return a + b
}

func dividir(a, b int) (int, int){
	return a / b, a % b
}

func main() {
	fmt.Print("Soma dos valores: 2 + 3 = ")
	fmt.Println((soma(2, 3)))

	quociente, resto := dividir(10, 3)
	fmt.Println("Divisão de 10 por 3:")
	fmt.Println("Divisão:", quociente, "Resto:", resto)
}