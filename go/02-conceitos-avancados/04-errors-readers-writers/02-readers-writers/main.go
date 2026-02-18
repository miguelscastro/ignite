package main

import (
	"errors"
	"io"
	"os"
	"strings"
)

// --- 1. WRITER CUSTOMIZADO (WRAPPING) ---
// O wrap de writers permite interceptar e transformar dados antes de passá-los adiante.
type MyWriter struct {
	w io.Writer // Recebe outro Writer (ex: os.Stdout)
}

func (mw MyWriter) Write(b []byte) (int, error) {
	// Exemplo de transformação: adiciona 10 a cada byte antes de escrever.
	for i, bb := range b {
		b[i] = bb + 10
	}
	return mw.w.Write(b) // Passa os dados transformados para o writer interno.
}

func main() {
	str := "hello, world\n"
	reader := strings.NewReader(str)

	// MyWriter embrulha o os.Stdout (que implementa io.Writer).
	writer := MyWriter{w: os.Stdout}

	// O buffer (pote) define quanto lemos por vez.
	// Ler em partes evita estourar a memória com arquivos gigantes.
	buffer := make([]byte, 2)

	// --- 2. O LOOP DE LEITURA (A "DANÇA" DO READER) ---
	for {
		n, err := reader.Read(buffer)
		if err != nil {
			// io.EOF (End Of File) é um erro sentinela que indica o fim da stream.
			if errors.Is(err, io.EOF) {
				break // Fim da leitura com sucesso.
			}
			panic(err) // Erros reais (ex: falha de disco) interrompem o programa.
		}

		// Importante: use buffer[:n] para pegar apenas os bytes preenchidos nesta volta.
		_, _ = writer.Write(buffer[:n])
	}

	// --- 3. MÉTODOS UTILITÁRIOS ---
	// io.ReadAll: Lê tudo de uma vez. CUIDADO: Pode esgotar a RAM em streams grandes.
	// r2 := strings.NewReader("outra stream")
	// data, _ := io.ReadAll(r2)

	// io.ReadFull: Preenche exatamente o tamanho do buffer passado, sem precisar de loop.
	// bufFull := make([]byte, 5)
	// io.ReadFull(reader, bufFull)
}
