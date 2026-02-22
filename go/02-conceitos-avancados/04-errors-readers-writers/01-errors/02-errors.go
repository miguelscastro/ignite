package main

import (
	"errors"
	"fmt"
	"math"
)

// --- 1. ERROS PERSONALIZADOS (CUSTOM ERRORS) ---
// Útil quando precisamos armazenar metadados além da mensagem.
type SqrtError struct {
	msg string
}

// Implementando interface error. Se usar Pointer Receiver (*SqrtError),
// apenas o ponteiro satisfará a interface.
func (s *SqrtError) Error() string { return s.msg }

func raizQuadrada(x float64) (float64, error) {
	if x < 0 {
		// Retornamos um ponteiro porque o receiver de Error() é um ponteiro.
		return 0, &SqrtError{msg: "Não existe raíz quadrada de número negativo"}
	}
	return math.Sqrt(x), nil
}

// --- 2. SENTINEL ERRORS E WRAPPING ---
var (
	ErrQualquer = errors.New("erro base")
	ErrNotFound = errors.New("not found")
)

func bar() error {
	return ErrQualquer
}

func fooWrapper() error {
	err := bar()
	if err != nil {
		// %w cria um WRAPPER: o erro original é "embrulhado" e não se perde.
		return fmt.Errorf("falha em foo: %w", err)
	}
	return nil
}

// --- 3. MULTI-ERRORS (JOIN) ---
var (
	ErrValidacao = errors.New("erro de validação")
	ErrBanco     = errors.New("erro de banco")
)

func executaMultiplasTarefas() error {
	var errs error
	// errors.Join agrupa erros mantendo a compatibilidade com Is/As.
	errs = errors.Join(errs, ErrValidacao)
	errs = errors.Join(errs, ErrBanco)
	return errs
}

func Errors2() {
	// --- TESTE AS (Para tipos customizados) ---
	_, errCustom := raizQuadrada(-1)
	var sErr *SqrtError
	// errors.As tenta converter o erro para o tipo específico.
	// Deve-se passar o ponteiro da variável destino.
	if errors.As(errCustom, &sErr) {
		fmt.Println("Capturado via As:", sErr.msg)
	}

	// --- TESTE IS (Para sentinel errors e wrappers) ---
	errW := fooWrapper()
	// errors.Is navega pela árvore de wraps até achar a causa raiz.
	if errors.Is(errW, ErrQualquer) {
		fmt.Println("Capturado via Is: A causa raiz é ErrQualquer")
	}

	// --- TESTE JOIN ---
	errJ := executaMultiplasTarefas()
	fmt.Println("Erros unidos:\n", errJ)
	fmt.Printf("Contém validação? %v\n", errors.Is(errJ, ErrValidacao))
}
