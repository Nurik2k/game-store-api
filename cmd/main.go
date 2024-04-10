package main

import (
	"context"
	"game-store-api/database"
	"game-store-api/handler"
	"game-store-api/service"
	"log"
	"net/http"

	"github.com/jackc/pgx/v4/pgxpool"
)

func main() {
	connect, err := ConnectDB()
	if err != nil {
		log.Fatal(err.Error())
	}

	favoritedb, _ := database.NewFavoriteDb(connect)
	genredb, _ := database.NewGenreDb(connect)
	librarydb, _ := database.NewLibraryDb(connect)
	platformdb, _ := database.NewPlatformDb(connect)
	reviewdb, _ := database.NewReviewDb(connect)
	gamedb, _ := database.NewGameDb(connect)
	publisherdb, _ := database.NewPublisherDb(connect)

	favoriteService := service.NewFavoriteService(favoritedb)
	genreService := service.NewGenreService(genredb)
	libraryService := service.NewLibraryService(librarydb)
	platformService := service.NewPlatformService(platformdb)
	reviewService := service.NewReviewService(reviewdb)
	gameService := service.NewGameService(gamedb)
	publisherService := service.NewPublisherService(publisherdb)

	handler := handler.NewHandler(
		handler.NewFavoriteHandler(favoriteService),
		handler.NewGenreHandler(genreService),
		handler.NewReviewHandler(reviewService),
		handler.NewGameHandler(gameService),
		handler.NewLibraryHandler(libraryService),
		handler.NewPlatformHandler(platformService),
		handler.NewPublisherHandler(publisherService),
	)

	r := handler.Routes()

	err = http.ListenAndServe(":8085", r)
	if err != nil {
		log.Fatal(err.Error())
	}


}

func ConnectDB() (*pgxpool.Pool, error) {
	ctx := context.Background()

	config, err := pgxpool.ParseConfig("postgres://game_user:p@55word@localhost:5432/game-data?sslmode=disable")
	if err != nil {
		return nil, err
	}
	config.MaxConns = 100

	pool, err := pgxpool.ConnectConfig(ctx, config)
	if err != nil {
		return nil, err
	}

	return pool, nil
}