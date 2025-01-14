package usecase

import (
	"fmt"

	"github.com/skip2/go-qrcode"
	"github.com/venture-technology/venture/internal/entity"
	"github.com/venture-technology/venture/internal/infra/contracts"
	"github.com/venture-technology/venture/internal/infra/persistence"
)

type CreateDriverUseCase struct {
	repositories *persistence.PostgresRepositories
	logger       contracts.Logger
	S3           contracts.S3Iface
}

func NewCreateDriverUseCase(
	repositories *persistence.PostgresRepositories,
	logger contracts.Logger,
	S3 contracts.S3Iface,
) *CreateDriverUseCase {
	return &CreateDriverUseCase{
		repositories: repositories,
		logger:       logger,
		S3:           S3,
	}
}

func (cduc *CreateDriverUseCase) CreateDriver(driver *entity.Driver) error {
	err := driver.ValidateCnh()
	if err != nil {
		return err
	}

	qrCode, err := qrcode.Encode(
		fmt.Sprintf(
			"https://venture-technology.xyz/driver/%s",
			driver.CNH,
		),
		qrcode.Medium,
		256,
	)
	if err != nil {
		return err
	}

	url, err := cduc.S3.Save(driver.CNH, "qrcode", qrCode)
	if err != nil {
		return err
	}

	driver.QrCode = url
	return cduc.repositories.DriverRepository.Create(driver)
}
