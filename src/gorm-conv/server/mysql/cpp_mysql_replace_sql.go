package mysql

import (
	"gorm-conv/common"
	"os"
	"strconv"
	"strings"
)

func CPPFieldsMapPackReplaceSQL_ForTables_DefineSQL(table common.TableInfo) string {
	var DefineSQL string = "#define " + strings.ToUpper(table.Name) + "REPLACESQL \"insert into "
	DefineSQL += table.Name
	DefineSQL += "_%d("

	for idx, col := range table.TableColumns {
		if idx != 0 {
			DefineSQL += ", "
		}
		DefineSQL += "`"
		DefineSQL += col.Name
		DefineSQL += "`"
	}

	DefineSQL += ") values ("
	for idx, col := range table.TableColumns {
		if idx != 0 {
			DefineSQL += ", "
		}
		DefineSQL += common.CPPFieldPackSQL_COL_FORMAT(col.Type)
	}
	DefineSQL += ") ON DUPLICATE KEY UPDATE `version`=`version`+1"
	for _, col := range table.TableColumns {
		if col.Name == "version" {
			continue
		}
		var match bool = false
		for _, cname := range table.SplitInfo.SplitCols {
			if col.Name == cname {
				match = true
				break
			}
		}
		if match {
			continue
		}
		DefineSQL += ","
		DefineSQL += " `"
		DefineSQL += col.Name
		DefineSQL += "`="
		DefineSQL += common.CPPFieldPackSQL_COL_FORMAT(col.Type)
	}
	DefineSQL += ";\"\n"
	return DefineSQL
}

func CPPFieldsMapPackReplaceSQL_ForTables_COL2SQL(table common.TableInfo, f *os.File) int {
	var len_str string = "    int iLen = iSqlLen + 128"
	var sprintf_str string = "    iLen = GORM_SafeSnprintf(szSQLBegin, iLen, " + strings.ToUpper(table.Name) + "REPLACESQL, iTableIndex"
	var release_str string = ""
	for _, col := range table.TableColumns {
		var col_type_name string = table.Name + "_" + col.Name
		f.WriteString("\n")
		var cppType string = common.CPPFieldsMapPackSQL_ForTables_COL2SQL_GetCPPType(col.Type)
		f.WriteString("    const " + cppType + " ")
		if cppType == "string" {
			f.WriteString("&")
		}
		f.WriteString(col_type_name + " = " + "table_" + table.Name + "." + col.Name + "();\n")
		if cppType == "string" {
			var mempoolstr string = "buffer_" + col_type_name
			f.WriteString("    const char *sz_" + col_type_name + " = \"\";\n")
			f.WriteString("    int len_" + col_type_name + " = 0;\n")
			f.WriteString("    GORM_MemPoolData *" + mempoolstr + " = nullptr;\n")
			f.WriteString("    if(" + col_type_name + ".size()>0)\n")
			f.WriteString("    {\n")

			var bufferName string = mempoolstr
			var bufferSize string = col_type_name + ".size()<<1"
			PrintGetBuffFromMemPool(f, bufferName, bufferSize)

			f.WriteString("        iTmpLen = mysql_real_escape_string(mysql, buffer_" + col_type_name)
			f.WriteString("->m_uszData, " + col_type_name + ".c_str(), " + col_type_name + ".size());\n")

			f.WriteString("        buffer_" + col_type_name + "->m_uszData[iTmpLen] = '\\0';\n")
			f.WriteString("        buffer_" + col_type_name + "->m_sUsedSize = iTmpLen;\n")
			f.WriteString("        sz_" + col_type_name + " = " + "buffer_" + col_type_name + "->m_uszData;\n")
			f.WriteString("        len_" + col_type_name + " = iTmpLen;\n")
			f.WriteString("    }\n")
		}

		if cppType == "string" {
			len_str += " + len_" + col_type_name
			sprintf_str += ", sz_" + col_type_name
			release_str += "    if (buffer_" + col_type_name + " != nullptr)\n"
			release_str += "        buffer_" + col_type_name + "->Release();\n"
		} else {
			len_str += " + 8"
			sprintf_str += ", " + col_type_name
		}
	}
	for _, col := range table.TableColumns {
		if col.Name == "version" {
			continue
		}
		var match bool = false
		for _, cname := range table.SplitInfo.SplitCols {
			if col.Name == cname {
				match = true
				break
			}
		}
		if match {
			continue
		}
		sprintf_str += ","
		var cppType string = common.CPPFieldsMapPackSQL_ForTables_COL2SQL_GetCPPType(col.Type)
		var col_type_name string = table.Name + "_" + col.Name
		if cppType == "string" {
			sprintf_str += " sz_" + col_type_name
		} else {
			sprintf_str += col_type_name
		}
	}
	len_str += ";\n"
	sprintf_str += ");\n"
	f.WriteString("\n")
	f.WriteString(len_str)
	f.WriteString("iLen *= 2;\n")

	var bufferName string = "pReqData"
	var bufferSize string = "iLen<<1"
	PrintGetBuffFromMemPool(f, bufferName, bufferSize)

	f.WriteString("    szSQLBegin = pReqData->m_uszData;\n")

	f.WriteString(sprintf_str)
	f.WriteString("    pReqData->m_sUsedSize = iLen;\n")
	f.WriteString("\n")
	f.WriteString(release_str)
	return 0
}

func CPPFieldsMapPackReplaceSQL_ForTables_One(table common.TableInfo, sqlLen string, f *os.File) int {
	f.WriteString("int GORM_PackReplaceSQL")
	f.WriteString(strings.ToUpper(table.Name))
	f.WriteString("_One(shared_ptr<GORM_MemPool> &pMemPool, MYSQL* mysql, int iTableIndex, const GORM_PB_Table_" + table.Name + " &table_" + table.Name + ", GORM_MemPoolData *&pReqData)\n")
	f.WriteString("{\n")
	f.WriteString("    char *szSQLBegin = nullptr;\n")
	f.WriteString("    int iSqlLen = " + sqlLen + ";\n")
	f.WriteString("    int iTmpLen = 0;\n")

	if 0 != CPPFieldsMapPackReplaceSQL_ForTables_COL2SQL(table, f) {
		return -1
	}
	f.WriteString("\n")
	f.WriteString("    return GORM_OK;\n")
	f.WriteString("}\n")

	return 0
}

func CPPFieldsMapPackReplaceSQL_ForTables(games []common.XmlCfg, f *os.File) int {
	for _, game := range games {
		for _, table := range game.DB.TableList {
			var DefineSQL string = CPPFieldsMapPackReplaceSQL_ForTables_DefineSQL(table)
			f.WriteString(DefineSQL)
			var DefineLen int = len(DefineSQL)
			// 函数头
			var BigTable string = strings.ToUpper(table.Name)
			if 0 != CPPFieldsMapPackReplaceSQL_ForTables_One(table, strconv.FormatInt(int64(DefineLen), 10), f) {
				return -1
			}
			f.WriteString("int GORM_PackReplaceSQL")
			f.WriteString(BigTable)
			f.WriteString("(shared_ptr<GORM_MemPool> &pMemPool, GORM_MySQLEvent *pMySQLEvent, MYSQL* mysql, int iTableIndex, const GORM_PB_REPLACE_REQ* pMsg, GORM_MemPoolData *&pReqData)\n")
			f.WriteString("{\n")
			f.WriteString("    int iTableNum = pMsg->tables_size();\n")
			f.WriteString("    if (iTableNum == 0)\n")
			f.WriteString("        return GORM_REQ_NO_RECORDS;\n")
			f.WriteString("    for (int i=0; i<iTableNum; i++)\n")
			f.WriteString("    {\n")
			f.WriteString("        const GORM_PB_TABLE &table = pMsg->tables(i);\n")
			f.WriteString("")
			f.WriteString("        if (!table.has_" + table.Name + "())\n")
			f.WriteString("            return GORM_REQ_NO_RECORDS;\n")
			f.WriteString("        const GORM_PB_Table_" + table.Name + " &table_" + table.Name + " = table." + table.Name + "();\n")

			/*
				f.WriteString("#ifdef GORM_DEBUG\n")
				f.WriteString("        GORM_MySQLUpdateTableSchema(pMySQLEvent, \"" + table.Name + "\", table.custom_columns());\n")
				f.WriteString("#endif\n")
			*/
			f.WriteString("        return GORM_PackReplaceSQL" + BigTable + "_One(pMemPool, mysql, iTableIndex, table_" + table.Name + ", pReqData);\n")
			f.WriteString("    }\n")
			f.WriteString("    return GORM_OK;\n")
			f.WriteString("}\n")
		}
	}
	return 0
}

func CPPFieldsMapPackReplaceSQL(games []common.XmlCfg, f *os.File) int {
	CPPFieldsMapPackReplaceSQL_ForTables(games, f)
	CPPFields_GORM_PackSQL_TEMPLATE("Replace", games, f)
	return 0
}
