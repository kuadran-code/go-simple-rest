package usecases

import (
	"log"

	"github.com/kuadran-code/go-simple-rest/internal/entities"
	repositories "github.com/kuadran-code/go-simple-rest/internal/respositories"
	"gorm.io/gorm"
)

type ProductUsecase struct {
	productRepo repositories.ProductRepoContract
	name        string
}

type ProductUsecaseContract interface {
	// Create new product
	Create(name, description string, stock int) (entities.Product, error)
	// List of product
	Read() ([]entities.Product, error)
	// Detail of product
	Detail(ID int) (entities.Product, error)
	// Update existing product
	Update(ID int, name, description string, stock int) (entities.Product, error)
	// Delete product
	Delete(ID int) error
}

func NewProductUsecase(db *gorm.DB) ProductUsecaseContract {
	return &ProductUsecase{
		productRepo: repositories.NewProductRepository(db),
		name:        "Product Usecase",
	}
}

func (p *ProductUsecase) Create(name, description string, stock int) (entities.Product, error) {
	log.Printf("[%s][Create] is executed\n", p.name)
	product := entities.Product{
		Name:        name,
		Description: description,
		Stock:       stock,
	}

	if err := p.productRepo.Create(&product); err != nil {
		log.Printf("Error : [%s][Create] %s \n", p.name, err.Error())
		return product, err
	}

	return product, nil
}

func (p *ProductUsecase) Read() ([]entities.Product, error) {
	log.Printf("[%s][Read] is executed\n", p.name)

	products, _, err := p.productRepo.List()
	if err != nil {
		log.Printf("Error : [%s][Read] %s \n", p.name, err.Error())
		return products, err
	}

	return products, nil
}

func (p *ProductUsecase) Detail(ID int) (entities.Product, error) {
	log.Printf("[%s][Detail] is executed\n", p.name)

	product, err := p.productRepo.Get(ID)
	if err != nil {
		log.Printf("Error : [%s][Detail] %s \n", p.name, err.Error())
		return product, err
	}

	return product, nil
}

func (p *ProductUsecase) Update(ID int, name, description string, stock int) (entities.Product, error) {
	log.Printf("[%s][Update] is executed\n", p.name)
	product := entities.Product{
		ID:          ID,
		Name:        name,
		Description: description,
		Stock:       stock,
	}

	if err := p.productRepo.Update(&product); err != nil {
		log.Printf("Error : [%s][Update] %s \n", p.name, err.Error())
		return product, err
	}

	return product, nil
}

func (p *ProductUsecase) Delete(ID int) error {
	log.Printf("[%s][Delete] is executed\n", p.name)

	err := p.productRepo.Delete(ID)
	if err != nil {
		log.Printf("Error : [%s][Delete] %s \n", p.name, err.Error())
		return err
	}

	return nil
}
