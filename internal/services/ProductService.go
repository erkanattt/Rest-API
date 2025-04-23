package services

import (
	"github.com/erkanattt/go-rest-crud-project/internal/models"
	"github.com/erkanattt/go-rest-crud-project/internal/repository"
)

type ProductService struct {
	repo *repository.ProductRepository
}

func NewProductService(repo *repository.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) GetAllProducts() ([]models.Product, error) {
	return s.repo.GetAll()
}

func (s *ProductService) GetProductsByUserID(userID uint) ([]models.Product, error) {
	return s.repo.GetByUserID(userID)
}

func (s *ProductService) GetProductByID(id int) (*models.Product, error) {
	return s.repo.GetById(id)
}

func (s *ProductService) CreateProduct(product *models.Product) error {
	return s.repo.Create(product)
}

func (s *ProductService) UpdateProduct(id int, product *models.Product) error {
	return s.repo.Update(id, product)
}

func (s *ProductService) DeleteProduct(id int) error {
	return s.repo.Delete(id)
}
