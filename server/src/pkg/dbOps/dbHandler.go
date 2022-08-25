package dbops

import (
	"context"
	"fmt"
	l "github.com/FKouhai/OpenCaaServer/src/pkg/logger"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.uber.org/zap"
)

func DbWrite(ctx context.Context, kv clientv3.KV, key string, value string) {
	newLog := l.NewLogger()
	_, err := kv.Put(ctx, key, value)

	if err != nil {
		logHelper(newLog, err)
	}
}

func logHelper(loggin *zap.SugaredLogger, err error) {
	l.LoggErr(loggin, err)
}
func DbRead(ctx context.Context, kv clientv3.KV, key string) clientv3.GetResponse {
	newLog := l.NewLogger()
	pr, err := kv.Get(ctx, key)
	fmt.Println(pr)
	if err != nil {
		logHelper(newLog, err)
	}
	return *pr
}

func DbRm(ctx context.Context, kv clientv3.KV, key string) *clientv3.DeleteResponse {
	newLog := l.NewLogger()
	pr, err := kv.Delete(ctx, key)
	if err != nil {
		logHelper(newLog, err)
	}
	return pr
}
