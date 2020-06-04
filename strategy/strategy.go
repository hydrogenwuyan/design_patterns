/*
	行为型模式：策略模式
	1.模式动机：
		在现实生活中常常遇到实现某种目标存在多种策略可供选择的情况，
		例如，出行旅游可以乘坐飞机、乘坐火车、骑自行车或自己开私家车等，超市促销可以釆用打折、送商品、送积分等方法。
		在软件开发中也常常遇到类似的情况，当实现某一个功能存在多种算法或者策略，
		我们可以根据环境或者条件的不同选择不同的算法或者策略来完成该功能，
		如数据排序策略有冒泡排序、选择排序、插入排序、二叉树排序等。
		如果使用多重条件转移语句实现（即硬编码），不但使条件语句变得很复杂，
		而且增加、删除或更换算法要修改原代码，不易维护，违背"开闭原则"。如果采用策略模式就能很好解决该问题。
	2.模式定义：
		该模式定义了一系列算法，并将每个算法封装起来，使它们可以相互替换，且算法的变化不会影响使用算法的客户。
		策略模式属于对象行为模式，它通过对算法进行封装，把使用算法的责任和算法的实现分割开来，并委派给不同的对象对这些算法进行管理。
	3.优点：
		多重条件语句不易维护，而使用策略模式可以避免使用多重条件语句。
		策略模式提供了一系列的可供重用的算法族，恰当使用继承可以把算法族的公共代码转移到父类里面，从而避免重复的代码。
		策略模式可以提供相同行为的不同实现，客户可以根据不同时间或空间要求选择不同的。
		策略模式提供了对"开闭原则"的完美支持，可以在不修改原代码的情况下，灵活增加新算法。
		策略模式把算法的使用放到环境类中，而算法的实现移到具体策略类中，实现了二者的分离。
	4.缺点：
		客户端必须理解所有策略算法的区别，以便适时选择恰当的算法类。
		策略模式造成很多的策略类。
	5.适用场景：
		如果在一个系统里面有许多类，它们之间的区别仅在于它们的行为，那么使用策略模式可以动态地让一个对象在许多行为中选择一种行为。
		一个系统需要动态地在几种算法中选择一种。
		如果一个对象有很多的行为，如果不用恰当的模式，这些行为就只好使用多重的条件选择语句来实现。
		不希望客户端知道复杂的、与算法相关的数据结构，在具体策略类中封装算法和相关的数据结构，提高算法的保密性与安全性。
*/

package main

import "fmt"

func main() {
	ctx := &Context{}
	ctx.SetSort(&BubbleSort{})
	ctx.IntSortOfAsc([]int{2, 7, 6, 3, 8, 1})
}

// 环境类
type Context struct {
	sort Sort
}

func (c *Context) SetSort(s Sort) {
	c.sort = s
}

func (c *Context) IntSortOfAsc(list []int) {
	c.sort.IntSortOfAsc(list)
	fmt.Println(list)
}

// 抽象策略类：排序
type Sort interface {
	IntSortOfAsc(list []int)
}

// 具体策略类：冒泡
type BubbleSort struct {
}

func (s *BubbleSort) IntSortOfAsc(list []int) {
	l := len(list)
	for i := 0; i < l-1; i++ {
		for j := 0; j < l-i-1; j++ {
			if list[j] > list[j+1] {
				list[j], list[j+1] = list[j+1], list[j]
			}
		}
	}
}

// 具体策略类：快排
type QuickSort struct {
}

func (s *QuickSort) IntSortOfAsc(list []int) {
	s.quickSort(list, 0, len(list)-1)
}

func (s *QuickSort) quickSort(list []int, l, r int) {
	if l < r {
		base := list[l]
		i := l
		j := r
		for i < j {
			for list[j] >= base && i < j {
				j--
			}
			if i >= j {
				break
			}

			list[i] = list[j]
			for list[i] <= base && i < j {
				i++
			}
			if i >= j {
				break
			}

			list[j] = list[i]
		}
		list[i] = base

		s.quickSort(list, l, i-1)
		s.quickSort(list, i+1, r)
	}
}
