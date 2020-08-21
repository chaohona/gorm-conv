package mysql

import (
	"fmt"
	"gorm-conv/common"
	"gorm-conv/cpp"
	"os"
	"strconv"
	"strings"
)

func CPPFieldsMapPackDeleteSQL_ForTables_DefineSQL(table common.TableInfo) (string, int) {
	var DefineSQL string = "#define " + strings.ToUpper(table.Name) + "DELETESQL \"delete  from "
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
				DefineSQL += " and "
			}
			DefineSQL += " `"
			DefineSQL += preCol.Name + "`="
			DefineSQL += cpp.CPPFieldPackSQL_COL_FORMAT(preCol.Type)
		}
		if !match {
			fmt.Println("invalid splitinfo, table:", table.Name)
			return "", -1
		}
	}

	DefineSQL += ";\""
	return DefineSQL, 0
}

// 没有配置split信息，不分表语句的查询
func CPPFieldsMapPackDeleteSQL_ForTables_COL2SQL_NoSplit(table common.TableInfo, f *os.File) int {
	var ilenstr string = "    int iLen = iSqlLen + 256 + 13*" + strconv.FormatInt(int64(len(table.TableColumns)), 10) + " + pMsg->ByteSizeLong();\n"
	f.WriteString(ilenstr)
	f.WriteString("    pReqData = GORM_MemPool::Instance()->GetData(iLen);\n")
	f.WriteString("    szSQLBegin = pReqData->m_uszData;\n")
	f.WriteString("    memcpy(szSQLBegin, " + strings.ToUpper(table.Name) + "DELETESQL, iSqlLen);\n\n")

	f.WriteString("    int iNowLen = 0;\n")
	f.WriteString("    for(int i=0; i<vFields.size(); i++)\n")
	f.WriteString("    {\n")
	f.WriteString("        switch (vFields[i])\n")
	f.WriteString("        {\n")
	for _, col := range table.TableColumns {
		f.WriteString("        case GORM_PB_FIELD_" + strings.ToUpper(table.Name) + "_" + strings.ToUpper(col.Name) + ":\n")
		f.WriteString("        {\n")
		f.WriteString("            if(i==0)\n")
		var format string = cpp.CPPFieldPackSQL_COL_FORMAT(col.Type)
		var snprintfstr string = col.Name + "=" + format + "\", table_" + table.Name + "." + col.Name + "()"
		if cpp.CPPFieldsMapPackSQL_ForTables_COL2SQL_GetCPPType(col.Type) == "string" {
			snprintfstr += ".c_str()"
		}
		snprintfstr += ");\n"
		f.WriteString("                 iNowLen = snprintf(szSQLBegin+iSqlLen, iLen-iSqlLen, \" " + snprintfstr)
		f.WriteString("            else\n")
		f.WriteString("                iNowLen = snprintf(szSQLBegin+iSqlLen, iLen-iSqlLen, \" and " + snprintfstr)
		f.WriteString("            break;\n")
		f.WriteString("        }\n")
	}
	f.WriteString("        default:\n")
	f.WriteString("            return GORM_INVALID_FIELD;\n")
	f.WriteString("        }\n\n")
	f.WriteString("        iSqlLen += iNowLen;\n")
	f.WriteString("        if (iSqlLen >= iLen)\n")
	f.WriteString("            return GORM_REQ_TOO_LARGE;\n")
	f.WriteString("    }\n")

	f.WriteString("    szSQLBegin[iSqlLen] = ';';\n")
	return 0
}

func CPPFieldsMapPackDeleteSQL_ForTables_COL2SQL(table common.TableInfo, f *os.File) int {
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

	var ilenstr string = "    int iLen = iSqlLen + 128"
	var snprintfstr string = "    iLen = snprintf(szSQLBegin, iLen, " + strings.ToUpper(table.Name) + "DELETESQL, iTableIndex"
	var releasestr string = ""
	for _, cname := range splitInfo.SplitCols {
		for _, col := range table.TableColumns {
			if col.Name != cname {
				continue
			}
			var colType string = cpp.CPPFieldsMapPackSQL_ForTables_COL2SQL_GetCPPType(col.Type)
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
				f.WriteString("    char *sz_" + table_col + " = \"\";\n")
				f.WriteString("    int len_" + table_col + " = 0;\n")
				f.WriteString("    GORM_MemPoolData *buffer_" + table_col + " = nullptr;\n")
				f.WriteString("    if(" + table_col + ".size() > 0)\n")
				f.WriteString("    {\n")
				f.WriteString("        buffer_" + table_col + " = GORM_MemPool::Instance()->GetData(" + table_col + ".size()<<1);\n")
				f.WriteString("        iTmpLen=mysql_real_escape_string(mysql, buffer_" + table_col + "->m_uszData, " + table_col + ".c_str(), " + table_col + ".size());\n")
				f.WriteString("        buffer_" + table_col + "->m_uszData[iTmpLen] = 0;\n")
				f.WriteString("        buffer_" + table_col + "->m_sUsedSize = iTmpLen;\n")
				f.WriteString("        sz_" + table_col + " = buffer_" + table_col + "->m_uszData;\n")
				f.WriteString("        len_" + table_col + " = iTmpLen;\n")
				f.WriteString("    }\n")
			}

			if colType == "string" {
				ilenstr += " + len_" + table_col
				snprintfstr += ", sz_" + table_col
				releasestr += "    if(buffer_" + table_col + " != nullptr)\n"
				releasestr += "        buffer_" + table_col + "->Release();\n"
			} else {
				ilenstr += " + 8"
				snprintfstr += ", " + table_col
			}
		}
	}
	ilenstr += " + pMsg->ByteSizeLong();\n"
	snprintfstr += ");\n"
	f.WriteString(ilenstr)
	f.WriteString("    pReqData = GORM_MemPool::Instance()->GetData(iLen);\n")
	f.WriteString("    szSQLBegin = pReqData->m_uszData;\n")
	f.WriteString(snprintfstr)
	f.WriteString("    pReqData->m_sUsedSize = iLen;\n\n")
	f.WriteString(releasestr)
	return 0
}

func CPPFieldsMapPackDeleteSQL_ForTables(games []common.XmlCfg, f *os.File) int {
	for _, game := range games {
		for _, table := range game.DB.TableList {
			var DefineSQL string
			var iRet int
			DefineSQL, iRet = CPPFieldsMapPackDeleteSQL_ForTables_DefineSQL(table)
			if 0 != iRet {
				return -1
			}
			f.WriteString(DefineSQL)
			f.WriteString("\n")

			f.WriteString("int GORM_PackDeleteSQL")
			f.WriteString(strings.ToUpper(table.Name))
			f.WriteString("(GORM_MySQLEvent *pMySQLEvent, MYSQL* mysql, int iTableIndex, const GORM_PB_DELETE_REQ* pMsg, GORM_MemPoolData *&pReqData)\n")
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
			f.WriteString("    char *szSQLBegin = nullptr;\n")
			f.WriteString("    int iSqlLen = ")
			var DefineLen int = len(DefineSQL)
			f.WriteString(strconv.FormatInt(int64(DefineLen), 10))
			f.WriteString(";\n")
			f.WriteString("    int iTmpLen = 0;\n\n")

			f.WriteString("    GORM_PB_Table_" + table.Name + " table_" + table.Name + " = table." + table.Name + "();\n\n")

			if table.SplitInfo.Columns == "" {
				if 0 != CPPFieldsMapPackDeleteSQL_ForTables_COL2SQL_NoSplit(table, f) {
					return -1
				}
			} else {
				if 0 != CPPFieldsMapPackDeleteSQL_ForTables_COL2SQL(table, f) {
					return -1
				}
			}

			f.WriteString("    \n")

			f.WriteString("#ifdef GORM_DEBUG\n")
			f.WriteString("        GORM_MySQLUpdateTableSchema(pMySQLEvent, \"" + table.Name + "\", table.custom_columns());\n")
			f.WriteString("#endif\n")
			f.WriteString("    return GORM_OK;\n")
			f.WriteString("}\n")
		}
	}
	return 0
}

func CPPFieldsMapPackDeleteSQL(games []common.XmlCfg, f *os.File) int {
	CPPFieldsMapPackDeleteSQL_ForTables(games, f)
	CPPFields_GORM_PackSQL_TEMPLATE("Delete", games, f)
	return 0
}
