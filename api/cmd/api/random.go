//Filename: cmd/api/random.go

package main

import (
	"fmt"
	"net/http"
)

func (app *application) showRandonString(w http.ResponseWriter, r *http.Request) {

	id, err := app.readIDParam(r)

	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	fmt.Fprintf(w, "id %d", id)
	// //create a new instance of the school struct containing the id we extracted from our url and some sample data

	// school := data.School{
	// 	ID:       id,
	// 	CreateAt: time.Now(),
	// 	Name:     "Apple Tree",
	// 	Level:    "High School",
	// 	Contact:  "Anna Smith",
	// 	Phone:    "601-4412",
	// 	Address:  "14 Apple Street",
	// 	Mode:     []string{"blended", "online"},
	// 	Version:  1,
	// }

	// err = app.writeJSON(w, http.StatusOK, envelope{"school": school}, nil)

	// if err != nil {
	// 	app.serverErrorResponse(w, r, err)
	// }
}
