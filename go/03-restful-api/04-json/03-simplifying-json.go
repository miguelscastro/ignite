package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Error string `json:"error,omitempty"` // só faz diferença durante o Marshall do JSON, se seu valor for o do padrão do seu tipo ele não é incluido no JSON
	Data  any    `json:"data,omitempty"`
}

func sendJSON(w http.ResponseWriter, resp Response, status int) {
	data, err := json.Marshal(resp)
	if err != nil {
		fmt.Println("error while doing json marshall: ", err)
		sendJSON(w, Response{Error: "something went wrong"}, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(status)
	if _, err := w.Write(data); err != nil {
		fmt.Println("error while sending response")
		return
	}
}
