package agentflags

import (
	"flag"
	"log"

	"github.com/caarlos0/env"
)

type flags struct {
	memServerAddr  string
	reportInterval int
	pollInterval   int
}

type envFlags struct {
	Address        string `env:"ADDRESS"`
	ReportInterval int    `env:"REPORT_INTERVAL"`
	PollInterval   int    `env:"POLL_INTERVAL"`
}

func NewParsedFlags() *flags {
	flags := flags{}
	flag.StringVar(&flags.memServerAddr, "a", ":8080", "address and port of metrics collector server")
	flag.IntVar(&flags.reportInterval, "r", 10, "the interval of sending to the server")
	flag.IntVar(&flags.pollInterval, "p", 2, "the interval of collecting metrics")
	flag.Parse()

	var ef envFlags
	if err := env.Parse(&ef); err != nil {
		log.Fatal(err)
	}

	if ef.Address != "" {
		flags.memServerAddr = ef.Address
	}

	if ef.ReportInterval != 0 {
		flags.pollInterval = ef.ReportInterval
	}

	if ef.PollInterval != 0 {
		flags.pollInterval = ef.PollInterval
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
