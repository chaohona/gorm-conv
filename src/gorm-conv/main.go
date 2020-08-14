package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/golang/glog"
)

var (
	h = flag.Bool("h", false, "get helper")
	// 输入xml文件
	inputpath  = flag.String("I", "./", "input xml files path")
	outputpath = flag.String("O", "./", "output files path")
	cppoutpath = flag.String("cpppath", "", "cpp codes file out")
	gooutpath  = flag.String("gopath", "", "cpp codes file out")
	// 输出类型flatbuffers,还是protobuffers配置文件
	fbtype   = flag.String("fb", "", "general flatbuffers files")
	pbtype   = flag.String("pb", "", "general protobuffers files")
	sqltype  = flag.String("sql", "", "general sql files")
	codetype = flag.String("codetype", "client", "client codes or server codes")
)

func usage() {
	fmt.Fprintf(os.Stderr,
		`Usage: ggdb-conv [-h] [-input xmlconfigfile] 
	
	Options:
	`)
	flag.PrintDefaults()
}

func main() {
	flag.Parse()
	if h != nil && *h {
		flag.Usage()
		return
	}
	defer glog.Flush()

	if inputpath != nil {
		fmt.Println(*inputpath)
	}
	if outputpath != nil {
		fmt.Println(*outputpath)
	}

	fmt.Println("ggdb-conv begin to work.")
	games, ret := ParseXmls(*inputpath)
	if ret != 0 {
		fmt.Println("parse xml failed.")
		return
	}
	if sqltype != nil {
		ret = GeneralSQLFiles(games, *outputpath)
		if ret != 0 {
			fmt.Println("generate sql file got error.")
			return
		}
	}
	if fbtype != nil {
		/*ret = GeneralFBFiles(games, *outputpath)
		if ret != 0 {
			fmt.Println("generate fbs file got error.")
			return
		}*/
	}
	if pbtype != nil && *pbtype != "" {
		ret = GeneralPBFiles(games, *outputpath)
		if ret != 0 {
			fmt.Println("generate pb file got error.")
			return
		}
		fmt.Println("generate pb files success")
	}
	var bServerCodes bool = false
	if codetype != nil && *codetype == "server" {
		bServerCodes = true
	}
	// 自动生成代码
	if cppoutpath != nil && *cppoutpath != "" {
		ret = GeneralCppCodes(games, *cppoutpath, bServerCodes)
		if ret != 0 {
			fmt.Println("generate cpp codes got error.")
			return
		}
	}
	if gooutpath != nil && *gooutpath != "" {
		ret = GeneralGolangCodes(games, *gooutpath)
		if ret != 0 {
			fmt.Println("general golang codes got error.")
			return
		}
	}
	return
}
