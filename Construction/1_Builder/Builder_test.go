package main

import (
	"fmt"
	"testing"
)

func TestBuilder(t *testing.T) {
	builder := &ConcreteBuilder1{}
	director := &Director{
		builder: builder,
	}
	product := director.Construct()

	fmt.Printf("Brand: %s\n", product.Brand)
	fmt.Printf("Model: %s\n", product.Model)
	fmt.Printf("Color: %s\n", product.Color)
	fmt.Printf("Option: %s\n", product.Option)
}
