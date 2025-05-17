package repository

import (
	"database/sql"
	"fmt"
	"time"

	"certificate-ledger/domain"
	"github.com/google/uuid"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Save(user *domain.User) error {
	if user.ID == "" {
		user.ID = uuid.New().String()
	}
	if user.CreatedAt.IsZero() {
		user.CreatedAt = time.Now()
	}
	user.UpdatedAt = time.Now()

	query := `
		INSERT INTO users (id, name, email, password, role, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?, ?)
		ON DUPLICATE KEY UPDATE
			name = ?, email = ?, password = ?, role = ?, updated_at = ?`
	_, err := r.db.Exec(query,
		user.ID, user.Name, user.Email, user.Password, user.Role, user.CreatedAt, user.UpdatedAt,
		user.Name, user.Email, user.Password, user.Role, user.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("failed to save user: %v", err)
	}
	return nil
}

func (r *UserRepository) FindByID(id string) (*domain.User, error) {
	query := `SELECT id, name, email, password, role, created_at, updated_at
	          FROM users WHERE id = ?`
	row := r.db.QueryRow(query, id)

	var user domain.User
	err := row.Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.Role,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("user with ID %s not found", id)
	}
	if err != nil {
		return nil, fmt.Errorf("failed to find user: %v", err)
	}
	return &user, nil
}

func (r *UserRepository) FindByEmail(email string) (*domain.User, error) {
	query := `SELECT id, name, email, password, role, created_at, updated_at
	          FROM users WHERE email = ?`
	row := r.db.QueryRow(query, email)

	var user domain.User
	err := row.Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.Role,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("user with email %s not found", email)
	}
	if err != nil {
		return nil, fmt.Errorf("failed to find user: %v", err)
	}
	return &user, nil
}

func (r *UserRepository) FindAll() ([]*domain.User, error) {
	query := `SELECT id, name, email, password, role, created_at, updated_at
	          FROM users`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to query users: %v", err)
	}
	defer rows.Close()

	var users []*domain.User
	for rows.Next() {
		var user domain.User
		if err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.Password,
			&user.Role,
			&user.CreatedAt,
			&user.UpdatedAt,
		); err != nil {
			return nil, fmt.Errorf("failed to scan user: %v", err)
		}
		users = append(users, &user)
	}
	return users, nil
}

func (r *UserRepository) Delete(id string) error {
	query := `DELETE FROM users WHERE id = ?`
	result, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete user: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to check rows affected: %v", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("user with ID %s not found", id)
	}

	return nil
}