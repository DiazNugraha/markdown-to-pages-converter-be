package utils

import (
	"fmt"
	"os"
)

func ConnectionUrlBuilder(name string) (string, error) {
	var url string

	switch name {
	case "fiber":
		url = fmt.Sprintf(
			"%s:%s",
			os.Getenv("SERVER_HOST"),
			os.Getenv("SERVER_PORT"),
		)
	default:
		return "", fmt.Errorf("invalid connection name")
	}

	return url, nil
}
