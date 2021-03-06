package main

import "fmt"

type file struct {
	name string
	size int
}

func newFile(name string, size int) *file {
	return &file{
		name: name,
		size: size,
	}
}

func (f *file) String() string {
	return fmt.Sprintf("%s (%d)", f.getName(), f.getSize())
}

func (f *file) getName() string {
	return f.name
}

func (f *file) getSize() int {
	return f.size
}

func (f *file) add(e entry) {
	panic("file can NOT add")
}

func (f *file) accept(v visitor) {
	v.visit(f)
}
