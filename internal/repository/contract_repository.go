package repository

type IContractRepository interface {
	Create()
	Get()
	FindAllByCnpj()
	FindAllByCpf()
	FindAllByCnh()
	Update()
	Cancel()
}
