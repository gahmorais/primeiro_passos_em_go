package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

const COUNT_MONITORING = 5
const DELAY = 5 * time.Second

func main() {
	exibeSaudacoes()
	exibeOpcoes()
	// leSitesDoArquivo()
	comando := leComando()

	switch comando {
	case 1:
		iniciarMonitoramento()
	case 2:
		fmt.Println("Exibir logs")
	case 0:
		fmt.Println("Sair do programa")
		os.Exit(0)
	default:
		fmt.Println("Comando não reconhecido!")
		os.Exit(-1)
	}

}

func exibeSaudacoes() {
	nome := "Gabriel"
	versao := 1.1
	fmt.Println("Olá, sr.", nome)
	fmt.Println("A versão do programa é:", versao)
}

func exibeOpcoes() {
	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Exibir logs")
	fmt.Println("0 - Sair do programa")
}

func leComando() int {
	var comando int
	fmt.Scanf("%d", &comando)
	fmt.Println("O comando escolhido é", comando)
	return comando
}

func iniciarMonitoramento() {
	fmt.Println("Iniciando monitoramento")
	sites := leSitesDoArquivo()
	fmt.Println(sites)

	for i := 0; i < COUNT_MONITORING; i++ {
		for _, site := range sites {
			fmt.Println("Testando site: ", site)
			testaSite(site)
		}
		fmt.Println("Aguardando próximo teste...")
		time.Sleep(DELAY)
	}

}

func testaSite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site: ", site, "foi carregado com sucesso")
	} else {
		fmt.Println("Site: ", site, "esta com problemas. Status Code: ", resp.StatusCode)
	}
}

func leSitesDoArquivo() []string {
	arquivo, err := os.Open("sites.txt")
	var sites []string
	if err != nil {
		fmt.Println("Ocorreu um erro", err)
	}

	leitor := bufio.NewReader(arquivo)
	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)
		sites = append(sites, linha)
		if err == io.EOF {
			break
		}
		fmt.Println(linha)
	}

	arquivo.Close()

	return sites
}
