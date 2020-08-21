package cpp

import (
	"fmt"
	"gorm-conv/common"
)

func GeneralCppCodes(games []common.XmlCfg, outpath string) int {
	// 1、生成表的列名和宏映射关系
	// 文件gorm_fields_map_define.cc
	if 0 != CppFieldsMapDefine(games, outpath) {
		fmt.Println("general cpp codes file gorm_fields_map_define.cc failed:", outpath)
		return -1
	}
	if 0 != CppGormTables_H(games, outpath) {
		fmt.Println("CppGormTables_H failed")
		return -1
	}

	return 0
}
