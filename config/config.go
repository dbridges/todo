package config

import (
	"io/ioutil"
	"os"
	"path"
	"regexp"
	"strings"

	"github.com/dbridges/todo/util"
)

var keyRegex = regexp.MustCompile(`^\s*\[(?P<Key>[A-Za-z]+)\]\s*$`)
var keyValueRegex = regexp.MustCompile(`^\s*(?P<Key>[A-Za-z]+) = "(?P<Value>.+)"$`)

type Config struct {
	Path        string
	LabelColors map[string]string
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

	fullPath := path.Join(configDir, "todo", "config.toml")
	bytes, err := ioutil.ReadFile(fullPath)
	if err != nil {
		return nil, err
	}

	c := &Config{}
	c.LabelColors = make(map[string]string)

	currentKey := ""

	for _, line := range strings.Split(string(bytes), "\n") {
		if keyRegex.MatchString(line) {
			matches := keyRegex.FindStringSubmatch(line)
			currentKey = matches[1]
			continue
		} else if keyValueRegex.MatchString(line) {
			matches := keyValueRegex.FindStringSubmatch(line)
			switch currentKey {
			case "core":
				if matches[1] == "path" {
					p, err := util.ExpandUser(matches[2])
					if err != nil {
						return nil, err
					}
					c.Path = p
				}
			case "labels":
				c.LabelColors[matches[1]] = matches[2]
			}
		}
	}

	return c, nil
}

func (cfg *Config) IsValid() bool {
	return cfg.Path != ""
}
