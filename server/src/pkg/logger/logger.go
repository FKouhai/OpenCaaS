package logger

import (
	"go.uber.org/zap"
)

func Logger(err error) {
  loggerMgr, _:= NewLogger()
  zap.ReplaceGlobals(loggerMgr)
  logger := loggerMgr.Sugar()
  defer logger.Sync()
  logger.Info("[ERROR] -> ", err)
}

func NewLogger() (*zap.Logger, error) {
  cfg := zap.NewProductionConfig()
  cfg.OutputPaths = []string {
    "./opencaas.log",
    "stderr",
  }
  return cfg.Build()
}

