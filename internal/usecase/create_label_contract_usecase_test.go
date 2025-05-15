package usecase

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"os"
	"testing"
	"text/template"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/venture-technology/venture/internal/domain/service/adapters"
	"github.com/venture-technology/venture/internal/value"
	"github.com/venture-technology/venture/mocks"
	"github.com/venture-technology/venture/pkg/realtime"
	"github.com/venture-technology/venture/pkg/stringcommon"
)

var (
	body = value.CreateContractParams{
		AmountCents:      500000,
		AmountAnualCents: 6000000,
		DriverAmount:     100000,
		UUID:             "a1b2c3d4-e5f6-7890-abcd-1234567890ef",
		ResponsibleCPF:   "12345678901",
		ResponsibleName:  "John Doe",
		ResponsibleAddr:  "123 Main St, Springfield, IL 62701",
		ResponsibleEmail: "john.doe@example.com",
		ResponsiblePhone: "+1-555-123-4567",
		KidRG:            "987654321",
		KidName:          "Emma Smith",
		KidShift:         "Morning",
		DriverName:       "Michael Johnson",
		DriverCNH:        "XYZ123456789",
		SchoolCNPJ:       "12345678000195",
		SchoolName:       "Springfield Elementary",
		SchoolAddr:       "456 School Rd, Springfield, IL 62702",
		FileURL:          "https://s3.example.com/contracts/a1b2c3d4-e5f6-7890-abcd-1234567890ef.pdf",
	}
	html = []byte("<html><body>Contrato</body></html>")
	pdf  = []byte("%PDF-1.4\n%FakePDFContent\n")
)

func TestCreateLabelContractUsecase_Execute(t *testing.T) {
	tests := []struct {
		name          string
		requestParams value.CreateContractParams
		setup         func(t *testing.T) *CreateLabelContractUsecase
		wantErr       bool
		msgErr        string
	}{
		{
			name:          "when build contract return error",
			requestParams: body,
			setup: func(t *testing.T) *CreateLabelContractUsecase {
				logger := mocks.NewLogger(t)
				as := mocks.NewAgreementService(t)
				s3 := mocks.NewS3Iface(t)
				queue := mocks.NewQueue(t)
				converter := mocks.NewConverters(t)

				as.On("BuildContract", mock.Anything).Return(nil, errors.New("build contract error"))

				return NewCreateLabelContractUsecase(
					logger,
					adapters.Adapters{
						AgreementService: as,
					},
					s3,
					queue,
					converter,
				)
			},
			wantErr: true,
			msgErr:  "build contract error",
		},
		{
			name:          "when converter return error",
			requestParams: body,
			setup: func(t *testing.T) *CreateLabelContractUsecase {
				logger := mocks.NewLogger(t)
				as := mocks.NewAgreementService(t)
				s3 := mocks.NewS3Iface(t)
				queue := mocks.NewQueue(t)
				converter := mocks.NewConverters(t)

				as.On("BuildContract", mock.Anything).
					Return(html, nil)

				converter.On("ConvertHTMLtoPDF", html, mock.Anything).
					Return(nil, errors.New("pdf create error"))

				return NewCreateLabelContractUsecase(
					logger,
					adapters.Adapters{
						AgreementService: as,
					},
					s3,
					queue,
					converter,
				)
			},
			wantErr: true,
			msgErr:  "pdf create error",
		},
		{
			name:          "when s3 return error",
			requestParams: body,
			setup: func(t *testing.T) *CreateLabelContractUsecase {
				logger := mocks.NewLogger(t)
				as := mocks.NewAgreementService(t)
				s3 := mocks.NewS3Iface(t)
				queue := mocks.NewQueue(t)
				converter := mocks.NewConverters(t)

				as.On("BuildContract", mock.Anything).
					Return(html, nil)

				converter.On("ConvertHTMLtoPDF", html, mock.Anything).
					Return(pdf, nil)

				s3.On(
					"Save",
					mock.Anything,
					mock.Anything,
					mock.Anything,
					mock.Anything,
					mock.Anything,
				).Return("", errors.New("save error"))

				return NewCreateLabelContractUsecase(
					logger,
					adapters.Adapters{
						AgreementService: as,
					},
					s3,
					queue,
					converter,
				)
			},
			wantErr: true,
			msgErr:  "save error",
		},
		{
			name:          "when sqs return error",
			requestParams: body,
			setup: func(t *testing.T) *CreateLabelContractUsecase {
				logger := mocks.NewLogger(t)
				as := mocks.NewAgreementService(t)
				s3 := mocks.NewS3Iface(t)
				queue := mocks.NewQueue(t)
				converter := mocks.NewConverters(t)

				as.On("BuildContract", mock.Anything).
					Return(html, nil)

				converter.On("ConvertHTMLtoPDF", html, mock.Anything).
					Return(pdf, nil)

				s3.On(
					"Save",
					mock.Anything,
					mock.Anything,
					mock.Anything,
					mock.Anything,
					mock.Anything,
				).Return("path", nil)

				queue.On("SendMessage", mock.Anything, mock.Anything).
					Return(errors.New("sqs error"))

				return NewCreateLabelContractUsecase(
					logger,
					adapters.Adapters{
						AgreementService: as,
					},
					s3,
					queue,
					converter,
				)
			},
			wantErr: true,
			msgErr:  "sqs error",
		},
		{
			name:          "when there is success",
			requestParams: body,
			setup: func(t *testing.T) *CreateLabelContractUsecase {
				logger := mocks.NewLogger(t)
				as := mocks.NewAgreementService(t)
				s3 := mocks.NewS3Iface(t)
				queue := mocks.NewQueue(t)
				converter := mocks.NewConverters(t)

				as.On("BuildContract", mock.Anything).
					Return(html, nil)

				converter.On("ConvertHTMLtoPDF", html, mock.Anything).
					Return(pdf, nil)

				s3.On(
					"Save",
					mock.Anything,
					mock.Anything,
					mock.Anything,
					mock.Anything,
					mock.Anything,
				).Return("path", nil)

				queue.On("SendMessage", mock.Anything, mock.Anything).
					Return(nil)

				logger.On("Infof", mock.Anything, mock.Anything).Return().Once()

				return NewCreateLabelContractUsecase(
					logger,
					adapters.Adapters{
						AgreementService: as,
					},
					s3,
					queue,
					converter,
				)
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		msg, err := stringcommon.RawMessage(tt.requestParams)
		if err != nil {
			fmt.Println(err)
		}
		uc := tt.setup(t)
		err = uc.Execute(msg)
		if tt.wantErr {
			assert.Error(t, err)
			assert.EqualError(t, err, tt.msgErr)
			continue
		}
		assert.NoError(t, err)
	}

	/*
		Use this function to CreateLabel PDF on your machine, and verify the label output.

		err := createLabel(body)
		if err != nil {
			log.Fatalf("failed to create label: %v", err)
		}
	*/
}

func CreateLabel(input value.CreateContractParams) error {
	time := realtime.Now()
	input.Time = time
	input.DateTime = time.Format("02/01/2006")

	path, err := GetHtml()
	if err != nil {
		return err
	}

	htmlFile, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	tmpl, err := template.New("webpage").Funcs(template.FuncMap{
		"centsToReais": func(cents int64) float64 {
			return float64(cents) / 100
		},
	}).Parse(string(htmlFile))
	if err != nil {
		return err
	}

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, input)
	if err != nil {
		return err
	}

	pdf, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		return err
	}

	pdf.AddPage(wkhtmltopdf.NewPageReader(bytes.NewReader([]byte(buf.String()))))
	err = pdf.Create()
	if err != nil {
		return err
	}

	pdfData := pdf.Bytes()

	err = os.WriteFile("output.pdf", pdfData, 0644)
	if err != nil {
		log.Fatalf("failed to write file: %v", err)
	}

	log.Println("PDF saved to output.pdf")
	return nil
}
