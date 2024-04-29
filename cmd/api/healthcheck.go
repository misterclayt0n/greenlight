package main

import (
	"net/http"
)

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	env := envelope{
		"status": "available",
		"system_info": map[string]string{
			"enviroment": app.config.env,
			"version":    version,
		},
	}

	err := app.writeJson(w, http.StatusOK, env, nil)

	if err != nil {
		app.logger.Print(err)
		http.Error(w, "the server encountered a problem and could not process your request", http.StatusInternalServerError)
	}
}
