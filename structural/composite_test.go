// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package structural

import "testing"

func Test_composite(t *testing.T){
	file1 := &file{name: "File1"}
	file2 := &file{name: "File2"}
	file3 := &file{name: "File3"}

	folder1 := &folder{
		name: "Folder1",
	}

	folder1.add(file1)

	folder2 := &folder{
		name: "Folder2",
	}
	folder2.add(file2)
	folder2.add(file3)
	folder2.add(folder1)

	folder2.search("rose")
}
