package main

import (
	"fmt"
)

type folder struct {
	childrens []inode
	name      string
}

func (f *folder) print(identation string) {
	fmt.Println(identation + f.name + "_clone")
	for _, i := range f.childrens {
		i.print(identation + identation)
	}
}

func (f *folder) clone() inode {
	cloneFolder := &folder{name: f.name}
	var tempChildrens []inode
	for _, i := range f.childrens {
		copy := i.clone()
		tempChildrens = append(tempChildrens, copy)
	}
	cloneFolder.childrens = tempChildrens
	return cloneFolder
}
