package handler

import (
	"encoding/json"
	"io/ioutil"
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

type TokenInfo struct {
    Status  string   `json:"status"`
    Message string   `json:"message"`
    UserID  string   `json:"user_id"`
    Roles   []string `json:"roles"`
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

func RequireAuth(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Authorization") == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		tokenInfo := verifyToken(r.Header.Get("Authorization"))
		if tokenInfo.Status != "success" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (h *Handler) Routes() *mux.Router {
	r:= mux.NewRouter()

	r.Use(LogRequest)
	
	subRouterGame := r.PathPrefix("/api/v1/game").Subrouter()
	subRouterGenre := r.PathPrefix("/api/v1/genre").Subrouter()
	subRouterLibrary := r.PathPrefix("/api/v1/library").Subrouter()
	subRouterPlatform := r.PathPrefix("/api/v1/platform").Subrouter()
	subRouterPublisher := r.PathPrefix("/api/v1/publisher").Subrouter()
	subRouterFavorite := r.PathPrefix("/api/v1/favorite").Subrouter()
	subRouterReview := r.PathPrefix("/api/v1/review").Subrouter()

	h.game.Routes(subRouterGame)
	h.genre.Routes(subRouterGenre)
	h.library.Routes(subRouterLibrary)
	h.platform.Routes(subRouterPlatform)
	h.publisher.Routes(subRouterPublisher)
	h.favorite.Routes(subRouterFavorite)
	h.review.Routes(subRouterReview)

	return r
}

func verifyToken(token string) TokenInfo {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://localhost:8080/api/v1/verify-token", nil)
	if err != nil {
		log.Println("Failed to create request for token verification:", err)
		return TokenInfo{Status: "error", Message: "Failed to create request for token verification"}
	}

	req.Header.Add("Authorization", token)

	resp, err := client.Do(req)
	if err != nil {
		log.Println("Failed to verify token:", err)
		return TokenInfo{Status: "error", Message: "Failed to verify token"}
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Failed to read token verification response body:", err)
		return TokenInfo{Status: "error", Message: "Failed to read token verification response body"}
	}

	var tokenInfo TokenInfo
	err = json.Unmarshal(body, &tokenInfo)
	if err != nil {
		log.Println("Failed to unmarshal token verification response:", err)
		return TokenInfo{Status: "error", Message: "Failed to unmarshal token verification response"}
	}

	return tokenInfo
}