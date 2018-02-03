package manager

import "gopkg.in/urfave/cli.v1"

type Config struct {
	Debug  bool         `toml:"debug"`
	Server ServerConfig `toml:"server"`
	Files  FileConfig   `toml:"files"`
}

type ServerConfig struct {
	Addr string `toml:"addr"`
	Port int    `toml:"port"`
}

type FileConfig struct {
	StatusFile string `toml:"status_file"`
	DBFile     string `toml:"db_file"`
	DBType     string `toml:"db_type"`
}

func LoadConfig(cfgFile string, c *cli.Context) (*Config, error) {
	cfg := new(Config)
	cfg.Server.Addr = "0.0.0.0"
	cfg.Server.Port = 14242
	cfg.Debug = false
	cfg.Files.StatusFile = ""

	return cfg, nil
}
