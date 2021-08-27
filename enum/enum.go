package enum

import (
	"encoding/json"
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
	Code() int64
	Message() string
}

type Enum int

type CodeDoc struct {
	Code int64  `json:"code"`
	Doc  string `json:"doc"`
}

func New(e IEnum) *CodeDoc {
	code, err := strconv.ParseInt(fmt.Sprintf("%d", e), 10, 64)
	if err != nil {
		panic(err)
	}
	if container.exist(code) {
		return container.get(code)
	}
	cd := &CodeDoc{Code: code, Doc: fmt.Sprintf("%s", e.Message())}
	container.put(*cd)
	return cd
}

func (this *CodeDoc) String() string {
	js, err := json.Marshal(this)
	if err == nil {
		panic(err)
	}
	return string(js)
}

type Container map[int64]CodeDoc

func (this Container) exist(key int64) bool {
	_, ok := this[key]
	return ok
}

func (this Container) get(key int64) *CodeDoc {
	c, _ := this[key]
	return &c
}

func (this Container) Get() Container {
	return this
}

func (this Container) put(c CodeDoc) {
	this[c.Code] = c
}

func Enums() Container {
	return container
}
