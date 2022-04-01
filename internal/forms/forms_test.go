package forms

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestForm_Valid(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	isValid := form.Valid()
	if !isValid {
		t.Error("got invalid when expected valid")
	}
}

func TestForm_Required(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	form.Required("a", "b", "c")
	isValid := form.Valid()
	if isValid {
		t.Error("got valid when required fields missing ")
	}

	postedData := url.Values{}
	postedData.Add("a", "a")
	postedData.Add("b", "a")
	postedData.Add("c", "a")

	r, _ = http.NewRequest("POST", "/whatever", nil)

	r.PostForm = postedData
	form = New(r.PostForm)
	form.Required("a", "b", "c")
	if !form.Valid() {
		t.Error("shows required fields not filled when they really are")
	}
}

//still need tests for has, is email, and new
func TestForm_IsEmail(t *testing.T) {
	postedData := url.Values{}
	form := New(postedData)

	form.IsEmail("a")
	if form.Valid() {
		t.Error("got valid when an email was not entered")
	}

	postedData = url.Values{}
	postedData.Add("a", "cole@cole.com")
	postedData.Add("b", "uyvgbvg@iugheiu")
	form = New(postedData)

	form.IsEmail("a")
	if !form.Valid() {
		t.Error("got invalid when an email was entered")
	}

	form.IsEmail("b")
	if form.Valid() {
		t.Error("got valid when an invalid email was entered")
	}
}

func TestForm_Has(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	has := form.Has("whatever")
	if has {
		t.Error("form says it has something it does not")
	}

	postedData := url.Values{}
	postedData.Add("a", "a")
	form = New(postedData)

	has = form.Has("a")
	if !has {
		t.Error("form shows that it does not have something that it does have")
	}
}

func TestForm_MinLength(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	form.MinLength("x", 10)
	if form.Valid() {
		t.Error("shows minlength satisfied when no entry")
	}

	isError := form.Errors.Get("x")
	if isError == "" {
		t.Error("should have an error in errors")
	}

	postedValues := url.Values{}
	postedValues.Add("some_field", "some value")
	form = New(postedValues)

	form.MinLength("some_field", 100)
	if form.Valid() {
		t.Error("shows minlength satisfied when data shorter")
	}

	postedValues = url.Values{}
	postedValues.Add("another_field", "abc123")
	form = New(postedValues)

	form.MinLength("another_field", 1)
	if !form.Valid() {
		t.Error("minlength not satisfied when it should be")
	}

	isError = form.Errors.Get("x")
	if isError != "" {
		t.Error("should have an error in errors")
	}
}
