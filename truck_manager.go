package main

import (
	"errors"
	"sync"
)

var ErrTruckNotFound = errors.New("truck not found")

type FleetManager interface {
	AddTruck(id string, cargo int) error
	GetTruck(id string) (Truck, error)
	RemoveTruck(id string) error
	UpdateTruckCargo(id string, cargo int) error
}

type Truck struct {
	ID    string
	Cargo int
}

type truckManager struct {
	trucks map[string]*Truck
	sync.RWMutex
}

func NewTruckManager() truckManager {
	return truckManager{
		trucks: make(map[string]*Truck),
	}
}

func (t *truckManager) AddTruck(id string, cargo int) error {
	t.Lock()
	defer t.Unlock()

	t.trucks[id] = &Truck{ID: id, Cargo: cargo}
	return nil
}

func (t *truckManager) GetTruck(id string) (Truck, error) {
	t.RLock()
	defer t.RUnlock()

	truck, exists := t.trucks[id]
	if !exists {
		return Truck{}, ErrTruckNotFound
	}
	return *truck, nil
}

func (t *truckManager) RemoveTruck(id string) error {
	t.Lock()
	defer t.Unlock()

	delete(t.trucks, id)
	return nil
}

func (t *truckManager) UpdateTruckCargo(id string, cargo int) error {
	t.Lock()
	defer t.Unlock()

	t.trucks[id].Cargo = cargo
	return nil
}
