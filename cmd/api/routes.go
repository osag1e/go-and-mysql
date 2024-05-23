package main

import (
	"database/sql"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/osag1e/go-and-mysql/handlers"
	"github.com/osag1e/go-and-mysql/query"
)

func initializeRouter(dbConn *sql.DB) http.Handler {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	booksRepo := query.NewBooksStore(dbConn)
	booksHandler := handlers.NewBooksHandler(booksRepo)

	router.Post("/book", booksHandler.HandleCreateBook)
	router.Get("/books", booksHandler.HandleFetchBooks)
	router.Delete("/book/{bookID}", booksHandler.HandleDeleteBook)

	return router
}
