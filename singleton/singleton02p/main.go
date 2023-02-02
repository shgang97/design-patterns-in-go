package main

import (
	"fmt"
	"sync"
	"time"
)

/*
@author: shg
@since: 2023/2/3 12:49 AM
@mail: shgang97@163.com
*/
func main() {
	var wg sync.WaitGroup
	var singletons []*singleton
	var mutex sync.Mutex
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			instance := GetInstance()
			// slice 是非线程安全的，因此需要加锁来保证单例对象能够被加入到 slice 中
			mutex.Lock()
			singletons = append(singletons, instance)
			mutex.Unlock()
		}()
	}
	wg.Wait()
	for i, instance := range singletons {
		fmt.Printf("第%d个单例为%p\n", i, instance)
	}
}

type singleton struct {
	name string
}

var instance *singleton

func GetInstance() *singleton {
	// 只有首次GetInstance()方法被调用，才会创建单例对象
	if instance == nil {
		// 为了能够提高创建多个不同对象的该类，加入休眠时间
		time.Sleep(500)
		instance = new(singleton)
	}
	// 接下来的GetInstance直接返回已经申请的实例即可
	return instance
}

/*
其他有用的单例方法
*/
func (s *singleton) someFunc() {
	fmt.Printf("单例类的某个对象方法。单例名称为：%s\n", s.name)
}
