package service

import (
	"fmt"
	"os"
	"text/tabwriter"

	structs "library-management/struct"

	"github.com/google/uuid"
)

func ListBookService(books map[uuid.UUID]structs.BookStruct) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', tabwriter.Debug)
	fmt.Fprintln(w, "ID\tTitle\tAuthor\tStatus")
	fmt.Fprintln(w, "--\t-----\t------\t------")
	for _, book := range books {
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\n", book.ID, book.Name, book.Author, book.Status)
	}
	w.Flush()
}
