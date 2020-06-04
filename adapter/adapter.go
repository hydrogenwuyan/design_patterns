/*
	结构性模式：适配器模式
	【合成复用原则：软件复用时，要尽量先使用组合或者聚合等关联关系来实现，其次才考虑使用继承关系来实现】
	1.模式动机：
		适配器提供客户类需要的接口，适配器的实现就是把客户类的请求转化为对适配者的相应接口的调用。
		也就是说：当客户类调用适配器的方法时，在适配器类的内部将调用适配者类的方法，而这个过程对客户类是透明的，
		客户类并不直接访问适配者类。因此，适配器可以使由于接口不兼容而不能交互的类可以一起工作。这就是适配器模式的模式动机。
	2.模式定义：
		将一个类的接口转换成客户希望的另外一个接口，使得原本由于接口不兼容而不能一起工作的那些类能一起工作。
	3.优点：
		客户端通过适配器可以透明地调用目标接口。
		复用了现存的类，程序员不需要修改原有代码而重用现有的适配者类。
		将目标类和适配者类解耦，解决了目标类和适配者类接口不一致的问题。
	4.缺点：
		对类适配器来说，更换适配器的实现过程比较复杂。
	5.适用场景：
		系统需要使用现有的类，而这些类的接口不符合系统的需要。
		想要建立一个可以重复使用的类，用于与一些彼此之间没有太大关联的一些类，包括一些可能在将来引进的类一起工作。
*/

package main

import "fmt"

func main() {
	motor := &OpticalMotor{}
	adapter := CreateOpticalAdapter(motor)
	adapter.Drive()
}

/*
	新能源汽车的发动机有电能发动机（Electric Motor）和光能发动机（Optical Motor）等，各种发动机的驱动方法不同，
	例如，电能发动机的驱动方法 electricDrive() 是用电能驱动，而光能发动机的驱动方法 opticalDrive() 是用光能驱动，
	它们是适配器模式中被访问的适配者。
	客户端希望用统一的发动机驱动方法 drive() 访问这两种发动机，
	所以必须定义一个统一的目标接口 Motor，然后再定义电能适配器（Electric Adapter）和光能适配器（Optical Adapter）去适配这两种发动机。
*/

// 适配器：发动机
type Motor interface {
	Drive()
}

// 适配者1：电能发动机
type ElectricMotor struct {
}

func (*ElectricMotor) ElectricDrive() {
	fmt.Println("电能发动机驱动汽车")
}

// 适配者2：光能发动机
type OpticalMotor struct {
}

func (*OpticalMotor) OpticalDrive() {
	fmt.Println("光能发动机驱动汽车")
}

// 电能适配器
type ElectricAdapter struct {
	motor *ElectricMotor
}

func (e *ElectricAdapter) Drive() {
	e.motor.ElectricDrive()
}

func CreateElectricAdapter(motor *ElectricMotor) *ElectricAdapter {
	return &ElectricAdapter{motor: motor}
}

// 光能适配器
type OpticalAdapter struct {
	motor *OpticalMotor
}

func (e *OpticalAdapter) Drive() {
	e.motor.OpticalDrive()
}

func CreateOpticalAdapter(motor *OpticalMotor) *OpticalAdapter {
	return &OpticalAdapter{motor: motor}
}
