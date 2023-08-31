package main

import (
	"fmt"
	"testing"
)

func TestEsayFactory(t *testing.T) {
	printer := NewPrinter("en")
	fmt.Println(printer.Print("bob"))
}
