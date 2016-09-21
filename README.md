# SwiftlyDeft's FillPDF Extension
This is an extension of https://github.com/desertbit/fillpdf, thank you to desertbit's work on fillpdf.
I needed to add some more functionality around FillPDF:
* Adding a password to the filled out pdf...

* Add a password to a filled out pdf *

# FillPDF

FillPDF is a golang library to easily fill PDF forms. This library uses the pdftk utility to fill the PDF forms with fdf data.
Currently this library only supports PDF text field values. Feel free to add support to more form types.


## Documentation 

Check the Documentation at [GoDoc.org](https://godoc.org/github.com/desertbit/fillpdf).


## Sample

There is an example in the sample directory:

```go
package main

import (
	"log"

	"github.com/desertbit/fillpdf"
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
}
```

Run the example as following:

```
cd sample
go build
./sample
```