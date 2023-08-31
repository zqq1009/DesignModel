package main

import "fmt"

// 抽象工厂接口
type AbstractFactory interface {
	CreateTelevision() ITelevision         // 创建电视的方法
	CreateAirConditioner() IAirConditioner // 创建空调的方法
}

// 电视接口
type ITelevision interface {
	Watch() // 观看电视的方法
}

// 空调接口
type IAirConditioner interface {
	SetTemperature(int) // 设置温度的方法
}

// 华为工厂
type HuaWeiFactory struct{}

// 华为电视
type HuaWeiTV struct{}

// 华为空调
type HuaWeiAirConditioner struct{}

func (hf *HuaWeiFactory) CreateTelevision() ITelevision {
	return &HuaWeiTV{} // 创建华为电视对象
}

func (ht *HuaWeiTV) Watch() {
	fmt.Println("Watch HuaWei TV") // 观看华为电视
}

func (hf *HuaWeiFactory) CreateAirConditioner() IAirConditioner {
	return &HuaWeiAirConditioner{} // 创建华为空调对象
}

func (ha *HuaWeiAirConditioner) SetTemperature(temp int) {
	fmt.Println("HuaWei AirConditioner set temperature to ", temp) // 设置华为空调温度
}

// 小米工厂
type MiFactory struct{}

// 小米电视
type MITV struct{}

// 小米空调
type MIAirConditioner struct{}

func (m *MiFactory) CreateTelevision() ITelevision {
	return &MITV{} // 创建小米电视对象
}

func (mt *MITV) Watch() {
	fmt.Println("Watch MI TV") // 观看小米电视
}

func (m *MiFactory) CreateAirConditioner() IAirConditioner {
	return &MIAirConditioner{} // 创建小米空调对象
}

func (ma *MIAirConditioner) SetTemperature(temp int) {
	fmt.Println("MI AirConditioner set temperature to ", temp) // 设置小米空调温度
}
