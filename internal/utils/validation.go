package utils

import (
	"regexp"
	"strings"
)

// ValidCurrencies contains the list of supported currencies
var ValidCurrencies = map[string]bool{
	"SEK": true,
	"USD": true,
	"EUR": true,
	"GBP": true,
	"COP": true,
	"MXN": true,
	"BRL": true,
	"ARS": true,
	"CLP": true,
	"PEN": true,
	"NOK": true,
	"DKK": true,
	"CHF": true,
	"JPY": true,
	"CNY": true,
	"AUD": true,
	"CAD": true,
	"NZD": true,
	"INR": true,
}

// ValidAccountTypes contains the list of supported account types
var ValidAccountTypes = map[string]bool{
	"Checking Account":     true,
	"Savings Account":      true,
	"Credit Card":          true,
	"Debit Card":           true,
	"Cash":                 true,
	"Investment Account":   true,
	"Digital Wallet":       true,
	"Money Market Account": true,
	"Loan":                 true,
	"Mortgage":             true,
}

// ValidMainCategories contains the list of supported main categories
var ValidMainCategories = map[string]bool{
	"Needs":    true,
	"Wants":    true,
	"Savings":  true,
	"Income":   true,
	"Transfer": true,
}

// ValidTransactionTypes contains the list of supported transaction types
var ValidTransactionTypes = map[string]bool{
	"Income":   true,
	"Expense":  true,
	"Savings":  true,
	"Transfer": true,
}

// IsValidEmail validates email format
func IsValidEmail(email string) bool {
	if email == "" {
		return false
	}
	// Basic email regex pattern
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	matched, _ := regexp.MatchString(pattern, email)
	return matched
}

// IsValidCurrency checks if the currency is supported
func IsValidCurrency(currency string) bool {
	return ValidCurrencies[strings.ToUpper(currency)]
}

// IsValidAccountType checks if the account type is supported
func IsValidAccountType(accountType string) bool {
	return ValidAccountTypes[accountType]
}

// IsValidMainCategory checks if the main category is supported
func IsValidMainCategory(category string) bool {
	return ValidMainCategories[category]
}

// IsValidTransactionType checks if the transaction type is supported
func IsValidTransactionType(transactionType string) bool {
	return ValidTransactionTypes[transactionType]
}
