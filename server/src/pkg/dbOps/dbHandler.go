package dbops

import (
	"context"
	"time"

	c "github.com/FKouhai/OpenCaaServer/src/pkg/config"
	l "github.com/FKouhai/OpenCaaServer/src/pkg/logger"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.uber.org/zap"
)

type NewClient struct {
}

func DbWrite(ctx context.Context, kv clientv3.KV, key string, value string) {
	newLog := l.NewLogger()
	pr, err := kv.Put(ctx, key, value)

	if err != nil {
		logHelper(newLog, err)
	}
}

func logHelper(loggin *zap.SugaredLogger, err error) {
	l.LoggErr(loggin, err)
}
func DbRead(ctx context.Context, kv clientv3.KV, key string) {
	newLog := l.NewLogger()
	pr, err := kv.Get(ctx, key)
	if err != nil {
		logHelper(newLog, err)
	}
}

func DbConn(ctx context.Context) clientv3.kv {
	newLog := l.NewLogger()
	config, err := c.NewConfig()

	if err != nil {
		l.LoggErr(newLog, err)
		return nil
	}
	conn, err := clientv3.New(clientv3.Config{
		DialTimeout: 120 * time.Second,
		Endpoints:   []string{config.Etcd},
	})
	if err != nil {
		logHelper(newLog, err)
	}
	defer conn.Close()
  return conn
}

func DbRm(ctx context.Context, kv clientv3.KV, key string) {
	newLog := l.NewLogger()
	pr, err := kv.Delete(ctx, key)
	if err != nil {
		logHelper(newLog, err)
	}
}
