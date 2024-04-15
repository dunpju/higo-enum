package EnumCategory

import (
	"fmt"
)

var (
	enums m
)

const (
	Dept    Category = 1 //部门
	Company Category = 2 //公司
	Team    Category = 3 //团队
	Group   Category = 4 //组
)

func init() {
	enums = e().
		append(Dept, newEnum(int(Dept), "部门")).
		append(Company, newEnum(int(Company), "公司")).
		append(Team, newEnum(int(Team), "团队")).
		append(Group, newEnum(int(Group), "组"))
}

type enum struct {
	code    int
	message string
}

func newEnum(code int, message string) *enum {
	return &enum{code: code, message: message}
}

type m map[Category]*enum

func e() m {
	return make(map[Category]*enum)
}

func (this m) append(key Category, value *enum) m {
	this[key] = value
	return this
}

func Inspect(value int) error {
	_, err := Category(value).inspect()
	if err != nil {
		return err
	}
	return nil
}

func Get(code int) Category {
	_ = Category(code).get()
	return Category(code)
}

// Category 种类
type Category int

func (this Category) inspect() (*enum, error) {
	if e, ok := enums[this]; ok {
		return e, nil
	}
	return nil, fmt.Errorf("%d enum undefined", this)
}

func (this Category) get() *enum {
	e, err := this.inspect()
	if err != nil {
		panic(err)
	}
	return e
}

func (this Category) Code() int {
	return this.get().code
}

func (this Category) Message() string {
	return this.get().message
}
