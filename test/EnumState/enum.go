package EnumState

import (
	"fmt"
)

var (
	enums map[state]*impl
)

const (
	Del  state = 1
	Undo state = 2
)

func init() {
	enums = make(map[state]*impl)
	enums[Del] = newEnum(int(Del), "已删除", "#f24b4b", "已删除")
	enums[Undo] = newEnum(int(Undo), "未处理", "#f47f28", "去处理")
}

func Enums() map[state]*impl {
	return enums
}

func Inspect(value int) error {
	_, err := state(value).inspect()
	if err != nil {
		return err
	}
	return nil
}

type state int

func (this state) inspect() (*impl, error) {
	if e, ok := enums[this]; ok {
		return e, nil
	}
	return nil, fmt.Errorf("%d enum undefined", this)
}

func (this state) get() *impl {
	e, err := this.inspect()
	if err != nil {
		panic(err)
	}
	return e
}

func (this state) Code() int {
	return this.get().code
}

func (this state) Message() string {
	return this.get().message
}

func (this state) Color() string {
	return this.get().color
}

func (this state) Button() string {
	return this.get().button
}

type impl struct {
	code                   int
	message, color, button string
}

func newEnum(code int, message string, color string, button string) *impl {
	return &impl{code: code, message: message, color: color, button: button}
}
