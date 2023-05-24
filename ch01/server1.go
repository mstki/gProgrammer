//Server1 is a minimal "echo" server.
package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var (
	mutex sync.Mutex
	count int
)

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, req *http.Request) {
	mutex.Lock()
	count++
	mutex.Unlock()
	fmt.Fprintf(w, "URL PATH = %q\n", req.URL.Path)
}

func counter(w http.ResponseWriter, req *http.Request) {
	mutex.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mutex.Unlock()
}
