package value

import (
	"time"

	"github.com/venture-technology/venture/internal/entity"
)

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
