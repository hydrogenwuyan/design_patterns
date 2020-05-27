/*
	结构型模式：外观模式
	1.模式动机：
		在现实生活中，常常存在办事较复杂的例子，如办房产证或注册一家公司，有时要同多个部门联系，这时要是有一个综合部门能解决一切手续问题就好了。
		软件设计也是这样，当一个系统的功能越来越强，子系统会越来越多（"单一职责原则"），客户对系统的访问也变得越来越复杂。
		这时如果系统内部发生改变，调用者也要跟着改变，这违背了"开闭原则"，也违背了"迪米特法则"，
		所以有必要为多个子系统提供一个统一的接口，从而降低系统的耦合度，这就是外观模式的目标。
		【迪米特法则：一个类对于其他类知道的越少越好。如果两个软件实体无须直接通信，那么就不应当发生直接的相互调用，可以通过第三方转发该调用。】
	2.模式定义：
		是一种通过为多个复杂的子系统提供一个一致的接口，而使这些子系统更加容易被访问的模式。
		该模式对外有一个统一接口，外部应用程序不用关心内部子系统的具体的细节，这样会大大降低应用程序的复杂度，提高了程序的可维护性。
	3.优点：
		降低了子系统与调用者之间的耦合度，使得子系统的变化不会影响调用它的客户类。
		对客户屏蔽了子系统组件，减少了客户处理的对象数目，并使得子系统使用起来更加容易。
		降低了大型软件系统中的编译依赖性，简化了系统在不同平台之间的移植过程，因为编译一个子系统不会影响其他的子系统，也不会影响外观对象。
	4.缺点：
		不能很好地限制客户使用子系统类。
		增加新的子系统可能需要修改外观类或客户端的源代码，违背了"开闭原则"。
	5.适用场景：
		对分层结构系统构建时，使用外观模式定义子系统中每层的入口点可以简化子系统之间的依赖关系。
		当一个复杂系统的子系统很多时，外观模式可以为系统设计一个简单的接口供外界访问。
		当调用者与多个子系统之间存在很大的联系时，引入外观模式可将它们分离，从而提高子系统的独立性和可移植性。
*/

package main

import "fmt"

func main() {
	menu := CreateMenu(&Beef{}, &Mutton{}, &Pork{}, &MaoDu{})
	menu.Show()
}

// 外观角色：菜单
type Menu struct {
	beef   *Beef
	mutton *Mutton
	pork   *Pork
	maoDu  *MaoDu
}

func (b *Menu) Show() {
	b.beef.Show()
	b.mutton.Show()
	b.pork.Show()
	b.maoDu.Show()
}

func CreateMenu(beef *Beef, mutton *Mutton, pork *Pork, maoDu *MaoDu) *Menu {
	return &Menu{
		beef:   beef,
		mutton: mutton,
		pork:   pork,
		maoDu:  maoDu,
	}
}

// 子系统角色：牛肉
type Beef struct {
}

func (b *Beef) Show() {
	fmt.Println("牛肉")
}

// 子系统角色：羊肉
type Mutton struct {
}

func (b *Mutton) Show() {
	fmt.Println("羊肉")
}

// 子系统角色：猪肉
type Pork struct {
}

func (b *Pork) Show() {
	fmt.Println("猪肉")
}

// 子系统角色：毛肚
type MaoDu struct {
}

func (b *MaoDu) Show() {
	fmt.Println("毛肚")
}
