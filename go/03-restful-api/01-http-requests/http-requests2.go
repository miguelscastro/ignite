package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

func HTTPRequests2() {
	// 1. Configuração do Timeout via Contexto
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 2. Criação da Requisição Customizada
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://google.com", nil)
	if err != nil {
		panic(err)
	}

	// 3. Configuração de Headers
	req.Header.Set("authorization", "456")   // Sobrescrito de "123" para "456"
	req.Header.Add("authorization", "extra") // Agora possui ["456", "extra"]
	req.Header.Set("accept", "application/json")

	// 4. Execução e Gerenciamento do Body
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	// IMPORTANTE: O fechamento garante o reuso da conexão TCP (Keep-Alive)
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Status: %d, Bytes lidos: %d\n", resp.StatusCode, len(data))
}
