package service

import (
	"fmt"
	"os"
	"text/tabwriter"

	structs "library-management/struct"

	"github.com/google/uuid"
)

func ListPersonService(persons map[uuid.UUID]structs.PersonStruct) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', tabwriter.Debug)
	fmt.Fprintln(w, "ID\tName\tEmail")
	fmt.Fprintln(w, "--\t----\t-----")
	for _, person := range persons {
		fmt.Fprintf(w, "%s\t%s\t%s\n", person.ID, person.Name, person.Email)
	}
	w.Flush()
}
