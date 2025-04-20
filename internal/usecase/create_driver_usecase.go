package usecase

import (
	"fmt"

	"github.com/skip2/go-qrcode"
	"github.com/venture-technology/venture/internal/entity"
	"github.com/venture-technology/venture/internal/infra/contracts"
	"github.com/venture-technology/venture/internal/infra/persistence"
	"github.com/venture-technology/venture/internal/value"
	"github.com/venture-technology/venture/pkg/utils"
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
	ok, errors := utils.ValidatePassword(driver.Password)
	if !ok {
		return fmt.Errorf(errors)
	}

	err := driver.ValidateCnh()
	if err != nil {
		return err
	}

	err = driver.ValidateCapacity()
	if err != nil {
		return err
	}

	driver, err = fillSeatCapacity(driver)
	if err != nil {
		return fmt.Errorf("error filling seat capacity: %w", err)
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

	url, err := cduc.S3.Save(value.GetBucketQRCode(), driver.CNH, "qrcode", qrCode)
	if err != nil {
		return err
	}

	driver.QrCode = url
	return cduc.repositories.DriverRepository.Create(driver)
}

// fillSeatCapacity fills the driver's seat capacity according to the schedule and his car capacity
func fillSeatCapacity(driver *entity.Driver) (*entity.Driver, error) {
	schedule, err := getScheduleDriver(driver)
	if err != nil {
		return nil, fmt.Errorf("error getting schedule: %w", err)
	}
	driver.Schedule = schedule
	fillSeats(driver)
	return driver, nil
}

func getScheduleDriver(driver *entity.Driver) (string, error) {
	schedule, exists := value.Schedules[driver.Schedule]
	if !exists {
		return "", fmt.Errorf("schedule not found")
	}
	return schedule, nil
}

func fillSeats(driver *entity.Driver) {
	driver.Seats.Remaining = driver.Car.Capacity
	var seatMap = map[string]func(driver *entity.Driver){
		"1": func(driver *entity.Driver) {
			driver.Seats.Morning = driver.Car.Capacity
		},
		"2": func(driver *entity.Driver) {
			driver.Seats.Afternoon = driver.Car.Capacity
		},
		"3": func(driver *entity.Driver) {
			driver.Seats.Night = driver.Car.Capacity
		},
		"4": func(driver *entity.Driver) {
			driver.Seats.Morning = driver.Car.Capacity / 2
			driver.Seats.Afternoon = driver.Car.Capacity / 2
		},
		"5": func(driver *entity.Driver) {
			driver.Seats.Morning = driver.Car.Capacity / 2
			driver.Seats.Night = driver.Car.Capacity / 2
		},
		"6": func(driver *entity.Driver) {
			driver.Seats.Afternoon = driver.Car.Capacity / 2
			driver.Seats.Night = driver.Car.Capacity / 2
		},
		"7": func(driver *entity.Driver) {
			driver.Seats.Morning = driver.Car.Capacity / 3
			driver.Seats.Afternoon = driver.Car.Capacity / 3
			driver.Seats.Night = driver.Car.Capacity / 3
		},
	}
	seatMap[driver.Schedule](driver)
}
