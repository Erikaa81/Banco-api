package logger

import (
	"log"
	"os"
)

// Armazena as informações de logging
var (
	Info  = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	Error = log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
)
