package payments

import (
	"context"
	"fmt"

	"github.com/mercadopago/sdk-go/pkg/preapproval"
	"github.com/venture-technology/venture/internal/value"
	"github.com/venture-technology/venture/pkg/stringcommon"
)

const (
	// pending Status serves for when new contract is created.
	//
	// It symbolizes the creation of a subscription by any means of payment
	pendingStatus = "pending"

	// authorized Status server for when new contract is created. But instead pending. The client choose pay with Credit Card.
	authorizedStatus = "authorized"
	canceledStatus   = "canceled"

	// Brazilian Currency
	brl = "BRL"

	// The frequency is the cycle of billing.
	//
	// In our language, the frequency symbolizes the year of contract.
	frequency = 1
	// The frequency type is the way of billing, and their ocurrency.
	frequencyType = "month"

	// The urlBack will used after the payment happens.
	urlBack = "https://venture.com.br"
)

type Payments interface {
	// NewPreApproval can be used to request payments to subscription.
	//
	// Create creates a new pre-approval. It is a post request to the endpoint: https://api.mercadopago.com/preapproval
	NewPreApproval(ctx context.Context, params value.CreateContractParams) (*preapproval.Response, error)

	// Cancel a pre-approval contract.
	CancelPreApproval(ctx context.Context, id string) error
}

type payments struct {
	preApproval preapproval.Client
}

func NewPayment(preApproval preapproval.Client) *payments {
	return &payments{
		preApproval: preApproval,
	}
}

func (p *payments) NewPreApproval(ctx context.Context, params value.CreateContractParams) (*preapproval.Response, error) {
	request := preapproval.Request{
		ExternalReference: params.UUID,
		PayerEmail:        params.ResponsibleEmail,
		BackURL:           urlBack,
		AutoRecurring: &preapproval.AutoRecurringRequest{
			FreeTrial:     &preapproval.FreeTrialRequest{},
			CurrencyID:    brl,
			StartDate:     params.CreatedAt,
			EndDate:       params.ExpiredAt,
			Frequency:     frequency,
			FrequencyType: frequencyType,
		},
		Status: pendingStatus,
		Reason: fmt.Sprintf("Assinatura Mensal - Motorista: %s", params.DriverName),
	}

	if !stringcommon.Empty(params.ResponsibleCardTokenID) {
		request.CardTokenID = params.ResponsibleCardTokenID
		request.Status = authorizedStatus
	}

	resource, err := p.preApproval.Create(ctx, request)
	if err != nil {
		return nil, err
	}

	return resource, nil
}

func (p *payments) CancelPreApproval(ctx context.Context, id string) error {
	request := preapproval.UpdateRequest{
		Status: canceledStatus,
	}

	_, err := p.preApproval.Update(ctx, id, request)
	if err != nil {
		return err
	}

	return err
}
