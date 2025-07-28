package agentflags

import (
	"flag"
	"log"
	"time"

	"github.com/caarlos0/env"
)

type flags struct {
	memServerAddr  string
	reportInterval int
	pollInterval   int
}

type envFlags struct {
	Address         string        `env:"ADDRESS"`
	Report_interval time.Duration `env:"REPORT_INTERVAL"`
	Poll_interval   time.Duration `env:"POLL_INTERVAL"`
}

func NewParsedFlags() *flags {
	flags := flags{}
	flag.StringVar(&flags.memServerAddr, "a", ":8080", "address and port to run server")
	flag.IntVar(&flags.reportInterval, "r", 10, "the interval of sending to the server")
	flag.IntVar(&flags.pollInterval, "p", 2, "the interval of collecting metrics")
	flag.Parse()

	ef := envFlags{}
	if err := env.Parse(&ef); err != nil {
		log.Fatal(err)
	}

	if ef.Address != "" {
		flags.memServerAddr = ef.Address
	}

	if ef.Poll_interval != 0 {
		flags.pollInterval = int(ef.Poll_interval)
	}

	if ef.Report_interval != 0 {
		flags.pollInterval = int(ef.Report_interval)
	}

	return &flags
}

func (f *flags) GetMemServerAddr() string {
	return f.memServerAddr
}

func (f *flags) GetReportInterval() int {
	return f.reportInterval
}

func (f *flags) GetPollInterval() int {
	return f.pollInterval
}
