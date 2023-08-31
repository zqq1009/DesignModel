package main

import (
	"fmt"
)

// 原型接口
type Prototype interface {
	Clone() Prototype
}

// 具体原型A
type ConcretePrototypeA struct {
	name string
}

// 深拷贝
func (c *ConcretePrototypeA) Clone() Prototype {
	clone := *c
	return &clone
}

// 具体原型B
type ConcretePrototypeB struct {
	number int
}

// 浅拷贝
func (c *ConcretePrototypeB) Clone() Prototype {
	return c
}

func main() {
	// 创建具体原型A对象
	prototypeA := &ConcretePrototypeA{name: "1_Prototype A"}
	// 使用深拷贝创建克隆对象
	cloneA := prototypeA.Clone().(*ConcretePrototypeA)
	fmt.Println(cloneA.name)

	// 创建具体原型B对象
	prototypeB := &ConcretePrototypeB{number: 100}
	// 使用浅拷贝创建克隆对象
	cloneB := prototypeB.Clone().(*ConcretePrototypeB)
	fmt.Println(cloneB.number)

	// 修改原型对象的值
	prototypeA.name = "New 1_Prototype A"
	prototypeB.number = 200

	// 打印克隆对象的值，观察深拷贝和浅拷贝的影响
	fmt.Println(cloneA.name)   // 深拷贝不受原型对象修改的影响，仍然保持原来的值
	fmt.Println(cloneB.number) // 浅拷贝会受原型对象修改的影响，值也会发生改变
}
