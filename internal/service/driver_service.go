package service

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/skip2/go-qrcode"
	"github.com/venture-technology/venture/config"
	"github.com/venture-technology/venture/internal/repository"
	"github.com/venture-technology/venture/models"
	"github.com/venture-technology/venture/utils"
)

type DriverService struct {
	driverrepository repository.IDriverRepository
	awsrepository    repository.IAWSRepository
}

func NewDriverService(driverrepository repository.IDriverRepository, awsrepository repository.IAWSRepository) *DriverService {
	return &DriverService{
		driverrepository: driverrepository,
		awsrepository:    awsrepository,
	}
}

func (d *DriverService) CreateDriver(ctx context.Context, driver *models.Driver) error {

	driver.Password = utils.HashPassword(driver.Password)

	statusCnh := driver.ValidateCnh()

	if !statusCnh {
		return fmt.Errorf("cnh invalid")
	}

	return d.driverrepository.CreateDriver(ctx, driver)
}

func (d *DriverService) GetDriver(ctx context.Context, cnh *string) (*models.Driver, error) {
	log.Printf("param read school -> cnh: %s", *cnh)
	return d.driverrepository.GetDriver(ctx, cnh)
}

func (d *DriverService) UpdateDriver(ctx context.Context, driver *models.Driver) error {
	log.Printf("input received to update school -> name: %s, cnh: %s, email: %s", driver.Name, driver.CNH, driver.Email)
	return d.driverrepository.UpdateDriver(ctx, driver)
}

func (d *DriverService) DeleteDriver(ctx context.Context, cnh *string) error {
	log.Printf("trying delete your infos --> %v", *cnh)
	return d.driverrepository.DeleteDriver(ctx, cnh)
}

func (d *DriverService) AuthDriver(ctx context.Context, driver *models.Driver) (*models.Driver, error) {
	driver.Password = utils.HashPassword((driver.Password))
	return d.driverrepository.AuthDriver(ctx, driver)
}

func (d *DriverService) CreateAndSaveQrCode(ctx context.Context, cnh string) (string, error) {

	url := fmt.Sprintf("https://venture-technology.xyz/driver/%s", cnh)

	image, err := qrcode.Encode(url, qrcode.Medium, 256)

	if err != nil {
		return "", err
	}

	return d.awsrepository.SaveImageOnAWSBucket(ctx, image, cnh)
}

func (d *DriverService) ParserJwtDriver(ctx *gin.Context) (interface{}, error) {

	cnh, found := ctx.Get("cnh")

	if !found {
		return nil, fmt.Errorf("error while veryfing token")
	}

	return cnh, nil

}

func (d *DriverService) CreateTokenJWTDriver(ctx context.Context, driver *models.Driver) (string, error) {

	conf := config.Get()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"cnh": driver.CNH,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	jwt, err := token.SignedString([]byte(conf.Server.Secret))

	if err != nil {
		return "", err
	}

	return jwt, nil

}
