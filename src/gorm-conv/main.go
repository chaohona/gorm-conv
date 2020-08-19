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
	fbtype       = flag.String("fb", "", "general flatbuffers files")
	pbtype       = flag.String("pb", "", "general protobuffers files")
	sqltype      = flag.String("sql", "", "general sql files")
	codetype     = flag.String("codetype", "client", "client codes or server codes")
	gopackage    = flag.String("gopackage", "gorm/gorm", "option go_package")
	protoversion = flag.String("protoversion", "3", "protobuff files version, 2 or 3")
)

var proto_optional string = ""

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
	if outputpath != nil && *outputpath != "" {
		_, err := os.Stat(*outputpath)
		if os.IsNotExist(err) {
			if err = os.Mkdir(*outputpath, os.ModePerm); err != nil {
				fmt.Println("make outputpath failed, path:", *outputpath, ", errinfo:", err)
				return
			}
		}
		fmt.Println(*outputpath)
	}

	if cppoutpath != nil && *cppoutpath != "" {
		_, err := os.Stat(*cppoutpath)
		if os.IsNotExist(err) {
			if err := os.Mkdir(*cppoutpath, os.ModePerm); err != nil {
				fmt.Println("make outputpath failed, path:", *cppoutpath, ", errinfo:", err)
				return
			}
		}
		fmt.Println(*cppoutpath)
	}

	if gooutpath != nil && *gooutpath != "" {
		_, err := os.Stat(*gooutpath)
		if os.IsNotExist(err) {
			if err := os.Mkdir(*gooutpath, os.ModePerm); err != nil {
				fmt.Println("make outputpath failed, path:", *gooutpath, ", errinfo:", err)
				return
			}
		}
		fmt.Println(*gooutpath)
	}

	fmt.Println("gorm-conv begin to work.")
	games, ret := ParseXmls(*inputpath)
	if ret != 0 {
		fmt.Println("parse xml failed.")
		return
	}
	if sqltype != nil && *sqltype == "true" {
		ret = GeneralSQLFiles(games, *outputpath)
		if ret != 0 {
			fmt.Println("generate sql file got error.")
			return
		}
	}
	if fbtype != nil && *fbtype == "true" {
		/*ret = GeneralFBFiles(games, *outputpath)
		if ret != 0 {
			fmt.Println("generate fbs file got error.")
			return
		}*/
	}
	if pbtype != nil && *pbtype == "true" {
		if *protoversion == "2" {
			proto_optional = "optional "
		}
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
