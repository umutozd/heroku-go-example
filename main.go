package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/hello", handleHello)
	mux.HandleFunc("/", handleIndex)

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("ListenAndServe returned error: %v", err)
	}
}

const indexPageHtml = `
<!DOCTYPE html>
<html>
<head>
    <meta charset='utf-8'>
    <meta http-equiv='X-UA-Compatible' content='IE=edge'>
    <title>Hello World</title>
    <meta name='viewport' content='width=device-width, initial-scale=1'>
</head>
<body style="width: 100%; display: flex; flex-direction: column; align-items: center;">
    <h1>Hello World!</h1>
</body>
</html>
`

func handleIndex(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(indexPageHtml))
}

func handleHello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	data := map[string]string{
		"message": "Hello World",
	}
	b, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Printf("error marshaling api response: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "unable to write api response"}`))
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write(b)
	}
}
