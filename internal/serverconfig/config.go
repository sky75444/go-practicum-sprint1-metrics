package serverconfig

type Config struct {
	RunAddr string
}

func NewConfig(runServerAddr string) *Config {
	return &Config{
		RunAddr: runServerAddr,
	}
}
