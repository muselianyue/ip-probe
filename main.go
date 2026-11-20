package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func fetch(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		return "none"
	}
	defer resp.Body.Close()
	b, _ := io.ReadAll(resp.Body)
	return string(b)
}

func handler(w http.ResponseWriter, r *http.Request) {
	host, _ := os.Hostname()
	fmt.Fprintf(w, "hostname: %s\negress IPv4: %s\negress IPv6: %s\n", host, fetch("https://api.ipify.org"), fetch("https://api64.ipify.org"))
}

func main() {
	http.HandleFunc("/", handler)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	http.ListenAndServe(":"+port, nil)
}
