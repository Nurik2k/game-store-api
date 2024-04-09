package handler

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct {
	favorite *FavoriteHandler
	genre *GenreHandler
	review *ReviewHandler
	game *GameHandler
	library *LibraryHandler
	platform *PlatformHandler
	publisher *PublisherHandler
}

func NewHandler(
	favorite *FavoriteHandler, 
	genre *GenreHandler, 
	review *ReviewHandler, 
	game *GameHandler, 
	library *LibraryHandler, 
	platform *PlatformHandler,
	publisher *PublisherHandler) *Handler {
	return &Handler{
		favorite: favorite,
		genre: genre,
		review : review,
		game: game,
		library: library,
		platform: platform,
		publisher: publisher,
	}
}

func enableCors(w http.ResponseWriter, method string) {
	w.Header().Set("Access-Control-Allow-Methods", method)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
}

func LogRequest(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        log.Printf("Request: %s %s", r.Method, r.URL.Path)
        next.ServeHTTP(w, r)
    })
}
func (h *Handler) Routes() *mux.Router {
	r:= mux.NewRouter()

	r.Use(LogRequest)

	r.HandleFunc("/api/v1/games", h.game.GetGames).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc("/api/v1/game/{id}", h.game.GetGame).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc("/api/v1/create-game", h.game.CreateGame).Methods(http.MethodPost, http.MethodOptions)
	r.HandleFunc("/api/v1/update-game/{id}", h.game.UpdateGame).Methods(http.MethodPut, http.MethodOptions)
	r.HandleFunc("/api/v1/delete-game/{id}", h.game.DeleteGame).Methods(http.MethodDelete, http.MethodOptions)

	r.HandleFunc("/api/v1/genres", h.genre.GetGenres).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc("/api/v1/genre/{id}", h.genre.GetGenre).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc("/api/v1/create-genre", h.genre.CreateGenre).Methods(http.MethodPost, http.MethodOptions)
	r.HandleFunc("/api/v1/update-genre/{id}", h.genre.UpdateGenre).Methods(http.MethodPut, http.MethodOptions)
	r.HandleFunc("/api/v1/delete-genre/{id}", h.genre.DeleteGenre).Methods(http.MethodDelete, http.MethodOptions)

	r.HandleFunc("/api/v1/library/{user_id}", h.library.GetLibrariesByUser).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc("/api/v1/create-library", h.library.CreateLibrary).Methods(http.MethodPost, http.MethodOptions)
	r.HandleFunc("/api/v1/add-library", h.library.AddGameToLibraryFromUser).Methods(http.MethodPost, http.MethodOptions)
	r.HandleFunc("/api/v1/delete-library/{id}", h.library.DeleteLibrary).Methods(http.MethodDelete, http.MethodOptions)

	r.HandleFunc("/api/v1/publishers", h.publisher.GetPublishers).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc("/api/v1/publisher/{id}", h.publisher.GetPublisher).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc("/api/v1/create-publisher", h.publisher.CreatePublisher).Methods(http.MethodPost, http.MethodOptions)
	r.HandleFunc("/api/v1/update-publisher/{id}", h.publisher.UpdatePublisher).Methods(http.MethodPut, http.MethodOptions)
	r.HandleFunc("/api/v1/delete-publisher/{id}", h.publisher.DeletePublisher).Methods(http.MethodDelete, http.MethodOptions)

	r.HandleFunc("/api/v1/platforms", h.platform.GetPlatforms).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc("/api/v1/platform/{id}", h.platform.GetPlatform).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc("/api/v1/create-platform", h.platform.CreatePlatform).Methods(http.MethodPost, http.MethodOptions)
	r.HandleFunc("/api/v1/update-platform/{id}", h.platform.UpdatePlatform).Methods(http.MethodPut, http.MethodOptions)
	r.HandleFunc("/api/v1/delete-platform/{id}", h.platform.DeletePlatform).Methods(http.MethodDelete, http.MethodOptions)

	r.HandleFunc("/api/v1/reviews", h.review.GetReviews).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc("/api/v1/review/{id}", h.review.GetReview).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc("/api/v1/create-review", h.review.CreateReview).Methods(http.MethodPost, http.MethodOptions)
	r.HandleFunc("/api/v1/update-review/{id}", h.review.UpdateReview).Methods(http.MethodPut, http.MethodOptions)
	r.HandleFunc("/api/v1/delete-review/{id}", h.review.DeleteReview).Methods(http.MethodDelete, http.MethodOptions)

	r.HandleFunc("/api/v1/favorite", h.favorite.GetFavorites).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc("/api/v1/favorite/{id}", h.favorite.AddFavorite).Methods(http.MethodPost, http.MethodOptions)
	r.HandleFunc("/api/v1/delete-favorite/{id}", h.favorite.DeleteFavorite).Methods(http.MethodDelete, http.MethodOptions)

	return r
}