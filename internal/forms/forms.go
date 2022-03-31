package forms

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/asaskevich/govalidator"
)

type Form struct {
	url.Values
	Errors errors
}

// create a new form
func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

// checks that all of the given fields have entries
func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		value := f.Get(field)
		if (strings.TrimSpace(value)) == "" {
			f.Errors.Add(field, "This field cannot be empty")
		}
	}
}

// checks if form field is in post and nonempty
func (f *Form) Has(field string, r *http.Request) bool {
	x := r.Form.Get(field)
	if x == "" {
		return false
	}
	return true
}

// returns true if there are no errors, false o/w
func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}

// checks that a string is longer than the given length
func (f *Form) MinLength(field string, length int, r *http.Request) bool {
	x := r.Form.Get(field)
	if len(x) < length {
		f.Errors.Add(field, fmt.Sprintf("This field must be at least %d characters long", length))
		return false
	}
	return true
}

// checks that a given email is valid
func (f *Form) IsEmail(field string) {
	if !govalidator.IsEmail(f.Get(field)) {
		f.Errors.Add(field, "invalid email address")
	}
}
