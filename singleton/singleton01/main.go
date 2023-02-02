package main

import "fmt"

/*
@author: shg
@since: 2023/2/3 12:20 AM
@mail: shgang97@163.com
*/
func main() {
	s := GetInstance()
	s.name = "singleton"
	s.someFunc()
}

/*
1、保证这个类非公有化，外界不能通过这个类直接创建一个对象,类名称首字母要小写
*/
type singleton struct {
	name string // 有用的单例数据
}

/*
2、但是还要有一个指针可以指向这个唯一对象，但是这个指针永远不能改变方向
Golang中没有常指针概念，所以只能通过将这个指针私有化不让外部模块访问
*/
var instance *singleton = new(singleton)

/*
GetInstance
3、如果全部为私有化，那么外部模块将永远无法访问到这个类和对象，
所以需要对外提供一个方法来获取这个唯一实例对象
注意：这个方法不可以定义为 singleton 的一个成员方法：

	因为如果为成员方法就必须要先访问对象、再访问函数
	但是类和对象目前都已经私有化，外界无法访问，所以这个方法一定是一个全局普通函数
*/
func GetInstance() *singleton {
	return instance
}

/*
其他有用的单例方法
*/
func (s *singleton) someFunc() {
	fmt.Printf("单例类的某个对象方法。单例名称为：%s\n", s.name)
}