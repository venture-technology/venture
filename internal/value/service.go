package value

import (
	"time"

	"github.com/google/uuid"
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
}

type GetChild struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	RG              string `json:"rg"`
	ResponsibleName string `json:"responsible_name"`
	Address         string `json:"address"`
	Period          string `json:"period"`
	ProfileImage    string `json:"profile_image"`
}

type ListChild struct {
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
}

type ListSchool struct {
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	Address      string    `json:"address"`
	Phone        string    `json:"phone"`
	ProfileImage string    `json:"profile_image"`
	CreatedAt    time.Time `json:"created_at"`
}

type GetDriver struct {
	ID           int            `json:"id"`
	Name         string         `json:"name"`
	Email        string         `json:"email"`
	QrCode       string         `json:"qrcode"`
	Amount       float64        `json:"amount"`
	Phone        string         `json:"phone"`
	Car          string         `json:"car"`
	ProfileImage string         `json:"profile_image"`
	CreatedAt    time.Time      `json:"created_at"`
	Gallery      map[int]string `json:"gallery"`
}

type ListDriver struct {
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	QrCode       string    `json:"qrcode"`
	Amount       float64   `json:"amount"`
	Phone        string    `json:"phone"`
	Car          string    `json:"car"`
	ProfileImage string    `json:"profile_image"`
	CreatedAt    time.Time `json:"created_at"`
}

type SchoolListInvite struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	Phone        string `json:"phone"`
	ProfileImage string `json:"profile_image"`
}

type DriverListInvite struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	Phone        string `json:"phone"`
	ProfileImage string `json:"profile_image"`
	Address      string `json:"address"`
}

type SchoolListPartners struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	QrCode       string `json:"qrcode"`
	Phone        string `json:"phone"`
	Car          string `json:"car"`
	ProfileImage string `json:"profile_image"`
}

type DriverListPartners struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	Address      string `json:"address"`
	Phone        string `json:"phone"`
	ProfileImage string `json:"profile_image"`
}

type GetContract struct {
	ID          int                           `json:"id"`
	Status      string                        `json:"status"`
	ChildName   string                        `json:"child_name"`
	Period      string                        `json:"period"`
	Amount      float64                       `json:"amount"`
	Record      uuid.UUID                     `json:"record"`
	CreatedAt   time.Time                     `json:"created_at"`
	ExpireAt    time.Time                     `json:"expire_at"`
	Responsible GetParentContract             `json:"responsible"`
	Driver      GetDriverContract             `json:"driver"`
	School      GetSchoolContract             `json:"school"`
	Invoices    map[string]entity.InvoiceInfo `json:"invoices"`
}

type DriverListContracts struct {
	ID          int               `json:"id"`
	Record      uuid.UUID         `json:"record"`
	Status      string            `json:"status"`
	Amount      float64           `json:"amount"`
	School      GetSchoolContract `json:"school"`
	Responsible GetParentContract `json:"responsible"`
	Child       GetChildContract  `json:"child"`
	CreatedAt   time.Time         `json:"created_at"`
	ExpireAt    time.Time         `json:"expire_at"`
}

type SchoolListContracts struct {
	ID          int               `json:"id"`
	Status      string            `json:"status"`
	ChildName   string            `json:"child_name"`
	Period      string            `json:"period"`
	Amount      float64           `json:"amount"`
	Record      uuid.UUID         `json:"record"`
	CreatedAt   time.Time         `json:"created_at"`
	ExpireAt    time.Time         `json:"expire_at"`
	Driver      GetDriverContract `json:"driver"`
	Responsible GetParentContract `json:"responsible"`
}

type ResponsibleListContracts struct {
	ID        int               `json:"id"`
	Status    string            `json:"status"`
	ChildName string            `json:"child_name"`
	Period    string            `json:"period"`
	Amount    float64           `json:"amount"`
	Record    uuid.UUID         `json:"record"`
	CreatedAt time.Time         `json:"created_at"`
	ExpireAt  time.Time         `json:"expire_at"`
	Driver    GetDriverContract `json:"driver"`
	School    GetSchoolContract `json:"school"`
}

type GetDriverContract struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	Address      string `json:"address"`
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

type GetChildContract struct {
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
