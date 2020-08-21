package redis

import (
	"fmt"
	"gorm-conv/common"
	"gorm-conv/cpp"
	"os"
	"strconv"
	"strings"
)

func CppRedisDefine_Insert_Table(table common.TableInfo, f *os.File) int {
	var bigtable string = strings.ToUpper(table.Name)
	var definekey string = "GORM_REDIS_INSERT_" + bigtable + "_LUA"
	var define string = "#define " + definekey + " \" \\\n"
	define += "if redis.call(\\\"EXISTS\\\",KEYS[1]) == 0 then \\\n"
	define += "    redis.call(\\\"hmset\\\", KEYS[1]"
	for idx, c := range table.TableColumns {
		define += ", \\\""
		define += c.Name
		define += "\\\""
		define += ", ARGV["
		define += strconv.FormatInt(int64(idx+1), 10)
		define += "]"
	}
	define += ") \\\n"
	define += "end \\\n"
	define += "return 0\"\n"
	f.WriteString(define)

	f.WriteString("int GORM_REDIS_INSERT_" + bigtable + "(const GORM_PB_Table_" + table.Name + " &inTable, redisContext *pRedisConn)\n")
	f.WriteString("{\n")
	f.WriteString("    char szKey[2048];\n")
	f.WriteString("    snprintf(szKey, 2048, \"" + table.Name)
	for _, c := range table.SplitInfo.SplitCols {
		f.WriteString("_")
		tc := table.GetColumn(c)
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
	f.WriteString("    redisReply* r = (redisReply*)redisCommand(pRedisConn, \"EVAL %s 1 %s")
	for _, tc := range table.TableColumns {
		f.WriteString(" ")
		f.WriteString(cpp.CPPFieldPackRedis_COL_FORMAT(tc.Type))
	}
	f.WriteString("\", " + definekey + ", szKey, \n")

	f.WriteString("    ")
	for idx, tc := range table.TableColumns {
		if idx != 0 {
			f.WriteString(", ")
		}
		f.WriteString("inTable." + tc.Name + "()")
		if cpp.CPPField_CPPType(tc.Type) == "string" {
			f.WriteString(".c_str()")
		}
	}
	f.WriteString(");\n")
	f.WriteString("    if (r == nullptr || r->type == REDIS_REPLY_ERROR)\n")
	f.WriteString("        return GORM_CACHE_ERROR;\n")
	f.WriteString("    int iResult = -1;\n")
	f.WriteString("    GORM_REDIS_REPLY_LONG(r,  iResult);\n")
	f.WriteString("    if (iResult != 0)\n")
	f.WriteString("        return GORM_CACHE_ERROR;\n")

	f.WriteString("\n    return GORM_OK;\n")
	f.WriteString("}\n")
	return 0
}

func CppRedisDefine_Insert(games []common.XmlCfg, f *os.File) int {
	for _, game := range games {
		for _, table := range game.DB.TableList {
			if 0 != CppRedisDefine_Insert_Table(table, f) {
				fmt.Println("CppRedisDefine_Hmget_Table failed, for table:", table.Name)
			}
		}
	}

	CppRedisDefine_Opt("insert", games, f)
	return 0
}
