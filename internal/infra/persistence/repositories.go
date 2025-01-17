package persistence

import "github.com/venture-technology/venture/internal/domain/repository"

type PostgresRepositories struct {
	ChildRepository       repository.ChildRepository
	ResponsibleRepository repository.ResponsiblesibleRepository
	SchoolRepository      repository.SchoolRepository
	DriverRepository      repository.DriverRepository
	PartnerRepository     repository.PartnerRepository
	ContractRepository    repository.ContractRepository
	InviteRepository      repository.InviteRepository
}
