package main

import "testing"

func TestAbstractFactory(t *testing.T) {
	var factory AbstractFactory
	var tv ITelevision
	var air IAirConditioner

	factory = &HuaWeiFactory{}           // 使用华为工厂创建产品
	tv = factory.CreateTelevision()      // 创建华为电视
	air = factory.CreateAirConditioner() // 创建华为空调
	tv.Watch()                           // 观看华为电视
	air.SetTemperature(25)               // 设置华为空调温度

	factory = &MiFactory{}               // 使用小米工厂创建产品
	tv = factory.CreateTelevision()      // 创建小米电视
	air = factory.CreateAirConditioner() // 创建小米空调
	tv.Watch()                           // 观看小米电视
	air.SetTemperature(30)               // 设置小米空调温度
}
