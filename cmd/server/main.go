package main

import (
	"github.com/sky75444/go-practicum-sprint1-metrics/cmd/server/serverflags"
	"github.com/sky75444/go-practicum-sprint1-metrics/internal/app"
	"github.com/sky75444/go-practicum-sprint1-metrics/internal/logger"
	"github.com/sky75444/go-practicum-sprint1-metrics/internal/serverconfig"
)

func main() {
	f := serverflags.NewParsedFlags()
	c := serverconfig.NewConfig(f.GetRunAddr(), f.GetLogLevel())

	di := app.NewDI()
	di.Init(c)

	if err := logger.Init(c.LogLevel); err != nil {
		panic(err)
	}

	di.Router.Start()
}
