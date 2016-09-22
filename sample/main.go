package main

import (
	"log"
	"../"
)

func main() {
	// Create the form values.
	form := fillpdf.Form{
		"field_1": "Hello",
		"field_2": "World",
	}

	// Fill the form PDF with our values.
	err := fillpdf.Fill(form, "form.pdf", "filled.pdf", true)
	if err != nil {
		log.Fatal(err)
	}


	err = fillpdf.FillAndEncode(form, "form.pdf", "filled-pass.pdf", "Password", "UserPassword", true)
	if err != nil {
		log.Fatal(err)
	}

}
