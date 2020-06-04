/*
	行为型模式：中介者模式
	【迪米特法则：如果两个软件实体无须直接通信，那么就不应当发生直接的相互调用，可以通过第三方转发该调用。
	其目的是降低类之间的耦合度，提高模块的相对独立性。】
	1.模式动机：
		在用户与用户直接聊天的设计方案中，用户对象之间存在很强的关联性，将导致系统出现如下问题：
		系统结构复杂：对象之间存在大量的相互关联和调用，若有一个对象发生变化，则需要跟踪和该对象关联的其他所有对象，并进行适当处理。
		对象可重用性差：由于一个对象和其他对象具有很强的关联，若没有其他对象的支持，一个对象很难被另一个系统或模块重用，
		这些对象表现出来更像一个不可分割的整体，职责较为混乱。
		系统扩展性低：增加一个新的对象需要在原有相关对象上增加引用，增加新的引用关系也需要调整原有对象，系统耦合度很高，
		对象操作很不灵活，扩展性差。
		在面向对象的软件设计与开发过程中，根据"单一职责原则"，我们应该尽量将对象细化，使其只负责或呈现单一的职责。
		对于一个模块，可能由很多对象构成，而且这些对象之间可能存在相互的引用，为了减少对象两两之间复杂的引用关系，
		使之成为一个松耦合的系统，我们需要使用中介者模式，这就是中介者模式的模式动机。
	2.模式定义：
		定义一个中介对象来封装一系列对象之间的交互，使原有对象之间的耦合松散，且可以独立地改变它们之间的交互。
		中介者模式又叫调停模式，它是"迪米特法则"的典型应用。
	3.优点：
		降低了对象之间的耦合性，使得对象易于独立地被复用。
		将对象间的一对多关联转变为一对一的关联，提高系统的灵活性，使得系统易于维护和扩展。
	4.缺点：
		当同事类太多时，中介者的职责将很大，它会变得复杂而庞大，以至于系统难以维护。
	5.适用场景：
		系统中对象之间存在复杂的引用关系，产生的相互依赖关系结构混乱且难以理解。
		一个对象由于引用了其他很多对象并且直接和这些对象通信，导致难以复用该对象。
		想通过一个中间类来封装多个类中的行为，而又不想生成太多的子类。可以通过引入中介者类来实现，在中介者中定义对象。
		交互的公共行为，如果需要改变行为则可以增加新的中介者类。
*/

package main

import (
	"fmt"
)

func main() {
	es := CreateEstateMedium()
	sellerA := CreateSeller("sellerA", es)
	buyerA := CreateBuyer("buyerA", es)
	es.Register(sellerA, buyerA)
	sellerA.Send("卖房")
	buyerA.Send("我买")
}

// 抽象中介者：中介公司
type Medium interface {
	Register(c ...Customer)
	Relay(from string, ad string) // 转发
}

// 具体中介者：房产中介公司
type EstateMedium struct {
	customers []Customer
}

func (e *EstateMedium) Register(c ...Customer) {
	e.customers = append(e.customers, c...)
}

func (e *EstateMedium) Relay(from string, ad string) {
	for _, v := range e.customers {
		if v.GetName() != from {
			v.Receive(from, ad)
		}
	}
}

func CreateEstateMedium() *EstateMedium {
	return &EstateMedium{customers: []Customer{}}
}

// 抽象同事类：客户
type Customer interface {
	Receive(from string, ad string)
	GetName() string
}

// 具体同时类：卖方
type Seller struct {
	name   string
	medium Medium
}

func (s *Seller) SetMedium(m Medium) {
	s.medium = m
}

func (s *Seller) Receive(from string, ad string) {
	fmt.Printf("%s: %s\n", from, ad)
}

func (s *Seller) GetName() string {
	return s.name
}

func (s *Seller) Send(ad string) {
	s.medium.Relay(s.name, ad)
}

func CreateSeller(name string, m Medium) *Seller {
	return &Seller{name: name, medium: m}
}

// 具体同时类：买方
type Buyer struct {
	name   string
	medium Medium
}

func (s *Buyer) SetMedium(m Medium) {
	s.medium = m
}

func (s *Buyer) Receive(from string, ad string) {
	fmt.Printf("%s: %s\n", from, ad)
}

func (s *Buyer) GetName() string {
	return s.name
}

func (s *Buyer) Send(ad string) {
	s.medium.Relay(s.name, ad)
}

func CreateBuyer(name string, m Medium) *Buyer {
	return &Buyer{name: name, medium: m}
}
