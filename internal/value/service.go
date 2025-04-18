package value

import (
	"time"

	"github.com/spf13/viper"
	"github.com/venture-technology/venture/internal/entity"
	"github.com/venture-technology/venture/pkg/utils"
)

const (
	// contract statuses

	ContractCurrently = "currently"
	ContractCanceled  = "canceled"
	ContractExpired   = "expired"

	// temp contract statuses

	TempContractPending  = "pending"
	TempContractAccepted = "accepted"
	TempContractCanceled = "canceled"
	TempContractExpired  = "expired"

	// shifts

	MorningShift   = "morning"
	AfternoonShift = "afternoon"
	NightShift     = "night"

	// invite

	InviteAccepted = "accepted"
	InvitePending  = "pending"
)

func GetJWTSecret() string {
	return viper.GetString("JWT_SECRET")
}

func GetBucketGallery() string {
	return viper.GetString("AWS_BUCKET_GALLERY")
}

func GetBucketQRCode() string {
	return viper.GetString("AWS_BUCKET_QRCODE")
}

func GetBucketContract() string {
	return viper.GetString("AWS_BUCKET_CONTRACTS")
}

func GetPathHTML() string {
	return "internal/domain/service/agreements/template/agreement.html"
}

type GetResponsible struct {
	ID              int       `json:"id"`
	Name            string    `json:"name"`
	Email           string    `json:"email"`
	CustomerId      string    `json:"customer_id"`
	Phone           string    `json:"phone"`
	Address         string    `json:"address"`
	ProfileImage    string    `json:"profile_image"`
	PaymentMethodId string    `json:"payment_method_id"`
	CreatedAt       time.Time `json:"created_at"`
	City            string    `json:"city"`
	States          string    `json:"states"`
}

type GetKid struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	RG              string `json:"rg"`
	ResponsibleName string `json:"responsible_name"`
	Address         string `json:"address"`
	Period          string `json:"period"`
	ProfileImage    string `json:"profile_image"`
}

type ListKid struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Period       string `json:"period"`
	ProfileImage string `json:"profile_image"`
}

type GetSchool struct {
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	Address      string    `json:"address"`
	Phone        string    `json:"phone"`
	ProfileImage string    `json:"profile_image"`
	CreatedAt    time.Time `json:"created_at"`
	City         string    `json:"city"`
	States       string    `json:"states"`
}

type ListSchool struct {
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	Address      string    `json:"address"`
	Phone        string    `json:"phone"`
	ProfileImage string    `json:"profile_image"`
	CreatedAt    time.Time `json:"created_at"`
	City         string    `json:"city"`
	States       string    `json:"states"`
}

type GetDriver struct {
	ID            int            `json:"id"`
	Name          string         `json:"name"`
	Email         string         `json:"email"`
	QrCode        string         `json:"qrcode"`
	Amount        float64        `json:"amount"`
	Phone         string         `json:"phone"`
	Car           string         `json:"car"`
	ProfileImage  string         `json:"profile_image"`
	CreatedAt     time.Time      `json:"created_at"`
	Gallery       map[int]string `json:"gallery"`
	Accessibility bool           `json:"accessibility"`
	City          string         `json:"city"`
	States        string         `json:"states"`
	Descriptions  string         `json:"descriptions"`
	Biography     string         `json:"biography"`
}

type ListDriver struct {
	ID            int       `json:"id"`
	Name          string    `json:"name"`
	Email         string    `json:"email"`
	QrCode        string    `json:"qrcode"`
	Amount        float64   `json:"amount"`
	Phone         string    `json:"phone"`
	Car           string    `json:"car"`
	ProfileImage  string    `json:"profile_image"`
	CreatedAt     time.Time `json:"created_at"`
	Accessibility bool      `json:"accessibility"`
	City          string    `json:"city"`
	States        string    `json:"states"`
}

type ListDriverToCalcPrice struct {
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	Amount       float64   `json:"amount"`
	Phone        string    `json:"phone"`
	ProfileImage string    `json:"profile_image"`
	CreatedAt    time.Time `json:"created_at"`
	PriceTotal   float64   `json:"price_total"` // this field is used to calculate the total price of the driver getting distance from responsible and school
}

type SchoolListInvite struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	Phone        string `json:"phone"`
	ProfileImage string `json:"profile_image"`
}

type DriverListInvite struct {
	Name         string `json:"name"`
	Email        string `json:"email"`
	Phone        string `json:"phone"`
	ProfileImage string `json:"profile_image"`
	Address      string `json:"address"`
}

type SchoolListPartners struct {
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	QrCode       string    `json:"qrcode"`
	Phone        string    `json:"phone"`
	Car          string    `json:"car"`
	ProfileImage string    `json:"profile_image"`
	CreatedAt    time.Time `json:"created_at"`
}

type DriverListPartners struct {
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	Address      string    `json:"address"`
	Phone        string    `json:"phone"`
	ProfileImage string    `json:"profile_image"`
	CreatedAt    time.Time `json:"created_at"`
}

type GetContractOutput struct {
	ID                   int               `json:"id"`
	UUID                 string            `json:"uuid"`
	Status               string            `json:"status"`
	SigningURL           string            `json:"signing_url"`
	StripeSubscriptionID string            `json:"subscription_id"`
	CreatedAt            int64             `json:"created_at"`
	ExpiredAt            int64             `json:"expired_at"`
	Driver               GetDriverContract `json:"driver"`
	Responsible          GetParentContract `json:"responsible"`
	Kid                  GetKidContract    `json:"kid"`
	School               GetSchoolContract `json:"school"`
}

type GetContract struct {
	Contract GetContractOutput             `json:"contract"`
	Invoices map[string]entity.InvoiceInfo `json:"invoices"`
}

type DriverListContracts struct {
	ID          int               `json:"id"`
	UUID        string            `json:"record"`
	Status      string            `json:"status"`
	Amount      float64           `json:"amount"`
	School      GetSchoolContract `json:"school"`
	Responsible GetParentContract `json:"responsible"`
	Kid         GetKidContract    `json:"kid"`
	CreatedAt   int64             `json:"created_at"`
	ExpireAt    int64             `json:"expire_at"`
}

type SchoolListContracts struct {
	ID          int               `json:"id"`
	Status      string            `json:"status"`
	KidName     string            `json:"kid_name"`
	Period      string            `json:"period"`
	UUID        string            `json:"record"`
	Amount      float64           `json:"amount"`
	CreatedAt   int64             `json:"created_at"`
	ExpireAt    int64             `json:"expire_at"`
	Driver      GetDriverContract `json:"driver"`
	Responsible GetParentContract `json:"responsible"`
}

type ResponsibleListContracts struct {
	ID        int               `json:"id"`
	Status    string            `json:"status"`
	KidName   string            `json:"kid_name"`
	Period    string            `json:"period"`
	UUID      string            `json:"record"`
	Amount    float64           `json:"amount"`
	CreatedAt int64             `json:"created_at"`
	ExpireAt  int64             `json:"expire_at"`
	Driver    GetDriverContract `json:"driver"`
	School    GetSchoolContract `json:"school"`
}

type GetDriverContract struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	Phone        string `json:"phone"`
	ProfileImage string `json:"profile_image"`
}

type GetSchoolContract struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Address      string `json:"address"`
	Phone        string `json:"phone"`
	ProfileImage string `json:"profile_image"`
}

type GetParentContract struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	Address      string `json:"address"`
	Phone        string `json:"phone"`
	ProfileImage string `json:"profile_image"`
}

type GetKidContract struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Period       string `json:"period"`
	ProfileImage string `json:"profile_image"`
}

type AuthParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Kind     string `json:"kind"`
}

var Schedules = map[string]string{
	"morning":            "1",
	"afternoon":          "2",
	"night":              "3",
	"morning, afternoon": "4",
	"morning, night":     "5",
	"afternoon, night":   "6",
	"all":                "7",
}

var ResponsibleAllowedKeys = map[string]bool{
	"name":          true,
	"email":         true,
	"password":      true,
	"street":        true,
	"number":        true,
	"zip":           true,
	"complement":    true,
	"phone":         true,
	"city":          true,
	"states":        true,
	"profile_image": true,
}

// para alterar schedule ou municipal_record deve entrar em contato com atendimento
var DriverAllowedKeys = map[string]bool{
	"name":          true,
	"email":         true,
	"password":      true,
	"street":        true,
	"number":        true,
	"zip":           true,
	"complement":    true,
	"phone":         true,
	"city":          true,
	"states":        true,
	"pix_key":       true,
	"amount":        true,
	"profile_image": true,
	"car_name":      true,
	"car_year":      true,
	"car_capacity":  true,
}

var SchoollAllowedKeys = map[string]bool{
	"name":          true,
	"email":         true,
	"password":      true,
	"profile_image": true,
}

var KidAllowedKeys = map[string]bool{
	"shift":                 true,
	"attendance_permission": true,
	"profile_image":         true,
}

var Shifts = map[string]string{
	"morning":   MorningShift,
	"afternoon": AfternoonShift,
	"night":     NightShift,
}

type CreateContractRequestParams struct {
	DriverCNH      string `json:"driver_cnh"`
	KidRG          string `json:"kid_rg"`
	ResponsibleCPF string `json:"responsible_cpf"`
	SchoolCNPJ     string `json:"school_cnpj"`
}

type CalculatePriceDriverOutput struct {
	Price  float64       `json:"price"`
	Driver entity.Driver `json:"driver"`
}

type GetTempContracts struct {
	ID                    int    `gorm:"primary_key;auto_increment" json:"id"`
	SigningURL            string `json:"signing_url"`
	Status                string `json:"status"`
	DriverCNH             string `json:"driver_cnh"`
	SchoolCNPJ            string `json:"school_cnpj"`
	KidRG                 string `json:"kid_rg"`
	ResponsibleCPF        string `json:"responsible_cpf"`
	SignatureRequestID    string `json:"signature_request_id"`
	UUID                  string `json:"uuid"`
	CreatedAt             int64  `json:"created_at"`               // epoch time
	ExpiredAt             int64  `json:"expires_at"`               // epoch time
	DriverAssignedAt      int64  `json:"driver_assigned_at"`       // epoch time
	ResponsibleAssignedAt int64  `json:"responsibles_assigned_at"` // epoch time
}

type CreateDriver struct {
	ID              int          `json:"id"`
	Name            string       `json:"name,omitempty" validate:"required"`
	Email           string       `json:"email,omitempty" validate:"required"`
	CNH             string       `json:"cnh,omitempty" validate:"required"`
	QrCode          string       `json:"qrcode,omitempty"`
	PixKey          string       `json:"pix_key,omitempty"`
	Address         string       `json:"address,omitempty" validate:"required"`
	Amount          float64      `json:"amount,omitempty" validate:"required"`
	Phone           string       `json:"phone,omitempty" validate:"required" example:"+55 11 123456789"`
	MunicipalRecord string       `json:"municipal_record,omitempty" validate:"required"`
	Car             entity.Car   `json:"car,omitempty" validate:"required"`
	ProfileImage    string       `json:"profile_image,omitempty"`
	CreatedAt       time.Time    `json:"created_at,omitempty"`
	UpdatedAt       time.Time    `json:"updated_at,omitempty"`
	Schedule        string       `json:"schedule,omitempty"`
	Seats           entity.Seats `json:"seats,omitempty"`
	Accessibility   bool         `json:"accessibility"`
	City            string       `json:"city"`
	States          string       `json:"states"`
}

type CreateResponse struct {
	ID           int       `json:"id,omitempty"`
	Name         string    `json:"name,omitempty" validate:"required"`
	Email        string    `json:"email,omitempty" validate:"required"`
	CPF          string    `json:"cpf,omitempty" validate:"required" example:"44000000000"` // sem pontuação
	Address      string    `json:"address,omitempty" validate:"required"`
	Phone        string    `json:"phone,omitempty" validate:"required" example:"+55 11 123456789"` // o telefone deve seguir este mesmo formato
	ProfileImage string    `json:"profile_image,omitempty"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
	UpdatedAt    time.Time `json:"updated_at,omitempty"`
	City         string    `json:"city"`
	States       string    `json:"states"`
}

type CreateSchool struct {
	ID           int       `json:"id,omitempty"`
	Name         string    `json:"name,omitempty" validate:"required"`
	CNPJ         string    `json:"cnpj,omitempty" validate:"required"`
	Email        string    `json:"email,omitempty" validate:"required"`
	Address      string    `json:"address,omitempty" validate:"required"`
	Phone        string    `json:"phone,omitempty" validate:"required" example:"+55 11 123456789"`
	ProfileImage string    `json:"profile_image,omitempty"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
	UpdatedAt    time.Time `json:"updated_at,omitempty"`
	City         string    `json:"city"`
	States       string    `json:"states"`
}

func MapSchoolEntityToResponse(school entity.School) CreateSchool {
	return CreateSchool{
		ID:    school.ID,
		Name:  school.Name,
		CNPJ:  school.CNPJ,
		Email: school.Email,
		Address: utils.BuildAddress(
			school.Address.Street,
			school.Address.Number,
			school.Address.Complement,
			school.Address.Zip,
		),
		Phone:        school.Phone,
		ProfileImage: school.ProfileImage,
		City:         school.City,
		States:       school.States,
	}
}

func MapResponsibleEntityToResponse(responsible entity.Responsible) CreateResponse {
	return CreateResponse{
		ID:    responsible.ID,
		Name:  responsible.Name,
		Email: responsible.Email,
		CPF:   responsible.CPF,
		Address: utils.BuildAddress(
			responsible.Address.Street,
			responsible.Address.Number,
			responsible.Address.Complement,
			responsible.Address.Zip,
		),
		Phone:        responsible.Phone,
		ProfileImage: responsible.ProfileImage,
		City:         responsible.City,
		States:       responsible.States,
	}
}

func MapDriverEntityToResponse(driver entity.Driver) CreateDriver {
	return CreateDriver{
		ID:     driver.ID,
		Name:   driver.Name,
		Email:  driver.Email,
		CNH:    driver.CNH,
		QrCode: driver.QrCode,
		Address: utils.BuildAddress(
			driver.Address.Street,
			driver.Address.Number,
			driver.Address.Complement,
			driver.Address.Zip,
		),
		Amount:          driver.Amount,
		Phone:           driver.Phone,
		MunicipalRecord: driver.MunicipalRecord,
		Car:             driver.Car,
		ProfileImage:    driver.ProfileImage,
		CreatedAt:       driver.CreatedAt,
		UpdatedAt:       driver.UpdatedAt,
		Schedule:        driver.Schedule,
		Seats:           driver.Seats,
		Accessibility:   driver.Accessibility,
		City:            driver.City,
		States:          driver.States,
	}
}
