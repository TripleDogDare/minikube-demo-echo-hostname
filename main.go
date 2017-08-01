package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

const DEFAULT_PORT = "80"

func main() {
	port := flag.String("port", DEFAULT_PORT, "Set server port")
	flag.Parse()
	host := fmt.Sprintf(":%s", *port)
	log.Fatal(http.ListenAndServe(host, http.HandlerFunc(handler)))
}

func handler(rw http.ResponseWriter, req *http.Request) {
	name, err := os.Hostname()
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(err.Error()))
	} else {
		rw.Write([]byte(name))
	}
}
