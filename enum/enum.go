package enum

import (
	"fmt"
	"strconv"
	"sync"
)

var (
	container Container
	onceCode  sync.Once
)

func init() {
	onceCode.Do(func() {
		container = make(Container, 0)
	})
}

type IEnum interface {
	Inspect(value interface{}) error
	Register() Message
	Name() string
	Message() string
}

type Message map[IEnum]string

func (this Message) Len() int {
	return len(this)
}

func (this Message) Put(key IEnum, value string) Message {
	if _, ok := this[key]; ok {
		panic(key.Name() + " Value Duplicate Definition")
	}
	this[key] = value
	return this
}

func (this Message) Get(key IEnum) string {
	if s, ok := this[key]; ok {
		return s
	}
	return "Undefined"
}

func (this Message) Exist(key interface{}) bool {
	for k, _ := range this {
		if i, ok := key.(int); ok {
			d, err := strconv.Atoi(fmt.Sprintf("%d", k))
			if nil != err {
				panic(err)
			}
			if d == i {
				return true
			}
		} else if i64, ok := key.(int64); ok {
			d64, err := strconv.ParseInt(fmt.Sprintf("%d", k), 10, 64)
			if nil != err {
				panic(err)
			}
			if d64 == i64 {
				return true
			}
		} else if f, ok := key.(float64); ok {
			f64, err := strconv.ParseFloat(fmt.Sprintf("%f", k), 32)
			if nil != err {
				panic(err)
			}
			if f64 == f {
				return true
			}
		} else if s, ok := key.(string); ok {
			if fmt.Sprintf("%s", k) == s {
				return true
			}
		}
	}
	return false
}

type Container map[string]interface{}

func (this Container) get(key string) (interface{}, bool) {
	c, ok := this[key]
	return c, ok
}

func (this Container) Get() Container {
	return this
}

func (this Container) Put(key string, value interface{}) {
	this[key] = value
}

func Get(e IEnum) Message {
	c, ok := container.get(e.Name())
	if !ok {
		c := e.Register()
		container.Put(e.Name(), c)
		return c
	}
	return c.(Message)
}

func Inspect(e IEnum, value interface{}) error {
	if c, ok := container.get(e.Name()); ok {
		if m, ok := c.(Message); ok {
			if m.Exist(value) {
				return nil
			}
			return fmt.Errorf("%s not in enum value", value)
		}
		return fmt.Errorf("%sEnum Message Type Error", "")
	}
	return fmt.Errorf("%sUndefined Enum", "")
}

func Enums() Container {
	return container
}
