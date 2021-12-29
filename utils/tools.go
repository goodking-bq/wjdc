package utils

import "os"

func GetEnvString(name string, defaults ...string) string {
	val := os.Getenv(name)
	if val == "" {
		return defaults[0]
	}
	return val
}
