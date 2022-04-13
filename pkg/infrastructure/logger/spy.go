package logger

import "go.uber.org/zap/zapcore"

type LoggerSpy struct{}

func (spy LoggerSpy) Info(msg string, fields ...zapcore.Field) {

}

func (spy LoggerSpy) Error(msg string, fields ...zapcore.Field) {

}

func (spy LoggerSpy) Warn(msg string, fields ...zapcore.Field) {

}

func (spy LoggerSpy) Debug(msg string, fields ...zapcore.Field) {

}

func (spy LoggerSpy) Fatal(msg string, fields ...zapcore.Field) {

}
