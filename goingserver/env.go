package main

import (
	"fmt"
	"os"
	"strconv"
)

func getPort() (uint16, error) {
	const (
		defaultPort = 30051
		envKey      = "GOING_PORT"
	)
	envValue, ok := os.LookupEnv(envKey)
	if !ok {
		return defaultPort, nil
	}
	port, err := strconv.ParseUint(envValue, 10, 16)
	if err != nil {
		return 0, fmt.Errorf("can't parse environment variable %q: %v", envKey, err)
	}
	return uint16(port), nil
}
