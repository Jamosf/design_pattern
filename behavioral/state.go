// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package behavioral

import "fmt"

// context封装类
type DayContext struct {
	t weekday
}

func NewDayContent() *DayContext {
	return &DayContext{
		t: &Monday{},
	}
}

func (d *DayContext) printCurrent(){
	d.t.today()
}

func (d *DayContext) stateTrans(){
	d.t.next(d)
}

// 状态类接口，所有的状态实现该接口
type weekday interface {
	today()
	next(content *DayContext)
}

type Monday struct {

}

func (m *Monday) today(){
	fmt.Println("Monday")
}

func (m *Monday) next(d *DayContext){
	d.t = &Tuesday{}
}

type Tuesday struct {

}

func (t *Tuesday) today(){
	fmt.Println("Tuesday")
}

func (t *Tuesday) next(d *DayContext){
	d.t = &Monday{}
}
