// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package creational

// 工厂方法模式：变种较多

// Operator 是被封装的实际类接口
type Operator interface {
	SetA(int)
	SetB(int)
	Result() int
}

// PlusOperator Operator 的实际加法实现
type PlusOperator struct {
	a, b int
}

// Result 获取结果
func (o PlusOperator) Result() int {
	return o.a + o.b
}

// SetA 设置 A
func (o *PlusOperator) SetA(a int) {
	o.a = a
}

// SetB 设置 B
func (o *PlusOperator) SetB(b int) {
	o.b = b
}

// MinusOperator Operator 的实际减法实现
type MinusOperator struct {
	a, b int
}

// Result 获取结果
func (o MinusOperator) Result() int {
	return o.a - o.b
}

// SetA 设置 A
func (o *MinusOperator) SetA(a int) {
	o.a = a
}

// SetB 设置 B
func (o *MinusOperator) SetB(b int) {
	o.b = b
}

// OperatorFactory 是工厂接口
type OperatorFactory interface {
	Create() Operator
}

// PlusOperatorFactory 是 PlusOperator 的工厂类
type PlusOperatorFactory struct{}

func (PlusOperatorFactory) Create() Operator {
	return &PlusOperator{}
}

// MinusOperatorFactory 是 MinusOperator 的工厂类
type MinusOperatorFactory struct{}

func (MinusOperatorFactory) Create() Operator {
	return &MinusOperator{}
}