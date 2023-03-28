package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/kuadran-code/go-simple-rest/internal/params"
	"github.com/kuadran-code/go-simple-rest/internal/usecases"
	jsonutils "github.com/kuadran-code/go-simple-rest/utils/json"
	"gorm.io/gorm"
)

type ProductHandler struct {
	productUsecase usecases.ProductUsecaseContract
	name           string
	handler        Handler
}

type ProductHandlerContract interface {
	// Create new product
	Create(w http.ResponseWriter, r *http.Request)
	// Update product
	Update(w http.ResponseWriter, r *http.Request)
	// List of product
	List(w http.ResponseWriter, r *http.Request)
	// Detail of product
	Detail(w http.ResponseWriter, r *http.Request)
	// Delete product
	Delete(w http.ResponseWriter, r *http.Request)
}

func NewProductHandler(db *gorm.DB) ProductHandlerContract {
	return &ProductHandler{
		productUsecase: usecases.NewProductUsecase(db),
		name:           "Product Handler",
	}
}

func (p *ProductHandler) Create(w http.ResponseWriter, r *http.Request) {
	log.Printf("[%s][Create] is executed\n", p.name)

	var param params.ProductCreateParam
	if err := jsonutils.Decode(r.Body, &param); err != nil {
		log.Printf("Error : [%s][Create] %s \n", p.name, err.Error())
		p.handler.ResponseError(w, err.Error(), http.StatusBadRequest)
		return
	}

	product, err := p.productUsecase.Create(param.Name, param.Description, param.Stock)
	if err != nil {
		p.handler.ResponseError(w, err.Error(), http.StatusBadRequest)
		return
	}

	p.handler.Response(w, product, http.StatusOK)
}

func (p *ProductHandler) Update(w http.ResponseWriter, r *http.Request) {
	log.Printf("[%s][Update] is executed\n", p.name)
	urlParam := mux.Vars(r)
	id := urlParam["id"]
	idx, _ := strconv.Atoi(id)

	var param params.ProductCreateParam
	if err := jsonutils.Decode(r.Body, &param); err != nil {
		log.Printf("Error : [%s][Update] %s \n", p.name, err.Error())
		p.handler.ResponseError(w, err.Error(), http.StatusBadRequest)
		return
	}

	product, err := p.productUsecase.Update(idx, param.Name, param.Description, param.Stock)
	if err != nil {
		p.handler.ResponseError(w, err.Error(), http.StatusBadRequest)
		return
	}

	p.handler.Response(w, product, http.StatusOK)
}

func (p *ProductHandler) List(w http.ResponseWriter, r *http.Request) {
	log.Printf("[%s][List] is executed\n", p.name)

	products, err := p.productUsecase.Read()
	if err != nil {
		p.handler.ResponseError(w, err.Error(), http.StatusBadRequest)
		return
	}

	p.handler.Response(w, products, http.StatusOK)
}

func (p *ProductHandler) Detail(w http.ResponseWriter, r *http.Request) {
	log.Printf("[%s][Detail] is executed\n", p.name)
	urlParam := mux.Vars(r)
	id := urlParam["id"]
	idx, _ := strconv.Atoi(id)

	product, err := p.productUsecase.Detail(idx)
	if err != nil {
		p.handler.ResponseError(w, err.Error(), http.StatusBadRequest)
		return
	}

	p.handler.Response(w, product, http.StatusOK)
}

func (p *ProductHandler) Delete(w http.ResponseWriter, r *http.Request) {
	log.Printf("[%s][Delete] is executed\n", p.name)
	urlParam := mux.Vars(r)
	id := urlParam["id"]
	idx, _ := strconv.Atoi(id)

	err := p.productUsecase.Delete(idx)
	if err != nil {
		p.handler.ResponseError(w, err.Error(), http.StatusBadRequest)
		return
	}

	p.handler.Response(w, "Product deleted successfully", http.StatusOK)
}
