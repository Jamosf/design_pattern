// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package creational

import "fmt"

type inode interface {
    print(string)
    clone() inode
}

type file struct {
    name string
}

func (f *file) print(indentation string) {
    fmt.Println(indentation + f.name)
}

func (f *file) clone() inode {
    return &file{name: f.name + "_clone"}
}

type folder struct {
    children []inode
    name      string
}

func (f *folder) print(indentation string) {
    fmt.Println(indentation + f.name)
    for _, i := range f.children {
        i.print(indentation + indentation)
    }
}

func (f *folder) clone() inode {
    cloneFolder := &folder{name: f.name + "_clone"}
    var tempChildren []inode
    for _, i := range f.children {
        copy := i.clone()
        tempChildren = append(tempChildren, copy)
    }
    cloneFolder.children = tempChildren
    return cloneFolder
}