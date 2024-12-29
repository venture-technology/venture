package value

import (
	"time"

	"github.com/google/uuid"
)

type CreateResponsibleInput struct {
}

type CreateResponsibleOutput struct {
}

type SaveCardInput struct {
}

type UpdateResponsibleInput struct {
}

type UpdateResponsibleOutput struct {
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
}

type LoginResponsible struct {
}

type CreateChildInput struct {
}

type CreateChildOutput struct {
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

type UpdateChildInput struct {
}

type UpdateChildOutput struct {
}

type CreateDriverInput struct {
}

type CreateDriverOutput struct {
}

type CreateSchoolInput struct {
}

type CreateSchoolOutput struct {
}

type LoginSchool struct {
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
}

type UpdateSchoolInput struct {
}

type UpdateSchoolOutput struct {
}

type LoginDriver struct {
}

type GetDriver struct {
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

type ListDriver struct {
}

type ListGallery struct {
}

type CreateInviteInput struct {
}

type CreateInviteOutput struct {
}

type GetInvite struct {
}

type ListInvite struct {
}

type GetPrice struct {
}

type CreatePartnerInput struct {
}

type CreatePartnerOutput struct {
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
	ID          int               `json:"id"`
	Status      string            `json:"status"`
	ChildName   string            `json:"child_name"`
	Period      string            `json:"period"`
	Amount      float64           `json:"amount"`
	Record      uuid.UUID         `json:"record"`
	CreatedAt   time.Time         `json:"created_at"`
	ExpireAt    time.Time         `json:"expire_at"`
	Driver      GetDriverContract `json:"driver"`
	School      GetSchoolContract `json:"school"`
	Responsible GetParentContract `json:"responsible"`
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
}

type ResponsibleListContracts struct {
}

type ListPartnerDriver struct {
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
