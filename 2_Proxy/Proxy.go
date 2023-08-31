package main

import "fmt"

type Car struct{}

type Vehicle interface {
	Drive()
}

type Driver struct {
	Age int
}

func (c *Car) Drive() {
	fmt.Println("Car is being driven")
}

type CarProxy struct {
	vehicle Vehicle
	driver  *Driver
}

func NewCarProxy(driver *Driver) *CarProxy {
	return &CarProxy{&Car{}, driver}
}

func (c *CarProxy) Drive() {
	if c.driver.Age >= 16 {
		c.vehicle.Drive()
	} else {
		fmt.Println("Driver too young!")
	}
}
