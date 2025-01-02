package usecase

import (
	"fmt"

	"github.com/venture-technology/venture/internal/infra/contracts"
	"github.com/venture-technology/venture/internal/infra/persistence"
	"github.com/venture-technology/venture/internal/value"
)

type GetDriverUseCase struct {
	repositories *persistence.PostgresRepositories
	logger       contracts.Logger
	S3           contracts.S3Iface
}

func NewGetDriverUseCase(
	repositories *persistence.PostgresRepositories,
	logger contracts.Logger,
	S3 contracts.S3Iface,
) *GetDriverUseCase {
	return &GetDriverUseCase{
		repositories: repositories,
		logger:       logger,
		S3:           S3,
	}
}

func (gduc *GetDriverUseCase) GetDriver(cnh string) (value.GetDriver, error) {
	driver, err := gduc.repositories.DriverRepository.Get(cnh)
	if err != nil {
		return value.GetDriver{}, err
	}

	gallery := make(map[int]string)

	images, err := gduc.S3.List(fmt.Sprintf("%s/gallery", cnh))
	for idx, image := range images {
		gallery[idx] = image
	}

	return value.GetDriver{
		ID:     driver.ID,
		Name:   driver.Name,
		Email:  driver.Email,
		QrCode: driver.QrCode,
		Amount: driver.Amount,
		Phone:  driver.Phone,
		Car: fmt.Sprintf(
			"%s, %s",
			driver.Car.Model,
			driver.Car.Year,
		),
		ProfileImage: driver.ProfileImage,
		CreatedAt:    driver.CreatedAt,
		Gallery:      gallery,
	}, nil
}
