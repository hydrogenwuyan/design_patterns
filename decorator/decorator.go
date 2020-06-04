/*
	结构型模式：装饰模式
	1.模式动机：
		一般有两种方式可以实现给一个类或对象增加行为：
			(1)继承机制，使用继承机制是给现有类添加功能的一种有效途径，通过继承一个现有类可以使得子类在拥有自身方法的同时还拥有父类的方法。
			但是这种方法是静态的，用户不能控制增加行为的方式和时机。
			(2)关联机制，即将一个类的对象嵌入另一个对象中，由另一个对象来决定是否调用嵌入对象的行为以便扩展自己的行为，
			我们称这个嵌入的对象为装饰器。
		装饰模式以对客户透明的方式动态地给一个对象附加上更多的责任，换言之，客户端并不会觉得对象在装饰前和装饰后有什么不同。
		装饰模式可以在不需要创造更多子类的情况下，将对象的功能加以扩展。这就是装饰模式的模式动机。
	2.模式定义：
		动态地给一个对象增加一些额外的职责，就增加对象功能来说，装饰模式比生成子类实现更为灵活。
		其别名也可以称为包装器，与适配器模式的别名相同，但它们适用于不同的场合。
	3.优点：
		采用装饰模式扩展对象的功能比采用继承方式更加灵活。
		可以设计出多个不同的具体装饰类，创造出多个不同行为的组合。
		具体构件类与具体装饰类可以独立变化，用户可以根据需要增加新的具体构件类和具体装饰类，
		在使用时再对其进行组合，原有代码无须改变，符合"开闭原则"。
	4.缺点：
		装饰模式增加了许多子类，如果过度使用会使程序变得很复杂。
	5.适用场景：
		在不影响其他对象的情况下，以动态、透明的方式给单个对象添加职责。
		需要动态地给一个对象增加功能，这些功能也可以动态地被撤销。
		当不能采用继承的方式对系统进行扩充或者采用继承不利于系统扩展和维护时。
		不能采用继承的情况主要有：系统中存在大量独立的扩展，为支持每一种组合将产生大量的子类，使得子类数目呈爆炸性增长。
*/

package main

import "fmt"

func main() {
	m := &Original{}
	changer := CreateChanger(m)
	girl := CreateGirl(changer)
	girl.Display()
}

/*
	在《恶魔战士》中，游戏角色“莫莉卡·安斯兰”的原身是一个可爱少女，但当她变身时，
	会变成头顶及背部延伸出蝙蝠状飞翼的女妖，当然她还可以变为穿着漂亮外衣的少女。
*/

//抽象构件角色：莫莉卡
type Morrigan interface {
	Display()
}

// 具体构件角色：原身
type Original struct {
	decorator string
}

func (o *Original) SetDecorator(d string) {
	o.decorator = d
}

func (o *Original) Display() {
	fmt.Println(o.decorator)
}

// 抽象装饰角色：变形
type Changer struct {
	morrigan Morrigan
}

func (c *Changer) Changer(m Morrigan) {
	c.morrigan = m
}

func (c *Changer) DisPlay() {
	c.morrigan.Display()
}

func CreateChanger(m Morrigan) *Changer {
	return &Changer{m}
}

// 具体装饰角色：女妖
type Succubus struct {
	changer *Changer
}

func (s *Succubus) SetChanger() {
	o, ok := s.changer.morrigan.(*Original)
	if !ok {
		fmt.Println("s.changer.morrigan type is not *Original")
		return
	}
	o.SetDecorator("女妖")
}

func (s *Succubus) Display() {
	s.SetChanger()
	s.changer.DisPlay()
}

func CreateSuccubus(changer *Changer) *Succubus {
	return &Succubus{changer}
}

// 具体装饰角色：少女
type Girl struct {
	changer *Changer
}

func (s *Girl) SetChanger() {
	o, ok := s.changer.morrigan.(*Original)
	if !ok {
		fmt.Println("s.changer.morrigan type is not *Original")
		return
	}
	o.SetDecorator("少女")
}

func (s *Girl) Display() {
	s.SetChanger()
	s.changer.DisPlay()
}

func CreateGirl(changer *Changer) *Girl {
	return &Girl{changer}
}
