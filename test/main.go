package main

import (
	"fmt"
	"github.com/dengpju/higo-enum/enum"
)

type TestEnum int

func (this TestEnum) Name() string {
	return "TestEnum"
}

func (this TestEnum) Inspect(value interface{}) error {
	return enum.Inspect(this, value)
}

func (this TestEnum) Message() string {
	return enum.String(this)
}

const (
	None TestEnum = 1
	CPU  TestEnum = 2 // 中央处理器
	GPU  TestEnum = 3 // 图形处理器
)

func (this TestEnum) Register() enum.Message {
	return make(enum.Message).
		Put(None, "None").
		Put(CPU, "CPU").
		Put(GPU, "GPU")
}

func main() {
	fmt.Println(None, None.Message())
	fmt.Println(enum.Enums())
	fmt.Println(GPU.Message(), GPU)
	fmt.Println(GPU.Inspect("1"))
	fmt.Println(1)
}
