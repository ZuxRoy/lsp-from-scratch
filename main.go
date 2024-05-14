package main

import (
	"log"
	"os"
)

func main () {
    logger := getLogger("/home/zuxroy/Code/Projects/lsp-from-scratch/logs.txt")
    logger.Println("LSP logging initiated")
}

func getLogger(filename string) *log.Logger {
    logfile, err := os.OpenFile(filename, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)

    if err != nil {
        panic("Error with the log file")
    }

    return log.New(logfile, "[START] ", log.Ldate|log.Ltime|log.Llongfile)
}
