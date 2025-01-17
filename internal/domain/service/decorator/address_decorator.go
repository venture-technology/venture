package decorator

import (
	"fmt"
	"strconv"

	"github.com/venture-technology/venture/internal/domain/service/adapters"
	"github.com/venture-technology/venture/internal/infra/contracts"
)

type AddressDecorator struct {
	AddressAdapter adapters.AddressService
	Cache          contracts.Cacher
}

func NewAddressDecorator(addressAdapter adapters.AddressService, cache contracts.Cacher) AddressDecorator {
	return AddressDecorator{
		AddressAdapter: addressAdapter,
		Cache:          cache,
	}
}

func (d AddressDecorator) GetDistance(origin, destination string) (*float64, error) {
	price, err := d.getDistanceCache(origin, destination)
	if err != nil {
		return d.getDistanceAdapter(origin, destination)
	}

	return price, nil
}

func (d AddressDecorator) getDistanceCache(origin, destination string) (*float64, error) {
	priceStr, err := d.Cache.Get(fmt.Sprintf("%s:%s", origin, destination))
	if err != nil {
		return nil, err
	}

	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		return nil, err
	}

	return &price, nil
}

func (d AddressDecorator) getDistanceAdapter(origin, destination string) (*float64, error) {
	price, err := d.AddressAdapter.GetDistance(origin, destination)
	if err != nil {
		return nil, err
	}

	err = d.setDistanceCache(origin, destination, fmt.Sprintf("%f", *price))
	if err != nil {
		return nil, err
	}

	return price, nil
}

func (d AddressDecorator) setDistanceCache(origin, destination, price string) error {
	err := d.Cache.Set(fmt.Sprintf("%s:%s", origin, destination), price, 0)
	if err != nil {
		return err
	}

	err = d.Cache.Set(fmt.Sprintf("%s:%s", destination, origin), price, 0)
	if err != nil {
		return err
	}

	return nil
}
