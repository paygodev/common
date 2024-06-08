package utils

import (
	"strconv"
	"syscall"
)

func EnvString(key, fallback string) string {
	if val, ok := syscall.Getenv(key); ok {
		return val
	}

	return fallback
}

func EnvInt(key string, fallback int64) int64 {
	if val, ok := syscall.Getenv(key); ok {
		i, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			return fallback
		}
		return i
	}

	return fallback
}
