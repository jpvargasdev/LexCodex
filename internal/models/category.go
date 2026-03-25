package models

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5"
)

type Category struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	MainCategory string `json:"main_category"`
	UserID       string `json:"user_id,omitempty"`
}

func GetCategories(uid string) ([]Category, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Get default categories (no user_id) AND user-specific categories
	rows, err := db.Query(ctx, "SELECT id, name, main_category FROM categories WHERE user_id IS NULL OR user_id = $1", uid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []Category
	for rows.Next() {
		var category Category
		if err := rows.Scan(&category.ID, &category.Name, &category.MainCategory); err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return categories, nil
}

func AddCategory(category Category) (Category, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := "INSERT INTO categories (name, main_category, user_id) VALUES ($1, $2, $3) RETURNING id"
	err := db.QueryRow(ctx, query, category.Name, category.MainCategory, category.UserID).Scan(&category.ID)
	if err != nil {
		return Category{}, err
	}

	return category, nil
}

// UpdateCategory updates an existing category
func UpdateCategory(category Category) (Category, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := db.Exec(ctx,
		"UPDATE categories SET name = $1, main_category = $2 WHERE id = $3 AND user_id = $4",
		category.Name, category.MainCategory, category.ID, category.UserID,
	)

	if err != nil {
		log.Println("Error updating category:", err)
		return Category{}, err
	}

	return category, nil
}

// DeleteCategory removes a category from the database
func DeleteCategory(category Category) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := db.Exec(ctx,
		"DELETE FROM categories WHERE id = $1 AND user_id = $2",
		category.ID, category.UserID,
	)

	if err != nil {
		log.Println("Error deleting category:", err)
		return err
	}

	return nil
}

// GetMainCategory returns the main category based on the category ID
func GetMainCategory(id string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var mainCategory string
	err := db.QueryRow(ctx, "SELECT main_category FROM categories WHERE id = $1", id).Scan(&mainCategory)
	if err != nil {
		if err == pgx.ErrNoRows {
			return "", fmt.Errorf("category '%s' not found in categories table", id)
		}
		return "", err
	}
	return mainCategory, nil
}

// GetSubCategory returns the subcategory name based on the category ID
func GetSubCategory(id string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var subCategory string
	err := db.QueryRow(ctx, "SELECT name FROM categories WHERE id = $1", id).Scan(&subCategory)
	if err != nil {
		if err == pgx.ErrNoRows {
			return "", fmt.Errorf("subcategory '%s' not found in categories table", id)
		}
		return "", err
	}
	return subCategory, nil
}
