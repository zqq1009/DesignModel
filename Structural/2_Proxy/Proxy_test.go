package main

import "testing"

func TestCarProxy_Drive(t *testing.T) {
	car := NewCarProxy(&Driver{12})
	car.Drive() // 输出 Driver too young!
	car2 := NewCarProxy(&Driver{22})
	car2.Drive() // 输出 Car is being driven
}
