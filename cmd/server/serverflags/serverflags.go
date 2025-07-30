package serverflags

import (
	"flag"
	"log"

	"github.com/caarlos0/env"
)

type flags struct {
	runAddr string
}

type envFlags struct {
	Address string `env:"ADDRESS"`
}

func NewParsedFlags() *flags {
	flags := flags{}
	flag.StringVar(&flags.runAddr, "a", ":8080", "address and port to run server")
	flag.Parse()

	var ef envFlags
	if err := env.Parse(&ef); err != nil {
		log.Fatal(err)
	}

	if ef.Address != "" {
		flags.runAddr = ef.Address
	}

	return &flags
}

func (f *flags) GetRunAddr() string {
	return f.runAddr
}
