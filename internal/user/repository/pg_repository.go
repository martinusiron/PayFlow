package repository

import (
	"context"
	"database/sql"

	"github.com/martinusiron/PayFlow/internal/user/domain"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *userRepository {
	return &userRepository{db: db}
}

func (r *userRepository) GetByUsername(ctx context.Context, username string) (*domain.User, error) {
	query := `SELECT id, username, password, salary, role FROM users WHERE username = $1`
	user := &domain.User{}
	err := r.db.QueryRowContext(ctx, query, username).Scan(&user.ID, &user.Username, &user.Password, &user.Salary, &user.Role)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) GetByID(ctx context.Context, id int) (*domain.User, error) {
	query := `SELECT id, username, password, salary, role FROM users WHERE id = $1`
	user := &domain.User{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(&user.ID, &user.Username, &user.Password, &user.Salary, &user.Role)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) GetAllEmployees(ctx context.Context) ([]domain.User, error) {
	query := `SELECT id, username, salary FROM users WHERE role = 'employee'`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []domain.User
	for rows.Next() {
		var u domain.User
		err := rows.Scan(&u.ID, &u.Username, &u.Salary)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}

func (r *userRepository) Create(ctx context.Context, user *domain.User) error {
	query := `
		INSERT INTO users (username, password, salary, role, created_at, updated_at, created_by, updated_by, ip_address, request_id)
		VALUES ($1, $2, $3, $4, NOW(), NOW(), $5, $6, $7, $8)
		RETURNING id
	`
	err := r.db.QueryRowContext(ctx, query, user.Username, user.Password, user.Salary, user.Role,
		user.CreatedBy, user.UpdatedBy, user.IPAddress, user.RequestID,
	).Scan(&user.ID)
	return err
}
