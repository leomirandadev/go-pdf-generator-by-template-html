package main

import (
	"fmt"
	"generatorPDF/htmlParser"
	"generatorPDF/pdfGenerator"
	"os"
)

type Data struct {
	Name string
}

func main() {

	h := htmlParser.New("tmp")
	wk := pdfGenerator.NewWkHtmlToPDF("tmp")

	dataHTML := Data{
		Name: "Leonardo",
	}

	htmlGenerated, err := h.Create("templates/example.html", dataHTML)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer os.Remove(htmlGenerated)
	fmt.Println("HTML gerado", htmlGenerated)

	filePDFName, err := wk.Create(htmlGenerated)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("PDF gerado", filePDFName)

}
