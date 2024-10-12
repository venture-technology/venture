package payments

import (
	"log"
	"testing"

	"github.com/venture-technology/venture/internal/entity"
)

func TestPayments_ListInvoice(t *testing.T) {
	contract := entity.Contract{
		StripeSubscription: entity.StripeSubscription{
			ID: "sub_1Q7VhYLfFDLpePGLJyByBbuo",
		},
	}

	contracts := []entity.Contract{contract}

	useCase := NewStripeContract()

	for _, contract := range contracts {
		invoices, err := useCase.ListInvoices(&contract)
		if err != nil {
			t.Errorf("TestPayments_ListInvoice: %s", err)
		}

		contract.Invoices = invoices
	}

	log.Print(contract)

}
