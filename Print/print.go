package main

import (
	"fmt"       // FORMATA E IMPRIME MENSAGENS NA TELA
	"log"       // REGISTRA MENSAGENS DE LOG (ERROS, WARNINGS, INFOS)
	"os"        // INTERAGE COM O SISTEMA OPERACIONAL (ARQUIVOS, AMBIENTE, ETC)
)

func main(){
	// PRINTS/MOSTRAR MENSAGENS NA TELA USANDO FMT | RECOMENDADA PARA USO PRÁTICO 
    //     /* ALGUMAS CONSIDERAÇÕES SOBRE O COMANDO:
    //         São os mais recomendados para mensagens normais ao usuário (ex.: prints em console, debug simples).
    //         São claros, legíveis e idiomáticos em Go.
    //         Se você está escrevendo "Olá Mundo", prints de debug ou exibindo informações formatadas, use fmt.
    //     */

        fmt.Print("Olá ")                           // mostra a mensagem sem quebra de linha
        fmt.Println("Mundo!")                       // mostra a mensagem com quebra de linha
        fmt.Println("Seu número da sorte é:", 7)    // concatenação de valores após a mensagem

        nomeDev := "Leonardo Abel"
        nome := "Leonardo"
        idade := 22 
        vivo := true
        
        fmt.Println("\nNome do desenvolvedor:", nomeDev)
        fmt.Println("Nome:", nome)
        fmt.Println("Idade:", idade)
        fmt.Println("Vivo?:", vivo)

        fmt.Printf("\nNome do desenvolvedor: %s \nNome: %s \nIdade: %d\nVivo? %t\n", nomeDev, nome, idade, vivo)      // impressão formatada (tipo printf C) 

        msg := fmt.Sprintf("\nNome do desenvolvedor: %s \nNome: %s \nIdade: %d\nVivo? %t\n", nomeDev, nome, idade, vivo) // cria a string formatada, mas não imprime (útil para logs ou juntar texto).
        fmt.Println(msg)

        // MELHORES: fmt.Println() e fmt.Printf() 
            // são os mais recomendados para mensagens normais ao usuário (ex.: prints em console, debug simples)

        //PLACEHOLDERS (FORMAT VERBS): usados para formatar a saída em fmt.Printf e fmt.Sprintf
            /*
                %s: Representa uma string.
                %d: Representa inteiros.
                %f: Representa um ponto flutuante (float), números em formato decimal.
                %v: Representa um formato padrão para algum valor, provendo uma representação generalizada. É comumente usada quando o o tipo exato não é crítico ou quando você quer uma representação simples.
                %t: Representa um valor boolean (true or false).
                %p: Representa o endereço de um ponteiro.
                %T: Representa o tipo de um valor.
                %x: or %X: Representa um valor hexadecimal (lowercase ou uppercase).
            */

        fmt.Fprintf(os.Stdout, "Escrevendo em Stdout %d\n", 123) // Escrever diretamente em um destino (ex: arquivo, buffer ou tela):
        fmt.Fprintln(os.Stdout, "Olá Mundo")                       // utiliza as bibliotecas OS e FMT
            // Aqui o destino é os.Stdout (a tela), mas poderia ser um arquivo.
    
    // PACKAGE LOG | RECOMENDAÇÕES DE USO POR BOAS PRÁTICAS
        /*  ALGUMAS CONSIDERAÇÕES SOBRE O COMANDO:
            Usado para mensagens de log (erros, warnings, infos).
            Inclui data e hora automaticamente, o que é essencial em aplicações reais.
            Melhor para diferenciar mensagens de log de simples saída para usuário.
        */

        log.Println("Mensagem de log")                   // Log simples (com quebra de linha)
        
        erro := "L0st C0nn3ct10n"
        log.Printf("Erro: %s - %s", "Falha de conexão", erro) // Log formatado
        log.Fatal("Erro grave, encerrando...")           //Log fatal (mostra e encerra o programa):
        log.Panic("Falha inesperada")                    // Log panic (mostra e dispara panic):

        // UTILIZANDO OS PACOTES OS E FMT PARA IMPRESSÃO | Isso é útil quando você quer controle mais baixo nível da saída.
            /* ALGUMAS CONSIDERAÇÕES SOBRE O COMANDO:
                É a mais performática porque escreve diretamente no stdout, sem formatação.
                Porém, é mais verbosa, menos legível e raramente necessária em aplicações comuns.
                Faz sentido em programas de baixo nível, ferramentas CLI que escrevem grandes volumes de dados ou benchmarks.
            */

        // os.Stdout.Write([]byte("Olá Mundo\n"))
}