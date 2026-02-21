package main

import (
	"errors"
	"fmt"
	"net/http"
	"time"
)

func HTTPServers() {
	mux := http.NewServeMux()

	// Multiplexer: associa rotas a funções (handlers).
	mux.HandleFunc("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "hello, world")
	})

	srv := &http.Server{
		// --- CONFIGURAÇÕES BÁSICAS ---
		// Endereço e porta. Ex: ":8080" escuta em todos endereços.
		Addr: ":8080",

		// O roteador que define como as chamadas são distribuídas.
		Handler: mux,

		// Se true, o servidor ignora o tratamento automático de requests OPTIONS *.
		DisableGeneralOptionsHandler: false,

		// --- SEGURANÇA E TIMEOUTS (ESSENCIAIS) ---
		// Tempo máximo para ler a requisição inteira, incluindo o corpo.
		ReadTimeout: 10 * time.Second,

		// Tempo máximo apenas para ler os cabeçalhos (ajuda contra ataques Slowloris).
		ReadHeaderTimeout: 5 * time.Second,

		// Tempo máximo para o servidor escrever a resposta de volta ao cliente.
		WriteTimeout: 10 * time.Second,

		// Tempo que uma conexão ociosa (keep-alive) espera por um novo pedido.
		IdleTimeout: 1 * time.Minute,

		// Limite de bytes nos headers. 0 usa o padrão (geralmente 1MB).
		MaxHeaderBytes: 1 << 20, // 1048576 bytes (1MB)

		// --- HTTPS E PROTOCOLOS ---
		// Configurações de certificados TLS para habilitar HTTPS.
		// TLSConfig: &tls.Config{
		// 	MinVersion: tls.VersionTLS12, // Exemplo de configuração de segurança
		// },

		// Mapeia protocolos para funções de upgrade (usado internamente para HTTP/2).
		// TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler)),

		// Configurações específicas para o protocolo HTTP/2.
		// HTTP2: &http.HTTP2Config{},

		// Define quais protocolos o servidor deve suportar (HTTP/1.1, HTTP/2, etc).
		// Protocols: &http.Protocols{},

		// --- OBSERVAÇÃO E ESTADO ---
		// Callback chamada sempre que uma conexão TCP muda de estado (New, Active, Idle, Closed).
		// ConnState: func(c net.Conn, state http.ConnState) {
		// 	// Útil para monitoramento de conexões ativas.
		// },

		// Logger para erros internos do servidor.
		// CORREÇÃO: Usar nil para o padrão ou log.New para um destino válido.
		// ErrorLog: log.New(os.Stderr, "SERVER_ERROR: ", log.LstdFlags),

		// --- CONTEXTOS (GERENCIAMENTO DE VIDA) ---
		// Define o contexto raiz (pai) de todos os requests. Recebe o Listener atual.
		// BaseContext: func(l net.Listener) context.Context {
		// 	return context.Background()
		// },

		// Permite modificar o contexto de cada conexão individualmente.
		// ConnContext: func(ctx context.Context, c net.Conn) context.Context {
		// 	return ctx
		// },
	}

	fmt.Println("Servidor iniciado em http://localhost:8080")

	// ListenAndServe: Método bloqueante que inicia o servidor.
	if err := srv.ListenAndServe(); err != nil {
		// ErrServerClosed é retornado quando paramos o servidor via código.
		if !errors.Is(err, http.ErrServerClosed) {
			panic(fmt.Sprintf("Erro fatal: %v", err))
		}
	}
}
