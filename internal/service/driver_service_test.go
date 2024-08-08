package service

import (
	"context"
	"database/sql"
	"log"
	"reflect"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/venture-technology/venture/config"
	"github.com/venture-technology/venture/internal/repository"
	"github.com/venture-technology/venture/models"

	_ "github.com/lib/pq"
)

func mockDriver() *models.Driver {
	return &models.Driver{
		Name:       "João Silva",
		CNH:        "15750721547",
		Email:      "gustavorodrigueslima2004@gmail.com",
		Password:   "123teste", //P@ssw0rd123
		CPF:        "00233173021",
		Street:     "Rua das Flores",
		Number:     "123",
		ZIP:        "11088440",
		Complement: "Apt 101",
		Amount:     117.46,
	}
}

func setupTesteDB(t *testing.T) (*sql.DB, *DriverService) {

	t.Helper()

	config, err := config.Load("../../config/config.yaml")
	if err != nil {
		t.Fatalf("falha ao carregar a configuração: %v", err)
	}

	db, err := sql.Open("postgres", newPostgres(config.Database))
	if err != nil {
		t.Fatalf("falha ao conectar ao banco de dados: %v", err)
	}

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(config.Cloud.Region),
		Credentials: credentials.NewStaticCredentials(config.Cloud.AccessKey, config.Cloud.SecretKey, config.Cloud.Token),
	})

	if err != nil {
		log.Fatalf("failed to create session at aws: %v", err)
	}

	driverRepository := repository.NewDriverRepository(db)
	awsRepository := repository.NewAWSRepository(sess)

	driverService := NewDriverService(driverRepository, awsRepository)

	return db, driverService

}

func TestCreateDriver(t *testing.T) {

	db, driverService := setupTesteDB(t)
	defer db.Close()

	driverMock := mockDriver()

	qrcode, err := driverService.CreateAndSaveQrCode(context.Background(), driverMock.CNH)
	if err != nil {
		t.Errorf("erro ao criar qrcode")
	}

	log.Print(qrcode)

	driverMock.QrCode = qrcode

	err = driverService.CreateDriver(context.Background(), driverMock)
	if err != nil {
		t.Errorf("Erro ao criar motorista: %v", err)
	}

}

func TestGetDriver(t *testing.T) {

	db, driverService := setupTesteDB(t)
	defer db.Close()

	driverMock := mockDriver()

	driverData, err := driverService.GetDriver(context.Background(), &driverMock.CNH)

	if err != nil {
		t.Errorf("Erro ao fazer leitura da escola: %v", err.Error())
	}

	driverMock.Password = ""
	driverMock.ID = driverData.ID
	driverMock.QrCode = driverData.QrCode

	log.Print("driverData: ", driverData, "driverMock: ", driverMock)

	if *driverMock != *driverData {
		t.Error("Mock é diferente do user retornado do banco")
	}

}

func TestUpdateDriver(t *testing.T) {

	db, driverService := setupTesteDB(t)
	defer db.Close()

	newDriver := models.Driver{
		Name:       "João Silva",
		CNH:        "15750721547",
		Email:      "gustavorodrigueslima2004@gmail.com",
		Password:   "123teste",
		CPF:        "00233173021",
		Street:     "Rua das Flores",
		Number:     "123",
		ZIP:        "11088440",
		Complement: "Apt 101",
	}

	err := driverService.UpdateDriver(context.Background(), &newDriver)
	if err != nil {
		t.Errorf("Erro ao atualizar motorista: %v", err)
	}

}

func TestAuthDriver(t *testing.T) {

	db, driverService := setupTesteDB(t)
	defer db.Close()

	driverMock := mockDriver()

	driverData, err := driverService.AuthDriver(context.Background(), driverMock)

	if err != nil {
		t.Errorf("Erro ao fazer login de motorista: %v", err.Error())
	}

	if reflect.TypeOf(driverData) != reflect.TypeOf(driverMock) {
		t.Errorf("Não foi retornado os dados de login da escola: %v", err.Error())
	}

}

func TestDeleteDriver(t *testing.T) {

	db, driverService := setupTesteDB(t)
	defer db.Close()

	driverMock := mockDriver()

	err := driverService.DeleteDriver(context.Background(), &driverMock.CNH)

	if err != nil {
		t.Errorf("Erro ao deletar motorista: %v", err.Error())
	}

}
