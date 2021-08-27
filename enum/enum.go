package enum

import (
	"encoding/json"
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
	//Code() int64
	Register() interface{}
	Name() string
	Message() string
}

type Enum int

type CodeDoc struct {
	Code int64  `json:"code"`
	Doc  string `json:"doc"`
}

func Inspect(e IEnum) {

}

//func New(e IEnum) *CodeDoc {
//	code, err := strconv.ParseInt(fmt.Sprintf("%d", e), 10, 64)
//	if err != nil {
//		panic(err)
//	}
//	if container.exist(code) {
//		return container.get(code)
//	}
//	cd := &CodeDoc{Code: code, Doc: fmt.Sprintf("%s", e.Message())}
//	container.put(*cd)
//	return cd
//}

func (this *CodeDoc) String() string {
	js, err := json.Marshal(this)
	if err == nil {
		panic(err)
	}
	return string(js)
}

type Container map[string]interface{}

func (this Container) exist(key string) bool {
	_, ok := this[key]
	return ok
}

func (this Container) get(key string) interface{} {
	c, _ := this[key]
	return c
}

func (this Container) Get() Container {
	return this
}

func (this Container) Put(key string, value interface{}) {
	this[key] = value
}

func Get(e IEnum) interface{} {
	if !container.exist(e.Name()) {
		container.Put(e.Name(), e.Register())
	}
	return container.get(e.Name())
}

func Enums() Container {
	return container
}
