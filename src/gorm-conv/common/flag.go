package common

import (
	"flag"
	"fmt"
	"os"
)

var (
	H = flag.Bool("h", false, "get helper")
	// 输入xml文件
	Inputpath  = flag.String("I", "./", "input xml files path")
	Outputpath = flag.String("O", "./", "output files path")
	Cppoutpath = flag.String("cpppath", "", "cpp codes file out")
	Gooutpath  = flag.String("gopath", "", "cpp codes file out")
	// 输出类型flatbuffers,还是protobuffers配置文件
	Fbtype       = flag.String("fb", "", "general flatbuffers files")
	Pbtype       = flag.String("pb", "", "general protobuffers files")
	Sqltype      = flag.String("sql", "", "general sql files")
	Codetype     = flag.String("codetype", "client", "client codes or server codes")
	Gopackage    = flag.String("gopackage", "gorm/gorm", "option go_package")
	Protoversion = flag.String("protoversion", "3", "protobuff files version, 2 or 3")
)

var Proto_optional string = ""

func usage() {
	fmt.Fprintf(os.Stderr,
		`Usage: ggdb-conv [-h] [-input xmlconfigfile] 
	
	Options:
	`)
	flag.PrintDefaults()
}

func InitFlag() {
	flag.Parse()
}
