/*
	结构型模式：享元模式
	1.模式动机：
		在面向对象程序设计过程中，有时会面临要创建大量相同或相似对象实例的问题。
		创建那么多的对象将会耗费很多的系统资源，它是系统性能提高的一个瓶颈。
		例如，围棋和五子棋中的黑白棋子，图像中的坐标点或颜色，局域网中的路由器、交换机和集线器，教室里的桌子和凳子等。
		这些对象有很多相似的地方，如果能把它们相同的部分提取出来共享，则能节省大量的系统资源，这就是享元模式的产生背景。
	2.模式定义：
		运用共享技术来有効地支持大量细粒度对象的复用。
		它通过共享已经存在的对象来大幅度减少需要创建的对象数量、避免大量相似类的开销，从而提高系统资源的利用率。
	3.优点：
		相同对象只要保存一份，这降低了系统中对象的数量，从而降低了系统中细粒度对象给内存带来的压力。
	4.缺点：
		为了使对象可以共享，需要将一些不能共享的状态外部化，这将增加程序的复杂性。
		读取享元模式的外部状态会使得运行时间稍微变长。
	5.适用场景：
		一个系统有大量相同或者相似的对象，由于这类对象的大量使用，造成内存的大量耗费。
		对象的大部分状态都可以外部化，可以将这些外部状态传入对象中。
		使用享元模式需要维护一个存储享元对象的享元池，而这需要耗费资源，因此，应当在多次重复使用享元对象时才值得使用享元模式。
*/

package main

import "fmt"

func main() {
	f := CreateGoFactory()
	f.GetChessPieces(ChessPiecesTypeBlack).DownPieces(Point{1, 1})
	f.GetChessPieces(ChessPiecesTypeWhite).DownPieces(Point{1, 2})
}

type ChessPiecesType int

const (
	ChessPiecesTypeWhite ChessPiecesType = iota
	ChessPiecesTypeBlack
)

// 非享元角色：点位置
type Point struct {
	X int
	Y int
}

// 抽象享元角色：棋子
type ChessPieces interface {
	DownPieces(p Point) // 下子
}

// 具体享元角色：白子
type WhitePieces struct {
}

// 下子
func (w *WhitePieces) DownPieces(p Point) {
	fmt.Println(p.X, " ", p.Y)
}

// 具体享元角色：黑子
type BlackPieces struct {
}

// 下子
func (b *BlackPieces) DownPieces(p Point) {
	fmt.Println(p.X, " ", p.Y)
}

// 享元工厂角色：围棋
type GoFactory struct {
	ChessPiecesMap map[ChessPiecesType]ChessPieces
}

func (g *GoFactory) GetChessPieces(typ ChessPiecesType) ChessPieces {
	return g.ChessPiecesMap[typ]
}

func CreateGoFactory() *GoFactory {
	g := &GoFactory{ChessPiecesMap: make(map[ChessPiecesType]ChessPieces)}
	g.ChessPiecesMap[ChessPiecesTypeWhite] = &WhitePieces{}
	g.ChessPiecesMap[ChessPiecesTypeBlack] = &BlackPieces{}
	return g
}
