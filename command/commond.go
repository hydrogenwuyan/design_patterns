/*
	行为型模式：命令模式
	1.模式动机：
		在软件设计中，我们经常需要向某些对象发送请求，但是并不知道请求的接收者是谁，也不知道被请求的操作是哪个，
		我们只需在程序运行时指定具体的请求接收者即可，此时，可以使用命令模式来进行设计，
		使得请求发送者与请求接收者消除彼此之间的耦合，让对象之间的调用关系更加灵活。
		命令模式可以对发送者和接收者完全解耦，发送者与接收者之间没有直接引用关系，
		发送请求的对象只需要知道如何发送请求，而不必知道如何完成请求。这就是命令模式的模式动机。
	2.模式定义：
		将一个请求封装为一个对象，使发出请求的责任和执行请求的责任分割开。
		这样两者之间通过命令对象进行沟通，这样方便将命令对象进行储存、传递、调用、增加与管理。
	3.优点：
		降低系统的耦合度。命令模式能将调用操作的对象与实现该操作的对象解耦。
		增加或删除命令非常方便。采用命令模式增加与删除命令不会影响其他类，它满足"开闭原则"，对扩展比较灵活。
		可以比较容易地设计一个命令队列和宏命令（组合命令）。
		方便实现Undo(撤销)和Redo(恢复)操作。命令模式可以与后面介绍的备忘录模式结合，实现命令的撤销与恢复。
	4.缺点：
		可能产生大量具体命令类。因为对每一个具体操作都需要设计一个具体命令类，这将增加系统的复杂性。
	5.适用场景：
		系统需要将请求调用者和请求接收者解耦，使得调用者和接收者不直接交互。
		系统需要在不同的时间指定请求、将请求排队和执行请求。
		系统需要支持命令的Undo操作和Redo操作。
		系统需要将一组操作组合在一起，即支持宏命令。
*/

package main

import "fmt"

func main() {
	w := CreateWaiter()
	f := CreateFoodA(&FoodAChef{})
	w.SetFoods(f)
	w.Cmd()
}

/*
	客户去餐馆点餐，客户可向服务员选择以上早餐中的若干种，
	服务员将客户的请求交给相关的厨师去做。这里的点餐相当于“命令”，服务员相当于“调用者”，厨师相当于“接收者”，所以用命令模式实现比较合适。
*/

// 调用者：服务员
type Waiter struct {
	foods []Food
}

// 点餐
func (w *Waiter) SetFoods(food Food) {
	w.foods = append(w.foods, food)
}

// 调用命令
func (w *Waiter) Cmd() {
	for _, v := range w.foods {
		v.Cooking()
	}
}

func CreateWaiter() *Waiter {
	return &Waiter{[]Food{}}
}

// 抽象命令：食物
type Food interface {
	Cooking()
}

// 具体命令：食物A
type FoodA struct {
	foodAChef *FoodAChef
}

func (f *FoodA) Cooking() {
	f.foodAChef.Cooking()
}

func CreateFoodA(chef *FoodAChef) *FoodA {
	return &FoodA{chef}
}

// 接受者：食物A厨师
type FoodAChef struct {
}

func (f *FoodAChef) Cooking() {
	fmt.Println("Food A!")
}
