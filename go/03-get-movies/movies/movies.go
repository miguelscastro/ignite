package movies

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Movie struct {
	ID      int64  `json:"id"`
	Nome    string `json:"nome"`
	Sinopse string `json:"sinopse"`
	Foto    string `json:"foto"`
}

func Search() ([]Movie, error) {
	// url.Values é um mapa (map[string][]string) especializado para Query Strings.
	// Ele é útil porque trata automaticamente caracteres especiais (espaços viram %20, etc).
	// var v url.Values
	// v.Set("apyKey", apiKey)       // Adiciona um par chave=valor
	// v.Add("tags", "action")       // Permite adicionar múltiplos valores para a mesma chave
	// http.Get("url" + v.Encode())  // v.Encode() transforma o mapa em string: "apyKey=123&tags=action"

	// CUIDADO: Você está usando http.NewRequest, que APENAS CRIA a requisição.
	// Para executá-la, você precisaria de um http.DefaultClient.Do(req).
	// Do jeito que está, 'resp' aqui é um objeto *http.Request, não a resposta do servidor.
	resp, err := http.NewRequest(http.MethodGet, "https://sujeitoprogramador.com/r-api/?api=filmes", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch movies: %w", err)
	}

	resp.Header.Set("Accept", "application/json")
	resp.Header.Set("User-Agent", "Go-Movie-App/1.0")

	// ATENÇÃO: Em um *http.Request (criado por NewRequest), o campo Body geralmente é nil
	// ou contém os dados que VOCÊ enviaria para o servidor.
	// Tentar dar Decode no Body de uma 'Request' não retornará os filmes da API.
	if resp.Body != nil {
		defer resp.Body.Close()
	}

	var movies []Movie
	// O Decode aqui provavelmente vai falhar ou travar, pois resp.Body não é a resposta da rede ainda.
	if err := json.NewDecoder(resp.Body).Decode(&movies); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}
	return movies, nil
}
