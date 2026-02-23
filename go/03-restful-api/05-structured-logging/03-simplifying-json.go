package main

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

type Response struct {
	Error string `json:"error,omitempty"`
	Data  any    `json:"data,omitempty"`
}

func sendJSON(w http.ResponseWriter, resp Response, status int) {
	data, err := json.Marshal(resp)
	if err != nil {
		// recebe a msg a ser exibida, e um chave:valor que identifica o conteudo so log e seu stack trace
		slog.Error("error while doing json marshall", "error", err)
		sendJSON(w, Response{Error: "something went wrong"}, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(status)
	if _, err := w.Write(data); err != nil {
		slog.Error("error while sending response", "error", err)
		return
	}
}
