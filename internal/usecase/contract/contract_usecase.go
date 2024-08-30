package contract

import "github.com/venture-technology/venture/internal/repository"

type ContractUseCase struct {
	contractRepository repository.IContractRepository
}

func NewContractUseCase(cr repository.IContractRepository) *ContractUseCase {
	return &ContractUseCase{
		contractRepository: cr,
	}
}

func (cu *ContractUseCase) Create() {

}

func (cu *ContractUseCase) Get() {

}

func (cu *ContractUseCase) FindAllByCnh() {

}

func (cu *ContractUseCase) FindAllByCnpj() {

}

func (cu *ContractUseCase) FindAllByCpf() {

}

func (cu *ContractUseCase) Cancel() {

}

func CalculateContract() {

}

func CreateProduct() {

}

func CreatePrice() {

}

func CreateSubscription() {

}

func GetSubscription() {

}

func ListSubscription() {

}

func GetInvoices() {

}

func ListInvoices() {

}

func CreateIntent() {

}
