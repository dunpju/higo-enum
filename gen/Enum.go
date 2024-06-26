package gen

import (
	"fmt"
	"github.com/dunpju/higo-enum/gen/stubs"
	"github.com/dunpju/higo-utils/utils/dirutil"
	"github.com/dunpju/higo-utils/utils/fileutil"
	"github.com/dunpju/higo-utils/utils/stringutil"
	"log"
	"os"
	"regexp"
	"strings"
)

var (
	name  string
	out   string
	force bool
)

func InitEnum() {
	EnumCommand.Flags().StringVarP(&name, "name", "n", "", `枚举Token,--name=bin\\enum_cmd.md或--name="-e=state -f=状态:issue-1-发布,draft-2-草稿"`)
	err := EnumCommand.MarkFlagRequired("name")
	if err != nil {
		panic(err)
	}
	EnumCommand.Flags().StringVarP(&out, "out", "o", "", "生成目录,如:app\\Enums")
	err = EnumCommand.MarkFlagRequired("out")
	if err != nil {
		panic(err)
	}
	EnumCommand.Flags().BoolVarP(&force, "force", "f", false, "强制更新")
}

const enum = "Enum"

type EnumMap struct {
	Key   string
	Value interface{}
	Doc   string
}

func NewEnumMap(key string, value interface{}, doc string) *EnumMap {
	return &EnumMap{Key: stringutil.Ucfirst(stringutil.CaseToCamel(key)), Value: value, Doc: doc}
}

type Enum struct {
	Package   string
	Name      string
	OutStruct string
	OutDir    string
	File      string
	Doc       string
	RealName  string
	EnumType  string
	LenMap    int
	EnumMap   []*EnumMap
	Enums     []*Enum
}

func NewEnum(name string, file string) *Enum {
	reg := regexp.MustCompile(`(-e=[a-zA-Z_]+\s*-f=).*`)
	if reg == nil {
		log.Fatalln("regexp err")
	}
	e := &Enum{}
	if fs := reg.FindString(name); fs != "" {
		e.Enums = append(e.Enums, newEnum(name, file))
	} else {
		outfile := fileutil.ReadFile(name)
		if !outfile.Exist() {
			log.Fatalln(name + " configure file non-exist")
		}
		err := outfile.ForEach(func(line int, b []byte) bool {
			s := string(b)
			s = strings.Replace(s, "\\", "", -1)
			s = strings.Trim(s, "\n")
			s = strings.Trim(s, "\r")
			s = strings.Trim(s, "")
			if "" != s {
				e.Enums = append(e.Enums, newEnum(s, file))
			}
			return true
		})
		if err != nil {
			log.Fatalln(err)
		}
	}
	return e
}

func newEnum(name string, file string) *Enum {
	reg := regexp.MustCompile(`(-e=[a-zA-Z_]+\s*-f=).*`)
	if reg == nil {
		log.Fatalln("regexp err")
	}
	name = strings.Replace(name, "\\", "", -1)
	name = strings.Trim(name, "\n")
	name = strings.Trim(name, "\r")
	name = strings.Trim(name, "")
	E := &Enum{}
	if fs := reg.FindString(name); fs != "" {
		name = strings.Trim(fs, "")
		name = strings.Trim(name, "-e=")
		names := strings.Split(name, "-f=")
		if len(names) != 2 {
			log.Fatalln("name err")
		}
		name = strings.Trim(names[0], "")
		docs := strings.Split(names[1], ":")
		doc := strings.Trim(docs[0], "")
		E.Doc = doc
		characterReg := regexp.MustCompile(`([a-zA-Z_]).*`)
		if characterReg == nil {
			log.Fatalln("character regexp err")
		}
		es := strings.Split(docs[1], ",")
		for _, v := range es {
			em := strings.Split(v, "-")
			k := strings.Trim(em[0], "")
			v := strings.Trim(em[1], "")
			d := strings.Trim(strings.Trim(strings.Trim(em[2], "\n"), "\r"), "")
			E.EnumMap = append(E.EnumMap, NewEnumMap(k, v, d))
			if valueMatch := characterReg.FindString(v); valueMatch != "" {
				E.EnumType = "string"
			} else {
				E.EnumType = "int"
			}
		}
		E.LenMap = len(E.EnumMap) - 1
		name = stringutil.Ucfirst(stringutil.CaseToCamel(name))
		E.Name = name
		E.RealName = name
		E.OutDir = file + dirutil.PathSeparator() + enum + E.RealName
		E.OutStruct = E.OutDir + dirutil.PathSeparator() + enum + strings.Replace(name, enum, "", -1)
		E.File = E.OutDir + dirutil.PathSeparator() + "enum.go"
		E.Package = enum + name
		return E
	} else {
		log.Fatalln(`name format error: ` + name)
	}
	return E
}

func (this *Enum) Template(tplfile string) *stubs.Stub {
	return stubs.NewStub(tplfile)
}

func (this *Enum) Generate() {
	for _, e := range this.Enums {
		e.generate()
	}
}

func (this *Enum) generate() {
	dirutil.Dir(this.OutDir).Create()
	if fileutil.FileExist(this.File) && !force {
		log.Println(this.File + " already existed")
		return
	}
	outFile := fileutil.NewFile(this.File, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0755)
	if !outFile.Exist() {
		outFile.Create()
	}
	defer outFile.Close()
	tmpl, err := this.Template("enum.stub").Parse()
	if err != nil {
		panic(err)
	}
	//生成enum
	err = tmpl.Execute(outFile.File(), this)
	if err != nil {
		panic(err)
	}
	if force {
		fmt.Println("enum: " + this.OutStruct + " was forced updated.")
	} else {
		fmt.Println("enum: " + this.OutStruct + " generate success!")
	}
}
