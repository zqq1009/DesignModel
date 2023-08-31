package main

import (
	"fmt"
	"sync"
)

type Singleton struct {
	message string
}

var (
	instance *Singleton
	once     sync.Once
)

func GetInstance(name string) *Singleton {
	once.Do(func() {
		instance = &Singleton{
			message: name + "  Hello, I am a singleton instance!",
		}
	})
	return instance
}

func main() {
	s1 := GetInstance("s1")
	s2 := GetInstance("s2")

	fmt.Println(s1.message)
	fmt.Println(s2.message)

	fmt.Println(s1 == s2) // Output: true，s1和s2是同一个实例
}
