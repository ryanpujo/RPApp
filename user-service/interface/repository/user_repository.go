package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/spriigan/RPMedia/domain"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *userRepository {
	return &userRepository{db: db}
}

var ErrNoUserFound = errors.New("user is not registered yet")

func (repo *userRepository) Create(user *domain.UserPayload) (int, error) {
	ctx, cancel := context.WithTimeout(context.TODO(), 2*time.Second)
	defer cancel()

	statement := "insert into users (first_name, last_name, username, password, email) values ($1, $2, $3, $4, $5) returning id"
	var id int

	err := repo.db.QueryRowContext(ctx, statement,
		user.Fname,
		user.Lname,
		user.Username,
		user.Password,
		user.Email,
	).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (repo *userRepository) FindUsers() ([]*domain.User, error) {
	ctx, cancel := context.WithTimeout(context.TODO(), 2*time.Second)
	defer cancel()

	statement := `select id, first_name, last_name, username, password, email from users order by first_name`

	rows, err := repo.db.QueryContext(ctx, statement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []*domain.User

	for rows.Next() {
		var user domain.User
		err = rows.Scan(
			&user.Id,
			&user.Fname,
			&user.Lname,
			&user.Username,
			&user.Password,
			&user.Email,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	return users, nil
}

func (repo *userRepository) FindByUsername(username string) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(context.TODO(), 2*time.Second)
	defer cancel()

	statement := `select id, first_name, last_name, username, password, email from users where username=$1`
	var user domain.User

	err := repo.db.QueryRowContext(ctx, statement, username).Scan(
		&user.Id,
		&user.Fname,
		&user.Lname,
		&user.Username,
		&user.Password,
		&user.Email,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoUserFound
		}
		return nil, err
	}
	return &user, nil
}

func (repo *userRepository) DeleteByUsername(username string) error {
	ctx, cancel := context.WithTimeout(context.TODO(), 2*time.Second)
	defer cancel()

	statement := "delete from users where username=$1"

	_, err := repo.db.ExecContext(ctx, statement, username)
	if err != nil {
		return err
	}
	return nil
}

func (repo *userRepository) Update(user domain.UserPayload) error {
	ctx, cancel := context.WithTimeout(context.TODO(), 2*time.Second)
	defer cancel()

	statement := `update users set
			first_name=$1,
			last_name=$2,
			username=$3,
			password=$4,
			email=$5
			where id=$6
	`

	_, err := repo.db.ExecContext(ctx, statement,
		user.Fname,
		user.Lname,
		user.Username,
		user.Password,
		user.Email,
		user.Id,
	)
	if err != nil {
		return err
	}
	return nil
}
