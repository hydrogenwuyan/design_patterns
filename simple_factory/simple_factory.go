/*
	创建型模式：简单工厂模式
   	抽象加减乘除
	1.模式动机：
		希望在使用这些运算类时，不需要知道这些具体运算类的名字，只需要知道表示该运算类的一个参数，
		并提供一个调用方便的方法，把该参数传入方法即可返回一个相应的运算对象，此时，就可以使用简单工厂模式。
	2.模式定义：
		在简单工厂模式中，可以根据参数的不同返回不同类的实例。
		简单工厂模式专门定义一个类来负责创建其他类的实例，被创建的实例通常都具有共同的父类。
	3.优点：
		调用者无需知道具体的类名，只需知道对应的参数即可。
		实现对象的创建和使用相分离。
	4.缺点：
		不够灵活，新增具体的对象需要修改工厂类内部逻辑即违反"开闭原则"。
	5.适用场景：
		工厂类负责创建的对象比较少：由于创建的对象较少，不会造成工厂方法中的业务逻辑太过复杂。
		客户端只知道传入工厂类的参数，对于如何创建对象不关心。
*/

package main

import "fmt"

type CalculateType int8

const (
	CalculateTypeAdd CalculateType = iota + 1 // 加
	CalculateTypeSub                          // 减
	CalculateTypeMul                          // 乘
	CalculateTypeDiv                          // 除
)

func main() {
	manager := calculateFactoryEntity.CreateCalculateManager(CalculateTypeAdd)
	if manager == nil {
		fmt.Println("manager is nil")
		return
	}
	fmt.Println(manager.Calculate(4, 5))
}

// 抽象运算类接口
type CalculateManager interface {
	Calculate(a, b int) int
}

var (
	calculateFactoryEntity = &CalculateFactory{}
)

// 运算类工厂
type CalculateFactory struct {
}

// 根据类型获取具体的运算类
func (f *CalculateFactory) CreateCalculateManager(typ CalculateType) CalculateManager {
	switch typ {
	case CalculateTypeAdd:
		return CreateCalculateAddManager()
	case CalculateTypeSub:
		return CreateCalculateSubManager()
	case CalculateTypeMul:
		return CreateCalculateMulManager()
	case CalculateTypeDiv:
		return CreateCalculateDivManager()
	}
	return nil
}

// 运算类具体实现
// 加法
type CalculateAddManager struct {
}

func (m *CalculateAddManager) Calculate(a, b int) int {
	return a + b
}

func CreateCalculateAddManager() *CalculateAddManager {
	return &CalculateAddManager{}
}

// 减法
type CalculateSubManager struct {
}

func (m *CalculateSubManager) Calculate(a, b int) int {
	return a - b
}

func CreateCalculateSubManager() *CalculateSubManager {
	return &CalculateSubManager{}
}

// 乘法
type CalculateMulManager struct {
}

func (m *CalculateMulManager) Calculate(a, b int) int {
	return a * b
}

func CreateCalculateMulManager() *CalculateMulManager {
	return &CalculateMulManager{}
}

// 除法
type CalculateDivManager struct {
}

func (m *CalculateDivManager) Calculate(a, b int) int {
	return a / b
}

func CreateCalculateDivManager() *CalculateDivManager {
	return &CalculateDivManager{}
}
