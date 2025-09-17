package main    // pacote principal
// package → palavra-chave usada para declarar a qual pacote este arquivo pertence.
// main → é o pacote especial que indica que este programa pode ser executado diretamente.

import "fmt"

func main(){  
    /* OBSERVAÇÕES IMPORTANTES:
        GO NÃO PERMITE RODAR/COMPILAR O CÓDIGO CASO HAJA VARIÁVEIS, IMPORTS, FUNÇÕES OU QUALQUER OUTRA COISA QUE NÃO ESTEJA SENDO USADA
    */

    // DECLARAÇÃO DE VARIÁVEIS
        /* ALGUMAS CONSIDERAÇÕES SOBRE O COMANDO:
            Declaração mais recomendada: declaração curta com :=
                . É a forma mais usada e idiomática dentro de funções.
                . Clara, concisa e evita repetição desnecessária do tipo.
                . O compilador infere o tipo automaticamente a partir do valor.
                . Recomendado na maioria dos casos em que a variável já pode ser inicializada.
                Ex: nome := "Leonardo"

            Usada em boas práticas de produção: var explícito com ou sem valor inicial
                Usado quando:
                . Você quer declarar no nível de pacote (fora de funções, := não funciona).
                . Você precisa dar zero-value (valor padrão do tipo) e inicializar depois.
                . Quer deixar explícito o tipo por clareza.
                . Esse estilo é comum em códigos mais sérios porque aumenta a clareza em alguns contextos.
                Ex: var idade int = 22 / var idade int

            Forma a mais “performática” (baixo nível): Declaração com var sem inicialização
                . O Go automaticamente atribui o zero-value (ex.: 0, "", false, nil).
                . Não há diferença real de performance em relação ao := com inicialização, mas essa forma evita a atribuição inicial desnecessária.
                . Útil em cenários como buffers, arrays grandes ou structs complexas.
                Ex: var buffer [1024]byte // alocado zerado

            Resumindo boas práticas:
                . No dia a dia (recomendado/idiomático): := (declaração curta).
                . Em produção (quando precisa clareza/tipo explícito ou fora de funções): var.
                . Em baixo nível / otimização específica: var sem inicialização (zero-value).
                . Em 99% dos casos, gophers usam := dentro de funções e var fora delas
        */

        // var name type = value            // forma padrão de declaração de variáveis

        const nomeDev = "Leonardo Abel"     // variável constante, logo, ela não será alterada ao decorrer do programa
        // var nome string = "Leonardo"     // caso nenhum valor seja declarado, inicializará como ""
        nome := "Leonardo"                  // := inferência automática de tipo
        // var idade int = 22               // caso nenhum valor seja declarado, inicializará como 0
        idade := 22 
        // var count = 10                      //  count é inferido como int
        // var vivo bool = true
        vivo := true

        fmt.Println("Nome do desevolvedor:", nomeDev, "\nNome:", nome, "\nIdade:", idade, "\nVivo?:", vivo)
}