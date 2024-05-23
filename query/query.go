package query

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/osag1e/go-and-mysql/model"
)

type BooksRepository interface {
	InsertBook(book *model.Book) (*model.Book, error)
	GetBooks() ([]model.Book, error)
	DeleteBookByID(bookID uuid.UUID) error
}

type BooksStore struct {
	DB *sql.DB
}

func NewBooksStore(db *sql.DB) BooksRepository {
	return &BooksStore{
		DB: db,
	}
}

func (b *BooksStore) InsertBook(book *model.Book) (*model.Book, error) {
	bookID := model.NewUUID()
	query := "INSERT INTO books (id, title, price) VALUES (?, ?, ?)"
	_, err := b.DB.Exec(query, bookID.String(), book.Title, book.Price)
	if err != nil {
		return nil, fmt.Errorf("failed to insert book: %v", err)
	}
	book.ID = bookID
	return book, nil
}

func (b *BooksStore) GetBooks() ([]model.Book, error) {
	query := "SELECT id, title, price FROM books"
	rows, err := b.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to get books: %v", err)
	}
	defer rows.Close()

	var books []model.Book
	for rows.Next() {
		var book model.Book
		if err := rows.Scan(&book.ID, &book.Title, &book.Price); err != nil {
			return nil, fmt.Errorf("fialed to scan book: %v", err)
		}
		books = append(books, book)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %v", err)
	}
	return books, nil
}

func (b *BooksStore) DeleteBookByID(bookID uuid.UUID) error {
	query := "DELETE FROM books WHERE id = ?"
	_, err := b.DB.Exec(query, bookID)
	if err != nil {
		return fmt.Errorf("failed to delete book: %v", err)
	}
	return nil
}
