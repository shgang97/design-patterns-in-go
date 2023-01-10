package main

import "fmt"

/*
@author: shg
@since: 2023/1/10
@desc: //TODO
*/

func main() {
	var barrier Shield
	barrier = new(Barrier)
	fmt.Printf("未被加强的屏障增加了%d点防御力\n", barrier.EnhanceDefense())

	fmt.Printf("%s\n", "---------------------------------")
	barrier = NewSonaW(barrier)
	fmt.Printf("被娑娜W加强的屏障增加了%d点防御力\n", barrier.EnhanceDefense())

	fmt.Printf("%s\n", "---------------------------------")
	barrier = NewJannaE(barrier)
	fmt.Printf("又被迦娜E加强的屏障增加了%d点防御力\n", barrier.EnhanceDefense())

	fmt.Printf("%s\n", "---------------------------------")
	barrier = NewLuluE(barrier)
	fmt.Printf("又被露露E加强的屏障增加了%d点防御力\n", barrier.EnhanceDefense())

	fmt.Printf("%s\n", "---------------------------------")
	barrier = NewLuxannaW(barrier)
	fmt.Printf("又被光辉W加强的屏障增加了%d点防御力\n", barrier.EnhanceDefense())
}

// 抽象层

// Shield 抽象组件
type Shield interface {
	EnhanceDefense() int
}

// SupportShield 抽象装饰器层,持有被装饰组件的引用
type SupportShield struct {
	Shield Shield
}

func (shield *SupportShield) EnhanceDefense() int {
	return 0
}

// 实现层

// Barrier 英雄自身的防御技能，如召唤师技能屏障、自身的护盾技能
type Barrier struct {
}

func (barrier *Barrier) EnhanceDefense() int {
	return 10
}

// 具体的装饰器护盾技能

// SonaW 娑娜的W技能
type SonaW struct {
	SupportShield SupportShield // 继承抽象装饰器
}

func (sonaW *SonaW) EnhanceDefense() int {
	return 1 + sonaW.SupportShield.Shield.EnhanceDefense()
}

func NewSonaW(shield Shield) Shield {
	return &SonaW{SupportShield{shield}}
}

// JannaE 迦娜的E技能
type JannaE struct {
	SupportShield // 继承抽象装饰器
}

func (jannaE *JannaE) EnhanceDefense() int {
	return 2 + jannaE.SupportShield.Shield.EnhanceDefense()
}

func NewJannaE(shield Shield) Shield {
	return &JannaE{SupportShield{shield}}
}

// LuluE 露露的E技能
type LuluE struct {
	SupportShield // 继承抽象装饰器
}

func (luluE *LuluE) EnhanceDefense() int {
	return 3 + luluE.SupportShield.Shield.EnhanceDefense()
}

func NewLuluE(shield Shield) Shield {
	return &LuluE{SupportShield{shield}}
}

// LuxannaW 光辉的W技能
type LuxannaW struct {
	SupportShield // 继承抽象装饰器
}

func (luxannaW *LuxannaW) EnhanceDefense() int {
	return 4 + luxannaW.SupportShield.Shield.EnhanceDefense()
}

func NewLuxannaW(shield Shield) Shield {
	return &LuxannaW{SupportShield{shield}}
}
