package util

import (
	"testing"
	"time"
)

func TestFormatTimeToISO(t *testing.T) {
	got := FormatTimeToISO(time.Date(2022, 5, 18, 9, 36, 0, 0, time.UTC))
	want := "2022-05-18T09:36:00Z"

	if want != got {
		t.Errorf("Expected '%s', but got '%s'", want, got)
	}

}
