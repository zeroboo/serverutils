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

// Return environment value or default value if not exists
func GetEnvOrDefault(key string, defaultValue string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		value = defaultValue
	}
	return strings.TrimSpace(value)
}

// Return environment value or default value if not exists
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

func GetHostName() string {
	hostName, _ := os.Hostname()
	return hostName
}

func PrintEnv() {
	for _, env := range os.Environ() {
		log.Printf("Environment: %v", env)
	}
}
