package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func handler(hostname string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		log.Printf("%+#v\n", req)
		fmt.Fprintf(w, "Hello from %s\n", hostname)
	}
}

func main() {
	hostname, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", handler(hostname))

	if err := http.ListenAndServe(":80", nil); err != nil {
		log.Fatalln(err)
	}
}
