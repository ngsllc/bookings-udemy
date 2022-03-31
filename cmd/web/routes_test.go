package main

import (
	"bookings-udemy/internal/config"
	"fmt"
	"testing"

	"github.com/go-chi/chi"
)

func TestRoutes(t *testing.T) {
	var app config.AppConfig

	mux := routes(&app)

	switch v := mux.(type) {
	case *chi.Mux:
		// do nothing, this is what we want
	default:
		t.Error(fmt.Sprintf("type is not chi.Mux, type is %v", v))
	}
}
