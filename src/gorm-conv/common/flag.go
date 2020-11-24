package common

import (
	"flag"
	"fmt"
	"os"
)

var (
	H = flag.Bool("h", false, "get helper")
	// 输入xml文件
	Inputpath  = flag.String("I", "", "input xml files path")     // xml配置文件路径
	Inputfile  = flag.String("xml", "", "input xml file")         // xml文件全路径，这个参数和上面的-I只能二选一使用
	Outputpath = flag.String("O", "./", "output files path")      // 输出文件路径，包括sql语句与proto配置文件
	Cppoutpath = flag.String("cpppath", "", "cpp codes file out") // 输出的cpp代码路径
	Gooutpath  = flag.String("gopath", "", "cpp codes file out")  // 输出的golang代码路径(目前只用于golang版本客户端)
	// 输出类型flatbuffers,还是protobuffers配置文件
	Fbtype        = flag.String("fb", "", "general flatbuffers files")                          // 输出为flatbuffer格式，目前不使用
	Pbtype        = flag.String("pb", "", "general protobuffers files")                         // 输出问protobuffer格式，目前主要是使用pb作为传输协议
	Sqltype       = flag.String("sql", "", "general sql files")                                 // 是否输出创建表的sql语句文件
	Codetype      = flag.String("codetype", "client", "client codes or server codes")           // 输出的代码为客户端代码试试服务器代码,client/server
	Gopackage     = flag.String("gopackage", "gorm/gorm", "option go_package")                  // golang客户端版本，package 包名
	Protoversion  = flag.String("protoversion", "3", "protobuff files version, 2 or 3")         // 使用的proto版本为2还是3
	GeneralDBCfg  = flag.String("generalcfg", "false", "general db route config true/false")    // 是否生成gorm-db.yml配置文件
	CPP_Coroutine = flag.String("cpp_coroutine", "false", "general db route config true/false") // C++版本的客户端是否支持协程
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
