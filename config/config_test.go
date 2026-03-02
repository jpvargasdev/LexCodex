package config

import (
	"os"
	"testing"
)

func TestGetEnv(t *testing.T) {
	tests := []struct {
		name       string
		key        string
		defaultVal string
		envValue   string
		setEnv     bool
		want       string
	}{
		{
			name:       "returns env value when set",
			key:        "TEST_VAR_1",
			defaultVal: "default",
			envValue:   "custom_value",
			setEnv:     true,
			want:       "custom_value",
		},
		{
			name:       "returns default when env not set",
			key:        "TEST_VAR_2",
			defaultVal: "default_value",
			envValue:   "",
			setEnv:     false,
			want:       "default_value",
		},
		{
			name:       "returns empty string when env is empty",
			key:        "TEST_VAR_3",
			defaultVal: "default",
			envValue:   "",
			setEnv:     true,
			want:       "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Clean up before test
			os.Unsetenv(tt.key)

			if tt.setEnv {
				os.Setenv(tt.key, tt.envValue)
				defer os.Unsetenv(tt.key)
			}

			got := getEnv(tt.key, tt.defaultVal)
			if got != tt.want {
				t.Errorf("getEnv(%q, %q) = %q, want %q", tt.key, tt.defaultVal, got, tt.want)
			}
		})
	}
}

func TestLoad(t *testing.T) {
	// Save original env values
	origPort := os.Getenv("SERVER_PORT")
	origEnv := os.Getenv("ENV")
	origCurrency := os.Getenv("BASE_CURRENCY")

	// Restore after test
	defer func() {
		if origPort != "" {
			os.Setenv("SERVER_PORT", origPort)
		} else {
			os.Unsetenv("SERVER_PORT")
		}
		if origEnv != "" {
			os.Setenv("ENV", origEnv)
		} else {
			os.Unsetenv("ENV")
		}
		if origCurrency != "" {
			os.Setenv("BASE_CURRENCY", origCurrency)
		} else {
			os.Unsetenv("BASE_CURRENCY")
		}
	}()

	t.Run("loads default values", func(t *testing.T) {
		os.Unsetenv("SERVER_PORT")
		os.Unsetenv("ENV")
		os.Unsetenv("BASE_CURRENCY")

		Load()

		if Config.ServerPort != "8080" {
			t.Errorf("ServerPort = %q, want %q", Config.ServerPort, "8080")
		}
		if Config.Env != "debug" {
			t.Errorf("Env = %q, want %q", Config.Env, "debug")
		}
		if Config.BaseCurrency != "SEK" {
			t.Errorf("BaseCurrency = %q, want %q", Config.BaseCurrency, "SEK")
		}
	})

	t.Run("loads custom values from env", func(t *testing.T) {
		os.Setenv("SERVER_PORT", "3000")
		os.Setenv("ENV", "production")
		os.Setenv("BASE_CURRENCY", "USD")

		Load()

		if Config.ServerPort != "3000" {
			t.Errorf("ServerPort = %q, want %q", Config.ServerPort, "3000")
		}
		if Config.Env != "production" {
			t.Errorf("Env = %q, want %q", Config.Env, "production")
		}
		if Config.BaseCurrency != "USD" {
			t.Errorf("BaseCurrency = %q, want %q", Config.BaseCurrency, "USD")
		}
	})
}

func TestGetters(t *testing.T) {
	// Set known values
	Config.ServerPort = "9000"
	Config.Env = "test"
	Config.BaseCurrency = "EUR"
	Config.SecretKey = "secret123"

	t.Run("GetServerPort", func(t *testing.T) {
		if got := GetServerPort(); got != "9000" {
			t.Errorf("GetServerPort() = %q, want %q", got, "9000")
		}
	})

	t.Run("GetEnv", func(t *testing.T) {
		if got := GetEnv(); got != "test" {
			t.Errorf("GetEnv() = %q, want %q", got, "test")
		}
	})

	t.Run("GetBaseCurrency", func(t *testing.T) {
		if got := GetBaseCurrency(); got != "EUR" {
			t.Errorf("GetBaseCurrency() = %q, want %q", got, "EUR")
		}
	})

	t.Run("GetSecretKey", func(t *testing.T) {
		if got := GetSecretKey(); got != "secret123" {
			t.Errorf("GetSecretKey() = %q, want %q", got, "secret123")
		}
	})
}
