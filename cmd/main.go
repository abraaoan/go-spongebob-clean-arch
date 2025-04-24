package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/lib/pq"

	deliveryHttp "github.com/abraaoan/go-spongebob-clean-arch/internal/delivery/http"
	"github.com/abraaoan/go-spongebob-clean-arch/internal/infrastructure/cache"
	"github.com/abraaoan/go-spongebob-clean-arch/internal/infrastructure/postgres"
	"github.com/abraaoan/go-spongebob-clean-arch/internal/usecase"
)

func main() {
	// Database
	connStr := "postgresql://bob:sponge@localhost:5432/bob_sponge?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Connect database error %v", err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Fatalf("Erro ao pingar banco: %v", err)
	}

	fmt.Println("✅ Banco conectado!")

	// Cache
	defaultCacheTimeout := 5 * time.Minute
	characterCache := cache.NewCharacterCache(defaultCacheTimeout)
	episodeCache := cache.NewEpisodeCache(defaultCacheTimeout)
	seasonCache := cache.NewSeasonCache(defaultCacheTimeout)
	quoteCache := cache.NewQuoteCache(defaultCacheTimeout)

	// Repo
	characterRepo := postgres.NewCharacterPostgres(db)
	episodeRepo := postgres.NewEpisodePostgres(db)
	seasonRepo := postgres.NewSeasonPostgres(db)
	quoteRepo := postgres.NewQuotePostgres(db)

	// Use case
	characterUC := usecase.NewCharacterUseCase(characterRepo, characterCache)
	episodeUC := usecase.NewEpisodeUseCase(episodeRepo, episodeCache)
	seasonUC := usecase.NewSeasonUseCase(seasonRepo, seasonCache)
	quoteUC := usecase.NewQuoteUseCase(quoteRepo, quoteCache)

	fmt.Println("✅ Use Cases prontos!")

	// Handler
	characterHandler := deliveryHttp.NewCharacterHandler(characterUC)
	episodeHandler := deliveryHttp.NewEpisodeHandler(episodeUC)
	seasonHandler := deliveryHttp.NewSeasonHandler(seasonUC)
	quoteHandler := deliveryHttp.NewQuoteHandler(quoteUC)

	// Router
	router := deliveryHttp.NewRouter(characterHandler, episodeHandler, seasonHandler, quoteHandler)
	http.ListenAndServe(":8080", router)
}
