package entity

type StripeSubscription struct {
	Title       string `json:"title_subscription"`
	ID          string `json:"subscription_id"`
	Description string `json:"description_subscription,omitempty"`
	Price       string `json:"price_id"`
	Product     string `json:"product_id"`
}

type DriverCP struct {
	Name  string `json:"driver_name"`
	CNH   string `json:"driver_cnh"`
	Email string `json:"driver_email"`
}

type ResponsibleCP struct {
	Name  string `json:"responsible_name"`
	CPF   string `json:"responsible_cnh"`
	Email string `json:"responsible_email"`
}

type KidCP struct {
	RG string `json:"kid_rg"`
}

type SchoolCP struct {
	CNPJ string `json:"school_cnpj"`
}

type Contract struct {
	ID                   int    `gorm:"primary_key;auto_increment" json:"id"`
	UUID                 string `json:"record,omitempty"`
	PreApprovalID        string `json:"pre_approval_id" validate:"required"`
	Status               string `json:"status" validate:"oneof='currently' 'canceled' 'expired'"`
	StripeSubscriptionID string `json:"stripe_subscription_id"`
	StripePriceID        string `json:"stripe_price_id"`
	StripeProductID      string `json:"stripe_product_id"`
	SigningURL           string `json:"dropbox_signing_url"`
	DriverCNH            string `json:"driver_cnh"`
	SchoolCNPJ           string `json:"school_cnpj"`
	KidRG                string `json:"kid_rg"`
	ResponsibleCPF       string `json:"responsible_cpf"`
	CreatedAt            int64  `json:"created_at,omitempty"`
	UpdatedAt            int64  `json:"updated_at,omitempty"`
	ExpireAt             int64  `json:"expire_at"`
	Amount               int64  `json:"amount" validate:"required"`
	AnualAmount          int64  `json:"anual_amount"`

	// Relations

	Driver      Driver      `gorm:"foreignKey:DriverCNH;references:CNH" json:"driver"`
	School      School      `gorm:"foreignKey:SchoolCNPJ;references:CNPJ" json:"school"`
	Kid         Kid         `gorm:"foreignKey:KidRG;references:RG" json:"kid"`
	Responsible Responsible `gorm:"foreignKey:ResponsibleCPF;references:CPF" json:"responsible"`
}

type SubscriptionInfo struct {
	ID     string `json:"subscription_id"`
	Status string `json:"subscription_status"`
}

func (c *Contract) ValidateAmount() bool {
	return c.Amount != 0
}
