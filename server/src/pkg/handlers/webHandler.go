package handler

import (
  "fmt"
  "log"
  "net/http"
  "strconv"
  "github.com/gorilla/mux"
)


func Index(w http.ResponseWriter, r *http.Request) {
fmt.Fprintf(w,"Working!")
}
