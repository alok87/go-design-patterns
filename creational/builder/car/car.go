package main

import (
	"fmt"
)

type Car interface {
	Stop()
	Move()
}

type car struct {
	reg   string
	brand string
}

func (c *car) Stop() {
	fmt.Println("Stopping car with reg# " + c.reg)
}

func (c *car) Move() {
	fmt.Println("Car running with reg# " + c.reg)
}

type CarBuilder interface {
	SetRegistration(string) CarBuilder
	SetBrand(string) CarBuilder
	Build() Car
}

type carBuilder struct {
	reg   string
	brand string
}

func (cb *carBuilder) SetRegistration(reg string) CarBuilder {
	cb.reg = reg
	return cb
}

func (cb *carBuilder) SetBrand(brand string) CarBuilder {
	cb.brand = brand
	return cb
}

//Builder pattern
func (cb *carBuilder) Build() Car {
	return &car{
		reg:   cb.reg,
		brand: cb.brand,
	}
}

//Factory method pattern
func NewCar() CarBuilder {
	return &carBuilder{}
}

func main() {
	builder := NewCar()
	car := builder.SetRegistration("KA 02E 5421").SetBrand("Hyundai").Build()
	car.Move()
	car.Stop()
}
