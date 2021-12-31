package EnumState

import "fmt"

const (
	Del  tenum = 1
	Undo tenum = 2
)

func (this tenum) register() *enum {
	switch (this) {
	case Del:
		return newEnum(int(Del), "已删除", "#f24b4b", "已删除")
	case Undo:
		return newEnum(int(Undo), "未处理", "#f47f28", "去处理")
	default:
		panic(fmt.Errorf("enum type undefined"))
	}
}

type tenum int

func (this tenum) Code() int {
	return this.register().code
}

func (this tenum) Message() string {
	return this.register().message
}

func (this tenum) Color() string {
	return this.register().color
}

func (this tenum) Button() string {
	return this.register().button
}

type enum struct {
	code                   int
	message, color, button string
}

func newEnum(code int, message string, color string, button string) *enum {
	return &enum{code: code, message: message, color: color, button: button}
}
