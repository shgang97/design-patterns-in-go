package main

import "fmt"

/*
@author: shg
@since: 2023/1/7
@desc: //TODO
*/

func main() {
	zed := &Zed{Hero{Name: "儿童劫"}}
	leBlanc := &LeBlanc{Hero{Name: "提款姬"}}
	leeSin := &LeeSin{Hero{Name: "小学僧"}}
	ezreal := &Ezreal{
		Hero: Hero{Name: "ez"},
	}

	fmt.Printf("%s进入红buff草丛\n", zed.Name)
	fmt.Printf("%s进入红buff草丛\n", leBlanc.Name)
	fmt.Printf("%s进入红buff草丛\n", leeSin.Name)
	ezreal.RegisterObserver(zed)
	ezreal.RegisterObserver(leBlanc)
	ezreal.RegisterObserver(leeSin)

	fmt.Printf("%s前来拿红buff，在草丛发现插眼，发现敌人在，立刻使用E技能逃跑\n", ezreal.Name)
	ezreal.NotifyObservers()
}

// 抽象层

type Hero struct {
	Name string
}

// Assassin 抽象的观察者——带位移的刺客英雄
type Assassin interface {
	Attack() // 当观察到ADC交了位移技能后发起的攻击
}

// ADCarry 抽象主题
type ADCarry interface {
	RegisterObserver(assassin Assassin)
	RemoveObserver(assassin Assassin)
	NotifyObservers()
}

// 实现层
// 具体的观察者

type Zed struct {
	Hero
}

func (hero *Zed) Attack() {
	fmt.Printf("敌方ADC在使用位移技能的一瞬间，%s开启大招，大招落地后WEQ接平A，2段R潇洒转身离开\n", hero.Name)
}

type LeBlanc struct {
	Hero
}

func (hero *LeBlanc) Attack() {
	fmt.Printf("敌方ADC在使用位移技能的一瞬间，%sW跟上QRE连招接平A后2段W潇洒转身离开\n", hero.Name)
}

type LeeSin struct {
	Hero
}

func (hero *LeeSin) Attack() {
	fmt.Printf("敌方ADC在使用位移技能的一瞬间，%sQ技能预判敌人落点命中2段Q跟上AAEAA最后来个神龙摆尾\n", hero.Name)
}

type Ezreal struct {
	Hero
	ObserverList []Assassin // 维护注册主题段观察者集合
}

func (hero *Ezreal) RegisterObserver(assassin Assassin) {
	hero.ObserverList = append(hero.ObserverList, assassin)
}

func (hero *Ezreal) RemoveObserver(assassin Assassin) {
	for i, observer := range hero.ObserverList {
		if assassin == observer {
			hero.ObserverList = append(hero.ObserverList[:i], hero.ObserverList[i+1:]...)
		}
	}
}

func (hero *Ezreal) NotifyObservers() {
	for _, observer := range hero.ObserverList {
		observer.Attack()
	}
}
