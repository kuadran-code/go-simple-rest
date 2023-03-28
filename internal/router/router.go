package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kuadran-code/go-simple-rest/internal/handler"
	"gorm.io/gorm"
)

type router struct {
	router *mux.Router
	db     *gorm.DB
}

type RouterContract interface {
	Route() http.Handler
}

func NewRouter(db *gorm.DB) RouterContract {
	return &router{
		router: mux.NewRouter(),
		db:     db,
	}
}

func (r *router) Route() http.Handler {
	productHandler := handler.NewProductHandler(r.db)
	r.router.HandleFunc("/products", productHandler.Create).Methods("POST")
	r.router.HandleFunc("/products", productHandler.List).Methods("GET")
	r.router.HandleFunc("/products/{id}", productHandler.Update).Methods("PUT")
	r.router.HandleFunc("/products/{id}", productHandler.Detail).Methods("GET")
	r.router.HandleFunc("/products/{id}", productHandler.Delete).Methods("DELETE")

	return r.router
}
