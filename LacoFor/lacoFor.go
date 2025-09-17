package main

import (
	"fmt"
)

func main(){
	// LAÇO FOR (FORMADO TRADICIONAL)
	for i := 0; i < 10; i++{
		fmt.Println("i =", i)
	}

	// LAÇO FOR (EM FORMATO WHILE)
	i := 0
	for i <= 10 {
		fmt.Println(i)
		i++
	}
}