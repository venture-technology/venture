package agreements

import (
	"github.com/venture-technology/venture/config"
	"github.com/venture-technology/venture/internal/entity"
)

type AgreementService struct {
	config config.Config
}

func NewAgreementService(
	config config.Config,
) *AgreementService {
	return &AgreementService{
		config: config,
	}
}

func (as *AgreementService) Auth() (string, error) {
	return "", nil
}

func (as *AgreementService) CreateAgreement(contract *entity.Contract) error {
	return nil
}
