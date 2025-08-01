package alogger

import (
	"go.uber.org/zap"
)

var AZLog *zap.Logger = zap.NewNop()

func Init(level string) error {
	lvl, err := zap.ParseAtomicLevel(level)
	if err != nil {
		return err
	}

	zcfg := zap.NewProductionConfig()
	zcfg.Level = lvl
	zl, err := zcfg.Build()
	if err != nil {
		return err
	}

	AZLog = zl
	return nil
}

func ZError(err error) zap.Field {
	return zap.Error(err)
}
