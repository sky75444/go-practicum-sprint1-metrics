package serverconfig

type Config struct {
	RestoreFileData bool
	RunAddr         string
	LogLevel        string
	FileName        string
	StoreInterval   int
}

func NewConfig(runServerAddr, logLevel, fileName string, restore bool, storeInterval int) *Config {
	return &Config{
		RunAddr:         runServerAddr,
		LogLevel:        logLevel,
		RestoreFileData: restore,
		FileName:        fileName,
		StoreInterval:   storeInterval,
	}
}
