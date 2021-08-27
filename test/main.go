package main

import (
	"fmt"
	"github.com/dengpju/higo-enum/enum"
)

type TestEnum enum.Enum

const (
	None TestEnum = iota
	CPU   // 中央处理器
	GPU   // 图形处理器
)

func (this TestEnum) Code() int64 {
	return enum.New(this).Code
}

func (this TestEnum) Message() string {
	switch this {
	case None:
		return "None"
	case CPU:
		return "CPU"
	case GPU:
		return "GPU"
	}
	return "N/A"
}

func main() {
	fmt.Println(None.Code(), None.Message())
	fmt.Println(CPU.Message(), CPU.Code())
	fmt.Println(enum.Enums())
	//fmt.Println(GPU.Message(), GPU.Code())
	fmt.Println(1)
}
