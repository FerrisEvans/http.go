package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

func handler(w http.ResponseWriter, r *http.Request) {
	dump, err := httputil.DumpRequest(r, true)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println(string(dump))
	fmt.Fprintln(w, "<html><body>hello</body></html>")
}

func main() {
	http.HandleFunc("/", handler)
	log.Println("start http listening")
	var httpServer http.Server
	httpServer.Addr = ":8080"
	log.Println(httpServer.ListenAndServe())
}
