package main

import (
  "log"
  "net/http"
)

func Log(handler http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
    handler.ServeHTTP(w, r)
  })
}
