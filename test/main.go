package main

import (
	"fmt"
	"github.com/dunpju/higo-enum/enum"
	"github.com/dunpju/higo-enum/gen"
	"github.com/dunpju/higo-enum/test/EnumState"
	"github.com/dunpju/higo-enum/test/app/enums/EnumCategory"
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
	fmt.Println(EnumState.Del.Message())
	fmt.Println(EnumState.Del.Code())
	fmt.Println(EnumState.Enums())
	err := EnumState.Inspect(3)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(EnumState.Undo.Message())
	fmt.Println(EnumState.Undo.Code())
	fmt.Println(EnumState.Undo.Color())
	fmt.Println(EnumState.Undo.Button())

	gen.Execute()

	fmt.Println(EnumCategory.Team)
	fmt.Println(EnumCategory.Team.Code())
	fmt.Println(EnumCategory.Team.Message())
	e := EnumCategory.Get(1)
	fmt.Println(e)
	fmt.Println(e.Code())
	fmt.Println(e.Message())
}
