package util

import (
	"fmt"
	"os"
	"path"
	"strings"
)

func ExitError(err error) {
	fmt.Printf("Error: %v\n", err)
	os.Exit(1)
}

func ExpandUser(p string) (string, error) {
	if !strings.HasPrefix(p, "~") {
		return p, nil
	}
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return path.Join(home, p[1:]), nil
}

func Clamp(val, min, max int) int {
	if val < min {
		return min
	} else if val > max {
		return max
	}
	return val
}
