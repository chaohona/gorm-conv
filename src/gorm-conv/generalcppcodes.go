package main

import (
	"fmt"
)

func GeneralCppCodes(games []XmlCfg, outpath string, bServerCode bool) int {
	// 1、生成表的列名和宏映射关系
	// 文件gorm_fields_map_define.cc
	if 0 != CppFieldsMapDefine(games, outpath) {
		fmt.Println("general cpp codes file gorm_fields_map_define.cc failed:", outpath)
		return -1
	}
	if 0 != CppGormTables_H(games, outpath) {
		return -1
	}
	if bServerCode {
		if 0 != gorm_server_codes_files(games, outpath) {
			return -1
		}
	}

	return 0
}
