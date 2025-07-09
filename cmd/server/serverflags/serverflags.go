package serverflags

import "flag"

type flags struct {
	runAddr string
}

func NewParsedFlags() *flags {
	flags := flags{}
	flag.StringVar(&flags.runAddr, "a", ":8080", "address and port to run server")
	flag.Parse()

	return &flags
}

func (f *flags) GetRunAddr() string {
	return f.runAddr
}
