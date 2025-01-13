package main

import (
	"fmt"
	"net/http"
	"os"
)

/*
NO GO EXISTEM TRÊS FORMAS DE DECLARAR VARIÁVEIS

#1 - DECLARAÇÕES COMPLETAS
var nome string = "Rafael"
var idade int = 35
var versao float32 = 1.1

#2 - DECLARAÇÕES COM O TIPO INFERIDO
var nome = "Rafael"
var idade = 35
var versao = 1.1

#3 - DECLARAÇÕES REDUZIDAS
nome := "Rafael"
idade := 35
versao := 1.1
*/
func main() {
	nome := receiveUserName()
	idade := receiveUserAge()

	showWelcome(nome, idade)
	for {
		showOptions()
		comando := receiveOption()
		executeOption(nome, comando)
	}
}

func receiveUserName() string {
	var nome string
	fmt.Print("Informe seu nome: ")
	fmt.Scan(&nome)

	return nome
}

func receiveUserAge() int {
	var idade int
	fmt.Print("Informe também sua idade: ")
	fmt.Scan(&idade)

	return idade
}

func showWelcome(nome string, idade int) {
	var versao float32 = 1.1
	fmt.Println("Bem-vindo ao uso do monitorador", versao)
	fmt.Println(nome, idade)

	//fmt.Println("Olá ao Mundo GoLang", nome, "!")
	//fmt.Println("Sua idade é", idade, "e o sistema está na versão", versao)
	fmt.Println()
	//fmt.Println("O tipo da variável nome é", reflect.TypeOf(nome))
	//fmt.Println("O tipo da variável idade é", reflect.TypeOf(idade))
	//fmt.Println("O tipo da variável versao é", reflect.TypeOf(versao))
}

func showOptions() {
	fmt.Println("Qual das operações abaixo você deseja realizar?")
	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Exibir logs monitorados")
	fmt.Println("0 - Encerrar monitoramento")
	fmt.Println()
}

func receiveOption() int {
	fmt.Print("Informe a opção à seguir: ")
	var comando int
	fmt.Scan(&comando)
	//ALÉM DA FUNÇÃO SCANF, HÁ A FUNÇÃO SCAN QUE IDENTIFICA AUTOMATICAMENTE O TIPO INFERIDO PRA INCLUIR NA VARIÁVEL

	return comando
}

func executeOption(nome string, comando int) {
	fmt.Println("A opção selecionada pelo usuário", nome, "foi a opção", comando)
	/*
		NO GO, OS PARÊNTESES DO IF NÃO SÃO NECESSÁRIOS:
		if (comando == 0) {
			...
		}

		COM A INSTRUÇÃO ACIMA, O GO VAI REMOVER OS PARÊNTESES
		if comando == 0 {
			exit()
		} else if comando == 1 {
			startMonitoring()
		} else if comando == 2 {
			showLogs()
		} else {
			unknownOption()
		}
	*/

	switch comando {
	case 0:
		exit()
	case 1:
		startMonitoring()
	case 2:
		showLogs()
	default:
		unknownOption()
	}
}

func exit() {
	fmt.Println("Desligando programa")
	os.Exit(0)
}

func startMonitoring() {
	fmt.Println("Iniciando monitoramento do sistema...")

	url := "https://httpbin.org/status/200"

	fmt.Print("Verificando site: ", url)
	response, _ := http.Get(url)
	if response.StatusCode == 200 {
		fmt.Println(" - ", response.StatusCode, " [Disponível]")
	} else {
		fmt.Println(" - ", response.StatusCode, " [Indisponível]")
	}
}

func showLogs() {
	fmt.Println("Abrindo logs do sistema...")
}

func unknownOption() {
	fmt.Println("Opção inválida!")
	os.Exit(-1)
}
