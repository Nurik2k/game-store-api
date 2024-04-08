package handler

import (
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

func (h *Handler) Routes() *mux.Router {
	r:= mux.NewRouter()

	r.PathPrefix("/api/v1/favorite-games").Handler(h.favorite.favoriteRoutes())
	r.PathPrefix("/api/v1/genres").Handler(h.genre.genreRoutes())
	r.PathPrefix("/api/v1/reviews").Handler(h.review.reviewRoutes())
	r.PathPrefix("/api/v1/games").Handler(h.game.gameRoutes())
	r.PathPrefix("/api/v1/library").Handler(h.library.libraryRoutes())
	r.PathPrefix("/api/v1/platforms").Handler(h.platform.platformRoutes())
	r.PathPrefix("/api/v1/publishers").Handler(h.publisher.publisherRoutes())
	return r
}