package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitoramentos = 3
const delay = 5

func main() {
	// showNames()
	showIntroduction()
	for { // Loop infinito
		showMenu()

		// _, idade := returnNameAndAge()
		// fmt.Println(idade)

		// if comando == 1 {
		// 	fmt.Println("Monitorando...")
		// } else if comando == 2 {
		// 	fmt.Println("Exibindo logs...")
		// } else if comando == 0 {
		// 	fmt.Println("Saindo...")
		// } else {
		// 	fmt.Println("Não conheço esse comando")
		// }

		comando := inputComand()
		switch comando {
		case 1:
			initMonitoring()
		case 2:
			showLogs()
		case 0:
			fmt.Println("Saindo...")
			os.Exit(0)
		default:
			fmt.Println("Não conheço esse comando")
			os.Exit(-1)
		}
	}
}

// func returnNameAndAge() (string, int) {
// 	name := "Karl"
// 	idade := 22
// 	return name, idade
// }

func showIntroduction() {
	var name = "Karl"
	version := 1.1 //Inferir o tipo da variável
	fmt.Println("Olá", name, "este programa está na versão", version)
}

func showMenu() {
	fmt.Println("1 - Iniciar o monitoramento")
	fmt.Println("2 - Exibir logs")
	fmt.Println("0 - Sair")
}

func inputComand() int {
	var command int
	// fmt.Scanf("%d", &comando) //o operador & é o endereço da variável
	fmt.Scan(&command) //Não precisa inferir o tipo do dado inputado
	fmt.Println("O comando escolhido foi", command)

	return command
}

func initMonitoring() {
	fmt.Println("Monitorando...")

	// sites := []string{
	// 	"https://random-status-code.herokuapp.com/",
	// 	"https://google.com",
	// 	"https://youtube.com"}

	// for i := 0; i < len(sites); i++ {
	// 	fmt.Println(sites[i])
	// }

	sites := readSitesFile()

	for i := 0; i < monitoramentos; i++ {
		for i, site := range sites {
			fmt.Println("Testando site", i, ":", site)
			testSite(site)

		}
		fmt.Println("")

		time.Sleep(delay * time.Second)
	}
	fmt.Println("Terminou o monitoramento")
	fmt.Println("")

}

func testSite(site string) {
	res, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	if res.StatusCode == 200 {
		fmt.Println("O site", site, "foi carregado com sucesso")
		registerLog(site, true)
	} else {
		fmt.Println("O site", site, "está com problemas")
		registerLog(site, false)
	}
}

func readSitesFile() []string {
	var sites []string

	arquivo, err := os.Open("sites.txt")
	// arquivo, err := ioutil.ReadFile("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	leitor := bufio.NewReader(arquivo)

	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)

		sites = append(sites, linha)

		if err == io.EOF {
			break
		}
	}

	arquivo.Close()

	return sites
}

func registerLog(site string, status bool) {
	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online: " + strconv.FormatBool(status) + "\n")

	arquivo.Close()
}

func showLogs() {
	fmt.Println("Exibindo logs...")

	arquivo, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	fmt.Println(string(arquivo))
}

// func showNames() {
// 	nomes := []string{"Karl", "Denise"}
// 	nomes = append(nomes, "Ivone")
// 	fmt.Println(nomes)
// }
