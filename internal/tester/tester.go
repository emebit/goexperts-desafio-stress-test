package tester

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

/*
===================================================================
  - Função: Run
  - Descrição : Função que roda o teste de carga
  - Parametros :
  - url - url a ser requisitada - tipo string
  - requests - Qtde de requisiçoes a serem executadas - tipo: int
  - concurrency - Qtde de requisiçoes a serem executadas
  - simultaneamente - tipo: int
  - Retorno:

===================================================================
*/
func Run(url string, requests int, concurrency int) {
	wChan := make(chan struct{}, requests)
	rChan := make(chan int, requests)

	var wg sync.WaitGroup
	startTime := time.Now()

	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go worker(&wg, url, wChan, rChan)
	}

	for i := 0; i < requests; i++ {
		wChan <- struct{}{}
	}
	close(wChan)

	wg.Wait()
	close(rChan)
	totalTime := time.Since(startTime)

	Report(rChan, totalTime)
}

/*
===================================================================
  - Função: worker
  - Descrição : Função que executa as chamadas HTTP do teste de carga
  - Parametros :
  - wg - ponteiro de WaitGroup - tipo: sync.WaitGroup
  - url - url a ser requisitada - tipo string
  - wChan - Canal de trabalho - tipo: chan struct{}
  - rChan - Canal de resultado - tipo: chan int
  - Retorno:

===================================================================
*/
func worker(wg *sync.WaitGroup, url string, wChan chan struct{}, rChan chan int) {
	defer wg.Done()
	for range wChan {
		resp, err := http.Get(url)
		if err != nil {
			rChan <- 0
			continue
		}
		rChan <- resp.StatusCode
	}
}

/*
===================================================================
  - Função: Report
  - Descrição : Função que exibe relatório do teste de carga
  - Parametros :
  - results - Canal de resultado - tipo: chan int
  - totalTime - Tempo de execução do teste - tipo time.Duration
  - Retorno:

===================================================================
*/
func Report(results chan int, totalTime time.Duration) {
	totalRequests := 0
	statusCount := make(map[int]int)

	for status := range results {
		totalRequests++
		statusCount[status]++
	}

	fmt.Printf("Relatório do Teste Carga\n")
	fmt.Printf("Tempo total: %v\n", totalTime)
	fmt.Printf("Total de requests: %d\n", totalRequests)
	fmt.Printf("Requests com status 200: %d\n", statusCount[200])
	fmt.Println("Distribuição de outros códigos de status HTTP:")
	for status, count := range statusCount {
		if status != 200 {
			fmt.Printf("Requests com status %d: %d\n", status, count)
		}
	}
}
