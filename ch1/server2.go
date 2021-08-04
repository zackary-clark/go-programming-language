package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mutualExclusionLock sync.Mutex
var count int

func main() {
	http.HandleFunc("/", handle2)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handle2(w http.ResponseWriter, r *http.Request) {
	mutualExclusionLock.Lock()
	count++
	mutualExclusionLock.Unlock()
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

func counter(w http.ResponseWriter, r *http.Request) {
	mutualExclusionLock.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mutualExclusionLock.Unlock()
}
