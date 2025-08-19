package infrastructure

import (
	"database/sql"

	"github.com/didinj/go-clean-architecture/internal/entity"
)

type PostgresUserRepository struct {
	db *sql.DB
}

func NewPostgresUserRepository(db *sql.DB) *PostgresUserRepository {
	return &PostgresUserRepository{db: db}
}

func (r *PostgresUserRepository) Create(user *entity.User) error {
	_, err := r.db.Exec("INSERT INTO users(name, email) VALUES($1, $2)", user.Name, user.Email)
	return err
}

func (r *PostgresUserRepository) GetByID(id int64) (*entity.User, error) {
	query := `SELECT id, name, email FROM users WHERE id = $1`
	row := r.db.QueryRow(query, id)

	var user entity.User
	err := row.Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *PostgresUserRepository) GetAll() ([]*entity.User, error) {
	query := "SELECT id, name, email FROM users"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*entity.User
	for rows.Next() {
		var user *entity.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}
