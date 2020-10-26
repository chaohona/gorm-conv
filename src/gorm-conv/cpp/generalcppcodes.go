package cpp

import (
	"fmt"
	"gorm-conv/common"
)

func GeneralCppServerCodes(games []common.XmlCfg, outpath string) int {
	// 1、生成表的列名和宏映射关系
	// 文件gorm_fields_map_define.cc,只有服务器代码需要生成
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
