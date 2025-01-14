package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
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

/*
	ARRAYS EM GO POSSUEM TAMANHO PRÉ-FIXADO, O QUE TORNA O TRABALHO COM ARRAYS, RUIM
	var sites [4]string
	sites[0] = "https://www.google.com.br"
	sites[1] = "https://www.alura.com.br"
	sites[2] = "https://httpbin.org/status/200"
	sites[3] = "https://httpbin.org/status/404"


	NO CASO, A VANTAGEM É UTILIZAR UM SLICE, QUE É MUTÁVEL E PERMITE CRESCER E REDUZIR O TAMANHO DA ESTRUTURA DE DADOS
	sites [] string {"https://www.google.com.br", "https://www.alura.com.br", "https://httpbin.org/status/200", "https://httpbin.org/status/404"}
*/

/*
NO GO HÁ DUAS FORMAS D ITERAR SOBRE UMA LISTA

#1 FOR TRADICIONAL COM USO DE ÍNDICE:

	for index := 0; index < len(sites); index++ {
		url := sites[index]
		fmt.Print("Verificando site: ", url)
		response, _ := http.Get(url)
		if response.StatusCode == 200 {
			fmt.Println(" - ", response.StatusCode, " [Disponível]")
		} else {
			fmt.Println(" - ", response.StatusCode, " [Indisponível]")
		}
	}

#2 FOR ONDE A ITERAÇÃO É SOBRE OS ITENS DA LISTA

	for index, element := range sites {
		url := element
		fmt.Print("Verificando site: ", url)
		response, _ := http.Get(url)
		if response.StatusCode == 200 {
			fmt.Println(" - ", response.StatusCode, " [Disponível]")
		} else {
			fmt.Println(" - ", response.StatusCode, " [Indisponível]")
		}
	}
*/

/*
A CRIAÇÃO DE CONSTANTES NO GO É FEITA USANDO A PALAVRA CHAVE CONST
const monitoramentos = 3
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

	switch comando {
	case 0:
		exit()
	case 1:
		startMonitoring(getContentToBeMonitored())
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

func startMonitoring(sites []string) {
	var times int
	fmt.Print("Informe o número de vezes que deseja realizar o monitoramento: ")
	fmt.Scan(&times)

	var interval int
	fmt.Print("Informe o intervalo entre monitoramentos (em segundos): ")
	fmt.Scan(&interval)

	fmt.Println("Você escolheu monitorar por", times, "vezes à cada", interval, "segundos")

	/*
		sites := []string{
			"https://www.google.com.br",
			"https://www.alura.com.br",
			"https://httpbin.org/status/200",
			"https://httpbin.org/status/404",
		}
	*/

	fmt.Println("Iniciando monitoramento do sistema...")
	for index := 0; index < times; index++ {
		for index, element := range sites {
			url := element
			testConnection(index, url)
		}
		time.Sleep(time.Duration(interval) * time.Second)
		fmt.Println()
	}
}

func getContentToBeMonitored() []string {
	fmt.Println("Obtendo arquivo de sites")
	sites := []string{}

	/* Utilizar a função os.Open basicamente é para identificarmos se um determinado arquivo/diretório existe. Essa função não retorna o conteúdo, apenas o ponteiro */
	/* Utilizar a função os.ReadFile é interessante para ler e imprimir o conteúdo do arquivo completamente, sem tratativas */
	arquivo, erro := os.Open("sites.txt")
	linha := bufio.NewReader(arquivo)
	if erro != nil {
		fmt.Println("Ocorreu um erro ao abrir o arquivo", erro)
	}

	for {
		conteudo, _, erro := linha.ReadLine()
		if erro == io.EOF {
			break
		} else if erro != nil {
			fmt.Println("Ocorreu um erro ao ler o conteúdo da linha", erro)
		}

		sites = append(sites, string(conteudo))
	}

	arquivo.Close()
	return sites
}

func testConnection(index int, url string) {
	fmt.Print("Verificando site ", index, ": ", url)
	resposta, erro := http.Get(url)
	if erro != nil {
		fmt.Println("Falha ao conectar ao site", url, ", erro", erro)
	} else {
		if resposta.StatusCode == 200 {
			fmt.Println(" -", resposta.StatusCode, "[Disponível]")
		} else {
			fmt.Println(" -", resposta.StatusCode, "[Indisponível]")
		}
	}

}

func showLogs() {
	fmt.Println("Abrindo logs do sistema...")
}

func unknownOption() {
	fmt.Println("Opção inválida!")
	os.Exit(-1)
}
