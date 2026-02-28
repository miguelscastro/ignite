package api

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

func sendJSON(w http.ResponseWriter, resp APIResponse, status int) {
	data, err := json.Marshal(resp)
	if err != nil {
		slog.Error("error while doing json marshall", "error", err)
		sendJSON(w, APIResponse{Error: "something went wrong"}, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(status)
	if _, err := w.Write(data); err != nil {
		slog.Error("error while sending response", "error", err)
		return
	}
}
