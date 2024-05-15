package main

import (
	"bufio"
	"encoding/json"
	"log"
	"lsp-from-scratch/analysis"
	"lsp-from-scratch/lsp"
	"lsp-from-scratch/rpc"
	"os"
)

func main () {
    logger := getLogger("/home/zuxroy/Code/Projects/lsp-from-scratch/logs.txt")
    logger.Println("LSP logging initiated")

    scanner := bufio.NewScanner(os.Stdin)
    scanner.Split(rpc.Split)
    
    state := analysis.NewState()

    for scanner.Scan() {
        mssg := scanner.Bytes()
        method, content, err := rpc.DecodeMessage(mssg)
        if err != nil {
            logger.Printf("Error: %s", err)
            continue
        }
        handleMessage(logger, state, method, content)
    }
}

func handleMessage(logger *log.Logger, state analysis.State, method string, content []byte) {
    logger.Printf("Received method: %s", method)

    switch method {
    case "initialize":
        var request lsp.InitializeRequest
        if err := json.Unmarshal(content, &request); err != nil {
            logger.Printf("Error: Unable to parse: %s", err)
        }
        logger.Printf("Connected to: %s %s", 
            request.Params.ClientInfo.Name,
            request.Params.ClientInfo.Version)

        mssg := lsp.NewInitializeResponse(request.ID)
        reply := rpc.EncodeMessage(mssg)

        writer := os.Stdout
        writer.Write([]byte(reply))

        logger.Printf("Response sent")
    case "textDocument/didOpen":
        var request lsp.DidOpenTextDocumentNotification
        if err := json.Unmarshal(content, &request); err != nil {
            logger.Printf("Error: textDocument/didOpen: %s", err)
            return
        }
        logger.Printf("Opened: %s", request.Params.TextDocument.URI)
        state.OpenDocument(request.Params.TextDocument.URI, request.Params.TextDocument.Text)
    case "textDocument/didChange":
        var request lsp.TextDocumentDidChangeNotification
        if err := json.Unmarshal(content, &request); err != nil {
            logger.Printf("Error: textDocument/didChange: %s", err)
            return
        }
        logger.Printf("Changed: %s", request.Params.TextDocument.URI)
        for _, change := range request.Params.ContentChanges {
            state.UpdateDocument(request.Params.TextDocument.URI, change.Text)
        }
    }
}

func getLogger(filename string) *log.Logger {
    logfile, err := os.OpenFile(filename, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)

    if err != nil {
        panic("Error with the log file")
    }

    return log.New(logfile, "[lsp-from-scratch] ", log.Ldate|log.Ltime|log.Llongfile)
}
