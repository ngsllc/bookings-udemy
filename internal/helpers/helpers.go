package helpers

import (
	"bookings-udemy/internal/config"
	"fmt"
	"net/http"
	"runtime/debug"
)

var app *config.AppConfig

// sets up appconfig for helpers
func NewHelpers(a *config.AppConfig) {
	app = a
}

// writes an error to the log
func ClientError(w http.ResponseWriter, status int) {
	app.InfoLog.Println("Client error with status:", status)
	http.Error(w, http.StatusText(status), status)
}

// writes an error to the log
func ServerError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.ErrorLog.Println(trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

// checks that the user_id has been added to the session
func IsAuthenticated(r *http.Request) bool {
	exists := app.Session.Exists(r.Context(), "user_id")
	return exists
}
