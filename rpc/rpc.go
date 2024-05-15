package rpc

import (
    "fmt"
    "encoding/json"
    "bytes"
    "errors"
    "strconv"
)

type BaseMessage struct {
    Method string `json:"method"`
}

func EncodeMessage(mssg any) string {
    content, err := json.Marshal(mssg)
    if err != nil {
        panic(err)
    }
    return fmt.Sprintf("Content-Length: %d\r\n\r\n%s", len(content), content)
}

func DecodeMessage(mssg []byte) (string, []byte, error){
    header, content, found := bytes.Cut(mssg, []byte{'\r', '\n', '\r', '\n'})
    if !found {
        return "", nil, errors.New("Error: Separator not found")
    }

    contentLengthBytes := header[len("Content-Length: "):]
    contentLength, err := strconv.Atoi(string(contentLengthBytes))
    if err != nil {
        return "", nil, err
    }

    var baseMssg BaseMessage 
    if err := json.Unmarshal(content[:contentLength], &baseMssg); err != nil { 
        return "", nil, err
    }

    return baseMssg.Method, content[:contentLength], nil
}

func Split(data []byte, _ bool) (advance int, token []byte, err error){
    header, content, found := bytes.Cut(data, []byte{'\r', '\n', '\r', '\n'})
    if !found {
        return 0, nil, nil
    }

    contentLengthBytes := header[len("Content-Length: "):]
    contentLength, err := strconv.Atoi(string(contentLengthBytes))
    if err != nil {
        return 0, nil, err
    }
    
    if len(content) < contentLength {
        return 0, nil, nil
    }

    totalLength := len(header) + 4 + contentLength
    return totalLength, data[:totalLength], nil
}
