/*
Common utilities for server
*/
package serverutils

import (
	"log"
	"os"
	"strconv"
	"strings"
)

var Time timeUtil
var Random randomUtil
var Hash hashUtil

// GetIntegerEnvOrDefault return environment variable value or default value if the variable not exists
func GetEnvOrDefault(key string, defaultValue string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		value = defaultValue
	}
	return strings.TrimSpace(value)
}

// GetIntegerEnvOrDefault returns environment variable parsed as int value. If failed to parse environment variables, return default value
func GetIntegerEnvOrDefault(key string, defaultValue int64) (int64, error) {
	var retValue int64 = defaultValue

	textValue := os.Getenv(key)
	if len(textValue) == 0 {
		return retValue, nil
	}
	textValue = strings.TrimSpace(textValue)
	parsedValue, errParse := strconv.ParseInt(textValue, 10, 64)
	if errParse != nil {
		retValue = defaultValue
	} else {
		retValue = parsedValue
	}

	return retValue, errParse
}

// GetIntEnvOrDefault returns environment variable parsed as int64 value. If failed to parse environment variables, return default value
func GetInt64EnvOrDefault(key string, defaultValue int64) int64 {
	var retValue int64 = defaultValue

	textValue := os.Getenv(key)
	if len(textValue) == 0 {
		return defaultValue
	}

	textValue = strings.TrimSpace(textValue)
	parsedValue, errParse := strconv.ParseInt(textValue, 10, 64)
	if errParse != nil {
		retValue = defaultValue
	} else {
		retValue = parsedValue
	}

	return retValue
}

// GetIntEnvOrDefault returns environment variable parsed as int value. If failed to parse environment variables, return default value
func GetIntEnvOrDefault(key string, defaultValue int) int {
	var retValue int = defaultValue

	textValue := os.Getenv(key)
	if len(textValue) == 0 {
		return defaultValue
	}

	textValue = strings.TrimSpace(textValue)
	parsedValue, errParse := strconv.Atoi(textValue)
	if errParse != nil {
		retValue = defaultValue
	} else {
		retValue = parsedValue
	}

	return retValue
}

// GetHostName return host name of running machine
func GetHostName() string {
	hostName, _ := os.Hostname()
	return hostName
}

// PrintEnv dumps environment to standard output
func PrintEnv() {
	for _, env := range os.Environ() {
		log.Printf("Environment: %v", env)
	}
}
