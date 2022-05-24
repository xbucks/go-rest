package util

import (
	"strings"
	"time"
)

func FormatTimeToISO(timeToFormat time.Time) string {
	return timeToFormat.Format(time.RFC3339)
}

func CurrentISOTime() string {
	return FormatTimeToISO(time.Now().UTC())
}

func IsNotEmptyString(input string) bool {
	return len(strings.TrimSpace(input)) != 0
}

// IsEmptyString - Check if the given string is empty or not
func IsEmptyString(input string) bool {
	return len(strings.TrimSpace(input)) == 0
}
