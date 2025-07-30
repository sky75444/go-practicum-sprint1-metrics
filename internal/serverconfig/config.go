package serverconfig

type Config struct {
	RunAddr  string
	LogLevel string
}

func NewConfig(runServerAddr string, logLevel string) *Config {
	return &Config{
		RunAddr:  runServerAddr,
		LogLevel: logLevel,
	}
}
