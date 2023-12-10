package repository

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type postgresqlCategoryRepository struct {
	connectionPool *sql.DB
}

func NewPostgresqlCategoryRepository(DB_URL string) *postgresqlCategoryRepository {
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

	return &postgresqlCategoryRepository{
		connectionPool: connectionPool,
	}
}