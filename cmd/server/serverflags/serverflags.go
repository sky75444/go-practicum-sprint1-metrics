package serverflags

import (
	"flag"
	"log"
	"os"
	"path/filepath"

	"github.com/caarlos0/env"
)

type flags struct {
	restoreFileData bool
	runAddr         string
	flagLogLevel    string
	fileName        string
	storeInterval   int
}

type envFlags struct {
	Restore         bool   `env:"RESTORE"`
	Address         string `env:"ADDRESS"`
	FileStoragePath string `env:"FILE_STORAGE_PATH"`
	StoreInterval   int    `env:"STORE_INTERVAL"`
}

func NewParsedFlags() *flags {
	flags := flags{}
	flag.StringVar(&flags.runAddr, "a", ":8080", "address and port to run server")
	flag.StringVar(&flags.flagLogLevel, "l", "info", "log level")

	currentDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	filepath := filepath.Join(currentDir, "savedMetricsData.json")

	flag.BoolVar(&flags.restoreFileData, "r", true, "restore saved data from file")
	flag.StringVar(&flags.fileName, "f", filepath, "save file dir")
	flag.IntVar(&flags.storeInterval, "i", 300, "interval of storing metric data to file")

	flag.Parse()

	var ef envFlags
	if err := env.Parse(&ef); err != nil {
		log.Fatal(err)
	}

	if ef.Address != "" {
		flags.runAddr = ef.Address
	}

	if ef.FileStoragePath != "" {
		flags.fileName = ef.FileStoragePath
	}

	if ef.Restore || !ef.Restore {
		flags.restoreFileData = ef.Restore
	}

	if ef.StoreInterval >= 0 {
		flags.storeInterval = ef.StoreInterval
	}

	return &flags
}

func (f *flags) GetRunAddr() string {
	return f.runAddr
}

func (f *flags) GetLogLevel() string {
	return f.flagLogLevel
}

func (f *flags) GetFileName() string {
	return f.fileName
}

func (f *flags) GetRestore() bool {
	return f.restoreFileData
}

func (f *flags) GetStoreInterval() int {
	return f.storeInterval
}
