/*
	创建型模式：建造者模式
	【建造者模式和工厂模式的关注点不同：建造者模式注重零部件的组装过程，而工厂方法模式更注重零部件的创建过程，两者可以结合使用。】
	1.模式动机：
		在软件开发中，存在大量的复杂对象，它们拥有一系列成员属性，这些成员属性中有些是引用类型的成员对象。
		而且在这些复杂对象中，还可能存在一些限制条件，如某些属性没有赋值则复杂对象不能作为一个完整的产品使用；
		有些属性的赋值必须按照某个顺序，一个属性没有赋值之前，另一个属性可能无法赋值等。
		复杂对象相当于一辆有待建造的汽车，而对象的属性相当于汽车的部件，建造产品的过程就相当于组合部件的过程。
		由于组合部件的过程很复杂，因此，这些部件的组合过程往往被"外部化"到一个称作建造者的对象里，
		建造者返还给的是一个已经建造完毕的完整产品对象，而调用者无须关心该对象所包含的属性以及它们的组装方式。
	2.模式定义：
		将一个复杂对象的构造与它的表示分离，使同样的构建过程可以创建不同的表示，这样的设计模式被称为建造者模式。
		它是将一个复杂的对象分解为多个简单的对象，然后一步一步构建而成。
		它将变与不变相分离，即产品的组成部分是不变的，但每一部分是可以灵活选择的。
	3.优点：
		各个具体的建造者相互独立，有利于系统的扩展。
		调用者不必知道产品内部组成的细节，便于控制细节风险。
	4.缺点：
		产品的组成部分必须相同，这限制了其使用范围。
		如果产品的内部变化复杂，该模式会增加很多的建造者类。
	5.适用场景：
		创建的对象较复杂，由多个部件构成，各部件面临着复杂的变化，但构件间的建造顺序是稳定的。
		创建复杂对象的算法独立于该对象的组成部分以及它们的装配方式，即产品的构建过程和最终的表示是独立的。
*/

package main

import (
	"fmt"
)

func main() {
	builder := CreateConcreteDecorator1()	// 建造者
	m := CreateProjectManager(builder)	// 指挥者
	m.Decorate()
}

// 产品：客厅
type Parlour struct {
	wall string // 墙
	tv   string // 电视
	sofa string // 沙发
}

func (p *Parlour) SetWall(wall string) {
	p.wall = wall
}

func (p *Parlour) SetTV(tv string) {
	p.tv = tv
}

func (p *Parlour) SetSofa(sofa string) {
	p.sofa = sofa
}

func (p *Parlour) Show() {
	fmt.Println("wall: ", p.wall)
	fmt.Println("tv: ", p.tv)
	fmt.Println("sofa: ", p.sofa)
}

// 抽象建造者：装修工人
type Decorator interface {
	BuildWall()
	BuildTV()
	BuildSofa()
	Show()
}

// 具体建造者：装修工人1
type ConcreteDecorator1 struct {
	parlour *Parlour
}

func (c *ConcreteDecorator1) BuildWall() {
	c.parlour.SetWall("wall1")
}

func (c *ConcreteDecorator1) BuildTV() {
	c.parlour.SetSofa("sofa1")
}

func (c *ConcreteDecorator1) BuildSofa() {
	c.parlour.SetTV("tv1")
}

func (c *ConcreteDecorator1) Show() {
	c.parlour.Show()
}

func CreateConcreteDecorator1() *ConcreteDecorator1{
	return &ConcreteDecorator1{parlour: &Parlour{}}
}

// 具体建造者：装修工人2
type ConcreteDecorator2 struct {
	parlour *Parlour
}

func (c *ConcreteDecorator2) NewParlour() {
	c.parlour = &Parlour{}
}

func (c *ConcreteDecorator2) BuildWall() {
	c.parlour.SetWall("wall2")
}

func (c *ConcreteDecorator2) BuildTV() {
	c.parlour.SetSofa("sofa2")
}

func (c *ConcreteDecorator2) BuildSofa() {
	c.parlour.SetTV("tv2")
}

func (c *ConcreteDecorator2) Show() {
	c.parlour.Show()
}

func CreateConcreteDecorator2() *ConcreteDecorator2{
	return &ConcreteDecorator2{parlour: &Parlour{}}
}

// 指挥者：项目经理
type ProjectManager struct {
	builder Decorator
}

func (p *ProjectManager) ProjectManager(builder Decorator) {
	p.builder = builder
}

//产品构建与组装方法
func (p *ProjectManager) Decorate() {
	p.builder.BuildSofa()
	p.builder.BuildTV()
	p.builder.BuildWall()
	p.builder.Show()
}

func CreateProjectManager(builder Decorator) *ProjectManager {
	return &ProjectManager{builder:builder}
}