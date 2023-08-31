package main

import (
	"fmt"
)

// Target 是目标接口，定义了客户端所期望的方法
type Target interface {
	Request()
}

// Adaptee 是被适配的接口，即需要被适配的旧代码
type Adaptee interface {
	SpecificRequest()
}

// NewAdaptee 是被适配接口的工厂函数
func NewAdaptee() Adaptee {
	return &ConcreteAdaptee{}
}

// ConcreteAdaptee 是具体的被适配者，即旧代码的实现
type ConcreteAdaptee struct{}

// SpecificRequest 是目标类的一个方法
func (c *ConcreteAdaptee) SpecificRequest() {
	fmt.Println("Specific Request")
}

// Adapter 是适配器，实现了目标接口并包含被适配者的实例
type Adapter struct {
	adaptee Adaptee
}

func NewAdapter(adaptee Adaptee) Target {
	return &Adapter{
		adaptee: adaptee,
	}
}

func (a *Adapter) Request() {
	// 在适配器的目标方法中调用被适配者的方法
	a.adaptee.SpecificRequest()
}
