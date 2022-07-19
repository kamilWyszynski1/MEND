package repository

import (
	"MEND/pkg/models"
	"context"
	"database/sql"
	"fmt"
)

const (
	insertUserQuery = "INSERT INTO user_table(id, name, surname) VALUES ($1, $2, $3)"
	getUserQuery    = "SELECT * FROM user_table WHERE id = $1"
	updateUserQuery = "UPDATE user_table SET name = $1, surname = $2 WHERE id = $3"
	deleteUserQuery = "DELETE FROM user_table WHERE id = $1"
)

// PSQL is a handler for PSQL operations.
type PSQL struct {
	client *sql.DB
}

// NewPSQL returns new PSQL instance.
func NewPSQL(client *sql.DB) *PSQL {
	return &PSQL{client: client}
}

// Create inserts new user into db.
func (p *PSQL) Create(ctx context.Context, user models.User) error {
	_, err := p.client.ExecContext(ctx, insertUserQuery, user.ID, user.Name, user.Surname)
	if err != nil {
		return fmt.Errorf("failed to insert user, %w", err)
	}
	return nil
}

// Get queries user by given id.
func (p *PSQL) Get(ctx context.Context, id int) (*models.User, error) {
	var user models.User
	err := p.client.QueryRowContext(ctx, getUserQuery, id).Scan(&user.ID, &user.Name, &user.Surname)

	if err != nil {
		return nil, fmt.Errorf("failed to query user, %w", err)
	}
	return &user, nil
}

// Update updates existing user.
func (p *PSQL) Update(ctx context.Context, id int, user models.User) error {
	_, err := p.client.ExecContext(ctx, updateUserQuery, user.Name, user.Surname, id)
	if err != nil {
		return fmt.Errorf("failed to update user, %w", err)
	}
	return nil
}

// Delete deletes existing user.
func (p *PSQL) Delete(ctx context.Context, id int) error {
	_, err := p.client.ExecContext(ctx, deleteUserQuery, id)
	if err != nil {
		return fmt.Errorf("failed to delete user, %w", err)
	}
	return nil
}
