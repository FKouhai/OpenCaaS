package logger

import (
	"log"

	"go.uber.org/zap"
)

func LoggErr(logger *zap.SugaredLogger, err error) {
	logger.Error("[ERROR] -> ", err)
}

func NewLogger() (*zap.SugaredLogger) {
	cfg := zap.NewProductionConfig()
	cfg.OutputPaths = []string{
		"./opencaas.log",
		"stderr",
	}
  logger, err := cfg.Build()
  if err != nil {
    log.Fatalf("Unable to construct logger: err=%v", err)
  }
	zap.ReplaceGlobals(logger)
	return logger.Sugar()
}
