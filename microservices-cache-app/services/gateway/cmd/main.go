package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/rwrwdss/microservices-cache-app/services/gateway/internal/builder"
	"github.com/rwrwdss/microservices-cache-app/services/gateway/internal/config"
	"github.com/rwrwdss/microservices-cache-app/services/gateway/internal/database"
	"github.com/rwrwdss/microservices-cache-app/services/gateway/internal/handler"
	"github.com/rwrwdss/microservices-cache-app/services/gateway/internal/repository"
	"github.com/rwrwdss/microservices-cache-app/services/gateway/internal/router"
	"github.com/rwrwdss/microservices-cache-app/services/gateway/internal/searchclient"
	"github.com/rwrwdss/microservices-cache-app/services/gateway/internal/service"
)

const shutdownTimeout = 5 * time.Second

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	log.Println("INFO Connecting to PostgreSQL...")
	db, err := database.Connect(cfg)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	log.Println("INFO Connected to PostgreSQL")

	log.Println("INFO Running migrations...")
	if err := repository.Migrate(context.Background(), db); err != nil {
		log.Fatalf("failed to run migrations: %v", err)
	}

	log.Println("INFO Seeding dictionary...")
	if err := repository.Seed(context.Background(), db); err != nil {
		log.Fatalf("failed to seed dictionary: %v", err)
	}

	keywordRepository := repository.NewPostgresKeywordRepository(db)
	dictionaryBuilder := builder.NewTextDictionaryBuilder(cfg.DictionaryPath)
	dictionaryService := service.NewDictionaryService(keywordRepository, dictionaryBuilder)

	log.Println("INFO Loading dictionary...")
	if err := dictionaryService.LoadDictionary(context.Background()); err != nil {
		log.Fatalf("failed to load dictionary: %v", err)
	}

	log.Println("INFO Connecting to search-core...")
	searchClient, err := searchclient.NewGRPCSearchClient(cfg.SearchCoreGRPCAddr)
	if err != nil {
		log.Fatalf("failed to create search-core client: %v", err)
	}
	if _, err := searchClient.Search(context.Background(), "", 1); err != nil {
		log.Fatalf("failed to reach search-core: %v", err)
	}
	log.Println("INFO Connected to search-core")

	searchService := service.NewSearchService(searchClient)
	searchHandler := handler.NewSearchHandler(searchService)
	r := router.New(searchHandler)

	server := &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: r,
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	go func() {
		log.Printf("INFO HTTP server started on :%s", cfg.Port)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("server failed: %v", err)
		}
	}()

	<-ctx.Done()
	log.Println("shutdown signal received, stopping server...")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("graceful shutdown failed: %v", err)
	}

	log.Println("server stopped gracefully")
}
