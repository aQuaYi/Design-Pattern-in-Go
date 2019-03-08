package main

import (
	"fmt"
	"strings"
)

type product interface {
	use(string)
	createClone() product
}

// messageBox implements product interface
type messageBox struct {
	decochar byte
}

func newMessageBox(char byte) *messageBox {
	return &messageBox{
		decochar: char,
	}
}

func (mb *messageBox) use(s string) {
	line := strings.Repeat(string(mb.decochar), len(s))
	fmt.Println(line)
	fmt.Printf("%c %s %c\n", mb.decochar, s, mb.decochar)
	fmt.Println(line)
}

func (mb *messageBox) createClone() *messageBox {
	clone := *mb
	return &clone
}

// underlinePen implements product interface
type underlinePen struct {
	ulchar byte
}

func newUnderlinePen(char byte) *underlinePen {
	return &underlinePen{
		ulchar: char,
	}
}

func (mb *underlinePen) use(s string) {
	fmt.Printf("\"%s\"\n", s)
	line := strings.Repeat(string(mb.ulchar), len(s))
	fmt.Println(" ", line)
}

func (mb *underlinePen) createClone() *underlinePen {
	clone := *mb
	return &clone
}