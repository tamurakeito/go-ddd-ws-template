package api_error

import "net/http"

func NewWebSocketHandshakeError() (code int, message map[string]string) {
	return http.StatusBadRequest, map[string]string{"error": "WebSocket handshake failed"}
}

func NewMessageHandlingError() (code int, message map[string]string) {
	return http.StatusInternalServerError, map[string]string{"error": "Failed to handle message"}
}
