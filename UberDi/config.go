package UberDi

type Config struct {
	Enabled      bool
	DatabasePath string
	Port         string
}

func NewConfig() *Config {
	return &Config{
		Enabled:      true,
		DatabasePath: "./example.db",
		Port:         "8000",
	}
}
