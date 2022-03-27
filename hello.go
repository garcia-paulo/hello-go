package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {

	showIntro()

	for {
		showMenu()

		opcao := readOpcao()

		switch opcao {
		case 1:
			startMonitoring()
		case 2:
			showLogs()
		case 0:
			fmt.Println("Saindo do programa.")
			os.Exit(0)
		default:
			fmt.Println("Comando não reconhecido.")
			os.Exit(-1)
		}
	}
}

func showIntro() {
	nome := "Paulo"
	version := "1.0"
	fmt.Println("Olá", nome)
	fmt.Println("Este programa está na versão: ", version)
}

func showMenu() {

	fmt.Println()

	fmt.Println("1- Iniciar monitoramento")
	fmt.Println("2- Exibir logs")
	fmt.Println("0- Sair do programa")
}

func readOpcao() int {
	var opcao int
	fmt.Scan(&opcao)
	fmt.Println("O comando escolhido foi: ", opcao)
	return opcao
}

func startMonitoring() {
	fmt.Println("Monitorando...")
	sites := []string{"https://www.google.com.br", "https://www.youtube.com"}

	for range [5]int{} {
		for _, site := range sites {
			testSite(site)
		}
		println()
		time.Sleep(time.Second * 5)
	}
}

func showLogs() {
	fmt.Println("Exibindo logs...")
}

func testSite(site string) {
	resp, _ := http.Get(site)
	if resp.StatusCode == 200 {
		fmt.Println(site, "foi carregado com sucesso.")
	} else {
		fmt.Println(site, "está com problemas. Status code: ", resp.StatusCode)
	}
}
