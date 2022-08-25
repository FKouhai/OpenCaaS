package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	c "github.com/FKouhai/OpenCaaServer/src/pkg/config"
	"github.com/FKouhai/OpenCaaServer/src/pkg/dbOps"
	l "github.com/FKouhai/OpenCaaServer/src/pkg/logger"
	clientv3 "go.etcd.io/etcd/client/v3"
	"net/http"
	"time"
)

type Example struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

var config, _ = c.NewConfig()
var conn, _ = clientv3.New(clientv3.Config{DialTimeout: 120 * time.Second, Endpoints: []string{config.Etcd}})
var kv = clientv3.NewKV(conn)

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
	dbops.DbWrite(ctx, kv, example.Key, example.Value)
	if err != nil {
		l.LoggErr(newLog, err)
		w.WriteHeader(500)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(example)
}

func Remove(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var example Example
	err := decoder.Decode(&example)
	newLog := l.NewLogger()
	if err != nil {
		l.LoggErr(newLog, err)
	}
	ctx := context.TODO()
	dbops.DbRm(ctx, kv, example.Key)
	if err != nil {
		l.LoggErr(newLog, err)
		w.WriteHeader(500)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "Key $s removed succesfully", string(example.Key))

}

func List(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var example Example
	err := decoder.Decode(&example)
	newLog := l.NewLogger()
	if err != nil {
		l.LoggErr(newLog, err)
	}
	ctx := context.TODO()
	values := dbops.DbRead(ctx, kv, example.Key)
	if err != nil {
		l.LoggErr(newLog, err)
		w.WriteHeader(500)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(example)
	fmt.Fprintf(w, "key vlaue %s", values)

}
