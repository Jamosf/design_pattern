// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package structural

import "fmt"

// 组合模式用于树形结构，通常具有递归的性质

// 接口
type component interface {
	search(string)
}

// 简单元素
type file struct {
	name string
}

func (f *file) search(keyword string) {
	fmt.Printf("Searching for keyword %s in file %s\n", keyword, f.name)
}

func (f *file) getName() string {
	return f.name
}

// 容器、复杂元素
type folder struct {
	components []component
	name       string
}

func (f *folder) search(keyword string) {
	fmt.Printf("Serching recursively for keyword %s in folder %s\n", keyword, f.name)
	for _, composite := range f.components {
		composite.search(keyword)
	}
}

func (f *folder) add(c component) {
	f.components = append(f.components, c)
}