package agentconfig

type Config struct {
	MemStorageServerAddr string
	LogLevel             string
	PollInterval         int
	ReportInterval       int
}

func NewConfig(memStorageServerAddr, LogLevel string, pollInterval, reportInterval int) *Config {
	return &Config{
		MemStorageServerAddr: memStorageServerAddr,
		LogLevel:             LogLevel,
		PollInterval:         pollInterval,
		ReportInterval:       reportInterval,
	}
}
