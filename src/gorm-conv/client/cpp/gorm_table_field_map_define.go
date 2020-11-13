package cpp

import (
	"fmt"
	"gorm-conv/common"
	gorm_cpp "gorm-conv/cpp"
	"os"
)

func CppFieldsMapDefine(games []common.XmlCfg, outpath string) int {
	outfile := outpath + "gorm_table_field_map_define.cc"
	f, err := os.OpenFile(outfile, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err.Error())
		return -1
	}
	defer f.Close()
	fmt.Println("begin to general file:" + outfile)
	err = f.Truncate(0)

	f.WriteString(common.NOT_EDIT_WARNING)

	// 1、输出固定的头/////////////////////////
	//#include "mysql.h"
	var header string = `#include "gorm_table_field_map_define.h"
#include "gorm_pb_tables_inc.pb.h"
#include "gorm_pb_proto.pb.h"

namespace gorm{

`
	_, err = f.WriteString(header)
	if err != nil {
		fmt.Println(err.Error())
		return -1
	}

	if 0 != gorm_cpp.GORM_InitTableSchemaInfo(games, f) {
		fmt.Println("GORM_InitTableSchemaInfo failed.")
		return -1
	}

	f.WriteString("\n}\n")
	return 0
}
