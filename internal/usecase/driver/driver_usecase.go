package driver

import (
	"context"
	"fmt"

	"github.com/skip2/go-qrcode"
	"github.com/venture-technology/venture/internal/entity"
	"github.com/venture-technology/venture/internal/repository"
)

type DriverUseCase struct {
	driverRepository repository.IDriverRepository
	awsRepository    repository.IAwsRepository
}

func NewDriverUseCase(dr repository.IDriverRepository, ar repository.IAwsRepository) *DriverUseCase {
	return &DriverUseCase{
		driverRepository: dr,
		awsRepository:    ar,
	}
}

func (du *DriverUseCase) Create(ctx context.Context, driver *entity.Driver) error {

	err := driver.ValidateCnh()

	if err != nil {
		return err
	}

	redirectUrl := fmt.Sprintf("https://venture-technology.xyz/driver/%s", driver.CNH)

	qrCode, err := qrcode.Encode(redirectUrl, qrcode.Medium, 256)

	if err != nil {
		return err
	}

	url, err := du.awsRepository.SaveAtS3(ctx, driver.CNH, "qrcode", qrCode)

	if err != nil {
		return err
	}

	driver.QrCode = url

	return du.driverRepository.Create(ctx, driver)
}

func (du *DriverUseCase) Get(ctx context.Context, cnh *string) (*entity.Driver, error) {
	return du.driverRepository.Get(ctx, cnh)
}

func (du *DriverUseCase) Update(ctx context.Context, driver *entity.Driver) error {
	return du.driverRepository.Update(ctx, driver)
}

func (du *DriverUseCase) Delete(ctx context.Context, cnh *string) error {
	return du.driverRepository.Delete(ctx, cnh)
}

func (du *DriverUseCase) SavePix(ctx context.Context, driver *entity.Driver) error {
	return du.driverRepository.SavePix(ctx, driver)
}

func (du *DriverUseCase) SaveBank(ctx context.Context, driver *entity.Driver) error {
	return du.driverRepository.SaveBank(ctx, driver)
}
