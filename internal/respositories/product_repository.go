package repositories

import (
	"log"

	"github.com/kuadran-code/go-simple-rest/internal/entities"
	"gorm.io/gorm"
)

type ProductRepo struct {
	db   *gorm.DB
	name string
}

type ProductRepoContract interface {
	// Create a new Product
	Create(product *entities.Product) error
	// Update Product
	Update(product *entities.Product) error
	// List Product
	List() ([]entities.Product, int, error)
	// Get Product
	Get(ID int) (entities.Product, error)
	// Delete Product
	Delete(ID int) error
}

// Create new role product instance
func NewProductRepository(db *gorm.DB) ProductRepoContract {
	return &ProductRepo{
		db:   db,
		name: "Product Repository",
	}
}

func (r *ProductRepo) Create(product *entities.Product) error {
	log.Printf("[%s][Create] is executed\n", r.name)

	if err := r.db.Create(&product).Error; err != nil {
		log.Printf("Error : [%s][Create] %s\n", r.name, err.Error())
		return err
	}

	return nil
}

func (r *ProductRepo) Get(ID int) (entities.Product, error) {
	log.Printf("[%s][Get] is executed\n", r.name)

	db := r.db
	var product entities.Product

	if err := db.Debug().First(&product, ID).Error; err != nil {
		log.Printf("Error : [%s][GET] %s", r.name, err.Error())
		return product, err
	}

	return product, nil
}

func (r *ProductRepo) List() ([]entities.Product, int, error) {
	log.Printf("[%s][List] is executed\n", r.name)

	var count int64
	var products []entities.Product

	db := r.db

	if err := db.Debug().Find(&products).Count(&count).Error; err != nil {
		log.Printf("Error : [%s][List] %s", r.name, err.Error())
		return products, int(count), err
	}

	return products, int(count), nil
}

func (r *ProductRepo) Update(product *entities.Product) error {
	log.Printf("[%s][Update] is executed\n", r.name)

	if err := r.db.Model(&product).Updates(&product).Error; err != nil {
		log.Printf("Error : [%s][Update] %s", r.name, err.Error())
		return err
	}

	return nil
}

func (r *ProductRepo) Delete(ID int) error {
	log.Printf("[%s][Delete] is executed\n", r.name)

	var product entities.Product

	if err := r.db.Delete(&product, ID).Error; err != nil {
		log.Printf("Error : [%s][Delete] %s", r.name, err.Error())
		return err
	}

	return nil
}
