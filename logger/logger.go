package logger

import "go.uber.org/zap"

type Logger interface {
	Error(string)
	Info(string)
	Debug(string)
}

func NewLogger() (Logger, error) {
	l, err := NewZapLogger()

	if err != nil {
		return nil, err
	}

	return l, nil
}

type ZapLogger struct {
	Logger
	zapLog *zap.Logger
}

func NewZapLogger() (*ZapLogger, error) {
	l, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}
	zLogger := ZapLogger{
		zapLog: l,
	}

	return &zLogger, nil
}

func (z *ZapLogger) Info(msg string) {
	z.zapLog.Info(msg)
}

func (z *ZapLogger) Error(msg string) {
	z.zapLog.Error(msg)
}

func (z *ZapLogger) Debug(msg string) {
	z.zapLog.Debug(msg)
}
