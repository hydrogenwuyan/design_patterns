/*
	创建型模式：抽象工厂模式
	1.模式动机：
		需要一个工厂可以提供多个产品对象，而不是单一的产品对象，
		并且产品对象位于不同产品等级结构中属于不同类型的具体产品时需要使用抽象工厂模式。
	2.模式定义：
		是一种为访问类提供一个创建一组相关或相互依赖对象的接口，
		且访问类无须指定所要产品的具体类就能得到同族的不同等级的产品的模式结构。
	3.优点：
		可以在类的内部对产品族中相关联的多等级产品共同管理，而不必专门引入多个新的类来进行管理。
		增加新的具体工厂和产品族很方便，无须修改已有系统，符合"开闭原则"。
	4.缺点：
		在添加新的产品对象时，难以扩展抽象工厂来生产新种类的产品。
		"开闭原则"的倾斜性（增加新的工厂和产品族容易，增加新的产品等级结构麻烦）。
	5.适用场景：
		一个系统不应当依赖于产品类实例如何被创建、组合和表达的细节，这对于所有类型的工厂模式都是重要的。
		系统中有多于一个的产品族，而每次只使用其中某一产品族。
		属于同一个产品族的产品将在一起使用，这一约束必须在系统的设计中体现出来。
		系统提供一个产品类的库，所有的产品以同样的接口出现（文中的Show函数）。
*/

package main

import (
	"fmt"
)

func main() {
	farm := CreateGoFarm()
	farm.CreateAnimal().Show()
	farm.CreatePlant().Show()
}

// 动物类接口
type Animal interface {
	Show()
}

// 具体实现：马类
type Horse struct {
}

func (h *Horse) Show() {
	fmt.Println("I am horse")
}

// 构造函数
func CreateHorse() *Horse{
	return &Horse{}
}

// 具体实现：牛类
type Cattle struct {
}

func (c *Cattle) Show() {
	fmt.Println("I am cattle")
}

func CreateCattle() *Cattle{
	return &Cattle{}
}

// 植物类接口
type Plant interface {
	Show()
}

// 具体实现：水果类
type Fruit struct {
}

func (c *Fruit) Show() {
	fmt.Println("I am fruit")
}

func CreateFruit() *Fruit{
	return &Fruit{}
}

// 具体实现：蔬菜类
type Vegetable struct {
}

func (c *Vegetable) Show() {
	fmt.Println("I am vegetable")
}

func CreateVegetable() *Vegetable{
	return &Vegetable{}
}

// 抽象工厂：农场类接口
type Farm interface {
	CreateAnimal() Animal
	CreatePlant() Plant
}

// 具体工厂：Go农场类
type GoFarm struct {
}

// 养马
func (f *GoFarm) CreateAnimal() Animal {
	return CreateHorse()
}

// 种水果
func (f *GoFarm) CreatePlant() Plant {
	return CreateFruit()
}

func CreateGoFarm() *GoFarm {
	return &GoFarm{}
}

// 具体工厂：C农场类
type CFarm struct {
}

// 养牛
func (f *CFarm) CreateAnimal() Animal {
	return CreateCattle()
}

// 种蔬菜
func (f *CFarm) CreatePlant() Plant {
	return CreateVegetable()
}

func CreateCFarm() *CFarm {
	return &CFarm{}
}

