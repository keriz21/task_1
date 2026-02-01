package repositories

import (
	"database/sql"
	"errors"
	"task_1/model"
)

type CategoriesRepository struct {
	db *sql.DB
}

func NewCategoryRepository(db *sql.DB) *CategoriesRepository {
	return &CategoriesRepository{db: db}
}

func (repo *CategoriesRepository) GetAll() ([]model.Categories, error) {
	query := "SELECT id, name, description FROM categories"
	rows, err := repo.db.Query(query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	categories := make([]model.Categories, 0)

	for rows.Next(){
		var p model.Categories
		err := rows.Scan(&p.ID, &p.Nama, &p.Description)
		if err != nil {
			return nil, err
		}
		categories = append(categories, p)
	}

	return categories, nil
}

func (repo *CategoriesRepository) Create(category *model.Categories) error {
	query := "INSERT INTO categories (name, description) VALUES ($1, $2) RETURNING id"
	err := repo.db.QueryRow(query, category.Nama, category.Description).Scan(&category.ID)
	return err
}

func (repo *CategoriesRepository) GetByID(id int) (*model.Categories, error) {
	query := "SELECT id, name, description FROM categories WHERE id = $1"

	var p model.Categories
	err := repo.db.QueryRow(query,id).Scan(&p.ID, &p.Nama, &p.Description)
	if err == sql.ErrNoRows{
		return nil, errors.New("produk tidak ditemukan")
	}

	if err != nil {
		return nil, err
	}

	return &p, nil
}

func (repo *CategoriesRepository) Update(category *model.Categories) error {
	query := "UPDATE categories SET name=$1, description=$2 WHERE id = $3"
	result, err := repo.db.Exec(query, category.Nama, category.Description, category.ID)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("produk tidak ditemukan")
	}

	return nil
}

func (repo *CategoriesRepository) Delete(id int) error {
	query := "DELETE FROM categories WHERE id = $1"
	result, err := repo.db.Exec(query, id)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("produk tidak ditemukan")
	}

	return err
}
