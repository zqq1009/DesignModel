package main

import (
	"fmt"
	"testing"
)

func TestFactoryModel(t *testing.T) {
	var factory OperatorFactory
	factory = &PlusOperatorFactory{}
	fmt.Println("加法：", Compute(factory, 3, 2))

	factory = &MultiOperatorFactory{}
	fmt.Println("乘法：", Compute(factory, 4, 2))
}

// 封装的函数,用来设置值以及返回答案
func Compute(factory OperatorFactory, numA, numB int) int {
	mathOp := factory.Create()
	mathOp.SetNumA(numA)
	mathOp.SetNumB(numB)
	return mathOp.ComputeResult()
}
