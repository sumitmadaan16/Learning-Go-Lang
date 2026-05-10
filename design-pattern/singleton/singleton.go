package singleton

import (
	"fmt"
	"sync"
)

type singleton struct {
	val int
}

var (
	instance *singleton
	once     sync.Once
)

func getInstance() *singleton {
	once.Do(func() {
		instance = &singleton{val: 50}
	})
	return instance
}

func SingletonPattern() {
	first := getInstance()
	fmt.Println("value of first is: ", first.val)
	second := getInstance()
	fmt.Println("value of first is: ", second.val)

	if first == second {
		fmt.Println("singleton pattern is same")
	}
}
