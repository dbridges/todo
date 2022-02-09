package config

import (
	"os"
	"path"

	"github.com/dbridges/todo/util"
	"gopkg.in/ini.v1"
)

type Config struct {
	Path string
}

func UserConfigDir() (string, error) {
	xdgConfig := os.Getenv("XDG_CONFIG_HOME")
	if len(xdgConfig) > 0 {
		return xdgConfig, nil
	}
	return os.UserConfigDir()
}

func UserCacheDir() (string, error) {
	xdgCache := os.Getenv("XDG_CACHE_HOME")
	if len(xdgCache) > 0 {
		return xdgCache, nil
	}
	return os.UserCacheDir()
}

func Load() (*Config, error) {
	configDir, err := UserConfigDir()
	if err != nil {
		return nil, err
	}

	cfg, err := ini.Load(path.Join(configDir, "todo", "config.ini"))
	if err != nil {
		return nil, err
	}

	p, err := util.ExpandUser(cfg.Section("core").Key("path").String())
	if err != nil {
		return nil, err
	}

	return &Config{Path: p}, nil
}

func (cfg *Config) IsValid() bool {
	return cfg.Path != ""
}
