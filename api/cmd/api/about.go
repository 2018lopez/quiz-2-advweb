//Filename /cmd/api/healthcheck.go

package main

import (
	"net/http"

	"quiz-2.imerlopez.net/internal/data"
)

func (app *application) showAboutMe(w http.ResponseWriter, r *http.Request) {
	//create a map to hold our healthcheck data

	data := data.Imer{
		Name:          "Imer Lopez",
		Lived:         "San Ignacio Town, Cayo",
		CurrentSchool: "University of Belize",
		Study:         " Bachelor's in Information Technology",
		Email:         "2018119251@ub.edu.bz",
		Hobbies:       []string{"Reading", "Football", "Outdoors adventures"},
		Interests:     []string{"Cyber security", "FullStack Development", "Robotics"},
	}

	err := app.writeJSON(w, http.StatusOK, envelope{"data": data}, nil)

	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

}
