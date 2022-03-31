package main

import (
	"fmt"
	"net/http"
	"testing"
)

func TestNoSurf(t *testing.T) {
	var myH myHandler
	h := NoSurf(&myH)
	switch v := h.(type) {
	case http.Handler:
		// do nothing, this is what we want
	default:
		t.Error(fmt.Sprintf("type of %t is not http.Handler", v))
	}
}

func TestSessionLoad(t *testing.T) {
	var myH myHandler
	h := SessionLoad(&myH)
	switch v := h.(type) {
	case http.Handler:
		// do nothing, this is what we want
	default:
		t.Error(fmt.Sprintf("type of %t is not http.Handler", v))
	}
}
