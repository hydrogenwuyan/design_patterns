/*
	行为型模式：观察者模式（发布订阅模式）
	1.模式动机：
		建立一种对象与对象之间的依赖关系，一个对象发生改变时将自动通知其他对象，其他对象将相应做出反应。
		在此，发生改变的对象称为观察目标，而被通知的对象称为观察者，一个观察目标可以对应多个观察者，
		而且这些观察者之间没有相互联系，可以根据需要增加和删除观察者，使得系统更易于扩展，这就是观察者模式的模式动机。
	2.模式定义：
		指多个对象间存在一对多的依赖关系，当一个对象的状态发生改变时，所有依赖于它的对象都得到通知并被自动更新。
		这种模式有时又称作发布-订阅模式，它是对象行为型模式。
	3.优点：
		降低了目标与观察者之间的耦合关系，两者之间是抽象耦合关系。
		目标与观察者之间建立了一套触发机制。
	4.缺点：
		目标与观察者之间的依赖关系并没有完全解除，而且有可能出现循环引用。
		当观察者对象很多时，通知的发布会花费很多时间，影响程序的效率。
	5.适用场景：
		一个抽象模型有两个方面，其中一个方面依赖于另一个方面。将这些方面封装在独立的对象中使它们可以各自独立地改变和复用。
		一个对象的改变将导致其他一个或多个对象也发生改变，而不知道具体有多少对象将发生改变，可以降低对象之间的耦合度。
		一个对象必须通知其他对象，而并不知道这些对象是谁。
		需要在系统中创建一个触发链，A对象的行为将影响B对象，B对象的行为将影响C对象等等，可以使用观察者模式创建一种链式触发机制。
*/

package main

import "fmt"

func main() {
	rmb := CreateRMBrate()
	rmb.Add(&ImportCompany{})
	rmb.Change(10)
}

// 抽象目标：汇率
type Rate interface {
	Add(c ...Company)
	Remove(c Company)
}

// 具体目标：人民币汇率
type RMBrate struct {
	companys []Company
}

func (r *RMBrate) Add(c ...Company) {
	r.companys = append(r.companys, c...)
}

func (r *RMBrate) Remove(c Company) {
	index := -1
	for i, v := range r.companys {
		if v == c {
			index = i
			break
		}
	}
	if index < 0 {
		return
	}
	companys := r.companys
	r.companys = companys[:index]
	r.companys = append(r.companys, companys[index+1:]...)
}

func (r *RMBrate) Change(num int) {
	for _, v := range r.companys {
		v.Response(num)
	}
}

func CreateRMBrate() *RMBrate {
	return &RMBrate{companys: []Company{}}
}

// 抽象观察者：公司
type Company interface {
	Response(rate int)
}

// 具体观察者：进口公司
type ImportCompany struct {
}

func (i *ImportCompany) Response(rate int) {
	fmt.Printf("汇率提高 %d\n", rate)
}
