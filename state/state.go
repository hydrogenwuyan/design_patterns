/*
	行为型模式：状态模式
	1.模式动机：
		在很多情况下，一个对象的行为取决于一个或多个动态变化的属性，这样的属性叫做状态，
		这样的对象叫做有状态对象，这样的对象状态是从事先定义好的一系列值中取出的。
		当一个这样的对象与外部事件产生互动时，其内部状态就会改变，从而使得系统的行为也随之发生变化。
	2.模式定义：
		对有状态的对象，把复杂的判断逻辑提取到不同的状态对象中，允许状态对象在其内部状态发生改变时改变其行为。
	3.优点：
		状态模式将与特定状态相关的行为局部化到一个状态中，并且将不同状态的行为分割开来，满足"单一职责原则"。
		减少对象间的相互依赖。将不同的状态引入独立的对象中会使得状态转换变得更加明确，且减少对象间的相互依赖。
		有利于程序的扩展。通过定义新的子类很容易地增加新的状态和转换。
	4.缺点：
		状态模式的使用必然会增加系统的类与对象的个数。
		状态模式的结构与实现都较为复杂，如果使用不当会导致程序结构和代码的混乱。
	5.适用场景：
		对象的行为依赖于它的状态（属性）并且可以根据它的状态改变而改变它的相关行为。
		代码中包含大量与对象状态有关的条件语句，这些条件语句的出现，会导致代码的可维护性和灵活性变差，
		不能方便地增加和删除状态，使客户类与类库之间的耦合增强。在这些条件语句中包含了对象的行为，而且这些条件对应于对象的各种状态。
*/

package main

import "fmt"

func main() {
	scoreCtx := CreateScoreContext()
	data := &StateData{scoreCtx, "one", 80}
	middle := CreateMiddleState(data)
	scoreCtx.SetState(middle)
	scoreCtx.AddScore(10)
}

// 环境类
type ScoreContext struct {
	state AbstractState
}

func (s *ScoreContext) SetState(state AbstractState) {
	s.state = state
	fmt.Println("状态改变, 分数：", state.GetStateData().Score)
}

func (s *ScoreContext) GetState() AbstractState {
	return s.state
}

func (s *ScoreContext) AddScore(x int) {
	s.state.AddScore(x)
}

func CreateScoreContext() *ScoreContext {
	return &ScoreContext{}
}

// 抽象状态类
type AbstractState interface {
	GetStateData() *StateData
	AddScore(x int)
}

// 状态数据
type StateData struct {
	ScoreContext *ScoreContext
	Name         string
	Score        int
}

// 具体状态类：不及格
type LowState struct {
	data *StateData
}

func (l *LowState) GetStateData() *StateData {
	return l.data
}

func (l *LowState) AddScore(x int) {
	l.data.Score += x
	l.checkState()
}

func (l *LowState) checkState() {
	if l.data.Score > 85 {
		l.data.ScoreContext.SetState(NewHighState(l))
	} else if l.data.Score > 60 {
		l.data.ScoreContext.SetState(NewMiddleState(l))
	}
}

func NewLowState(state AbstractState) AbstractState {
	return &LowState{
		data: state.GetStateData(),
	}
}

func CreateLowState(data *StateData) AbstractState {
	return &LowState{data}
}

// 具体状态类：中等
type MiddleState struct {
	data *StateData
}

func (l *MiddleState) GetStateData() *StateData {
	return l.data
}

func (l *MiddleState) AddScore(x int) {
	l.data.Score += x
	l.checkState()
}

func (l *MiddleState) checkState() {
	if l.data.Score > 85 {
		l.data.ScoreContext.SetState(NewHighState(l))
	} else if l.data.Score < 60 {
		l.data.ScoreContext.SetState(NewLowState(l))
	}
}

func NewMiddleState(state AbstractState) AbstractState {
	return &MiddleState{
		data: state.GetStateData(),
	}
}

func CreateMiddleState(data *StateData) AbstractState {
	return &MiddleState{data}
}

// 具体状态类：优秀
type HighState struct {
	data *StateData
}

func (l *HighState) GetStateData() *StateData {
	return l.data
}

func (l *HighState) AddScore(x int) {
	l.data.Score += x
	l.checkState()
}

func (l *HighState) checkState() {
	if l.data.Score < 60 {
		l.data.ScoreContext.SetState(NewLowState(l))
	} else if l.data.Score < 85 {
		l.data.ScoreContext.SetState(NewMiddleState(l))
	}
}

func NewHighState(state AbstractState) AbstractState {
	return &MiddleState{
		data: state.GetStateData(),
	}
}

func CreateHighState(data *StateData) AbstractState {
	return &HighState{data}
}
