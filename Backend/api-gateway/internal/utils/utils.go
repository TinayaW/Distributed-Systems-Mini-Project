package utils

import (
	"fmt"
	"strings"
)

func GetServiceName(path string) (string, error) {
	parts := strings.Split(path[1:], "/")
	if len(parts) > 1 {
		return parts[0] + "service", nil
	}
	return "", fmt.Errorf("invalid URL path")
}
