package web

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

type errorResponse struct {
	Code  int    `json:"code,omitempty"`
	Error string `json:"error,omitempty"`
}

func SendUnauthorized(w http.ResponseWriter) {
	sendPayload(w, 401, errorResponse{
		Code:  401,
		Error: "401 unauthorized",
	})
}

func SendBadRequest(w http.ResponseWriter, err error) {
	sendPayload(w, 401, errorResponse{
		Code:  400,
		Error: err.Error(),
	})
}

func SendPayload(w http.ResponseWriter, payload any) {
	sendPayload(w, 200, payload)
}

func sendPayload(w http.ResponseWriter, code int, payload any) {
	w.Header().Add("Content-Type", "application/json")
	b, err := json.Marshal(payload)
	if err != nil {
		slog.Error("unable to write response",
			slog.Any("http_code", code),
			slog.Any("payload", payload),
			slog.Any("error", err),
		)
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	_, err = w.Write(b)
	if err != nil {
		slog.Error("unable to write response",
			slog.Any("error", err),
		)
	}

}
