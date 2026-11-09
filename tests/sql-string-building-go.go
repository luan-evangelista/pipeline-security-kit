package rules

import (
	"context"
	"database/sql"
	"fmt"
)

func find(ctx context.Context, db *sql.DB, email string) {
	// ruleid: sql-string-building-go
	db.Query(fmt.Sprintf("SELECT id FROM users WHERE email = '%s'", email))

	// ruleid: sql-string-building-go
	db.QueryRowContext(ctx, "SELECT id FROM users WHERE email = '"+email+"'")

	q := fmt.Sprintf("DELETE FROM users WHERE email = '%s'", email)
	// ruleid: sql-string-building-go
	db.Exec(q)

	// ok: sql-string-building-go
	db.QueryRowContext(ctx, "SELECT id FROM users WHERE email = ?", email)
}
