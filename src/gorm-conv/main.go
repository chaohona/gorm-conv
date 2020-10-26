package main

import (
	"flag"
	"fmt"
	"os"

	"gorm-conv/client/golang"
	"gorm-conv/common"
	"gorm-conv/cpp"
	"gorm-conv/protobuf"
	"gorm-conv/server"
	"gorm-conv/server/config"
	"gorm-conv/sql"

	"github.com/golang/glog"

	clientCPP "gorm-conv/client/cpp"
)

func main() {
	common.InitFlag()

	if common.H != nil && *common.H {
		flag.Usage()
		return
	}
	defer glog.Flush()

	var inputfile string
	var inputpath string

	if common.Inputpath != nil && *common.Inputpath != "" {
		fmt.Println("config path >> ", *common.Inputpath)
		inputpath = *common.Inputpath
	} else if common.Inputfile != nil && *common.Inputfile != "" {
		fmt.Println("config file >> ", *common.Inputfile)
		inputfile = *common.Inputfile
	}

	if inputfile == "" && inputpath == "" {
		fmt.Println("-I/--xml there should one of them")
		return
	}

	if common.Outputpath != nil && *common.Outputpath != "" {
		_, err := os.Stat(*common.Outputpath)
		if os.IsNotExist(err) {
			if err = os.Mkdir(*common.Outputpath, os.ModePerm); err != nil {
				fmt.Println("make outputpath failed, path:", *common.Outputpath, ", errinfo:", err)
				return
			}
		}
		fmt.Println(*common.Outputpath)
	}

	if common.Cppoutpath != nil && *common.Cppoutpath != "" {
		_, err := os.Stat(*common.Cppoutpath)
		if os.IsNotExist(err) {
			if err := os.Mkdir(*common.Cppoutpath, os.ModePerm); err != nil {
				fmt.Println("make outputpath failed, path:", *common.Cppoutpath, ", errinfo:", err)
				return
			}
		}
		fmt.Println(*common.Cppoutpath)
	}

	if common.Gooutpath != nil && *common.Gooutpath != "" {
		_, err := os.Stat(*common.Gooutpath)
		if os.IsNotExist(err) {
			if err := os.Mkdir(*common.Gooutpath, os.ModePerm); err != nil {
				fmt.Println("make outputpath failed, path:", *common.Gooutpath, ", errinfo:", err)
				return
			}
		}
		fmt.Println(*common.Gooutpath)
	}

	fmt.Println("gorm-conv begin to work.")
	games, ret := common.ParseXmls(inputpath, inputfile)
	if ret != 0 {
		fmt.Println("parse xml failed.")
		return
	}
	if common.Sqltype != nil && *common.Sqltype == "true" {
		ret = sql.GeneralSQLFiles(games, *common.Outputpath)
		if ret != 0 {
			fmt.Println("generate sql file got error.")
			return
		}
	}
	if common.Fbtype != nil && *common.Fbtype == "true" {
		/*ret = flatbuffers.GeneralFBFiles(games, *outputpath)
		if ret != 0 {
			fmt.Println("generate fbs file got error.")
			return
		}*/
	}
	if common.Pbtype != nil && *common.Pbtype == "true" {
		if *common.Protoversion == "2" {
			common.Proto_optional = "optional "
		}
		ret = protobuf.GeneralPBFiles(games, *common.Outputpath)
		if ret != 0 {
			fmt.Println("generate pb file got error.")
			return
		}
		fmt.Println("generate pb files success")
	}
	var bServerCodes bool = false
	var bClientCodes bool = false
	if common.Codetype != nil && *common.Codetype == "server" {
		bServerCodes = true
	}
	if common.Codetype != nil && *common.Codetype == "client" {
		bClientCodes = true
	}
	// 自动生成代码
	if bServerCodes && common.Cppoutpath != nil && *common.Cppoutpath != "" {
		ret = cpp.GeneralCppServerCodes(games, *common.Cppoutpath)
		if ret != 0 {
			fmt.Println("generate cpp codes got error.")
			return
		}
	}
	if bServerCodes && common.Cppoutpath != nil && *common.Cppoutpath != "" {
		if 0 != server.GORM_ServerCodesFilesGeneral(games, *common.Cppoutpath) {
			fmt.Println("gorm_server_codes_files failed.")
			return
		}
	}
	// 生成客户端CPP代码
	if bClientCodes && common.Cppoutpath != nil && *common.Cppoutpath != "" {
		if 0 != clientCPP.GeneralClientCPPCodes(games, *common.Cppoutpath) {
			fmt.Println("gorm_client_codes_files failed.")
			return
		}
		fmt.Println("gorm_client_codes_files success")
	}
	if common.Gooutpath != nil && *common.Gooutpath != "" {
		ret = golang.GeneralClientGolangCodes(games, *common.Gooutpath)
		if ret != 0 {
			fmt.Println("general golang codes got error.")
			return
		}
	}
	if common.GeneralDBCfg != nil && *common.GeneralDBCfg == "true" {
		ret = config.GeneralDBCfg(games, *common.Outputpath)
		if ret != 0 {
			fmt.Println("general db route config failed.")
			return
		}
	}
	return
}
