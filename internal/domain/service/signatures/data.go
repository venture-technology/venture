package signatures

import (
	"github.com/venture-technology/venture/internal/entity"
)

const (
	TemporaryContractPending  = "pending"
	TemporaryContractExpired  = "expired"
	TemporaryContractCanceled = "canceled"
	TemporaryContractAccepted = "accepted"
)

type Signer struct {
	EmailAddress string `json:"email_address"`
	Name         string `json:"name"`
}

type Metadata struct {
	CustomID string `json:"custom_id"`
	Keys     struct {
		UUID               string `json:"uuid"`
		DriverCNH          string `json:"driver_cnh"`
		ResponsibleCPF     string `json:"responsible_cpf"`
		KidRG              string `json:"kid_rg"`
		SchoolCNPJ         string `json:"school_cnpj"`
		AmountContract     int64  `json:"amount_contract"`
		AnualContractValue int64  `json:"anual_contract_value"`
	} `json:"keys"`
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

type ContractParams struct {
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

type Signatures struct {
	SignatureID string `json:"signature_id"`
	SignerEmail string `json:"signer_email_address"`
	SignerName  string `json:"signer_name"`
	SignedAt    int64  `json:"signed_at"'`
}

type SignatureRequest struct {
	Metadata           Metadata     `json:"metadata"`
	CreatedAt          int64        `json:"created_at"`
	SigningURL         string       `json:"signing_url"`
	Signatures         []Signatures `json:"signatures"`
	SignatureRequestID string       `json:"signature_request_id"`
}

type SignatureRequestAllSigned struct {
	SignatureRequest SignatureRequest `json:"signature_request"`
	Event            Event            `json:"event"`
}

// Agreement Signature Request All Signed Input
//
// Receive body and mapping contract and temp_contract to do business logic
type ASRASOutput struct {
	Contract   entity.Contract `json:"contract"`
	Signatures []Signatures    `json:"signatures"`
}
