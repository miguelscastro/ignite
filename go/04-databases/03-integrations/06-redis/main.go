package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func main() {
	ExampleClient()
}

// O Redis em Go exige o uso de context para gerenciar cancelamentos e timeouts
var ctx = context.Background()

func ExampleClient() {
	// Inicialização do cliente. Addr é o endereço do container Redis (padrão: 6379)
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // Senha definida no docker-compose (se houver)
		DB:       0,  // O Redis possui 16 bancos lógicos (0 a 15), o 0 é o padrão
	})
	// Fechar a conexão ao final da execução
	defer rdb.Close()

	// SET: Salva uma chave e valor. O último parâmetro (0) é o TTL (Time To Live).
	// 0 significa que o dado nunca expira automaticamente
	err := rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	// GET: Recupera o valor da chave. .Result() retorna o valor e o erro simultaneamente
	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	// Tratamento de "Cache Miss" (Chave não encontrada):
	val2, err := rdb.Get(ctx, "key2").Result()
	if err == redis.Nil {
		// No Redis, uma chave inexistente retorna este erro específico 'redis.Nil'
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
}
