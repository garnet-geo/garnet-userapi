package env

import (
	"errors"
	"os"
	"strconv"

	log "github.com/sirupsen/logrus"
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

	log.Debugln("Got integer env", envName, "with value", intValue)

	return intValue
}

func GetStringEnv(envName string) string {
	value := os.Getenv(envName)
	if value == "" {
		panic(errors.New("missing string environment variable: " + envName))
	}

	log.Debugln("Got string env", envName, "with value", value)

	return value
}
