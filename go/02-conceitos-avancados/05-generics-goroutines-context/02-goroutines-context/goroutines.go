package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func Goroutines() {
	start := time.Now()
	const n = 10
	var wg sync.WaitGroup

	// Adicionamos o número de "jobs" que o scheduler deve rastrear
	wg.Add(n)

	for i := range n {
		// A keyword 'go' dispara a função concorrentemente
		go func() {
			// Sinaliza ao WaitGroup que esta tarefa terminou ao sair da função
			defer wg.Done()

			response, err := http.Get("https://google.com")
			if err != nil {
				// Em produção, evite panic em goroutines; trate o erro via channels
				fmt.Printf("Erro na goroutine %d: %v\n", i, err)
				return
			}

			// Fecha o Body para liberar recursos e permitir reuso da conexão TCP
			defer response.Body.Close()
			fmt.Printf("Goroutine %d: OK\n", i)
		}()
	}

	// Bloqueia até que todas as n tarefas chamem wg.Done()
	wg.Wait()

	fmt.Printf("Tempo total: %v\n", time.Since(start))
}
