package decorator

import (
	"fmt"
	"strconv"

	"github.com/venture-technology/venture/internal/domain/service/address"
	"github.com/venture-technology/venture/internal/infra/contracts"
)

type AddressDecorator struct {
	Address address.Address
	Cache   contracts.Cacher
}

func NewAddressDecorator(address address.Address, cache contracts.Cacher) AddressDecorator {
	return AddressDecorator{
		Address: address,
		Cache:   cache,
	}
}

func (d AddressDecorator) Distance(origin, destination string) (*float64, error) {
	price, err := d.distanceWithCache(origin, destination)
	if err != nil {
		return d.distanceWithoutCache(origin, destination)
	}

	return price, nil
}

func (d AddressDecorator) distanceWithCache(origin, destination string) (*float64, error) {
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

func (d AddressDecorator) distanceWithoutCache(origin, destination string) (*float64, error) {
	price, err := d.Address.Distance(origin, destination)
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
