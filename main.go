package main

import "net/http"

func main() {
        panic(http.ListenAndServe(":5188", http.FileServer(http.Dir("."))))
}

