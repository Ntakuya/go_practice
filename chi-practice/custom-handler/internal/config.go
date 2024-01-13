package internal

type Config struct {
	port int
}

func NewConfig() *Config {
	return &Config{
		port: 3333,
	}
}

func (c *Config) Port() int {
	return c.port
}
