package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strconv"
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
	sites := readSitesFromFile()

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

	file, err := os.Open("log.txt")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func testSite(site string) {
	resp, _ := http.Get("http://" + site)
	var response string

	if resp.StatusCode == 200 {
		response = site + " foi carregado com sucesso."
	} else {
		response = site + " está com problemas. Status code: " + strconv.Itoa(resp.StatusCode)
	}
	writeLog(response)

	fmt.Println(response)
}

func readSitesFromFile() []string {
	var sites []string
	file, err := os.Open("sites.txt")
	if err != nil {
		fmt.Println("Error: ", err)
		return sites
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		sites = append(sites, scanner.Text())
	}

	return sites
}

func writeLog(response string) {
	file, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	file.WriteString(time.Now().Local().Format("02/01/2006 15:04:05") + ": " + response + "\n")
	defer file.Close()
}
