// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package creational

import "sync"

var once sync.Once

type single struct {}

var singleInstance *single

func getInstance() *single{
	once.Do(func() {
		singleInstance =  &single{}
	})
	return singleInstance
}
