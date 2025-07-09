package agentflags

import (
	"flag"
	"fmt"
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

	if len(flags.memServerAddr) == 5 {
		//Если длина 5, это значит что хост не указан. А для агента важно знать хост
		flags.memServerAddr = fmt.Sprintf("http://localhost%s", flags.memServerAddr)
	}
	if flags.memServerAddr[:4] != "http:" {
		flags.memServerAddr = fmt.Sprintf("http://%s", flags.memServerAddr)
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
