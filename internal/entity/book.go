package entity

type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var Books []Book

func AddBook(book Book) {
	Books = append(Books, book)
}

func UpdateBook(updatedBook Book) {
	for i, book := range Books {
		if book.ID == updatedBook.ID {
			Books[i] = updatedBook
			break
		}
	}
}

func DeleteBook(id string) {
	for i, book := range Books {
		if book.ID == id {
			Books = append(Books[:i], Books[i+1:]...)
			break
		}
	}
}
