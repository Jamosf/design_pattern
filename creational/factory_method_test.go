// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package creational

import (
	"fmt"
	"testing"
)

func compute(op OperatorFactory, a, b int){
	f := op.Create()
	f.SetA(a)
	f.SetB(b)
	fmt.Println(f.Result())
}

func Test_factory_method(t *testing.T){
	compute(PlusOperatorFactory{}, 1,2)
	compute(MinusOperatorFactory{}, 1,2)
}