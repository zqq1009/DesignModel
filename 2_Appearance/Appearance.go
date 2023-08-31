package main

import (
	"fmt"
)

const (
	BOOT_ADDRESS = 0
	BOOT_SECTOR  = 0
	SECTOR_SIZE  = 0
)

// CPU 表示计算机的中央处理器
type CPU struct{}

// Freeze 模拟 CPU 的冻结操作
func (c *CPU) Freeze() {
	fmt.Println("CPU.Freeze()")
}

// Jump 模拟 CPU 的跳转操作
func (c *CPU) Jump(position int) {
	fmt.Println("CPU.Jump()")
}

// Execute 模拟 CPU 的执行操作
func (c *CPU) Execute() {
	fmt.Println("CPU.Execute()")
}

// Memory 表示计算机的内存
type Memory struct{}

// Load 模拟内存的加载操作
func (m *Memory) Load(position int, data []byte) {
	fmt.Println("Memory.Load()")
}

// HardDrive 表示计算机的硬盘
type HardDrive struct{}

// Read 模拟硬盘的读取操作
func (hd *HardDrive) Read(lba int, size int) []byte {
	fmt.Println("HardDrive.Read()")
	return make([]byte, 0)
}

// ComputerFacade 是计算机的外观类
type ComputerFacade struct {
	processor *CPU
	ram       *Memory
	hd        *HardDrive
}

// NewComputerFacade 创建并返回一个计算机外观对象
func NewComputerFacade() *ComputerFacade {
	return &ComputerFacade{new(CPU), new(Memory), new(HardDrive)}
}

// start 启动计算机
func (c *ComputerFacade) start() {
	ch := make(chan int)

	go func() {
		c.processor.Freeze()
		ch <- 0
	}()

	go func() {
		c.ram.Load(BOOT_ADDRESS, c.hd.Read(BOOT_SECTOR, SECTOR_SIZE))
	}()

	go func() {
		<-ch
		c.processor.Jump(BOOT_ADDRESS)
		c.processor.Execute()
	}()

}
