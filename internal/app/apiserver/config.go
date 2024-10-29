package apiserver

// Config ...
type Config struct {
	BindAddr string `toml:"bind_addr"` //представляет адрес, к которому сервер будет привязан
	LogLevel string `toml:"log_level"`
}

// NewConfig ...
func NewConfig() *Config {
	return &Config{
		BindAddr: "8080",
		LogLevel: "debug",
	}
}
