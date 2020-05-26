/*
	结构型模式：桥接模式
	1.模式动机：
		设想如果要绘制矩形、圆形、椭圆、正方形，我们至少需要4个形状类，但是如果绘制的图形需要具有不同的颜色，如红色、绿色、蓝色等，此时至少有如下两种设计方案：
		第一种设计方案是为每一种形状都提供一套各种颜色的版本。
		第二种设计方案是根据实际需要对形状和颜色进行组合
		对于有两个变化维度（即两个变化的原因）的系统，采用方案二来进行设计系统中类的个数更少，且系统扩展更为方便。
		设计方案二即是桥接模式的应用。桥接模式将继承关系转换为关联关系，从而降低了类与类之间的耦合，减少了代码编写量。
	2.模式定义：
		将抽象与实现分离，使它们可以独立变化。
	3.优点：
		由于抽象与实现分离，所以扩展能力强；
		其实现细节对客户透明。
	4.缺点：
		由于"聚合关系"建立在抽象层，要求开发者针对抽象化进行设计与编程，这增加了系统的理解与设计难度。
		【聚合关系：体现的是整体与部分、拥有的关系，此时整体与部分之间是可分离的，他们可以具有各自的生命周期，
		部分可以属于多个整体对象，也可以为多个整体对象共享；比如计算机与CPU、公司与员工的关系等】
	5.适用场景：
		当一个类存在两个独立变化的维度，且这两个维度都需要进行扩展时。
		当一个系统不希望使用继承或因为多层次继承导致系统类的个数急剧增加时。
		当一个系统需要在构件的抽象化角色和具体化角色之间增加更多的灵活性时。
*/

package main

import "fmt"

func main() {
	color := &Yellow{}
	bag := CreateBag(color)
	handBag := CreateHandBag(bag)
	fmt.Println(handBag.GetName())
}


// 实现化角色：颜色
type Color interface{
	GetColor() string
}

//具体实现化角色：黄色
type Yellow struct {
}

func (y *Yellow) GetColor() string {
	return "Yellow"
}

// 具体实现化角色：红色
type Red struct {
}

func (r *Red) GetColor() string {
	return "Red"
}

// 抽象化角色：包
type Bag struct {
	Color Color
}

func CreateBag(c Color) *Bag {
	return &Bag{c}
}

// 扩展抽象化角色：挎包
type HandBag struct {
	bag *Bag
}

func (h *HandBag) GetName() string {
	return h.bag.Color.GetColor()+" HandBag"
}

func CreateHandBag(bag *Bag) *HandBag {
	return &HandBag{bag}
}

// 扩展抽象化角色：挎包
type Wallet struct {
	bag *Bag
}

func (h *Wallet) GetName() string {
	return h.bag.Color.GetColor()+" Wallet"
}

func CreateWallet(bag *Bag) *Wallet {
	return &Wallet{bag}
}