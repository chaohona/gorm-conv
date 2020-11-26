package mysql

import (
	"fmt"
	"gorm-conv/common"
	"os"
	"strconv"
	"strings"
)

func CPPFieldsMapPackGetSQL_ForTables_DefineSQL_Where(table common.TableInfo) (string, int) {
	var DefineSQL string
	DefineSQL += " from "
	DefineSQL += table.Name
	DefineSQL += "_%d where "
	var splitInfo common.SplitInfo = table.SplitInfo
	if splitInfo.Columns == "" {
		DefineSQL += "\""
		return DefineSQL, 0
	}
	var matchNum int = 0
	for _, colname := range splitInfo.SplitCols {
		var match bool = false
		for _, preCol := range table.TableColumns {
			if colname != preCol.Name {
				continue
			}
			matchNum += 1
			match = true
			if matchNum != 1 {
				DefineSQL += " and"
			}
			DefineSQL += " `"
			DefineSQL += preCol.Name + "`="
			DefineSQL += common.CPPFieldPackSQL_COL_FORMAT(preCol.Type)
		}
		if !match {
			fmt.Println("invalid splitinfo, table:", table.Name)
			return "", -1
		}
	}

	return DefineSQL, 0
}

func CPPFieldsMapPackGetSQL_ForTables_DefineSQL(table common.TableInfo) (string, int) {
	var DefineSQL string = "#define " + strings.ToUpper(table.Name) + "GETSQL \"select "
	for idx, c := range table.TableColumns {
		if idx != 0 {
			DefineSQL += ", "
		}
		DefineSQL += "`"
		DefineSQL += c.Name
		DefineSQL += "`"
	}
	where, ret := CPPFieldsMapPackGetSQL_ForTables_DefineSQL_Where(table)
	if 0 != ret {
		return "", -1
	}
	DefineSQL += where
	DefineSQL += " limit 1"

	DefineSQL += ";\n"
	return DefineSQL, 0
}

func CPPFieldsMapPackGetSQL_ForTables_One(table common.TableInfo, strsqllen string, f *os.File) int {
	var bigtable string = strings.ToUpper(table.Name)
	f.WriteString("int GORM_PackGetSQL" + bigtable + "_ONE(shared_ptr<GORM_MemPool> &pMemPool, MYSQL* mysql, int iTableIndex, const GORM_PB_Table_" + table.Name + " &table_" + table.Name + ", GORM_MemPoolData *&pReqData)\n")
	f.WriteString("{\n")
	var ilenstr string = "    int iLen = iSqlLen + 128"
	var GORM_SafeSnprintfstr string = "    iLen = GORM_SafeSnprintf(szSQLBegin, iLen, " + strings.ToUpper(table.Name) + "GETSQL, iTableIndex"
	var releasestr string = ""

	f.WriteString("    char *szSQLBegin = nullptr;\n")
	f.WriteString("    int iSqlLen = ")
	f.WriteString(strsqllen)
	f.WriteString(";\n")
	f.WriteString("    int iTmpLen = 0;\n\n")

	for _, cname := range table.SplitInfo.SplitCols {
		for _, col := range table.TableColumns {
			if col.Name != cname {
				continue
			}
			var colType string = common.CPPFieldsMapPackSQL_ForTables_COL2SQL_GetCPPType(col.Type)
			f.WriteString("    const ")
			f.WriteString(colType)
			f.WriteString(" ")
			if colType == "string" {
				f.WriteString("&")
			}
			f.WriteString(table.Name + "_" + col.Name)
			f.WriteString(" = table_" + table.Name + "." + col.Name + "();\n")

			var table_col string = table.Name + "_" + col.Name
			if colType == "string" {
				f.WriteString("    const char *sz_" + table_col + " = \"\";\n")
				f.WriteString("    int len_" + table_col + " = 0;\n")
				f.WriteString("    GORM_MemPoolData *buffer_" + table_col + " = nullptr;\n")
				f.WriteString("    if(" + table_col + ".size() > 0)\n")
				f.WriteString("    {\n")

				var bufferName string = "buffer_" + table_col
				var bufferSize string = table_col + ".size()<<1"
				PrintGetBuffFromMemPool(f, bufferName, bufferSize)

				f.WriteString("        iTmpLen=mysql_real_escape_string(mysql, buffer_" + table_col + "->m_uszData, " + table_col + ".c_str(), " + table_col + ".size());\n")
				f.WriteString("        buffer_" + table_col + "->m_uszData[iTmpLen] = 0;\n")
				f.WriteString("        buffer_" + table_col + "->m_sUsedSize = iTmpLen;\n")
				f.WriteString("        sz_" + table_col + " = buffer_" + table_col + "->m_uszData;\n")
				f.WriteString("        len_" + table_col + " = iTmpLen;\n")
				f.WriteString("    }\n")
			}

			if colType == "string" {
				ilenstr += " + len_" + table_col
				GORM_SafeSnprintfstr += ", sz_" + table_col
				releasestr += "    if(buffer_" + table_col + " != nullptr)\n"
				releasestr += "        buffer_" + table_col + "->Release();\n"
			} else {
				ilenstr += " + 8"
				GORM_SafeSnprintfstr += ", " + table_col
			}
		}
	}
	ilenstr += " + table_" + table.Name + ".ByteSizeLong();\n"
	GORM_SafeSnprintfstr += ");\n"
	f.WriteString(ilenstr)

	var bufferName string = "pReqData"
	var bufferSize string = "iLen"
	PrintGetBuffFromMemPool(f, bufferName, bufferSize)

	f.WriteString("    szSQLBegin = pReqData->m_uszData;\n")
	f.WriteString(GORM_SafeSnprintfstr)
	f.WriteString("    pReqData->m_sUsedSize = iLen;\n\n")
	f.WriteString(releasestr)

	f.WriteString("\n    return GORM_OK;\n")
	f.WriteString("}\n")
	return 0
}

func CPPFieldsMapPackGetSQL_ForTables_One_Debug(table common.TableInfo, strsqllen string, f *os.File) int {
	return 0
	where, _ := CPPFieldsMapPackGetSQL_ForTables_DefineSQL_Where(table)
	var bigtable string = strings.ToUpper(table.Name)
	f.WriteString("#define " + bigtable + "GETSQL_DEBUG_WHERE \"" + where + "\n")

	f.WriteString("int GORM_PackGetSQL" + bigtable + "_ONE_DEBUG(shared_ptr<GORM_MemPool> &pMemPool, GORM_MySQLEvent *pMySQLEvent, int iTableIndex, const GORM_PB_Table_" + table.Name + " &table_" + table.Name + ", GORM_MemPoolData *&pReqData)\n")
	f.WriteString("{\n")
	f.WriteString(`
	MYSQL* mysql = pMySQLEvent->m_pMySQL;
    char *szSQLBegin = nullptr;
    int iSqlLen = 93;
    int iTmpLen = 0;
`)

	f.WriteString("    auto itr = pMySQLEvent->m_mapTablesColumnInfo.find(string(\"" + table.Name + "\"));\n")
	f.WriteString("    if (itr == pMySQLEvent->m_mapTablesColumnInfo.end())\n")
	f.WriteString("    {\n")
	f.WriteString("        GORM_LOGE(\"" + table.Name + " table info not inited\");\n")
	f.WriteString("        return GORM_ERROR;\n")
	f.WriteString("    }\n")
	f.WriteString("    unordered_map<string, GORM_PB_COLUMN_TYPE> &columnMap = itr->second;\n")
	f.WriteString("    if (columnMap.size() == 0)\n")
	f.WriteString("        return GORM_ERROR;\n")
	f.WriteString("    int iTotalLen = iSqlLen + 128 + 64*columnMap.size() + table_" + table.Name + ".ByteSizeLong();\n")
	f.WriteString("    int iLen = 7;\n")

	var bufferName string = "pReqData"
	var bufferSize string = "iTotalLen"
	PrintGetBuffFromMemPool(f, bufferName, bufferSize)

	f.WriteString("    szSQLBegin = pReqData->m_uszData;\n")
	f.WriteString("    memcpy(szSQLBegin, \"select \", 7);\n")
	f.WriteString("    int idx = 0;\n")
	f.WriteString("    vector<string> &vectorColumns = pMySQLEvent->m_mapTablesColumnOrder[\"" + table.Name + "\"];\n")
	f.WriteString("    if (vectorColumns.size() != columnMap.size())\n")
	f.WriteString("        return GORM_ERROR;\n")
	f.WriteString("    for (auto c : vectorColumns)\n")
	f.WriteString("    {\n")
	f.WriteString("        auto itr = columnMap.find(c);\n")
	f.WriteString("        if (itr == columnMap.end())\n")
	f.WriteString("            return GORM_ERROR;\n")
	f.WriteString("        if (idx++ != 0)\n")
	f.WriteString("            iLen += GORM_SafeSnprintf(szSQLBegin+iLen, iTotalLen-iLen, \",`%s`\", (char*)(c.c_str()));\n")
	f.WriteString("        else\n")
	f.WriteString("            iLen += GORM_SafeSnprintf(szSQLBegin+iLen, iTotalLen-iLen, \"`%s`\", (char*)(c.c_str()));\n")
	f.WriteString("    }\n")

	var GORM_SafeSnprintfstr string = "    iLen += GORM_SafeSnprintf(szSQLBegin+iLen, iTotalLen-iLen, " + bigtable + "GETSQL_DEBUG_WHERE, iTableIndex"
	var releasestr string = ""
	for _, cname := range table.SplitInfo.SplitCols {
		for _, col := range table.TableColumns {
			if col.Name != cname {
				continue
			}
			var colType string = common.CPPFieldsMapPackSQL_ForTables_COL2SQL_GetCPPType(col.Type)
			f.WriteString("    const ")
			f.WriteString(colType)
			f.WriteString(" ")
			if colType == "string" {
				f.WriteString("&")
			}
			f.WriteString(table.Name + "_" + col.Name)
			f.WriteString(" = table_" + table.Name + "." + col.Name + "();\n")

			var table_col string = table.Name + "_" + col.Name
			if colType == "string" {
				f.WriteString("    const char *sz_" + table_col + " = \"\";\n")
				f.WriteString("    int len_" + table_col + " = 0;\n")
				f.WriteString("    GORM_MemPoolData *buffer_" + table_col + " = nullptr;\n")
				f.WriteString("    if(" + table_col + ".size() > 0)\n")
				f.WriteString("    {\n")

				var bufferName string = "buffer_" + table_col
				var bufferSize string = table_col + ".size()<<1"
				PrintGetBuffFromMemPool(f, bufferName, bufferSize)

				f.WriteString("        iTmpLen=mysql_real_escape_string(mysql, buffer_" + table_col + "->m_uszData, " + table_col + ".c_str(), " + table_col + ".size());\n")
				f.WriteString("        buffer_" + table_col + "->m_uszData[iTmpLen] = 0;\n")
				f.WriteString("        buffer_" + table_col + "->m_sUsedSize = iTmpLen;\n")
				f.WriteString("        sz_" + table_col + " = buffer_" + table_col + "->m_uszData;\n")
				f.WriteString("        len_" + table_col + " = iTmpLen;\n")
				f.WriteString("    }\n")
			}

			if colType == "string" {
				GORM_SafeSnprintfstr += ", sz_" + table_col
				releasestr += "    if(buffer_" + table_col + " != nullptr)\n"
				releasestr += "        buffer_" + table_col + "->Release();\n"
			} else {
				GORM_SafeSnprintfstr += ", " + table_col
			}
		}
	}
	GORM_SafeSnprintfstr += ");\n"
	f.WriteString(GORM_SafeSnprintfstr)
	f.WriteString("    pReqData->m_sUsedSize = iLen;\n\n")
	f.WriteString(releasestr)

	f.WriteString("\n    return GORM_OK;\n")
	f.WriteString("}\n")
	return 0
}

func CPPFieldsMapPackGetSQL_ForTables_CheckSplitInfo(table common.TableInfo, f *os.File) int {
	var splitInfo common.SplitInfo = table.SplitInfo
	for _, cname := range splitInfo.SplitCols {
		f.WriteString("    bMatch = false;\n")
		f.WriteString("    for(int i=0; i<vFields.size(); i++)\n")
		f.WriteString("    {\n")
		f.WriteString("        if (GORM_PB_FIELD_" + strings.ToUpper(table.Name) + "_" + strings.ToUpper(cname) + " == vFields[i])\n")
		f.WriteString("        {\n")
		f.WriteString("            bMatch = true;\n")
		f.WriteString("            break;\n")
		f.WriteString("        }\n")
		f.WriteString("    }\n")
		f.WriteString("    if (!bMatch)\n")
		f.WriteString("        return GORM_REQ_NEED_SPLIT;\n\n")
	}
	return 0
}

func CPPFieldsMapPackGetSQL_ForTables(games []common.XmlCfg, f *os.File) int {
	for _, game := range games {
		for _, table := range game.DB.TableList {
			var DefineSQL string
			var iRet int
			DefineSQL, iRet = CPPFieldsMapPackGetSQL_ForTables_DefineSQL(table)
			if 0 != iRet {
				return -1
			}
			f.WriteString(DefineSQL)
			f.WriteString("\n")
			var DefineLen int = len(DefineSQL)
			var definesqllen string = strconv.FormatInt(int64(DefineLen), 10)
			if 0 != CPPFieldsMapPackGetSQL_ForTables_One(table, definesqllen, f) {
				return -1
			}

			f.WriteString("int GORM_PackGetSQL")
			f.WriteString(strings.ToUpper(table.Name))
			f.WriteString("(shared_ptr<GORM_MemPool> &pMemPool, GORM_MySQLEvent *pMySQLEvent, MYSQL* mysql, int iTableIndex, const GORM_PB_GET_REQ* pMsg, GORM_MemPoolData *&pReqData)\n")
			f.WriteString("{\n")
			f.WriteString("    if (!pMsg->has_header())\n")
			f.WriteString("        return GORM_REQ_MSG_NO_HEADER;\n")
			f.WriteString("    if (!pMsg->has_table())\n")
			f.WriteString("        return GORM_REQ_NO_RECORDS;\n\n")
			f.WriteString("    GORM_PB_TABLE table = pMsg->table();\n")
			f.WriteString("    if (!table.has_" + table.Name + "())\n")
			f.WriteString("        return GORM_REQ_NO_RECORDS;\n\n")
			f.WriteString("    const GORM_PB_REQ_HEADER &header = pMsg->header();\n")
			f.WriteString("    string fieldMode = header.fieldmode();\n")
			f.WriteString("    if (fieldMode == \"\")\n")
			f.WriteString("        return GORM_REQ_NO_RECORDS;\n")
			f.WriteString("    vector<int> vFields = GORM_FieldsOpt::GetFields(fieldMode.c_str(), fieldMode.size());\n\n")
			f.WriteString("    bool bMatch = false;\n")
			f.WriteString("    const GORM_PB_Table_" + table.Name + " &table_" + table.Name + " = table." + table.Name + "();\n\n")

			if 0 != CPPFieldsMapPackGetSQL_ForTables_CheckSplitInfo(table, f) {
				return -1
			}

			f.WriteString("    \n")
			f.WriteString("    return GORM_PackGetSQL" + strings.ToUpper(table.Name) + "_ONE(pMemPool, mysql, iTableIndex, table_" + table.Name + ", pReqData);\n")
			f.WriteString("}\n")
		}
	}
	return 0
}

func CPPFieldsMapPackGetSQL(games []common.XmlCfg, f *os.File) int {
	CPPFieldsMapPackGetSQL_ForTables(games, f)
	CPPFields_GORM_PackSQL_TEMPLATE("Get", games, f)
	return 0
}
