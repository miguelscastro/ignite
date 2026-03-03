package jsonutils

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/miguelscastro/ignite/go/05-gobid/internal/validator"
)

func EncodeJSON[T any](w http.ResponseWriter, r *http.Request, statusCode int, data T) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		return fmt.Errorf("failed to encode json: %w", err)
	}
	return nil
}

func DecodeValidJSON[T validator.Validator](r *http.Request) (T, map[string]string, error) {
	var data T
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		return data, nil, fmt.Errorf("failed to decode JSON: %w", err)
	}
	if problems := data.Valid(r.Context()); problems != nil {
		return data, problems, fmt.Errorf("invalid %T: %d problems", data, len(problems))
	}

	return data, nil, nil
}

func DecodeJSON[T any](r *http.Request) (T, error) {
	var data T
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		return data, fmt.Errorf("failed to decode JSON: %w", err)
	}

	return data, nil
}
