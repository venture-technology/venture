package agreements

import "time"

const (
	TemporaryContractPending  = "pending"
	TemporaryContractExpired  = "expired"
	TemporaryContractCanceled = "canceled"
	TemporaryContractAccepted = "accepted"
)

type Signer struct {
	EmailAddress string `json:"email_address"`
	Name         string `json:"name"`
	Order        int    `json:"order"`
}

type Metadata struct {
	CustomID string `json:"custom_id"`
	Keys     struct {
		UUID               string    `json:"uuid"`
		DriverID           string    `json:"driver_id"`
		DriverName         string    `json:"driver_name"`
		ResponsibleID      string    `json:"responsible_id"`
		ResponsibleName    string    `json:"responsible_name"`
		ResponsibleCPF     string    `json:"responsible_cpf"`
		ResponsibleEmail   string    `json:"responsible_email"`
		ResponsiblePhone   string    `json:"responsible_phone"`
		ResponsibleAddr    string    `json:"responsible_addr"`
		KidID              string    `json:"kid_id"`
		KidName            string    `json:"kid_name"`
		SchoolID           string    `json:"school_id"`
		SchoolName         string    `json:"school_name"`
		SchoolAddr         string    `json:"school_addr"`
		DateTime           string    `json:"date_time"`
		AmountContract     float64   `json:"amount_contract"`
		AnualContractValue float64   `json:"anual_contract_value"`
		Time               time.Time `json:"time"`
	} `json:"keys"`
}

type ContractRequest struct {
	Title            string   `json:"title"`
	Subject          string   `json:"subject"`
	Message          string   `json:"message"`
	Signers          []Signer `json:"signers"`
	CCEmailAddresses []string `json:"cc_email_addresses"`
	FileUrls         []string `json:"file_urls"`
	Metadata         Metadata `json:"metadata"`
	SigningOptions   struct {
		Draw        bool   `json:"draw"`
		Type        bool   `json:"type"`
		Upload      bool   `json:"upload"`
		Phone       bool   `json:"phone"`
		DefaultType string `json:"default_type"`
	} `json:"signing_options"`
	FieldOptions struct {
		DateFormat string `json:"date_format"`
	} `json:"field_options"`
	ExpiresAt int64 `json:"expires_at"`
	TestMode  bool  `json:"test_mode"`
}

type Event struct {
	EventTime     string        `json:"event_time"`
	EventType     string        `json:"event_type"`
	EventHash     string        `json:"event_hash"`
	EventMetadata EventMetadata `json:"event_metadata"`
}

type EventMetadata struct {
	ReportedForAccountID string `json:"reported_for_account_id"`
}

type EventWrapper struct {
	Event Event `json:"event"`
}

type SignatureResponse struct {
	SignatureRequest struct {
		Metadata           Metadata `json:"metadata"`
		SignatureRequestID string   `json:"signature_request_id"`
		SigningURL         string   `json:"signing_url"`
		CreatedAt          int64    `json:"created_at"`
		Signatures         []struct {
			SignatureID string `json:"signature_id"`
			SignerEmail string `json:"signer_email_address"`
			SignerName  string `json:"signer_name"`
		} `json:"signatures"`
	} `json:"signature_request"`
}
