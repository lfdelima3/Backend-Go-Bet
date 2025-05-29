package util

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Error   string `json:"error,omitempty"`
}

func NewErrorResponse(status int, message string, err error) ErrorResponse {
	errorResponse := ErrorResponse{
		Status:  status,
		Message: message,
	}
	if err != nil {
		errorResponse.Error = err.Error()
	}
	return errorResponse
}

func RespondWithError(w http.ResponseWriter, status int, message string, err error) {
	errorResponse := NewErrorResponse(status, message, err)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(errorResponse)
}

// Erros comuns
var (
	ErrInvalidInput       = "entrada inválida"
	ErrNotFound           = "recurso não encontrado"
	ErrUnauthorized       = "não autorizado"
	ErrForbidden          = "acesso proibido"
	ErrInternalServer     = "erro interno do servidor"
	ErrDatabase           = "erro no banco de dados"
	ErrInvalidToken       = "token inválido"
	ErrExpiredToken       = "token expirado"
	ErrInvalidCredentials = "credenciais inválidas"
)
