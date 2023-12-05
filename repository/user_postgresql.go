package repository

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/varomnrg/money-tracker/model"
)

type postgresqlUserRepository struct {
	connectionPool *sql.DB
}

func NewPostgresqlUserRepository(DB_URL string) *postgresqlUserRepository {
	connString := DB_URL

	connectionPool, err := sql.Open("postgres", connString)

	if err != nil {
		log.Fatal(err)
	}

	err = connectionPool.Ping()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected to database!")

	return &postgresqlUserRepository{
		connectionPool: connectionPool,
	}
}

func (p *postgresqlUserRepository) GetUsers() ([]model.UserResponse, error) {
	rows, err := p.connectionPool.Query("SELECT id, username, email, created_at FROM users")

	if err != nil {
		return []model.UserResponse{}, err
	}
	defer rows.Close()

	users := make([]model.UserResponse, 0)

	for rows.Next() {
		user := model.UserResponse{}
		err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.Created_At)
		if err != nil {
			return users, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (p *postgresqlUserRepository) GetUser(id string) (model.UserResponse, error) {
	user := model.UserResponse{}

	row := p.connectionPool.QueryRow("SELECT id, username, email, created_at FROM users WHERE id = $1", id)

	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Created_At)

	if err != nil {
		return user, err
	}

	return user, nil
}

func (p *postgresqlUserRepository) CreateUser(movie model.User) error {
	_, err := p.connectionPool.Exec(
		"INSERT INTO users (id, username, password, email, created_at) VALUES ($1, $2, $3, $4, $5)",
		movie.ID, movie.Username, movie.Password, movie.Email, movie.Created_At,
	)

	if err != nil {
		return err
	}

	return nil
}

func (p *postgresqlUserRepository) DeleteUser(id string) error {
	_, err := p.connectionPool.Exec(
		"DELETE FROM users WHERE id = $1",
		id,
	)

	if err != nil {
		return err
	}

	return nil
}

func (p *postgresqlUserRepository) UpdateUser(id string, movie model.UserRequest) error {
	_, err := p.connectionPool.Exec(
		"UPDATE users SET username = $1, password = $2, email = $3 WHERE id = $4",
		movie.Username, movie.Password, movie.Email, id,
	)

	if err != nil {
		return err
	}

	return nil
}

func (p *postgresqlUserRepository) IsUsernameExist(username string) bool {
	row := p.connectionPool.QueryRow("SELECT username FROM users WHERE username = $1", username)

	err := row.Scan(&username)

	return err == nil
}
