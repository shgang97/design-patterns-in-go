package main

import "fmt"

/*
@author: shg
@since: 2023/2/3 12:49 AM
@mail: shgang97@163.com
*/
func main() {
	s := GetInstance()
	s.name = "singleton"
	s.someFunc()
}

type singleton struct {
	name string
}

var instance *singleton

func GetInstance() *singleton {
	// 只有首次GetInstance()方法被调用，才会创建单例对象
	if instance == nil {
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
