package main

import "fmt"

type Pessoa struct {
    Nome string
    Idade int
}

// MÉTODO COM RECEPTOR POR VALOR
func (p Pessoa) falar(){
    // fmt.Println("Meu nome é", p.Nome)
    fmt.Println("Meu nome é", p.Nome, "e minha idade é", p.Idade)
}

func (p * Pessoa) envelhecer(){
    p.Idade++
}

func main() {
	// STRUCTS
	/* ALGUMAS DECLARAÇÕES SOBRE O COMANDO
		. É como uma "caixinha de atributos" (parecido com classes simples).
	*/
	// declarando e usando a struct
	p1 := Pessoa{"Carlos", 19}
	p2 := Pessoa{Nome: "Giovana", Idade: 25}

	fmt.Println(p1)
	fmt.Println("Nome:", p2.Nome, "\nIdade:", p2.Idade)

	// STRUCT COM MÉTODOS
	/* ALGUMAS DECLARAÇÕES SOBRE O COMANDO
		Importante: use *Pessoa quando o método precisa alterar o objeto.
	*/
	p := Pessoa{"João", 20}
	p.falar()
	p.envelhecer()
	fmt.Println("Nova idade", p.Idade)
}