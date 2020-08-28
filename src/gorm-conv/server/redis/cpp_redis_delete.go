package redis

import (
	"fmt"
	"gorm-conv/common"
	"gorm-conv/cpp"
	"os"
	"strings"
)

func CppRedisDefine_Delete_Table(table common.TableInfo, f *os.File) int {
	var bigtable string = strings.ToUpper(table.Name)
	f.WriteString("int GORM_REDIS_DELETE_" + bigtable + "(const GORM_PB_Table_" + table.Name + " &inTable, redisContext *pRedisConn)\n")
	f.WriteString("{\n")
	f.WriteString("    char szKey[2048];\n")
	f.WriteString("    GORM_SafeSnprintf(szKey, 2048, \"" + table.Name)
	for _, c := range table.SplitInfo.SplitCols {
		tc := table.GetColumn(c)
		f.WriteString("_")
		f.WriteString(cpp.CPPFieldPackRedis_COL_FORMAT(tc.Type))
	}
	f.WriteString("\"")

	for _, c := range table.SplitInfo.SplitCols {
		f.WriteString(", inTable." + c + "()")
		tc := table.GetColumn(c)
		if cpp.CPPField_CPPType(tc.Type) == "string" {
			f.WriteString(".c_str()")
		}
	}
	f.WriteString(");\n")

	f.WriteString("    redisReply* r = (redisReply*)redisCommand(pRedisConn, \"del %s\", szKey);\n")
	f.WriteString("    if (r == nullptr || r->type == REDIS_REPLY_ERROR)\n")
	f.WriteString("        return GORM_CACHE_ERROR;\n")
	f.WriteString("    int iResult = -1;\n")
	f.WriteString("    GORM_REDIS_REPLY_LONG(r,  iResult);\n")
	f.WriteString("    if (iResult == -1)\n")
	f.WriteString("        return GORM_CACHE_ERROR;\n")
	f.WriteString("    return GORM_OK;\n")
	f.WriteString("}\n")
	return 0
}

func CppRedisDefine_Delete(games []common.XmlCfg, f *os.File) int {
	for _, game := range games {
		for _, table := range game.DB.TableList {
			if 0 != CppRedisDefine_Delete_Table(table, f) {
				fmt.Println("CppRedisDefine_Delete_Table failed, for table:", table.Name)
			}
		}
	}

	CppRedisDefine_Opt("delete", games, f)
	return 0
}
