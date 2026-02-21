package main

import (
	"fmt"
	"io"
	"net/http"
)

func HTTPRequests1() {
	// --- EXEMPLO: HTTP POST ---
	// Usado para enviar dados ao servidor.
	// Parâmetros: URL, Content-Type e o Body (io.Reader).
	respPost, err := http.Post("https://httpbin.org/post", "application/json", nil)
	if err != nil {
		panic(err)
	}
	// O corpo deve ser fechado assim que o erro for verificado.
	defer respPost.Body.Close()

	dataPost, _ := io.ReadAll(respPost.Body)
	fmt.Println("POST Status:", respPost.Status)
	fmt.Println("POST Body:", string(dataPost))

	// --- EXEMPLO: HTTP GET ---
	// Usado para buscar dados.
	// Note o uso de UMA NOVA VARIÁVEL (respGet) para evitar o bug do defer.
	respGet, err := http.Get("https://google.com")
	if err != nil {
		panic(err)
	}
	defer respGet.Body.Close() // Agora este defer aponta para a conexão correta.

	dataGet, _ := io.ReadAll(respGet.Body)
	fmt.Println("GET Status:", respGet.Status)
	fmt.Printf("GET Body: %d bytes lidos\n", len(dataGet))
}
