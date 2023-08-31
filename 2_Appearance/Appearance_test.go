package main

import (
	"fmt"
	"testing"
	"time"
)

func TestAppearance(t *testing.T) {
	computer := NewComputerFacade()
	fmt.Printf("是否开机，开机请按 1，输入其他则结束: ")
	var num int
	fmt.Scan(&num)
	if num == 1 {
		computer.start()
		time.Sleep(2 * time.Second)
		fmt.Println("\n 欢迎进入系统页面")
	}

}
