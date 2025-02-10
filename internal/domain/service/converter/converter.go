package converter

import (
	"bytes"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"github.com/venture-technology/venture/internal/entity"
	"github.com/venture-technology/venture/internal/infra/contracts"
)

type Converter struct {
	Logger contracts.Logger
}

func NewConverter(logger contracts.Logger) Converter {
	return Converter{
		Logger: logger,
	}
}

func (c Converter) ContractToHTML(contract entity.Contract) (bytes.Buffer, error) {
	return bytes.Buffer{}, nil
}

func (c Converter) HTMLtoPDF(html bytes.Buffer) (*wkhtmltopdf.PDFGenerator, error) {
	return nil, nil
}

func (c Converter) PDFtoBase64(pdf *wkhtmltopdf.PDFGenerator) (string, error) {
	return "", nil}

func (c Converter) Base64toPDF(base64 string) (*wkhtmltopdf.PDFGenerator, error) {
	return nil, nil
}
