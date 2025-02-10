package main

import (
	"bytes"
	"encoding/base64"
	"html/template"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

type TemplateData struct {
	Driver      string
	Responsible string
	Child       string
	Address     string
}

func main() {
	// Dados para substituir no template
	data := TemplateData{
		Driver:      "John Doe",
		Responsible: "Jane Doe",
		Child:       "Baby Doe",
		Address:     "123 Main St, City, Country",
	}

	// Carrega o template HTML
	tmpl, err := template.ParseFiles("template.html")
	if err != nil {
		panic(err)
	}

	// Gera o HTML final com os dados substituídos
	var htmlBuffer bytes.Buffer
	err = tmpl.Execute(&htmlBuffer, data)
	if err != nil {
		panic(err)
	}

	// Cria um novo gerador de PDF
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		panic(err)
	}

	// Adiciona o HTML ao gerador de PDF
	page := wkhtmltopdf.NewPageReader(bytes.NewReader(htmlBuffer.Bytes()))
	pdfg.AddPage(page)

	// Gera o PDF
	err = pdfg.Create()
	if err != nil {
		panic(err)
	}

	// Converte o PDF para base64
	pdfBase64 := base64.StdEncoding.EncodeToString(pdfg.Bytes())
	println("PDF em Base64:", pdfBase64)
}
