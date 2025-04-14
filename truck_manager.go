package main

import (
	"errors"
	"log"
)

var ErrTruckNotFound = errors.New("truck not found")

type FleetManager interface {
	AddTruck(id string, cargo int) error
	GetTruck(id string) (*Truck, error)
	RemoveTruck(id string) error
	UpdateTruckCargo(id string, cargo int) error
}

type Truck struct {
	ID    string
	Cargo int
}

type truckManager struct {
	trucks map[string]*Truck
}

func NewTruckManager() truckManager {
	return truckManager{
		trucks: make(map[string]*Truck),
	}
}

func (t *truckManager) AddTruck(id string, cargo int) error {
	trucks := make(map[string]*Truck)
	if _, exists := trucks[id]; exists {
		log.Println("truck already exists")
	} else {
		trucks[id] = &Truck{ID: id, Cargo: cargo}
		log.Println("truck has been added")
	}
	return nil
}
