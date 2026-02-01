package services

import (
	"task_1/model"
	"task_1/repositories"
)

type CategoryService struct {
	repo *repositories.CategoriesRepository
}

func NewCategoryService(repo *repositories.CategoriesRepository) *CategoryService {
	return &CategoryService{repo: repo}
}

func (s *CategoryService) GetAll() ([]model.Categories, error){
	return s.repo.GetAll()
}

func (s *CategoryService) Create(data *model.Categories) error {
	return s.repo.Create(data)
}

func (s *CategoryService) GetByID(id int) (*model.Categories, error) {
	return s.repo.GetByID(id)
}

func (s *CategoryService) Update(category *model.Categories) error {
	return s.repo.Update(category)
}

func (s *CategoryService) Delete(id int) error {
	return s.repo.Delete(id)
}