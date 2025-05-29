package util

import (
	"fmt"
	"log"
	"os"
)

var (
	infoLogger  *log.Logger
	errorLogger *log.Logger
)

func init() {
	// Configura o arquivo de log
	logFile, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal("Erro ao abrir arquivo de log:", err)
	}

	// Inicializa os loggers
	infoLogger = log.New(logFile, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	errorLogger = log.New(logFile, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

// LogInfo registra uma mensagem de informação
func LogInfo(message string) {
	infoLogger.Output(2, message)
}

// LogError registra uma mensagem de erro
func LogError(message string, err error) {
	errorLogger.Output(2, fmt.Sprintf("%s: %v", message, err))
}
