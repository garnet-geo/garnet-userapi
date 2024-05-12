package env

import (
	"errors"
	"os"
	"strconv"
)

func GetIntegerEnv(envName string) int {
	value := os.Getenv(envName)
	if value == "" {
		panic(errors.New("missing integer environment variable: " + envName))
	}

	intValue, err := strconv.Atoi(value)
	if err != nil {
		panic(err)
	}

	return intValue
}

func GetStringEnv(envName string) string {
	value := os.Getenv(envName)
	if value == "" {
		panic(errors.New("missing string environment variable: " + envName))
	}

	return value
}
