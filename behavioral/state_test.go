// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package behavioral

import "testing"

func Test_state(t *testing.T){
	start := NewDayContent()
	dayAndNext := func(w *DayContext){
		w.printCurrent()
		w.stateTrans()
	}
	for i := 0; i < 2; i++{
		dayAndNext(start)
	}
}
