package mysql

import (
	"fmt"
	"gorm-conv/common"
	"os"
	"strconv"
	"strings"
)

func CPPFieldsMapPackIncreaseSQL_ForTables_DefineSQL(table common.TableInfo) (string, string, int) {
	var DefineSQL string = "#define " + strings.ToUpper(table.Name) + "INCREASESQL \"update "
	DefineSQL += table.Name
	DefineSQL += "_%d set "
	var WhereSQL string = "#define " + strings.ToUpper(table.Name) + "INCREASEWHERESQL \" where"
	var splitInfo common.SplitInfo = table.SplitInfo
	if splitInfo.Columns == "" {
		return "", "", -1
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
				WhereSQL += " and"
			}

			WhereSQL += " `"
			WhereSQL += preCol.Name + "`="
			WhereSQL += common.CPPFieldPackSQL_COL_FORMAT(preCol.Type)
		}
		if !match {
			fmt.Println("invalid splitinfo, table:", table.Name)
			return "", "", -1
		}
	}

	DefineSQL += "\""
	WhereSQL += "\""

	return DefineSQL, WhereSQL, 0
}

func CPPFieldsMapPackIncreaseSQL_ForTables_COL2SQL_FORVARIABLE(table common.TableInfo, col common.TableColumn, f *os.File) int {
	// 如果不在splitfields，并且是字符串类型则直接跳过
	var vtype string = common.CPPFieldsMapPackSQL_ForTables_COL2SQL_GetCPPType(col.Type)
	if vtype == "string" {
		var bmatch bool = false
		for _, c := range table.SplitInfo.SplitCols {
			if c == col.Name {
				bmatch = true
				break
			}
		}
		if !bmatch {
			return 0
		}
	}

	if vtype != "string" {
		f.WriteString("    const " + vtype + " " + table.Name + "_" + col.Name + " = table_" + table.Name + "." + col.Name + "();\n")
		return 0
	}

	var table_col string = table.Name + "_" + col.Name
	f.WriteString("    const string &" + table_col + " = table_" + table.Name + "." + col.Name + "();\n")
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
	f.WriteString("    }\n\n")
	return 0
}

func CPPFieldsMapPackIncreaseSQL_ForTables_COL2SQL(table common.TableInfo, f *os.File) int {
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

	for _, col := range table.TableColumns {
		ret := CPPFieldsMapPackIncreaseSQL_ForTables_COL2SQL_FORVARIABLE(table, col, f)
		if ret != 0 {
			fmt.Println("CPPFieldsMapPackIncreaseSQL_ForTables_COL2SQL_FORVARIABLE failed, for ", table.Name, ".", col.Name)
			return -1
		}
	}

	return 0
}

func CPPFieldsMapPackIncreaseSQL_ForTables_WhereSQL(table common.TableInfo, f *os.File) int {
	f.WriteString("\n")
	f.WriteString("    int iWhereLen = iSqlLen + 128 ")
	var vtype string
	var intLen int64 = 0
	for _, col := range table.TableColumns {
		vtype = common.CPPFieldsMapPackSQL_ForTables_COL2SQL_GetCPPType(col.Type)
		if vtype == "string" {
			var bmatch bool = false
			for _, c := range table.SplitInfo.SplitCols {
				if c == col.Name {
					bmatch = true
					break
				}
			}
			if !bmatch {
				continue
			}
			f.WriteString(" + len_" + table.Name + "_" + col.Name)
		} else {
			intLen += 8
		}
	}
	f.WriteString(" + " + strconv.FormatInt(intLen, 10))
	f.WriteString(";\n")
	var bufferName string = "buffer_" + table.Name + "_where"
	f.WriteString("    GORM_MemPoolData *" + bufferName + " = nullptr;\n")

	var bufferSize string = "iWhereLen"
	PrintGetBuffFromMemPool(f, bufferName, bufferSize)

	f.WriteString("    char *szWhereBegin = buffer_" + table.Name + "_where->m_uszData;\n")
	f.WriteString("    iWhereLen += GORM_SafeSnprintf(szWhereBegin+iWhereLen, buffer_" + table.Name + "_where->m_sCapacity, " + strings.ToUpper(table.Name) + "INCREASEWHERESQL ")
	for _, colname := range table.SplitInfo.SplitCols {
		for _, preCol := range table.TableColumns {
			if colname != preCol.Name {
				continue
			}
			vtype = common.CPPFieldsMapPackSQL_ForTables_COL2SQL_GetCPPType(preCol.Type)
			if vtype == "string" {
				f.WriteString(", sz_" + table.Name + "_" + preCol.Name)
			} else {
				f.WriteString(", " + table.Name + "_" + preCol.Name)
			}
		}
	}
	f.WriteString(");\n")
	f.WriteString("    iWhereLen += GORM_GETVERSION_WHERE(szWhereBegin+iWhereLen, buffer_" + table.Name + "_where->m_sCapacity-iWhereLen, GORM_CheckDataVerType(header.verpolice()), " + table.Name + "_version);\n")
	f.WriteString("    buffer_" + table.Name + "_where->m_sUsedSize = iWhereLen;\n")
	f.WriteString("\n")
	return 0
}

func CPPFieldsMapPackIncreaseSQL_ForTables_SetSQL(table common.TableInfo, f *os.File) int {
	var upTableName string = strings.ToUpper(table.Name)
	f.WriteString("    int iLen = iSqlLen + 128 + pMsg->ByteSizeLong() ")
	var intLen int64 = 0
	for _, col := range table.TableColumns {
		var vtype string = common.CPPFieldsMapPackSQL_ForTables_COL2SQL_GetCPPType(col.Type)
		if vtype == "string" {
			var bmatch bool = false
			for _, c := range table.SplitInfo.SplitCols {
				if c == col.Name {
					bmatch = true
					break
				}
			}
			intLen += int64(len(col.Name)) + 10
			if !bmatch {
				continue
			}
			f.WriteString("+ len_" + table.Name + "_" + col.Name)
		} else {
			intLen += 8
		}
	}
	f.WriteString("+")
	f.WriteString(strconv.FormatInt(intLen, 10))
	f.WriteString(";\n")

	var bufferName string = "pReqData"
	var bufferSize string = "iLen+iWhereLen+1"
	PrintGetBuffFromMemPool(f, bufferName, bufferSize)

	f.WriteString("    szSQLBegin = pReqData->m_uszData;\n")
	f.WriteString("    memcpy(szSQLBegin, " + upTableName + "UPDATESQL, iSqlLen);\n")
	f.WriteString("    int iDataLen = GORM_SafeSnprintf(szSQLBegin, iLen, " + upTableName + "INCREASESQL, iTableIndex);\n")
	f.WriteString("    szSQLBegin += iDataLen;\n")
	f.WriteString("    pReqData->m_sUsedSize = iDataLen;\n")
	f.WriteString("    int iSetField = 1;\n")
	f.WriteString("    iDataLen  = GORM_GETVERSION_SET(szSQLBegin, iLen-pReqData->m_sUsedSize, GORM_CheckDataVerType(header.verpolice()), " + table.Name + "_version);\n")
	f.WriteString("    szSQLBegin += iDataLen;\n")
	f.WriteString("    pReqData->m_sUsedSize += iDataLen;\n")
	f.WriteString("    iLen -= iDataLen;\n")
	f.WriteString("    iDataLen = 0;\n")
	f.WriteString("    for (int i=0; i<vFields.size(); i++)\n")
	f.WriteString("    {\n")
	f.WriteString("        int iFieldId = vFields[i];\n")
	f.WriteString("        if (")

	f.WriteString("GORM_PB_FIELD_" + upTableName + "_VERSION == iFieldId")
	for _, colname := range table.SplitInfo.SplitCols {
		f.WriteString(" || ")
		f.WriteString("GORM_PB_FIELD_" + upTableName + "_" + strings.ToUpper(colname) + " == iFieldId")
	}
	f.WriteString(")\n")
	f.WriteString("            continue;\n")
	f.WriteString("        iSetField += 1;\n")
	f.WriteString("        switch (iFieldId)\n")
	f.WriteString("        {\n")
	for _, col := range table.TableColumns {
		var vtype string = common.CPPFieldsMapPackSQL_ForTables_COL2SQL_GetCPPType(col.Type)
		if vtype == "string" {
			continue
		}
		f.WriteString("        case GORM_PB_FIELD_" + upTableName + "_" + strings.ToUpper(col.Name) + ":\n")
		f.WriteString("        {\n")
		f.WriteString("            char cOpt = '+';\n")
		f.WriteString("            if (GORM_FieldsOpt::FieldInMode(strMinusFieldsMode.c_str(), strMinusFieldsMode.size(), GORM_PB_FIELD_")
		f.WriteString(strings.ToUpper(table.Name) + "_" + strings.ToUpper(col.Name) + "))\n")
		f.WriteString("                cOpt = '-';\n")
		f.WriteString("            if (iSetField != 1)\n")
		f.WriteString("                iDataLen = GORM_SafeSnprintf(szSQLBegin, iLen, \", `" + col.Name + "`=`" + col.Name + "`%c")
		f.WriteString(common.CPPFieldPackSQL_COL_FORMAT(col.Type))
		f.WriteString("\", cOpt, ")
		if vtype == "string" {
			f.WriteString("sz_")
		}
		f.WriteString(table.Name + "_" + col.Name)
		f.WriteString(");\n")
		f.WriteString("            else\n")
		f.WriteString("                iDataLen = GORM_SafeSnprintf(szSQLBegin, iLen, \" `" + col.Name + "`=`" + col.Name + "`%c")
		f.WriteString(common.CPPFieldPackSQL_COL_FORMAT(col.Type))
		f.WriteString("\", cOpt, ")
		if vtype == "string" {
			f.WriteString("sz_")
		}
		f.WriteString(table.Name + "_" + col.Name)
		f.WriteString(");\n")
		f.WriteString("            iLen -= iDataLen;\n")
		f.WriteString("            szSQLBegin += iDataLen;\n")
		f.WriteString("            break;\n")
		f.WriteString("        }\n")
	}
	f.WriteString("        default:\n")
	f.WriteString("            continue;\n")
	f.WriteString("        }\n")
	f.WriteString("        pReqData->m_sUsedSize += iDataLen;\n")
	f.WriteString("        if (iLen <= 0)\n")
	f.WriteString("            break;\n")
	f.WriteString("    }\n")
	f.WriteString("    memcpy(pReqData->m_uszData+pReqData->m_sUsedSize, szWhereBegin, iWhereLen);\n")
	f.WriteString("    pReqData->m_sUsedSize += iWhereLen;\n")
	f.WriteString("    pReqData->m_uszData[pReqData->m_sUsedSize] = ';';\n")
	f.WriteString("    pReqData->m_sUsedSize += 1;\n")

	f.WriteString("\n")
	return 0
}

func CPPFieldsMapPackIncreaseSQL_ForTables(games []common.XmlCfg, f *os.File) int {
	for _, game := range games {
		for _, table := range game.DB.TableList {
			var bigTable string = strings.ToUpper(table.Name)
			var DefineSQL, WhereSQL string
			var iRet int
			DefineSQL, WhereSQL, iRet = CPPFieldsMapPackIncreaseSQL_ForTables_DefineSQL(table)
			if 0 != iRet {
				return -1
			}
			f.WriteString(DefineSQL)
			f.WriteString("\n")
			f.WriteString(WhereSQL)
			f.WriteString("\n")

			f.WriteString("int GORM_PackIncreaseSQL")
			f.WriteString(bigTable)
			f.WriteString("(shared_ptr<GORM_MemPool> &pMemPool, GORM_MySQLEvent *pMySQLEvent, MYSQL* mysql, int iTableIndex, const GORM_PB_INCREASE_REQ* pMsg, GORM_MemPoolData *&pReqData)\n")
			f.WriteString("{\n")
			f.WriteString("    if (!pMsg->has_header())\n")
			f.WriteString("        return GORM_REQ_MSG_NO_HEADER;\n")
			f.WriteString("    if (pMsg->tables_size() == 0)\n")
			f.WriteString("        return GORM_REQ_NO_RECORDS;\n\n")
			f.WriteString("    GORM_PB_TABLE table = pMsg->tables(0);\n")
			f.WriteString("    if (!table.has_" + table.Name + "())\n")
			f.WriteString("        return GORM_REQ_NO_RECORDS;\n\n")
			f.WriteString("    const GORM_PB_REQ_HEADER &header = pMsg->header();\n")
			f.WriteString("    const string &fieldMode = header.fieldmode();\n")
			f.WriteString("    if (fieldMode == \"\")\n")
			f.WriteString("        return GORM_REQ_NO_RECORDS;\n")
			f.WriteString("    vector<int> vFields = GORM_FieldsOpt::GetFields(fieldMode.c_str(), fieldMode.size());\n\n")
			f.WriteString("    const string &strMinusFieldsMode = pMsg->minuscolumns();\n")
			f.WriteString("    bool bMatch = false;\n")
			f.WriteString("    char *szSQLBegin = nullptr;\n")
			f.WriteString("    int iSqlLen = strlen(" + bigTable + "INCREASESQL);\n")
			//var DefineLen int = len(DefineSQL)
			//f.WriteString(strconv.FormatInt(int64(DefineLen), 10))
			//f.WriteString(";\n")
			f.WriteString("    int iTmpLen = 0;\n\n")

			f.WriteString("    GORM_PB_Table_" + table.Name + " table_" + table.Name + " = table." + table.Name + "();\n\n")

			if table.SplitInfo.Columns == "" {
				return -1
			} else {
				if 0 != CPPFieldsMapPackIncreaseSQL_ForTables_COL2SQL(table, f) {
					return -1
				}
			}
			if 0 != CPPFieldsMapPackIncreaseSQL_ForTables_WhereSQL(table, f) {
				return -1
			}
			if 0 != CPPFieldsMapPackIncreaseSQL_ForTables_SetSQL(table, f) {
				return -1
			}

			// 释放buffer
			for _, col := range table.TableColumns {
				var vtype string = common.CPPFieldsMapPackSQL_ForTables_COL2SQL_GetCPPType(col.Type)
				if vtype != "string" {
					continue
				}
				var bContinue = false
				for _, cname := range table.SplitInfo.SplitCols {
					if cname == col.Name {
						bContinue = true
						break
					}
				}
				if !bContinue {
					continue
				}
				f.WriteString("    if (buffer_" + table.Name + "_" + col.Name + " != nullptr)\n")
				f.WriteString("        buffer_" + table.Name + "_" + col.Name + "->Release();\n")
			}

			f.WriteString("    \n")
			/*
				f.WriteString("#ifdef GORM_DEBUG\n")
				f.WriteString("        GORM_MySQLUpdateTableSchema(pMySQLEvent, \"" + table.Name + "\", table.custom_columns());\n")
				f.WriteString("#endif\n")
			*/
			f.WriteString("    return GORM_OK;\n")
			f.WriteString("}\n")
		}
	}
	return 0
}

func CPPFieldsMapPackIncreaseSQL(games []common.XmlCfg, f *os.File) int {
	CPPFieldsMapPackIncreaseSQL_ForTables(games, f)
	CPPFields_GORM_PackSQL_TEMPLATE("Increase", games, f)
	return 0
}
