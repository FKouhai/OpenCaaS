package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/FKouhai/OpenCaaServer/src/pkg/dbOps"
	l "github.com/FKouhai/OpenCaaServer/src/pkg/logger"
)
type Example struct {
  Key string `json:"key"`
  Value string `json:"value"`
}
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Working!")
}
func Add(w http.ResponseWriter, r *http.Request) {
  decoder := json.NewDecoder(r.Body)
  var example Example
  err := decoder.Decode(&example)
  newLog := l.NewLogger()
  if err != nil {
    l.LoggErr(newLog, err)
  }
  ctx := context.TODO()
  kv := dbops.DbConn(ctx)
  dbops.DbWrite(ctx, kv, example.Key, example.Value )
  if err != nil {
    l.LoggErr(newLog, err)
    w.WriteHeader(500)
  }
  w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(example)
}

