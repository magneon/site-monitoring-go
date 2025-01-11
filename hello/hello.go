package main

import (
	"fmt"
)

func main() {

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

	var versao float32 = 1.1
	var nome string
	var idade int

	fmt.Println("Bem-vindo ao uso do monitorador", versao)
	fmt.Print("Informe seu nome: ")
	fmt.Scan(&nome)
	fmt.Print("Informe também sua idade: ")
	fmt.Scan(&idade)

	//fmt.Println("Olá ao Mundo GoLang", nome, "!")
	//fmt.Println("Sua idade é", idade, "e o sistema está na versão", versao)
	fmt.Println()
	//fmt.Println("O tipo da variável nome é", reflect.TypeOf(nome))
	//fmt.Println("O tipo da variável idade é", reflect.TypeOf(idade))
	//fmt.Println("O tipo da variável versao é", reflect.TypeOf(versao))

	fmt.Println("Qual das operações abaixo você deseja realizar?")
	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Exibir logs monitorados")
	fmt.Println("0 - Encerrar monitoramento")
	fmt.Println()

	fmt.Print("Informe a opção à seguir: ")
	var comando int
	fmt.Scan(&comando)
	//ALÉM DA FUNÇÃO SCANF, HÁ A FUNÇÃO SCAN QUE IDENTIFICA AUTOMATICAMENTE O TIPO INFERIDO PRA INCLUIR NA VARIÁVEL

	fmt.Println("A opção selecionada pelo usuário", nome, "foi a opção", comando)
}
