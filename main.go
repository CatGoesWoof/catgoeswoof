package main

import "net/http"

func init() {
  log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir("."))))
}
