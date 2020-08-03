package main

import (
	"fmt"
	"os"
)

func gorm_general_mysql_define(games []XmlCfg, outpath string) int {
	outfile := outpath + "server/gorm_server_table_define.cc"
	f, err := os.OpenFile(outfile, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err.Error())
		return -1
	}
	defer f.Close()
	fmt.Println("begin to general file:" + outfile)
	err = f.Truncate(0)

	// 1、输出固定的头/////////////////////////
	var header string = `#include "gorm_server_table_define.h"
#include "gorm_table_field_map_define.h"
#include "gorm_server_table_stable.h"
#include "gorm_mysql_conn_pool.h"
#include "gorm_pb_proto.pb.h"
#include "gorm_mempool.h"
#include "gorm_hash.h"
#include "mysql.h"

using namespace gorm;
`

	f.WriteString(header)

	if 0 != CPPFieldsMapPack_VERSION_SQL(games, f) {
		fmt.Println("CPPFieldsMapPack_VERSION_SQL failed.")
		return -1
	}
	if 0 != CPPFieldsMapPackInsertSQL(games, f) {
		fmt.Println("CPPFieldsMapPackInsertSQL failed.")
		return -1
	}
	if 0 != CPPFieldsMapPackGetSQL(games, f) {
		fmt.Println("CPPFieldsMapPackGetSQL failed.")
		return -1
	}
	if 0 != CPPFieldsMapPackDeleteSQL(games, f) {
		fmt.Println("CPPFieldsMapPackGetSQL failed.")
		return -1
	}
	if 0 != CPPFieldsMapPackUpdateSQL(games, f) {
		fmt.Println("CPPFieldsMapPackUpdateSQL failed.")
		return -1
	}
	if 0 != CPPFieldsMapPackIncreaseSQL(games, f) {
		fmt.Println("CPPFieldsMapPackIncreaseSQL failed.")
		return -1
	}
	if 0 != CPPFieldsMapPackReplaceSQL(games, f) {
		fmt.Println("CPPFieldsMapPackReplaceSQL failed.")
		return -1
	}
	if 0 != CPPFieldsMapPackBatchGetSQL(games, f) {
		fmt.Println("CPPFieldsMapPackBatchGetSQL failed.")
		return -1
	}
	if 0 != GORM_MySQLResult2PbMSG(games, f) {
		fmt.Println("GORM_MySQLResult2PbMSG failed.")
		return -1
	}

	return 0
}

// 生成gorm_server专用的代码文件
func gorm_server_codes_files(games []XmlCfg, outpath string) int {
	if 0 != CppRedisDefine(games, outpath) {
		fmt.Println("general cpp codes file gorm_redis_define.cc failed:", outpath)
		return -1
	}

	if 0 != gorm_general_mysql_define(games, outpath) {
		fmt.Println("gorm_general_mysql_define failed.")
		return -1
	}
	return 0
}
