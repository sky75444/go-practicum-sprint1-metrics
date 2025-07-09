package agentconfig

type Config struct {
	MemStorageServerAddr string
	PollInterval         int
	ReportInterval       int
}

func NewConfig(memStorageServerAddr string, pollInterval, reportInterval int) *Config {
	return &Config{
		MemStorageServerAddr: memStorageServerAddr,
		PollInterval:         pollInterval,
		ReportInterval:       reportInterval,
	}
}
