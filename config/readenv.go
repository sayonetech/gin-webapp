package config

import (
	"os"
	"strings"
)

func Read() map[string]interface{} {
	lines := os.Environ()
	key := make(map[string]interface{})
	for _, line := range lines {
		keyValue := strings.Split(line, "=")
		key[string(keyValue[0])] = keyValue[1]
	}
	return key
}
