/*
	创建型模式：工厂方法模式
	1.模式动机：
		工厂抽象化的结果使这种结构可以在不修改具体工厂类的情况下引进新的产品，
		如果出现新的运算类型，只需要为这种新类型的运算创建一个具体的工厂类就可以获得该新运算类的实例，
		这一特点更加符合"开闭原则"。
		【开闭原则：当应用的需求改变时，在不修改软件实体的源代码或者二进制代码的前提下，可以扩展模块的功能，使其满足新的需求。】
	2.模式定义：
		定义一个创建产品对象的工厂接口，将产品对象的实际创建工作推迟到具体子工厂类当中。
		这满足创建型模式中所要求的"创建与使用相分离"的特点。
	3.优点：
		用户只需要知道具体工厂的名称就可得到所要的产品，无须知道产品的具体创建过程；
		在系统增加新的产品时只需要添加具体产品类和对应的具体工厂类，无须对原工厂进行任何修改，满足"开闭原则"；
	4.缺点：
		每增加一个产品就要增加一个具体产品类和一个对应的具体工厂类，这增加了系统的复杂度。
	5.适用场景：
		客户只知道创建产品的工厂名，而不知道具体的产品名。
		创建对象的任务由多个具体子工厂中的某一个完成，而抽象工厂只提供创建产品的接口。
*/

package main

import "fmt"

func main() {
	manager := CreateCalculateAddManager()
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

// 运算类工厂接口
type CalculateFactory interface {
	CreateCalculateManager() CalculateManager
}

// 运算类工厂类型枚举
type CalculateFactoryFunc func() CalculateManager

func (f CalculateFactoryFunc) CreateCalculateManager() CalculateManager {
	return f()
}

// 运算类具体实现
// 加法
type CalculateAddManager struct {
}

func (m *CalculateAddManager) Calculate(a, b int) int {
	return a + b
}

// 加法类工厂具体实现
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
