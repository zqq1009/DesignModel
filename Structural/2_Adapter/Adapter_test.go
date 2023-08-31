package main

import "testing"

func TestAdapter(t *testing.T) {
	// 创建被适配者
	adaptee := NewAdaptee()

	// 创建适配器，并将被适配者传入
	adapter := NewAdapter(adaptee)

	// 调用适配器的目标方法
	adapter.Request()

	//不使用适配器的话
	/*
		// 创建具体的被适配者
		adaptee := &ConcreteAdaptee{}

		// 直接调用具体的被适配者的方法
		adaptee.SpecificRequest()
	*/
}
