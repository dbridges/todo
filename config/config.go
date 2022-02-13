package config

import (
	"io/ioutil"
	"os"
	"path"
	"regexp"
	"strings"

	"github.com/dbridges/todo/util"
)

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

	fullPath := path.Join(configDir, "todo", "config.ini")
	bytes, err := ioutil.ReadFile(fullPath)
	if err != nil {
		return nil, err
	}

	c := &Config{}
	c.LabelColors = make(map[string]string)

	keyRegex := regexp.MustCompile(`^\s*\[(?P<Key>[A-Za-z]+)\]\s*$`)
	currentKey := ""

	for _, line := range strings.Split(string(bytes), "\n") {
		if keyRegex.MatchString(line) {
			matches := keyRegex.FindStringSubmatch(line)
			currentKey = matches[1]
			continue
		}
		fields := strings.Fields(line)
		if len(fields) == 3 && fields[1] == "=" {
			switch currentKey {
			case "core":
				if fields[0] == "path" {
					p, err := util.ExpandUser(fields[2])
					if err != nil {
						return nil, err
					}
					c.Path = p
				}
			case "labels":
				c.LabelColors[fields[0]] = fields[2]
			}
		}
	}

	return c, nil
}

func (cfg *Config) IsValid() bool {
	return cfg.Path != ""
}
