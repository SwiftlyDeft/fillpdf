package fillpdf

import (
	"testing"
	"log"
)

func TestTheAbilityToAddAPasswordToAPDF(t *testing.T) {
	form := Form{
		"field_1": "Hello",
		"field_2": "World",
	}
	err := FillAndEncode(form, "sample/form.pdf", "sample/filled.pdf", "Password", "UserPassword", true)
	if err != nil {
		log.Fatal(err)
	}
}