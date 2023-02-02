package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/*
@author: shg
@since: 2023/2/3 12:20 AM
@mail: shgang97@163.com
*/
var (
	instance    *singleton
	m           sync.Mutex
	initialized uint32
)

func main() {
	s := GetInstance()
	s.name = "singleton"
	s.someFunc()
}

type singleton struct {
	name string // 有用的单例数据
}

func GetInstance() *singleton {
	if atomic.LoadUint32(&initialized) == 1 {
		return instance
	}
	m.Lock()
	defer m.Unlock()
	if instance == nil {
		defer atomic.StoreUint32(&initialized, 1)
		instance = new(singleton)
	}
	return instance
}

func (s *singleton) someFunc() {
	fmt.Printf("单例类的某个对象方法。单例名称为：%s\n", s.name)
}
