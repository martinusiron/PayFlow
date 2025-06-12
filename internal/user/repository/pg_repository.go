package repository

import (
	"context"
	"database/sql"
	"fmt"
	"math/rand"

	"github.com/martinusiron/PayFlow/internal/user/domain"
	"github.com/martinusiron/PayFlow/pkg/utils"
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

func (r *userRepository) SeedIfEmpty(ctx context.Context) error {
	var count int
	err := r.db.QueryRowContext(ctx, "SELECT COUNT(*) FROM users").Scan(&count)
	if err != nil {
		return err
	}
	if count > 0 {
		fmt.Println("Users already seeded.")
		return nil
	}

	for i := 1; i <= 100; i++ {
		username := fmt.Sprintf("employee%d", i)
		password := utils.HashPassword(fmt.Sprintf("pass%d", i))
		salary := rand.Intn(3000000) + 2000000
		_, err := r.db.ExecContext(ctx, `
			INSERT INTO users (username, password, role, salary)
			VALUES ($1, $2, 'employee', $3)`, username, password, salary)
		if err != nil {
			return err
		}
	}

	adminPass := utils.HashPassword("admin123")
	_, err = r.db.ExecContext(ctx, `
		INSERT INTO users (username, password, role, salary)
		VALUES ('admin', $1, 'admin', 10000000)`, adminPass)
	if err != nil {
		return err
	}

	fmt.Println("Seeded 100 employees and 1 admin successfully.")
	return nil
}
