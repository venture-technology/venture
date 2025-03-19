package converters

import (
	"bytes"
	"html/template"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"github.com/venture-technology/venture/internal/entity"
)

type Converter struct {
}

func NewConverter() Converter {
	return Converter{}
}

func (c Converter) ConvertPDFtoHTML(htmlFile []byte, contractProperty entity.ContractProperty) ([]byte, error) {
	tmpl, err := template.New("webpage").Parse(string(htmlFile))
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, contractProperty)
	if err != nil {
		return nil, err
	}

	pdf, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		return nil, err
	}

	pdf.AddPage(wkhtmltopdf.NewPageReader(bytes.NewReader([]byte(buf.String()))))
	err = pdf.Create()
	if err != nil {
		return nil, err
	}

	pdfData := pdf.Bytes()

	return pdfData, nil
}
