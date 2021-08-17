package main

import (
	"fmt"
	"github.com/dengpju/higo-enum/enum"
)

type Enum int

const (
	None Enum = iota
	CPU   // 中央处理器
	GPU   // 图形处理器
)

func (this Enum) String() string {
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

func (this Enum) Message() string {
	return enum.New(this).Doc
}

func (this Enum) Code() int64 {
	return enum.New(this).Code
}

func main() {
	fmt.Println(CPU.Message(), CPU.Code())
	fmt.Println(enum.Enums())
	fmt.Println(GPU.Message(), GPU.Code())
	fmt.Println(1)
}
