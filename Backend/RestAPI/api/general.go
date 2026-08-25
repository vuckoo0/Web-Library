package api

import (
	"fmt"
)

type Book struct {
	Id     int64  `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	ISBN   string `json:"isbn"`
}

func (b Book) String() string {
	return fmt.Sprintf("%d  | %s | %s | %s\n", b.Id, b.Title, b.Author, b.ISBN)
}
