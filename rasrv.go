package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func client(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[+] client connected")
	forwarded := r.Header.Get("x-forwarded-for")
	client := strings.Split(forwarded, ",")[0]
	w.Write([]byte(r.RemoteAddr))
	w.Write([]byte("\n"))
	w.Write([]byte(client + "\n"))
}

func main() {
	var port string
	if port = os.Getenv("PORT"); port == "" {
		port = "8080"
	}
	addr := ":" + port
	http.HandleFunc("/", client)
	fmt.Printf("[+] listening on http://127.0.0.1:%s\n", port)
	log.Fatal(http.ListenAndServe(addr, nil))
}

