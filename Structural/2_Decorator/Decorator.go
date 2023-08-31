package main

import (
	"fmt"
	"time"
)

// PS5 产品接口
type PS5 interface {
	StartGPUEngine()
	GetPrice() int64
}

// CD版本
type PS5WithCD struct{}

func (p PS5WithCD) StartGPUEngine() {
	fmt.Println("start engine")
}
func (p PS5WithCD) GetPrice() int64 {
	return 5000
}

// 数字版本
type PS5WithDigital struct{}

func (p PS5WithDigital) StartGPUEngine() {
	fmt.Println("start normal gpu engine")
}

func (p PS5WithDigital) GetPrice() int64 {
	return 3600
}

// Plus 版的装饰器
type PS5MachinePlus struct {
	ps5Machine PS5
}

// 将基础的信息传递给装饰器
func (p *PS5MachinePlus) SetPS5Machine(ps5 PS5) {
	p.ps5Machine = ps5
}

func (p PS5MachinePlus) StartGPUEngine() {
	p.ps5Machine.StartGPUEngine()
	fmt.Println("开启 plus 插件")
}

func (p PS5MachinePlus) GetPrice() int64 {
	return p.ps5Machine.GetPrice() + 500
}

// 主题色版的装饰器
type PS5WithTopicColor struct {
	ps5Machine PS5
}

// 将基础的信息传递给装饰器
func (p *PS5WithTopicColor) SetPS5Machine(ps5 PS5) {
	p.ps5Machine = ps5
}

func (p PS5WithTopicColor) StartGPUEngine() {
	p.ps5Machine.StartGPUEngine()
	fmt.Println("更改为主题色主机")
}
func (p PS5WithTopicColor) GetPrice() int64 {
	return p.ps5Machine.GetPrice() + 200
}

func main() {
	var num int
	fmt.Println("1、PS5 CD 版本：5000 元")
	fmt.Println("2、PS5 数字版本：3600 元")
	fmt.Print("请输入想要的版本: ")
	fmt.Scan(&num)
	switch num {
	case 1:
		PS5 := PS5WithCD{}
		decorator1(decorator(PS5))
	case 2:
		PS5 := PS5WithDigital{}
		decorator1(decorator(PS5))
	default:
		fmt.Println("输入错误，没有这种机器")
		return
	}

}

func decorator(PS5 PS5) PS5 {
	clearScreen()
	fmt.Println("当前价格为：", PS5.GetPrice())
	var num int
	fmt.Println("1、升级Plus插件：500 元")
	fmt.Println("2、不进行添加，继续")
	fmt.Print("请输入想要的版本: ")
	fmt.Scan(&num)
	switch num {
	case 1:
		clearScreen()
		ps5MachinePlus := PS5MachinePlus{}
		ps5MachinePlus.SetPS5Machine(PS5)
		ps5MachinePlus.StartGPUEngine()
		ps5MachinePlus.GetPrice()
		return ps5MachinePlus
	case 2:
		return PS5
	default:
		fmt.Println("输入错误，没有这种机器")
		return PS5
	}
}

func decorator1(PS5 PS5) {
	clearScreen()
	fmt.Println("当前价格为：", PS5.GetPrice())
	var num int
	fmt.Println("1、升级主题色：200 元")
	fmt.Println("2、不进行添加，继续")
	fmt.Print("请输入想要的版本: ")
	fmt.Scan(&num)
	switch num {
	case 1:
		clearScreen()
		ps5WithTopicColor := PS5WithTopicColor{}
		ps5WithTopicColor.SetPS5Machine(PS5)
		ps5WithTopicColor.StartGPUEngine()
		price := ps5WithTopicColor.GetPrice()
		if price == 5700 || price == 4300 {
			fmt.Printf("\nPS5 豪华Plus + 经典主题配色 版，价格: %d 元\n", price)
		} else {
			fmt.Printf("\nPS5  经典主题配色 版，价格: %d 元\n", price)

		}

	case 2:
		if PS5.GetPrice() == 4100 || PS5.GetPrice() == 5500 {
			fmt.Printf("\nPS5  豪华Plus版，价格: %d 元\n", PS5.GetPrice())
		} else {
			fmt.Printf("\nPS5  常规版，价格: %d 元\n", PS5.GetPrice())
		}
	default:
		fmt.Println("输入错误，没有这种输入")
	}
}

func clearScreen() {
	time.Sleep(1 * time.Second) // 等待 1 秒，给用户阅读之前的输出信息
	fmt.Print("\033[H\033[2J")  // 使用特殊的控制字符清理命令行界面
}
