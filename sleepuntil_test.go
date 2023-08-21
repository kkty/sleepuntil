package main

import (
	"testing"
	"time"
)

func TestParseTime(t *testing.T) {
	tests := []struct {
		input    string
		expected time.Time
		hasError bool
	}{
		{"7:00am", time.Date(0, 1, 1, 7, 0, 0, 0, time.UTC), false},
		{"7:00:30am", time.Date(0, 1, 1, 7, 0, 30, 0, time.UTC), false},
		{"19:00", time.Date(0, 1, 1, 19, 0, 0, 0, time.UTC), false},
		{"19:00:30", time.Date(0, 1, 1, 19, 0, 30, 0, time.UTC), false},
		{"invalid", time.Time{}, true},
	}

	for _, tt := range tests {
		got, err := parseTime(tt.input)
		if (err != nil) != tt.hasError {
			t.Fatalf("parseTime(%q) returned unexpected error: %v", tt.input, err)
		}
		if !got.Equal(tt.expected) {
			t.Fatalf("parseTime(%q) = %v; want %v", tt.input, got, tt.expected)
		}
	}
}

func TestGetNextTargetTime(t *testing.T) {
	now := time.Now()
	past := time.Date(0, 1, 1, now.Hour()-1, now.Minute(), 0, 0, time.UTC)
	future := time.Date(0, 1, 1, now.Hour()+1, now.Minute(), 0, 0, time.UTC)

	tests := []struct {
		input    time.Time
		expected time.Duration
	}{
		{past, 23 * time.Hour},
		{future, time.Hour},
	}

	for _, tt := range tests {
		got := getNextTargetTime(tt.input)
		duration := got.Sub(now)

		if duration > tt.expected+time.Minute || duration < tt.expected-time.Minute {
			t.Fatalf("getNextTargetTime(%v) expected to be roughly %v hours from now, but was %v", tt.input, tt.expected.Hours(), duration.Hours())
		}
	}
}
