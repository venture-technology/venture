package persistence

import "github.com/venture-technology/venture/internal/domain/repository"

type PostgresRepositories struct {
	ChildRepository       repository.IChildRepository
	ResponsibleRepository repository.IResponsibleRepository
	SchoolRepository      repository.ISchoolRepository
	DriverRepository      repository.IDriverRepository
	PartnerRepository     repository.IPartnerRepository
	ContractRepository    repository.IContractRepository
	InviteRepository      repository.IInviteRepository
}
