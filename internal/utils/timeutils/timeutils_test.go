package timeutils

import (
	"testing"
	"time"
)

func TestCalculatePeriodBoundaries(t *testing.T) {
	// Fixed date for testing: January 15, 2024
	testDate := time.Date(2024, 1, 15, 12, 0, 0, 0, time.UTC)

	tests := []struct {
		name      string
		date      time.Time
		offset    []int
		wantStart int64
		wantEnd   int64
	}{
		{
			name:      "default offset",
			date:      testDate,
			offset:    nil,
			wantStart: time.Date(2024, 1, 16, 0, 0, 0, 0, time.UTC).Unix(),
			wantEnd:   time.Date(2024, 2, 17, 23, 59, 59, 0, time.UTC).Unix(),
		},
		{
			name:      "custom offset 0",
			date:      testDate,
			offset:    []int{0},
			wantStart: time.Date(2024, 1, 15, 0, 0, 0, 0, time.UTC).Unix(),
			wantEnd:   time.Date(2024, 2, 16, 23, 59, 59, 0, time.UTC).Unix(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotStart, gotEnd := CalculatePeriodBoundaries(tt.date, tt.offset...)
			if gotStart != tt.wantStart {
				t.Errorf("CalculatePeriodBoundaries() start = %v, want %v", gotStart, tt.wantStart)
			}
			if gotEnd != tt.wantEnd {
				t.Errorf("CalculatePeriodBoundaries() end = %v, want %v", gotEnd, tt.wantEnd)
			}
		})
	}
}

func TestGetSalaryMonthRange(t *testing.T) {
	// This test verifies the function returns valid dates
	// The exact values depend on the current date, so we test structure

	t.Run("default days (25-24)", func(t *testing.T) {
		start, end := GetSalaryMonthRange("", "")

		if start.IsZero() {
			t.Error("start date should not be zero")
		}
		if end.IsZero() {
			t.Error("end date should not be zero")
		}
		if !end.After(start) {
			t.Errorf("end date (%v) should be after start date (%v)", end, start)
		}
		if start.Day() != 25 {
			t.Errorf("start day should be 25, got %d", start.Day())
		}
		if end.Day() != 24 {
			t.Errorf("end day should be 24, got %d", end.Day())
		}
	})

	t.Run("custom days (1-28)", func(t *testing.T) {
		start, end := GetSalaryMonthRange("1", "28")

		if start.Day() != 1 {
			t.Errorf("start day should be 1, got %d", start.Day())
		}
		if end.Day() != 28 {
			t.Errorf("end day should be 28, got %d", end.Day())
		}
	})

	t.Run("custom days (15-14)", func(t *testing.T) {
		start, end := GetSalaryMonthRange("15", "14")

		if start.Day() != 15 {
			t.Errorf("start day should be 15, got %d", start.Day())
		}
		if end.Day() != 14 {
			t.Errorf("end day should be 14, got %d", end.Day())
		}
	})
}
