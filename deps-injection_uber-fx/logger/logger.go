package logger

import (
	"go.uber.org/fx"
	"log"
	"os"
)

type Logger interface {
	Println(v ...interface{})
}

func NewLogger() Logger {
	return log.New(os.Stdout, "[ACME] ", 0)
}

// with fx

var Module = fx.Options(
	fx.Provide(
		NewLogger,
		),
	)