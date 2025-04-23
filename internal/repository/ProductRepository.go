package repository

import (
	"github.com/erkanattt/go-rest-crud-project/internal/models"
	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) GetAll() ([]models.Product, error) {
	var products []models.Product
	err := r.db.Find(&products).Error
	return products, err
}

func (r *ProductRepository) GetByUserID(userID uint) ([]models.Product, error) {
	var products []models.Product
	err := r.db.Where("user_id = ?", userID).Find(&products).Error
	return products, err
}

func (r *ProductRepository) GetById(id int) (*models.Product, error) {
	var product models.Product
	err := r.db.First(&product, id).Error
	return &product, err
}

func (r *ProductRepository) Create(product *models.Product) error {
	return r.db.Create(product).Error
}

func (r *ProductRepository) Update(id int, product *models.Product) error {
	return r.db.Model(&models.Product{}).Where("id = ?", id).Updates(product).Error
}

func (r *ProductRepository) Delete(id int) error {
	return r.db.Delete(&models.Product{}, id).Error
}
