package main

import (
	"fmt"
	"os"
	"strconv"
)

func getHost() (string, error) {
	const (
		defaultHost = "localhost"
		envKey      = "GOING_HOST"
	)
	envValue, ok := os.LookupEnv(envKey)
	if !ok {
		return defaultHost, nil
	}
	if envValue == "" {
		return "", fmt.Errorf("%q environment variable is empty", envKey)
	}
	return envValue, nil
}

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
		return 0, fmt.Errorf("could not parse %q environment variable: %v", envKey, err)
	}
	return uint16(port), nil
}
