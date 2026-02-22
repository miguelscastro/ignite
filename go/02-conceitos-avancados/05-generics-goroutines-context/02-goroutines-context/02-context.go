package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"sync"
	"time"
)

func Context() {
	start := time.Now()
	const n = 10
	var wg sync.WaitGroup
	wg.Add(n)

	// 1. Criamos o contexto raiz
	ctx := context.Background()

	// 2. Criamos um contexto com timeout. Se passar de 5s, as requisições param
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel() // Libera os recursos associados ao contexto

	// 3. Simulamos um servidor lento que demora 10s para responder
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(10 * time.Second)
		fmt.Fprintln(w, "hello, world")
	}))
	defer server.Close()

	for i := 0; i < n; i++ {
		go func(ctx context.Context, id int) {
			defer wg.Done()

			// Criamos a requisição amarrada ao nosso contexto
			req, err := http.NewRequestWithContext(ctx, "GET", server.URL, nil)
			if err != nil {
				panic(err)
			}

			// O cliente HTTP respeitará o sinal de cancelamento do contexto
			response, err := http.DefaultClient.Do(req)
			if err != nil {
				// Verificamos se o erro foi especificamente por causa do timeout do contexto
				if errors.Is(err, context.DeadlineExceeded) {
					fmt.Printf("Goroutine %d: timeout excedido\n", id)
					return
				}
				fmt.Printf("Goroutine %d: erro na requisição: %v\n", id, err)
				return
			}
			defer response.Body.Close()
			fmt.Printf("Goroutine %d: resposta recebida\n", id)
		}(ctx, i)
	}

	wg.Wait()
	fmt.Printf("Execução finalizada em: %v\n", time.Since(start))
}
