package main

import "fmt"

// 客户端的选择
func NewPrinter(lang string) Printer {
	switch lang {
	case "cn":
		return new(CnPrinter)
	case "en":
		return new(EnPrinter)
	default:
		return new(Normal)
	}
}

// 简单工厂返回的接口
type Printer interface {
	Print(name string) string
}

// 功能的实现（CnPrinter , EnPrinter , Normal）
type CnPrinter struct{}

func (*CnPrinter) Print(name string) string {
	return fmt.Sprintf("你好，%s", name)
}

type EnPrinter struct{}

func (*EnPrinter) Print(name string) string {
	return fmt.Sprintf("hello，%s", name)
}

type Normal struct{}

func (*Normal) Print(name string) string {
	return fmt.Sprintf("没有这种语言，%s", name)
}
