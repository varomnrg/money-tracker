package repository

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/varomnrg/money-tracker/model"
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

func (p *postgresqlCategoryRepository) GetCategories() ([]model.Category, error){
	rows, err := p.connectionPool.Query("SELECT id, name, user_id FROM categories")
	
	if err  != nil {
		return []model.Category{}, err
	}
	
	defer rows.Close()

	categories := make([]model.Category, 0)

	for rows.Next(){
		category := model.Category{}
		err := rows.Scan(&category.ID, &category.Name, &category.User_ID)
		if err != nil {
			return categories, err
		}
		categories = append(categories, category)
	}
	
	return categories, nil
}

func (p *postgresqlCategoryRepository) GetUserCategories(userID string) ([]model.Category, error){
	rows, err := p.connectionPool.Query("SELECT id, name, user_id FROM categories WHERE user_id = $1", userID)

	if err != nil {
		return []model.Category{}, err
	}

	defer rows.Close()

	categories := make([]model.Category, 0)

	for rows.Next() {
		category := model.Category{}
		err := rows.Scan(&category.ID, &category.Name, &category.User_ID)
		if err != nil {
			return categories, err
		}
		categories = append(categories, category)
	}

	return categories, nil
}

func (p *postgresqlCategoryRepository) GetCategory(id string) (model.Category, error){
	category := model.Category{}

	row := p.connectionPool.QueryRow("SELECT id, name, user_id FROM categories WHERE id = $1", id)

	err := row.Scan(&category.ID, &category.Name, &category.User_ID)

	if err != nil {
		return category, err
	}

	return category, nil
}

func (p *postgresqlCategoryRepository) CreateCategory(category model.Category) error{
	_, err := p.connectionPool.Exec(
		"INSERT INTO categories (id, name, user_id) VALUES ($1, $2, $3)",
		category.ID, category.Name, category.User_ID,
	)

	if err != nil {
		return err
	}

	return nil
}

func (p *postgresqlCategoryRepository) DeleteCategory(id string) error{
	_, err := p.connectionPool.Exec(
		"DELETE FROM categories WHERE id = $1",
		id,
	)

	if err != nil {
		return err
	}

	return nil
}

func (p *postgresqlCategoryRepository) IsUserCategoryExist(userID string, categoryName string) bool{
	row := p.connectionPool.QueryRow(
		"SELECT id FROM categories WHERE user_id = $1 AND name = $2", 
		userID, categoryName,
	)

	category := model.Category{}

	err := row.Scan(&category.ID)

	return err == nil
}