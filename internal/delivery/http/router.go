package http

import "net/http"

func NewRouter(
	characterHandler *CharacterHandler,
	episodeHandler *EpisodeHandler,
	seasonHandler *SeasonHandler,
	quoteHandler *QuoteHandler,
) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/characters", characterHandler.Create)
	mux.HandleFunc("/episodes", episodeHandler.Create)
	mux.HandleFunc("/seasons", seasonHandler.Create)
	mux.HandleFunc("/quotes", quoteHandler.Create)

	mux.HandleFunc("/characters/", characterHandler.GetByID)
	mux.HandleFunc("/episodes/", episodeHandler.GetByID)
	mux.HandleFunc("/seasons/", seasonHandler.GetByID)
	mux.HandleFunc("/quotes/", quoteHandler.GetByID)

	return mux
}
