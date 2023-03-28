package server

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"gorm.io/gorm"

	"time"

	"github.com/kuadran-code/go-simple-rest/internal/router"
)

type httpServer struct {
	router http.Handler
}

type ServerContract interface {
	Run(ctx context.Context, port int)
	Done()
}

func NewServer(db *gorm.DB) ServerContract {
	// Create new router
	rtr := router.NewRouter(db)
	return &httpServer{
		router: rtr.Route(),
	}
}

// Run server
func (h *httpServer) Run(ctx context.Context, port int) {
	log.Printf("Server running on port %d. Access it from http://127.0.0.1:%d\n", port, port)
	server := http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%d", port),
		Handler: h.router,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Printf("http server got %v", err.Error())
		}
	}()

	<-ctx.Done()

	ctxShutDown, cancel := context.WithTimeout(context.Background(), 6*time.Second)
	defer func() {
		cancel()
	}()

	if err := server.Shutdown(ctxShutDown); err != nil {
		log.Printf("server shutdown failed: %v", err)
		return
	}

	log.Printf("server existed properly")
}

func (h *httpServer) Done() {
	log.Fatal("Server closed")
}
