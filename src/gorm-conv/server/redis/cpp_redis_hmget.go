package redis

import (
	"fmt"
	"gorm-conv/common"
	"gorm-conv/cpp"
	"os"
	"strconv"
	"strings"
)

func CppRedisDefine_Hmget_Table(table common.TableInfo, f *os.File) int {
	var bigtable string = strings.ToUpper(table.Name)
	var define string = "#define GORM_" + bigtable + "_REDIS_HMGET \"hmget %s "
	for _, tc := range table.TableColumns {
		define += " " + tc.Name
	}
	define += "\"\n"
	f.WriteString(define)

	f.WriteString("int GORM_REDIS_HMGET_" + bigtable + "(const GORM_PB_Table_" + table.Name + " &inTable, redisContext *pRedisConn, GORM_PB_TABLE *pOutPbTable, bool &bGotResult)\n")
	f.WriteString("{\n")
	f.WriteString("    char szKey[2048];\n")
	f.WriteString("    GORM_SafeSnprintf(szKey, 2048, \"" + table.Name)
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

	f.WriteString("    redisReply* r = (redisReply*)redisCommand(pRedisConn, GORM_" + bigtable + "_REDIS_HMGET, szKey);\n")

	var columnsLen string = strconv.FormatInt(int64(len(table.TableColumns)), 10)
	f.WriteString("    if (r == nullptr)\n")
	f.WriteString("        return GORM_CACHE_ERROR;\n")
	f.WriteString("    if (r->type != REDIS_REPLY_ARRAY || r->elements != " + columnsLen + ")\n")
	f.WriteString("    {\n")
	f.WriteString("        freeReplyObject(r);\n")
	f.WriteString("        return GORM_CACHE_ERROR;\n")
	f.WriteString("    }\n")
	f.WriteString("    bGotResult = false;\n")
	f.WriteString("    for (int i=0; i<" + columnsLen + "; i++)\n")
	f.WriteString("    {\n")
	f.WriteString("        if ((*(r->element+i))->type != REDIS_REPLY_NIL)\n")
	f.WriteString("        {\n")
	f.WriteString("            bGotResult = true;\n")
	f.WriteString("            break;\n")
	f.WriteString("        }\n")
	f.WriteString("    }\n")
	f.WriteString("    if (!bGotResult)\n")
	f.WriteString("    {\n")
	f.WriteString("        freeReplyObject(r);\n")
	f.WriteString("        return GORM_OK;\n")
	f.WriteString("    }\n")
	//f.WriteString("    pOutPbTable = new GORM_PB_TABLE();\n")
	//f.WriteString("    if (pOutPbTable == nullptr)\n")
	//f.WriteString("    {\n")
	//f.WriteString("        freeReplyObject(r);\n")
	//f.WriteString("        return GORM_CACHE_ERROR;\n")
	//f.WriteString("    }\n")

	var table_ptr string = "p" + bigtable
	f.WriteString("    GORM_PB_Table_" + table.Name + " *" + table_ptr + " = pOutPbTable->mutable_" + table.Name + "();\n")
	f.WriteString("    if (" + table_ptr + " == nullptr)\n")
	f.WriteString("    {\n")
	//f.WriteString("        delete pOutPbTable;\n")
	//f.WriteString("        pOutPbTable = nullptr;\n")
	f.WriteString("        freeReplyObject(r);\n")
	f.WriteString("        return GORM_CACHE_ERROR;\n")
	f.WriteString("    }\n")
	f.WriteString("    redisReply *pElement = nullptr;\n")

	for idx, c := range table.TableColumns {
		f.WriteString("    pElement = *(r->element+" + strconv.FormatInt(int64(idx), 10) + ");\n")
		if cpp.CPPField_CPPType(c.Type) == "string" {
			f.WriteString("    if (pElement->len > 0)\n")
			f.WriteString("        " + table_ptr + "->set_" + c.Name + "(pElement->str, pElement->len);\n")
		} else {
			f.WriteString("    " + cpp.CPPFieldsMapPackSQL_ForTables_COL2SQL_GetCPPType(c.Type) + " " + table.Name + "_" + c.Name + " = 0;\n")
			if c.Type == "double" {
				f.WriteString("    GORM_REDIS_REPLY_DOUBLE(pElement, " + table.Name + "_" + c.Name + ");\n")
			} else if cpp.CPPFieldsMapPackSQL_ForTables_COL2SQL_GET_LONG_ULONG(c.Type) == "long" {
				f.WriteString("    GORM_REDIS_REPLY_LONG(pElement, " + table.Name + "_" + c.Name + ");\n")
			} else {
				f.WriteString("    GORM_REDIS_REPLY_ULONG(pElement, " + table.Name + "_" + c.Name + ");\n")
			}
			f.WriteString("    " + table_ptr + "->set_" + c.Name + "(" + table.Name + "_" + c.Name + ");\n")
		}
	}

	f.WriteString("\n    freeReplyObject(r);\n")
	f.WriteString("    bGotResult = true;\n")
	f.WriteString("    return GORM_OK;\n")
	f.WriteString("}\n")
	return 0
}

func CppRedisDefine_Hmget(games []common.XmlCfg, f *os.File) int {
	for _, game := range games {
		for _, table := range game.DB.TableList {
			if 0 != CppRedisDefine_Hmget_Table(table, f) {
				fmt.Println("CppRedisDefine_Hmget_Table failed, for table:", table.Name)
			}
		}
	}
	f.WriteString("\n")
	f.WriteString("int GORM_REDIS_HMGET(int iTableId, redisContext *pRedisConn, GORM_PB_TABLE *pPbTable, GORM_PB_TABLE *pOutPbTable, bool &bGotResult)\n")
	f.WriteString("{\n")
	f.WriteString("    bGotResult = false;\n")
	f.WriteString("    switch (iTableId)\n")
	f.WriteString("    {\n")
	for _, game := range games {
		for _, table := range game.DB.TableList {
			f.WriteString("    case GORM_PB_TABLE_IDX_" + strings.ToUpper(table.Name) + ":\n")
			f.WriteString("    {\n")
			f.WriteString("        if (!pPbTable->has_" + table.Name + "())\n")
			f.WriteString("            return GORM_REQ_NO_RECORDS;\n")
			f.WriteString("        return GORM_REDIS_HMGET_" + strings.ToUpper(table.Name) + "(pPbTable->" + table.Name + "(), pRedisConn, pOutPbTable, bGotResult);\n")
			f.WriteString("    }\n")
		}
	}
	f.WriteString("    default:\n")
	f.WriteString("        return GORM_OK;\n")
	f.WriteString("    }\n")
	f.WriteString("}\n")
	return 0
}
