package server

import (
	"fmt"
	"gorm-conv/common"
	"gorm-conv/server/mysql"
	"os"
	"strconv"
	"strings"
)

func CPPGeneralSplitTableName(games []common.XmlCfg, f *os.File) int {
	f.WriteString("int GORM_GetSplitTableName(int iTableId, uint32 uiHashCode, char *szOutTableName, int iInBuffLen, int &iUsedBuffLen)\n")
	f.WriteString("{\n")
	f.WriteString("    switch (iTableId)\n")
	f.WriteString("    {\n")
	for _, game := range games {
		for _, table := range game.DB.TableList {
			var bigTable string = strings.ToUpper(table.Name)
			f.WriteString("    case GORM_PB_TABLE_IDX_" + bigTable + ":\n")
			f.WriteString("    {\n")
			if table.SplitInfo.Num <= 1 {
				f.WriteString("        iUsedBuffLen = snprintf(szOutTableName, iInBuffLen, \" " + table.Name + " \");\n")
			} else {
				var strNum string = strconv.FormatInt(int64(table.SplitInfo.Num), 10)
				f.WriteString("        iUsedBuffLen = snprintf(szOutTableName, iInBuffLen, \" " + table.Name + "_%d \", uiHashCode%" + strNum + ");\n")
			}
			f.WriteString("        break;\n")
			f.WriteString("    }\n")
		}
	}
	f.WriteString("    default:\n")
	f.WriteString("        return GORM_INVALID_TABLE;\n")
	f.WriteString("    }\n")
	f.WriteString("    return GORM_OK;\n")
	f.WriteString("}\n")
	return 0
}

func gorm_general_mysql_define(games []common.XmlCfg, outpath string) int {
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

	if 0 != CPPGeneralSplitTableName(games, f) {
		fmt.Println("CPPGeneralSplitTableName failed.")
		return -1
	}

	if 0 != mysql.CPPFieldsMapPack_VERSION_SQL(games, f) {
		fmt.Println("CPPFieldsMapPack_VERSION_SQL failed.")
		return -1
	}
	if 0 != mysql.CPPFieldsMapPackInsertSQL(games, f) {
		fmt.Println("CPPFieldsMapPackInsertSQL failed.")
		return -1
	}
	if 0 != mysql.CPPFieldsMapPackGetSQL(games, f) {
		fmt.Println("CPPFieldsMapPackGetSQL failed.")
		return -1
	}
	if 0 != mysql.CPPFieldsMapPackDeleteSQL(games, f) {
		fmt.Println("CPPFieldsMapPackGetSQL failed.")
		return -1
	}
	if 0 != mysql.CPPFieldsMapPackUpdateSQL(games, f) {
		fmt.Println("CPPFieldsMapPackUpdateSQL failed.")
		return -1
	}
	if 0 != mysql.CPPFieldsMapPackIncreaseSQL(games, f) {
		fmt.Println("CPPFieldsMapPackIncreaseSQL failed.")
		return -1
	}
	if 0 != mysql.CPPFieldsMapPackReplaceSQL(games, f) {
		fmt.Println("CPPFieldsMapPackReplaceSQL failed.")
		return -1
	}
	if 0 != mysql.CPPFieldsMapPackBatchGetSQL(games, f) {
		fmt.Println("CPPFieldsMapPackBatchGetSQL failed.")
		return -1
	}
	if 0 != mysql.CPPFieldsMapPackGetByNonPrimaryKeySQL(games, f) {
		fmt.Println("CPPFieldsMapPackGetByNonPrimaryKeySQL_ForTables failed.")
		return -1
	}
	if 0 != mysql.GORM_MySQLResult2PbMSG(games, f) {
		fmt.Println("GORM_MySQLResult2PbMSG failed.")
		return -1
	}

	return 0
}

// 生成gorm_server专用的代码文件
func GORM_ServerCodesFilesGeneral(games []common.XmlCfg, outpath string) int {
	/*if 0 != CppRedisDefine(games, outpath) {
		fmt.Println("general cpp codes file gorm_redis_define.cc failed:", outpath)
		return -1
	}*/

	if 0 != gorm_general_mysql_define(games, outpath) {
		fmt.Println("gorm_general_mysql_define failed.")
		return -1
	}
	return 0
}
