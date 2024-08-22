package main

import (
    "fmt"
    "net/http"
    "time"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
      fmt.Fprint(w, "Hello, World! I am Transcriber")
    })

    http.ListenAndServe(":2222", nil)    
}

