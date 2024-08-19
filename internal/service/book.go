package service

import "database/sql"

type Book struct {
	ID     int
	Title  string
	Author string
	Genre  string
}

type BookService struct {
	db *sql.DB
}

func NewBookService(db *sql.DB) *BookService {
	return &BookService{db}
}

func (s *BookService) CreateBook(book *Book) error {
	query := "INSERT INTO books (title, author, genre) VALUES ($1, $2, $3)"
	result, err := s.db.Exec(query, book.Title, book.Author, book.Genre)

	if err != nil {
		return err
	}

	lastInsertId, err := result.LastInsertId()

	if err != nil {
		return err
	}

	book.ID = int(lastInsertId)
	return nil
}

func (s *BookService) GetBooks() ([]Book, error) {
	query := "SELECT * FROM books"
	rows, err := s.db.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	books := []Book{}

	for rows.Next() {
		book := Book{}
		// & is used to get the address of the variable, so that the value can be stored in the variable
		// pega os dados do banco e joga na struct
		err = rows.Scan(&book.ID, &book.Title, &book.Author, &book.Genre)

		if err != nil {
			return nil, err
		}

		books = append(books, book)
	}

	return books, nil
}

func (s *BookService) GetBookById(id int) (*Book, error) {
	query := "SELECT * FROM books WHERE id = $1"
	row := s.db.QueryRow(query, id)

	book := &Book{}

	err := row.Scan(&book.ID, &book.Title, &book.Author, &book.Genre)

	if err != nil {
		return nil, err
	}

	return book, nil
}

func (s *BookService) UpdateBook(book *Book) error {
	query := "UPDATE books SET title = $1, author = $2, genre = $3 WHERE id = $4"
	_, err := s.db.Exec(query, book.Title, book.Author, book.Genre, book.ID)

	if err != nil {
		return err
	}

	return nil
}

func (s *BookService) DeleteBook(id int) error {
	query := "DELETE FROM books WHERE id = $1"
	_, err := s.db.Exec(query, id)

	if err != nil {
		return err
	}

	return nil
}
