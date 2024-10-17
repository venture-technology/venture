package contract

import (
	"context"
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stripe/stripe-go/v79"
	"github.com/venture-technology/venture/internal/domain/adapter"
	"github.com/venture-technology/venture/internal/domain/payments"
	"github.com/venture-technology/venture/internal/entity"
	"github.com/venture-technology/venture/mocks"
)

func TestContractUseCase_Create(t *testing.T) {
	t.Run("when create is sucess", func(t *testing.T) {
		cou := mocks.NewIContractRepository(t)
		st := mocks.NewIStripe(t)
		googleAdapter := adapter.NewGoogleAdapter()

		contract := entity.Contract{
			Driver: entity.Driver{
				Name: "Motorista",
				Pix: entity.Pix{
					Key: "key",
				},
				Bank: entity.Bank{
					Account: "account",
					Agency:  "agency",
					Name:    "bank",
				},
				Car: entity.Car{
					Model: "model",
					Year:  "year",
				},
				Amount: 185.0,
				CNH:    "cnh",
			},
			School: entity.School{
				Name: "Escola",
				Address: entity.Address{
					Street: "Avenida Barão de Alagoas",
					Number: "223",
					ZIP:    "08120000",
				},
			},
			Child: entity.Child{
				Responsible: entity.Responsible{
					Name:            "Responsible",
					CPF:             "cpf",
					PaymentMethodId: "pm_1PxgSrLfFDLpePGLDi2tnFEm",
					CustomerId:      "cus_QpL9XTVfM6sBhD",
					Address: entity.Address{
						Street: "Rua Masato Sakai",
						Number: "180",
						ZIP:    "008538300",
					},
				},
				Name: "Child",
				RG:   "RG",
			},
		}

		cou.On("GetSimpleContractByTitle", context.Background(), &contract.StripeSubscription.Title).Return(&entity.Contract{}, nil)
		cou.On("Create", context.Background(), &contract).Return(nil)
		st.On("CreateProduct", &contract).Return(&stripe.Product{
			ID: "prod_QzgBXR7xS7nyQ4",
		}, nil)
		st.On("CreatePrice", &contract).Return(&stripe.Price{
			ID: "price_1Q7gooLfFDLpePGL9xAzzxKK",
		}, nil)
		st.On("CreateSubscription", &contract).Return(&stripe.Subscription{
			ID: "sub_1Q7gooLfFDLpePGLkcm2nPkC",
		}, nil)

		useCase := NewContractUseCase(cou, st, googleAdapter, nil)

		err := useCase.Create(context.Background(), &contract)
		if err != nil {
			t.Errorf("Error: %s", err)
		}

	})

	t.Run("when getsimple fails", func(t *testing.T) {
		cou := mocks.NewIContractRepository(t)
		st := mocks.NewIStripe(t)
		googleAdapter := adapter.NewGoogleAdapter()

		contract := entity.Contract{
			Driver: entity.Driver{
				Name: "Motorista",
				Pix: entity.Pix{
					Key: "key",
				},
				Bank: entity.Bank{
					Account: "account",
					Agency:  "agency",
					Name:    "bank",
				},
				Car: entity.Car{
					Model: "model",
					Year:  "year",
				},
				Amount: 185.0,
				CNH:    "cnh",
			},
			School: entity.School{
				Name: "Escola",
				Address: entity.Address{
					Street: "Avenida Barão de Alagoas",
					Number: "223",
					ZIP:    "08120000",
				},
			},
			Child: entity.Child{
				Responsible: entity.Responsible{
					Name:            "Responsible",
					CPF:             "cpf",
					PaymentMethodId: "pm_1PxgSrLfFDLpePGLDi2tnFEm",
					CustomerId:      "cus_QpL9XTVfM6sBhD",
					Address: entity.Address{
						Street: "Rua Masato Sakai",
						Number: "180",
						ZIP:    "008538300",
					},
				},
				Name: "Child",
				RG:   "RG",
			},
		}

		cou.On("GetSimpleContractByTitle", context.Background(), &contract.StripeSubscription.Title).Return(nil, fmt.Errorf("error"))

		useCase := NewContractUseCase(cou, st, googleAdapter, nil)

		err := useCase.Create(context.Background(), &contract)

		assert.Error(t, err)
	})

	t.Run("when title already exists", func(t *testing.T) {
		cou := mocks.NewIContractRepository(t)
		st := payments.NewStripeContract()
		googleAdapter := adapter.NewGoogleAdapter()

		contract := entity.Contract{
			Driver: entity.Driver{
				Name: "Motorista",
				Pix: entity.Pix{
					Key: "key",
				},
				Bank: entity.Bank{
					Account: "account",
					Agency:  "agency",
					Name:    "bank",
				},
				Car: entity.Car{
					Model: "model",
					Year:  "year",
				},
				Amount: 185.0,
				CNH:    "cnh",
			},
			School: entity.School{
				Name: "Escola",
				Address: entity.Address{
					Street: "Avenida Barão de Alagoas",
					Number: "223",
					ZIP:    "08120000",
				},
			},
			Child: entity.Child{
				Responsible: entity.Responsible{
					Name:            "Responsible",
					CPF:             "cpf",
					PaymentMethodId: "pm_1PxgSrLfFDLpePGLDi2tnFEm",
					CustomerId:      "cus_QpL9XTVfM6sBhD",
					Address: entity.Address{
						Street: "Rua Masato Sakai",
						Number: "180",
						ZIP:    "008538300",
					},
				},
				Name: "Child",
				RG:   "RG",
			},
		}

		cou.On("GetSimpleContractByTitle", context.Background(), &contract.StripeSubscription.Title).Return(&contract, nil)

		useCase := NewContractUseCase(cou, st, googleAdapter, nil)

		err := useCase.Create(context.Background(), &contract)
		assert.Error(t, err)
	})

	t.Run("when doesnt had payment method", func(t *testing.T) {
		cou := mocks.NewIContractRepository(t)
		st := payments.NewStripeContract()
		googleAdapter := adapter.NewGoogleAdapter()

		contract := entity.Contract{
			Driver: entity.Driver{
				Name: "Motorista",
				Pix: entity.Pix{
					Key: "key",
				},
				Bank: entity.Bank{
					Account: "account",
					Agency:  "agency",
					Name:    "bank",
				},
				Car: entity.Car{
					Model: "model",
					Year:  "year",
				},
				Amount: 185.0,
				CNH:    "cnh",
			},
			School: entity.School{
				Name: "Escola",
				Address: entity.Address{
					Street: "Avenida Barão de Alagoas",
					Number: "223",
					ZIP:    "08120000",
				},
			},
			Child: entity.Child{
				Responsible: entity.Responsible{
					Name:       "Responsible",
					CPF:        "cpf",
					CustomerId: "cus_QpL9XTVfM6sBhD",
					Address: entity.Address{
						Street: "Rua Masato Sakai",
						Number: "180",
						ZIP:    "008538300",
					},
				},
				Name: "Child",
				RG:   "RG",
			},
		}

		cou.On("GetSimpleContractByTitle", context.Background(), &contract.StripeSubscription.Title).Return(&entity.Contract{}, nil)

		useCase := NewContractUseCase(cou, st, googleAdapter, nil)

		err := useCase.Create(context.Background(), &contract)
		assert.Error(t, err)
	})

	t.Run("when doesnt had pix or bank", func(t *testing.T) {
		cou := mocks.NewIContractRepository(t)
		st := payments.NewStripeContract()
		googleAdapter := adapter.NewGoogleAdapter()

		contract := entity.Contract{
			Driver: entity.Driver{
				Name: "Motorista",
				Car: entity.Car{
					Model: "model",
					Year:  "year",
				},
				Amount: 185.0,
				CNH:    "cnh",
			},
			School: entity.School{
				Name: "Escola",
				Address: entity.Address{
					Street: "Avenida Barão de Alagoas",
					Number: "223",
					ZIP:    "08120000",
				},
			},
			Child: entity.Child{
				Responsible: entity.Responsible{
					Name:            "Responsible",
					CPF:             "cpf",
					PaymentMethodId: "pm_1PxgSrLfFDLpePGLDi2tnFEm",
					CustomerId:      "cus_QpL9XTVfM6sBhD",
					Address: entity.Address{
						Street: "Rua Masato Sakai",
						Number: "180",
						ZIP:    "008538300",
					},
				},
				Name: "Child",
				RG:   "RG",
			},
		}

		cou.On("GetSimpleContractByTitle", context.Background(), &contract.StripeSubscription.Title).Return(&entity.Contract{}, nil)

		useCase := NewContractUseCase(cou, st, googleAdapter, nil)

		err := useCase.Create(context.Background(), &contract)
		assert.Error(t, err)
	})

	t.Run("when doesnt had car and year", func(t *testing.T) {
		cou := mocks.NewIContractRepository(t)
		st := payments.NewStripeContract()
		googleAdapter := adapter.NewGoogleAdapter()

		contract := entity.Contract{
			Driver: entity.Driver{
				Name: "Motorista",
				Pix: entity.Pix{
					Key: "key",
				},
				Bank: entity.Bank{
					Account: "account",
					Agency:  "agency",
					Name:    "bank",
				},
				Amount: 185.0,
				CNH:    "cnh",
			},
			School: entity.School{
				Name: "Escola",
				Address: entity.Address{
					Street: "Avenida Barão de Alagoas",
					Number: "223",
					ZIP:    "08120000",
				},
			},
			Child: entity.Child{
				Responsible: entity.Responsible{
					Name:            "Responsible",
					CPF:             "cpf",
					PaymentMethodId: "pm_1PxgSrLfFDLpePGLDi2tnFEm",
					CustomerId:      "cus_QpL9XTVfM6sBhD",
					Address: entity.Address{
						Street: "Rua Masato Sakai",
						Number: "180",
						ZIP:    "008538300",
					},
				},
				Name: "Child",
				RG:   "RG",
			},
		}

		cou.On("GetSimpleContractByTitle", context.Background(), &contract.StripeSubscription.Title).Return(&entity.Contract{}, nil)

		useCase := NewContractUseCase(cou, st, googleAdapter, nil)

		err := useCase.Create(context.Background(), &contract)
		assert.Error(t, err)
	})

	t.Run("when get distance return error", func(t *testing.T) {
		cou := mocks.NewIContractRepository(t)
		st := payments.NewStripeContract()
		googleAdapter := adapter.NewGoogleAdapter()

		contract := entity.Contract{
			Driver: entity.Driver{
				Name: "Motorista",
				Pix: entity.Pix{
					Key: "key",
				},
				Bank: entity.Bank{
					Account: "account",
					Agency:  "agency",
					Name:    "bank",
				},
				Car: entity.Car{
					Model: "model",
					Year:  "year",
				},
				Amount: 185.0,
				CNH:    "cnh",
			},
			School: entity.School{
				Name: "Escola",
				Address: entity.Address{
					Street: "DAHSUIDAJHIe Xuyrhauds",
					Number: "783425893",
					ZIP:    "842763842",
				},
			},
			Child: entity.Child{
				Responsible: entity.Responsible{
					Name:            "Responsible",
					CPF:             "cpf",
					PaymentMethodId: "pm_1PxgSrLfFDLpePGLDi2tnFEm",
					CustomerId:      "cus_QpL9XTVfM6sBhD",
					Address: entity.Address{
						Street: "Rua Vander Marcelo Freitas Juvenesso",
						Number: "892374983",
						ZIP:    "472389472",
					},
				},
				Name: "Child",
				RG:   "RG",
			},
		}

		cou.On("GetSimpleContractByTitle", context.Background(), &contract.StripeSubscription.Title).Return(&entity.Contract{}, nil)

		useCase := NewContractUseCase(cou, st, googleAdapter, nil)

		err := useCase.Create(context.Background(), &contract)
		assert.Error(t, err)
	})

	t.Run("when stripe product return error", func(t *testing.T) {
		cou := mocks.NewIContractRepository(t)
		st := mocks.NewIStripe(t)
		googleAdapter := adapter.NewGoogleAdapter()

		contract := entity.Contract{
			Driver: entity.Driver{
				Name: "Motorista",
				Pix: entity.Pix{
					Key: "key",
				},
				Bank: entity.Bank{
					Account: "account",
					Agency:  "agency",
					Name:    "bank",
				},
				Car: entity.Car{
					Model: "model",
					Year:  "year",
				},
				Amount: 185.0,
				CNH:    "cnh",
			},
			School: entity.School{
				Name: "Escola",
				Address: entity.Address{
					Street: "Avenida Barão de Alagoas",
					Number: "223",
					ZIP:    "08120000",
				},
			},
			Child: entity.Child{
				Responsible: entity.Responsible{
					Name:            "Responsible",
					CPF:             "cpf",
					PaymentMethodId: "pm_1PxgSrLfFDLpePGLDi2tnFEm",
					CustomerId:      "cus_QpL9XTVfM6sBhD",
					Address: entity.Address{
						Street: "Rua Masato Sakai",
						Number: "180",
						ZIP:    "008538300",
					},
				},
				Name: "Child",
				RG:   "RG",
			},
		}

		cou.On("GetSimpleContractByTitle", context.Background(), &contract.StripeSubscription.Title).Return(&entity.Contract{}, nil)
		st.On("CreateProduct", &contract).Return(nil, fmt.Errorf("error"))

		useCase := NewContractUseCase(cou, st, googleAdapter, nil)

		err := useCase.Create(context.Background(), &contract)
		assert.Error(t, err)
	})

	t.Run("when stripe price return error", func(t *testing.T) {
		cou := mocks.NewIContractRepository(t)
		st := mocks.NewIStripe(t)
		googleAdapter := adapter.NewGoogleAdapter()

		contract := entity.Contract{
			Driver: entity.Driver{
				Name: "Motorista",
				Pix: entity.Pix{
					Key: "key",
				},
				Bank: entity.Bank{
					Account: "account",
					Agency:  "agency",
					Name:    "bank",
				},
				Car: entity.Car{
					Model: "model",
					Year:  "year",
				},
				Amount: 185.0,
				CNH:    "cnh",
			},
			School: entity.School{
				Name: "Escola",
				Address: entity.Address{
					Street: "Avenida Barão de Alagoas",
					Number: "223",
					ZIP:    "08120000",
				},
			},
			Child: entity.Child{
				Responsible: entity.Responsible{
					Name:            "Responsible",
					CPF:             "cpf",
					PaymentMethodId: "pm_1PxgSrLfFDLpePGLDi2tnFEm",
					CustomerId:      "cus_QpL9XTVfM6sBhD",
					Address: entity.Address{
						Street: "Rua Masato Sakai",
						Number: "180",
						ZIP:    "008538300",
					},
				},
				Name: "Child",
				RG:   "RG",
			},
		}
		cou.On("GetSimpleContractByTitle", context.Background(), &contract.StripeSubscription.Title).Return(&entity.Contract{}, nil)
		st.On("CreateProduct", &contract).Return(&stripe.Product{
			ID: "prod_QzgBXR7xS7nyQ4",
		}, nil)
		st.On("CreatePrice", &contract).Return(nil, fmt.Errorf("price error"))

		useCase := NewContractUseCase(cou, st, googleAdapter, nil)

		err := useCase.Create(context.Background(), &contract)
		assert.Error(t, err)
	})

	t.Run("when stripe subscription return error", func(t *testing.T) {
		cou := mocks.NewIContractRepository(t)
		st := mocks.NewIStripe(t)
		googleAdapter := adapter.NewGoogleAdapter()

		contract := entity.Contract{
			Driver: entity.Driver{
				Name: "Motorista",
				Pix: entity.Pix{
					Key: "key",
				},
				Bank: entity.Bank{
					Account: "account",
					Agency:  "agency",
					Name:    "bank",
				},
				Car: entity.Car{
					Model: "model",
					Year:  "year",
				},
				Amount: 185.0,
				CNH:    "cnh",
			},
			School: entity.School{
				Name: "Escola",
				Address: entity.Address{
					Street: "Avenida Barão de Alagoas",
					Number: "223",
					ZIP:    "08120000",
				},
			},
			Child: entity.Child{
				Responsible: entity.Responsible{
					Name:            "Responsible",
					CPF:             "cpf",
					PaymentMethodId: "pm_1PxgSrLfFDLpePGLDi2tnFEm",
					CustomerId:      "cus_QpL9XTVfM6sBhD",
					Address: entity.Address{
						Street: "Rua Masato Sakai",
						Number: "180",
						ZIP:    "008538300",
					},
				},
				Name: "Child",
				RG:   "RG",
			},
		}
		cou.On("GetSimpleContractByTitle", context.Background(), &contract.StripeSubscription.Title).Return(&entity.Contract{}, nil)
		st.On("CreateProduct", &contract).Return(&stripe.Product{
			ID: "prod_QzgBXR7xS7nyQ4",
		}, nil)
		st.On("CreatePrice", &contract).Return(&stripe.Price{
			ID: "price_1Q7gooLfFDLpePGL9xAzzxKK",
		}, nil)
		st.On("CreateSubscription", &contract).Return(nil, fmt.Errorf("subscription error"))

		useCase := NewContractUseCase(cou, st, googleAdapter, nil)

		err := useCase.Create(context.Background(), &contract)
		assert.Error(t, err)
	})

	t.Run("when repository return error", func(t *testing.T) {
		cou := mocks.NewIContractRepository(t)
		st := mocks.NewIStripe(t)
		googleAdapter := adapter.NewGoogleAdapter()

		contract := entity.Contract{
			Driver: entity.Driver{
				Name: "Motorista",
				Pix: entity.Pix{
					Key: "key",
				},
				Bank: entity.Bank{
					Account: "account",
					Agency:  "agency",
					Name:    "bank",
				},
				Car: entity.Car{
					Model: "model",
					Year:  "year",
				},
				Amount: 185.0,
				CNH:    "cnh",
			},
			School: entity.School{
				Name: "Escola",
				Address: entity.Address{
					Street: "Avenida Barão de Alagoas",
					Number: "223",
					ZIP:    "08120000",
				},
			},
			Child: entity.Child{
				Responsible: entity.Responsible{
					Name:            "Responsible",
					CPF:             "cpf",
					PaymentMethodId: "pm_1PxgSrLfFDLpePGLDi2tnFEm",
					CustomerId:      "cus_QpL9XTVfM6sBhD",
					Address: entity.Address{
						Street: "Rua Masato Sakai",
						Number: "180",
						ZIP:    "008538300",
					},
				},
				Name: "Child",
				RG:   "RG",
			},
		}

		cou.On("GetSimpleContractByTitle", context.Background(), &contract.StripeSubscription.Title).Return(&entity.Contract{}, nil)
		cou.On("Create", context.Background(), &contract).Return(fmt.Errorf("create repository error"))
		st.On("CreateProduct", &contract).Return(&stripe.Product{
			ID: "prod_QzgBXR7xS7nyQ4",
		}, nil)
		st.On("CreatePrice", &contract).Return(&stripe.Price{
			ID: "price_1Q7gooLfFDLpePGL9xAzzxKK",
		}, nil)
		st.On("CreateSubscription", &contract).Return(&stripe.Subscription{
			ID: "sub_1Q7gooLfFDLpePGLkcm2nPkC",
		}, nil)

		useCase := NewContractUseCase(cou, st, googleAdapter, nil)

		err := useCase.Create(context.Background(), &contract)
		if err == nil {
			t.Errorf("Error: %s", err)
		}
	})

}

func TestContract_Get(t *testing.T) {
	t.Run("when get return success", func(t *testing.T) {
		cou := mocks.NewIContractRepository(t)
		st := mocks.NewIStripe(t)
		googleAdapter := adapter.NewGoogleAdapter()

		id, _ := uuid.NewV7()

		contract := entity.Contract{
			Driver: entity.Driver{
				Name: "Motorista",
				Pix: entity.Pix{
					Key: "key",
				},
				Bank: entity.Bank{
					Account: "account",
					Agency:  "agency",
					Name:    "bank",
				},
				Car: entity.Car{
					Model: "model",
					Year:  "year",
				},
				Amount: 185.0,
				CNH:    "cnh",
			},
			School: entity.School{
				Name: "Escola",
				Address: entity.Address{
					Street: "Avenida Barão de Alagoas",
					Number: "223",
					ZIP:    "08120000",
				},
			},
			Child: entity.Child{
				Responsible: entity.Responsible{
					Name:            "Responsible",
					CPF:             "cpf",
					PaymentMethodId: "pm_1PxgSrLfFDLpePGLDi2tnFEm",
					CustomerId:      "cus_QpL9XTVfM6sBhD",
					Address: entity.Address{
						Street: "Rua Masato Sakai",
						Number: "180",
						ZIP:    "008538300",
					},
				},
				Name: "Child",
				RG:   "RG",
			},
		}

		cou.On("Get", context.Background(), id).Return(&contract, nil)
		st.On("ListInvoices", mock.Anything).Return([]entity.InvoiceInfo{}, nil)

		useCase := NewContractUseCase(cou, st, googleAdapter, nil)

		_, err := useCase.Get(context.Background(), id)
		if err != nil {
			t.Errorf("Error: %s", err)
		}

	})

	t.Run("when list invoices return fails", func(t *testing.T) {
		cou := mocks.NewIContractRepository(t)
		st := mocks.NewIStripe(t)
		googleAdapter := adapter.NewGoogleAdapter()

		id, _ := uuid.NewV7()

		contract := entity.Contract{
			Driver: entity.Driver{
				Name: "Motorista",
				Pix: entity.Pix{
					Key: "key",
				},
				Bank: entity.Bank{
					Account: "account",
					Agency:  "agency",
					Name:    "bank",
				},
				Car: entity.Car{
					Model: "model",
					Year:  "year",
				},
				Amount: 185.0,
				CNH:    "cnh",
			},
			School: entity.School{
				Name: "Escola",
				Address: entity.Address{
					Street: "Avenida Barão de Alagoas",
					Number: "223",
					ZIP:    "08120000",
				},
			},
			Child: entity.Child{
				Responsible: entity.Responsible{
					Name:            "Responsible",
					CPF:             "cpf",
					PaymentMethodId: "pm_1PxgSrLfFDLpePGLDi2tnFEm",
					CustomerId:      "cus_QpL9XTVfM6sBhD",
					Address: entity.Address{
						Street: "Rua Masato Sakai",
						Number: "180",
						ZIP:    "008538300",
					},
				},
				Name: "Child",
				RG:   "RG",
			},
		}

		cou.On("Get", context.Background(), id).Return(&contract, nil)
		st.On("ListInvoices", mock.Anything).Return(nil, fmt.Errorf("list invoices error"))

		useCase := NewContractUseCase(cou, st, googleAdapter, nil)

		_, err := useCase.Get(context.Background(), id)
		if err == nil {
			t.Errorf("Error: %s", err)
		}
	})

	t.Run("when get return fails", func(t *testing.T) {
		cou := mocks.NewIContractRepository(t)
		st := mocks.NewIStripe(t)
		googleAdapter := adapter.NewGoogleAdapter()

		id, _ := uuid.NewV7()

		cou.On("Get", context.Background(), id).Return(nil, fmt.Errorf("get error"))

		useCase := NewContractUseCase(cou, st, googleAdapter, nil)

		_, err := useCase.Get(context.Background(), id)
		if err == nil {
			t.Errorf("Error: %s", err)
		}
	})
}

func TestContract_FindAllByRg(t *testing.T) {
	t.Run("when get return success", func(t *testing.T) {
		cou := mocks.NewIContractRepository(t)
		st := mocks.NewIStripe(t)
		googleAdapter := adapter.NewGoogleAdapter()

		rg := "RG"

		contract := entity.Contract{
			Driver: entity.Driver{
				Name: "Motorista",
				Pix: entity.Pix{
					Key: "key",
				},
				Bank: entity.Bank{
					Account: "account",
					Agency:  "agency",
					Name:    "bank",
				},
				Car: entity.Car{
					Model: "model",
					Year:  "year",
				},
				Amount: 185.0,
				CNH:    "cnh",
			},
			School: entity.School{
				Name: "Escola",
				Address: entity.Address{
					Street: "Avenida Barão de Alagoas",
					Number: "223",
					ZIP:    "08120000",
				},
			},
			Child: entity.Child{
				ID: 1,
				Responsible: entity.Responsible{
					Name:            "Responsible",
					CPF:             "cpf",
					PaymentMethodId: "pm_1PxgSrLfFDLpePGLDi2tnFEm",
					CustomerId:      "cus_QpL9XTVfM6sBhD",
					Address: entity.Address{
						Street: "Rua Masato Sakai",
						Number: "180",
						ZIP:    "008538300",
					},
				},
				Name: "Child",
				RG:   "112223334",
			},
		}

		cou.On("FindAllByRg", context.Background(), &rg).Return(&contract, nil)

		useCase := NewContractUseCase(cou, st, googleAdapter, nil)

		_, err := useCase.FindAllByRg(context.Background(), &rg)
		if err != nil {
			t.Errorf("Error: %s", err)
		}

	})

	t.Run("when get return fails", func(t *testing.T) {
		cou := mocks.NewIContractRepository(t)
		st := mocks.NewIStripe(t)
		googleAdapter := adapter.NewGoogleAdapter()

		rg := "RG"

		cou.On("FindAllByRg", context.Background(), &rg).Return(nil, fmt.Errorf("get error"))

		useCase := NewContractUseCase(cou, st, googleAdapter, nil)

		_, err := useCase.FindAllByRg(context.Background(), &rg)
		if err == nil {
			t.Errorf("Error: %s", err)
		}
	})

}

func TestContract_FindAllByCnh(t *testing.T) {

	t.Run("when get return success", func(t *testing.T) {
		cou := mocks.NewIContractRepository(t)
		st := mocks.NewIStripe(t)
		googleAdapter := adapter.NewGoogleAdapter()

		cnh := "CNH"

		cou.On("FindAllByCnh", context.Background(), &cnh).Return([]entity.Contract{}, nil)
		useCase := NewContractUseCase(cou, st, googleAdapter, nil)
		_, err := useCase.FindAllByCnh(context.Background(), &cnh)
		if err != nil {
			t.Errorf("Error: %s", err)
		}
	})

	t.Run("when get return fails", func(t *testing.T) {
		cou := mocks.NewIContractRepository(t)
		st := mocks.NewIStripe(t)
		googleAdapter := adapter.NewGoogleAdapter()

		cnh := "CNH"

		cou.On("FindAllByCnh", context.Background(), &cnh).Return(nil, fmt.Errorf("get error"))

		useCase := NewContractUseCase(cou, st, googleAdapter, nil)

		_, err := useCase.FindAllByCnh(context.Background(), &cnh)
		if err == nil {
			t.Errorf("Error: %s", err)
		}
	})

}

func TestContract_FindAllByCpf(t *testing.T) {

	t.Run("when get return success", func(t *testing.T) {
		cou := mocks.NewIContractRepository(t)
		st := mocks.NewIStripe(t)
		googleAdapter := adapter.NewGoogleAdapter()

		cpf := "CPF"

		cou.On("FindAllByCpf", context.Background(), &cpf).Return([]entity.Contract{}, nil)
		useCase := NewContractUseCase(cou, st, googleAdapter, nil)
		_, err := useCase.FindAllByCpf(context.Background(), &cpf)
		if err != nil {
			t.Errorf("Error: %s", err)
		}
	})

	t.Run("when get return fails", func(t *testing.T) {
		cou := mocks.NewIContractRepository(t)
		st := mocks.NewIStripe(t)
		googleAdapter := adapter.NewGoogleAdapter()

		cpf := "Cpf"

		cou.On("FindAllByCpf", context.Background(), &cpf).Return(nil, fmt.Errorf("get error"))

		useCase := NewContractUseCase(cou, st, googleAdapter, nil)

		_, err := useCase.FindAllByCpf(context.Background(), &cpf)
		if err == nil {
			t.Errorf("Error: %s", err)
		}
	})

}

func TestContract_FindAllByCnpj(t *testing.T) {

	t.Run("when get return success", func(t *testing.T) {
		cou := mocks.NewIContractRepository(t)
		st := mocks.NewIStripe(t)
		googleAdapter := adapter.NewGoogleAdapter()

		cnpj := "CNPJ"

		cou.On("FindAllByCnpj", context.Background(), &cnpj).Return([]entity.Contract{}, nil)
		useCase := NewContractUseCase(cou, st, googleAdapter, nil)
		_, err := useCase.FindAllByCnpj(context.Background(), &cnpj)
		if err != nil {
			t.Errorf("Error: %s", err)
		}
	})

	t.Run("when get return fails", func(t *testing.T) {
		cou := mocks.NewIContractRepository(t)
		st := mocks.NewIStripe(t)
		googleAdapter := adapter.NewGoogleAdapter()

		cnpj := "CNPJ"

		cou.On("FindAllByCnpj", context.Background(), &cnpj).Return(nil, fmt.Errorf("get error"))

		useCase := NewContractUseCase(cou, st, googleAdapter, nil)

		_, err := useCase.FindAllByCnpj(context.Background(), &cnpj)
		if err == nil {
			t.Errorf("Error: %s", err)
		}
	})

}
