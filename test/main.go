package main

import (
	"fmt"
	"higo-enum/enum"
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

func (this Enum) Message() *enum.CodeDoc {
	return enum.New(this)
}

func main() {
	fmt.Println(CPU.Message().Code, CPU.Message().Doc)
	fmt.Println(enum.Enums())
	fmt.Println(GPU.Message().Code, GPU.Message().Doc)
	fmt.Println(1)
}
