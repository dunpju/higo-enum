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

func (this TestEnum) Name() string {
	return "TestEnum"
}

func (this TestEnum) Message() string {
	container := enum.Get(this)
	if e, ok := container.(map[TestEnum]string)[this]; ok {
		return e
	}
	return "Undefined"
}

func (this TestEnum) Register() interface{} {
	container := make(map[TestEnum]string)
	container[None] = "None"
	container[CPU] = "CPU"
	container[GPU] = "GPU"
	return container
}

func main() {
	fmt.Println(None, None.Message())
	//fmt.Println(CPU.Message(), CPU.Code())
	fmt.Println(enum.Enums())
	fmt.Println(GPU.Message(), GPU)
	fmt.Println(1)
}
