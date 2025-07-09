package main

import (
	"github.com/sky75444/go-practicum-sprint1-metrics/cmd/server/serverflags"
	"github.com/sky75444/go-practicum-sprint1-metrics/internal/app"
	"github.com/sky75444/go-practicum-sprint1-metrics/internal/serverconfig"
)

func main() {
	f := serverflags.NewParsedFlags()
	c := serverconfig.NewConfig(f.GetRunAddr())

	di := app.NewDI()
	di.Init(c)

	di.Router.Start()
}
