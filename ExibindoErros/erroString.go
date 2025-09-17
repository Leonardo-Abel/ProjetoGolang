package main

import (
	"errors" // BIBLIOTECA PARA MOSTRAR ERROS DE SISTEMA
	"fmt"
)

func main(){
	err := errors.New("algo deu errado") // declaração da variável que carrega a mensagem
	fmt.Println("Erro:", err)            // mostrando o erro via Terminal
}