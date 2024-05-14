package rpc_test

import (
    "lsp-from-scratch/rpc"
    "testing"
)

type EncodingExample struct {
    Testing bool
}

func TestEncoding(t *testing.T) {
    expected := "Content-Length: 16\r\n\r\n{\"Testing\":true}"
    actual := rpc.EncodeMessage(EncodingExample{Testing: true})
    if expected != actual {
        t.Fatalf("Error:\r\nExpected: %s\r\nGot: %s", expected, actual)
    }
}

func TestDecode(t *testing.T) {
	incomingMessage := "Content-Length: 15\r\n\r\n{\"Method\":\"15\"}"
	method, content, err := rpc.DecodeMessage([]byte(incomingMessage))
	contentLength := len(content)

	if err != nil {
		t.Fatal(err)
	}

	if contentLength != 15 {
		t.Fatalf("Expected: 15, Got: %d", contentLength)
	}

	if method != "15" {
		t.Fatalf("Expected: '15', Got: %s", method)
	}
}
