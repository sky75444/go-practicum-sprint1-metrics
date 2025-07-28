package agentflags

import (
	"flag"
)

type flags struct {
	memServerAddr  string
	reportInterval int
	pollInterval   int
}

func NewParsedFlags() *flags {
	flags := flags{}
	flag.StringVar(&flags.memServerAddr, "a", ":8080", "address and port to run server")
	flag.IntVar(&flags.reportInterval, "r", 10, "the interval of sending to the server")
	flag.IntVar(&flags.pollInterval, "p", 2, "the interval of collecting metrics")
	flag.Parse()

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
