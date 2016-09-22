package fillpdf

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"log"
)

func TestTheAbilityToAddAPassword(t *testing.T) {

	form := Form{
		"field_1": "Hello",
		"field_2": "World",
	}

 	// err := FillAndEncode(form, "./sample/form.pdf", "./outputs/filled-pass", "Password", "UserPassword", true)
	err := Fill(form, "./sample/form.pdf", "./outputs/filled-test.pdf", "Password", "UserPassword", true)
	// err := Fill(form, "./sample/form.pdf", "./outputs/filled-test.pdf", true)
	if err != nil {
		log.Fatal(err)
	}

	assert.Nil(t, 100, "Test")
}
