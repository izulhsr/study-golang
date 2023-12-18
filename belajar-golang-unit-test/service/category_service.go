package service

import (
	"belajar-golang-unit-test/entity"
	"belajar-golang-unit-test/repository"
	"errors"
)

type CategoryService struct {
	Repository repository.CategoryRepository
}

func (service CategoryService) get(id string) (*entity.Category, error) {
	category := service.Repository.FindById(id)
	if category == nil {
		return category, errors.New("Category tidak ditemukan")
	} else {
		return category, nil
	}
}
