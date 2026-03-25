package models

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5"
)

type User struct {
	ID          string `json:"id"`
	Email       string `json:"email"`
	DisplayName string `json:"display_name"`
	PhoneNumber string `json:"phone_number"`
	PhotoUrl    string `json:"photo_url"`
}

// Default categories to seed for new users
var defaultCategories = []struct {
	Name         string
	MainCategory string
}{
	{"Emergency Fund", "Savings"},
	{"Investments", "Savings"},
	{"Debt Repayment", "Savings"},
	{"Short Term", "Savings"},
	{"Savings", "Savings"},
	{"Interests Earned", "Savings"},

	{"Pets", "Needs"},
	{"House Services", "Needs"},
	{"Bank Fees", "Needs"},
	{"Groceries", "Needs"},
	{"Rent", "Needs"},
	{"Utilities", "Needs"},
	{"Transportation", "Needs"},
	{"Work Lunches", "Needs"},
	{"Family Support", "Needs"},

	{"Streaming Services", "Wants"},
	{"Health", "Wants"},
	{"Leisure", "Wants"},
	{"Self Care", "Wants"},
	{"Entertainment", "Wants"},
	{"Shopping", "Wants"},
	{"Hobbies", "Wants"},
	{"Taxi", "Wants"},
	{"Restaurants", "Wants"},
	{"Travels", "Wants"},

	{"Transfer", "Transfer"},
	{"Salary", "Income"},
	{"Interests", "Income"},
	{"Payments", "Income"},
}

// CreateUser inserts a new user if they do not exist
func CreateUser(user User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Check if the user already exists
	var exists bool
	err := db.QueryRow(ctx, "SELECT EXISTS(SELECT 1 FROM users WHERE id = $1)", user.ID).Scan(&exists)
	if err != nil {
		return fmt.Errorf("failed to check user existence: %w", err)
	}

	if exists {
		return nil // User already exists, nothing to do
	}

	// Insert new user
	query := `
		INSERT INTO users (id, email, display_name, phone_number, photo_url) 
		VALUES ($1, $2, $3, $4, $5)
	`
	_, err = db.Exec(ctx, query, user.ID, user.Email, user.DisplayName, user.PhoneNumber, user.PhotoUrl)
	if err != nil {
		return fmt.Errorf("failed to insert user: %w", err)
	}

	// Seed default categories for the new user
	if err := seedCategoriesForUser(ctx, user.ID); err != nil {
		log.Printf("Warning: failed to seed categories for user %s: %v", user.ID, err)
		// Don't fail user creation if category seeding fails
	}

	return nil
}

// seedCategoriesForUser creates default categories for a new user
func seedCategoriesForUser(ctx context.Context, userID string) error {
	batch := &pgx.Batch{}
	for _, category := range defaultCategories {
		batch.Queue(
			"INSERT INTO categories (name, main_category, user_id) VALUES ($1, $2, $3) ON CONFLICT DO NOTHING",
			category.Name, category.MainCategory, userID,
		)
	}

	br := db.SendBatch(ctx, batch)
	defer br.Close()

	for range defaultCategories {
		if _, err := br.Exec(); err != nil {
			return fmt.Errorf("failed to insert category: %w", err)
		}
	}

	log.Printf("Seeded %d categories for user %s", len(defaultCategories), userID)
	return nil
}

// DeleteUser removes a user by their UID
func DeleteUser(uid string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := db.Exec(ctx, "DELETE FROM users WHERE id = $1", uid)
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}

	return nil
}
