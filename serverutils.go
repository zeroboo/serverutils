/*
Common utilities
*/
package serverutils

import (
	"log"
	"os"
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

func GetHostName() string {
	hostName, _ := os.Hostname()
	return hostName
}

func PrintEnv() {
	for _, env := range os.Environ() {
		log.Printf("Environment: %v", env)
	}
}
