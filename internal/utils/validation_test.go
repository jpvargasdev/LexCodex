package utils

import "testing"

func TestIsValidEmail(t *testing.T) {
	tests := []struct {
		name  string
		email string
		want  bool
	}{
		{"valid email", "test@example.com", true},
		{"valid email with subdomain", "test@mail.example.com", true},
		{"valid email with plus", "test+tag@example.com", true},
		{"valid email with dots", "test.user@example.com", true},
		{"empty email", "", false},
		{"no at symbol", "testexample.com", false},
		{"no domain", "test@", false},
		{"no username", "@example.com", false},
		{"no tld", "test@example", false},
		{"spaces", "test @example.com", false},
		{"double at", "test@@example.com", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidEmail(tt.email); got != tt.want {
				t.Errorf("IsValidEmail(%q) = %v, want %v", tt.email, got, tt.want)
			}
		})
	}
}

func TestIsValidCurrency(t *testing.T) {
	tests := []struct {
		name     string
		currency string
		want     bool
	}{
		{"SEK uppercase", "SEK", true},
		{"USD uppercase", "USD", true},
		{"EUR uppercase", "EUR", true},
		{"COP uppercase", "COP", true},
		{"sek lowercase", "sek", true},
		{"usd lowercase", "usd", true},
		{"invalid currency", "XYZ", false},
		{"empty currency", "", false},
		{"partial currency", "US", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidCurrency(tt.currency); got != tt.want {
				t.Errorf("IsValidCurrency(%q) = %v, want %v", tt.currency, got, tt.want)
			}
		})
	}
}

func TestIsValidAccountType(t *testing.T) {
	tests := []struct {
		name        string
		accountType string
		want        bool
	}{
		{"Checking Account", "Checking Account", true},
		{"Savings Account", "Savings Account", true},
		{"Credit Card", "Credit Card", true},
		{"Debit Card", "Debit Card", true},
		{"Cash", "Cash", true},
		{"Investment Account", "Investment Account", true},
		{"Digital Wallet", "Digital Wallet", true},
		{"Money Market Account", "Money Market Account", true},
		{"Loan", "Loan", true},
		{"Mortgage", "Mortgage", true},
		{"invalid type", "Invalid", false},
		{"empty type", "", false},
		{"old Bank type", "Bank", false},
		{"old Checking type", "Checking", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidAccountType(tt.accountType); got != tt.want {
				t.Errorf("IsValidAccountType(%q) = %v, want %v", tt.accountType, got, tt.want)
			}
		})
	}
}

func TestIsValidMainCategory(t *testing.T) {
	tests := []struct {
		name     string
		category string
		want     bool
	}{
		{"Needs", "Needs", true},
		{"Wants", "Wants", true},
		{"Savings", "Savings", true},
		{"Income", "Income", true},
		{"Transfer", "Transfer", true},
		{"invalid category", "Invalid", false},
		{"empty category", "", false},
		{"lowercase needs", "needs", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidMainCategory(tt.category); got != tt.want {
				t.Errorf("IsValidMainCategory(%q) = %v, want %v", tt.category, got, tt.want)
			}
		})
	}
}

func TestIsValidTransactionType(t *testing.T) {
	tests := []struct {
		name            string
		transactionType string
		want            bool
	}{
		{"Income", "Income", true},
		{"Expense", "Expense", true},
		{"Savings", "Savings", true},
		{"Transfer", "Transfer", true},
		{"invalid type", "Invalid", false},
		{"empty type", "", false},
		{"lowercase income", "income", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidTransactionType(tt.transactionType); got != tt.want {
				t.Errorf("IsValidTransactionType(%q) = %v, want %v", tt.transactionType, got, tt.want)
			}
		})
	}
}
