package models

import (
	"context"
	"fmt"
)

func RecalculateAllAccountBalances(ctx context.Context, uid string) error {
	// Primero, obtenemos todas las cuentas
	rows, err := db.Query(ctx, `SELECT id FROM accounts WHERE user_id = $1`, uid)
	if err != nil {
		return fmt.Errorf("failed to fetch accounts: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var accountID string
		if err := rows.Scan(&accountID); err != nil {
			return fmt.Errorf("failed to scan account ID: %v", err)
		}

		// Recalcular el balance para esta cuenta
		var balance float64
		err = db.QueryRow(ctx, `
			SELECT COALESCE(SUM(amount), 0)
			FROM transactions
			WHERE account_id = $1`, accountID).Scan(&balance)
		if err != nil {
			return fmt.Errorf("failed to calculate balance for account %s: %v", accountID, err)
		}

		// Actualizar la cuenta con el nuevo balance
		_, err = db.Exec(ctx, `
			UPDATE accounts
			SET balance = $1
			WHERE id = $2`, balance, accountID)
		if err != nil {
			return fmt.Errorf("failed to update balance for account %s: %v", accountID, err)
		}
	}

	if err := rows.Err(); err != nil {
		return fmt.Errorf("error iterating over accounts: %v", err)
	}

	return nil
}
