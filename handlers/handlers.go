package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	"github.com/osag1e/go-and-mysql/model"
	"github.com/osag1e/go-and-mysql/query"
)

type BooksHandler struct {
	DB        *sql.DB
	booksRepo query.BooksRepository
}

func NewBooksHandler(booksRepo query.BooksRepository) *BooksHandler {
	return &BooksHandler{
		booksRepo: booksRepo,
	}
}

func (bh *BooksHandler) HandleCreateBook(w http.ResponseWriter, r *http.Request) {
	var params model.Book
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	book, err := bh.booksRepo.InsertBook(&params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	render.JSON(w, r, book)
}

func (bh *BooksHandler) HandleFetchBooks(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	offsetParam := query.Get("offset")
	limitParam := query.Get("limit")

	offset := 0
	limit := 10

	if limitParam != "" {
		if l, err := strconv.Atoi(limitParam); err == nil {
			limit = l
		}
	}
	if offsetParam != "" {
		if o, err := strconv.Atoi(offsetParam); err == nil {
			offset = o
		}
	}
	books, err := bh.booksRepo.GetBooks(offset, limit)
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Books not found"})
		return
	}
	render.JSON(w, r, books)
}

func (bh *BooksHandler) HandleDeleteBook(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "bookID")
	bookID, err := uuid.Parse(id)
	if err != nil {
		http.Error(w, "Invalid UUID", http.StatusBadRequest)
		return
	}

	if err := bh.booksRepo.DeleteBookByID(bookID); err != nil {
		render.JSON(w, r, err)
		return
	}
	render.JSON(w, r, map[string]string{"deleted": id})
}
