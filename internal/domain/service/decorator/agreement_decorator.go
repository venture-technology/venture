package decorator

import (
	"github.com/venture-technology/venture/internal/domain/service/adapters"
	"github.com/venture-technology/venture/internal/infra/contracts"
)

type AgreementDecorator struct {
	adapters.AgreementService
	contracts.Cacher
}

func NewAgreementDecorator(
	agreementService adapters.AgreementService,
	cache contracts.Cacher,
) AgreementDecorator {
	return AgreementDecorator{
		AgreementService: agreementService,
		Cacher:           cache,
	}
}

func (d AgreementDecorator) Auth() (string, error) {
	return d.AgreementService.Auth()
}

func (d AgreementDecorator) getCache() (string, error) {
	return "", nil
}

func (d AgreementDecorator) setCache() error {
	return nil
}
