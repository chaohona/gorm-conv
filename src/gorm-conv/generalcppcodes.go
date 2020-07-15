package main

import (
	"fmt"
)

func GeneralCppCodes(games []XmlCfg, outpath string) int {
	// 1、生成表的列名和宏映射关系
	// 文件gorm_fields_map_define.cc
	if 0 != CppFieldsMapDefine(games, outpath) {
		fmt.Println("general cpp codes file gorm_fields_map_define.cc failed:", outpath)
	}
	if 0 != CppRedisDefine(games, outpath) {
		fmt.Println("general cpp codes file gorm_redis_define.cc failed:", outpath)
	}

	return 0
}
