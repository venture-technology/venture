package usecase

import (
	"context"
	"fmt"
	"sync"

	"github.com/google/uuid"
	"github.com/venture-technology/venture/internal/entity"
	"github.com/venture-technology/venture/internal/infra/contracts"
	"github.com/venture-technology/venture/internal/infra/persistence"
	"github.com/venture-technology/venture/internal/value"
	"github.com/venture-technology/venture/pkg/realtime"
)

type CreateContractUsecase struct {
	repositories *persistence.PostgresRepositories
	queueWorker  contracts.WorkerCreateContract
	logger       contracts.Logger
}

func NewCreateContractUsecase(
	repositories *persistence.PostgresRepositories,
	queueWorker contracts.WorkerCreateContract,
	logger contracts.Logger,
) *CreateContractUsecase {
	return &CreateContractUsecase{
		repositories: repositories,
		queueWorker:  queueWorker,
		logger:       logger,
	}
}

func (ccuc *CreateContractUsecase) CreateContract(
	ctx context.Context,
	requestParams *value.CreateContractParams,
) error {
	requestParams, err := ccuc.handleParams(ctx, requestParams)
	if err != nil {
		return err
	}

	return ccuc.queueWorker.Enqueue(requestParams)
}

func (ccuc *CreateContractUsecase) handleParams(
	ctx context.Context,
	params *value.CreateContractParams,
) (*value.CreateContractParams, error) {
	var (
		wg      sync.WaitGroup
		errCh   = make(chan error, 5)
		results sync.Map
	)

	wg.Add(6)
	go func() {
		defer wg.Done()
		hasParent, err := ccuc.repositories.
			KidRepository.FindByResponsible(params.ResponsibleCPF, params.KidRG)
		if err != nil {
			errCh <- err
			return
		}
		if !hasParent {
			errCh <- fmt.Errorf("parents not found")
			return
		}
		results.Store("hasParent", hasParent)
	}()

	go func() {
		defer wg.Done()
		driver, err := ccuc.repositories.
			DriverRepository.Get(params.DriverCNH)
		if err != nil {
			errCh <- err
			return
		}
		results.Store("driver", driver)
	}()

	go func() {
		defer wg.Done()
		kid, err := ccuc.repositories.
			KidRepository.Get(&params.KidRG)
		if err != nil {
			errCh <- err
			return
		}
		results.Store("kid", kid)
	}()

	go func() {
		defer wg.Done()
		school, err := ccuc.repositories.
			SchoolRepository.Get(params.SchoolCNPJ)
		if err != nil {
			errCh <- err
			return
		}
		results.Store("school", school)
	}()

	go func() {
		defer wg.Done()
		responsible, err := ccuc.repositories.
			ResponsibleRepository.Get(params.ResponsibleCPF)
		if err != nil {
			errCh <- err
			return
		}
		results.Store("responsible", responsible)
	}()

	go func() {
		defer wg.Done()
		result, err := ccuc.repositories.
			TempContractRepository.HasTemporaryContract(&entity.TempContract{
			KidRG:          params.KidRG,
			ResponsibleCPF: params.ResponsibleCPF,
		})
		if err != nil {
			errCh <- err
			return
		}
		if result {
			errCh <- fmt.Errorf("parents already have temporary contract")
		}
	}()

	wg.Wait()
	close(errCh)

	for err := range errCh {
		return nil, err
	}

	if v, ok := results.Load("driver"); ok {
		params.DriverName = v.(*entity.Driver).Name
		params.DriverAmount = v.(*entity.Driver).Amount
		params.DriverEmail = v.(*entity.Driver).Email
		params.DriverCNH = v.(*entity.Driver).CNH
	}

	if v, ok := results.Load("responsible"); ok {
		params.ResponsibleName = v.(*entity.Responsible).Name
		params.ResponsibleEmail = v.(*entity.Responsible).Email
		params.ResponsibleCPF = v.(*entity.Responsible).CPF
		params.ResponsibleAddr = v.(*entity.Responsible).Address.GetFullAddress()
		params.ResponsiblePhone = v.(*entity.Responsible).Phone
	}

	if v, ok := results.Load("kid"); ok {
		params.KidName = v.(*entity.Kid).Name
		params.KidRG = v.(*entity.Kid).RG
		params.KidShift = v.(*entity.Kid).Shift
	}

	if v, ok := results.Load("school"); ok {
		params.SchoolName = v.(*entity.School).Name
		params.SchoolCNPJ = v.(*entity.School).CNPJ
		params.SchoolAddr = v.(*entity.School).Address.GetFullAddress()
	}

	time := realtime.Now()
	params.UUID = uuid.NewString()
	params.Time = time
	params.DateTime = time.Format("01/02/2006")

	return params, nil
}
