package main

import (
	"os"
	"strconv"
	"strings"
)

func CPPFieldsMapPackGetByNonPrimaryKeySQL_ForTables_One_GetDefineSql(table TableInfo, bDebug bool) string {
	var defineSQL string = "#define GetByNonPrimaySQL_" + strings.ToUpper(table.Name)
	if bDebug {
		defineSQL += "_DEBUG"
	}
	defineSQL += " \"select "
	for idx, col := range table.TableColumns {
		if idx != 0 {
			defineSQL += ","
		}
		defineSQL += "`"
		defineSQL += col.Name
		defineSQL += "`"
	}
	if !bDebug {
		defineSQL += " from " + table.Name + " where \"\n"
	} else {
		defineSQL += "\"\n"
	}

	return defineSQL
}

func CPPFieldsMapPackGetByNonPrimaryKeySQL_ForTables_One_COL(table TableInfo, f *os.File) int {
	var BigTable string = strings.ToUpper(table.Name)
	for _, col := range table.TableColumns {
		var bigCol string = strings.ToUpper(col.Name)
		f.WriteString("        case GORM_PB_FIELD_" + BigTable + "_" + bigCol + ":\n")
		f.WriteString("        {\n")
		var colType string = CPPField_CPPType(col.Type)
		if colType == "string" {
			f.WriteString("            char *szData = \"\";\n")
			f.WriteString("            const string &strData = table_" + table.Name + "." + col.Name + "();\n")
			f.WriteString("            GORM_MemPoolData *pDataBuffer = nullptr;\n")
			f.WriteString("            int iTmpLen = 0;\n")
			f.WriteString("            if (strData.size() > 0)\n")
			f.WriteString("            {\n")
			f.WriteString(`
				pDataBuffer = GORM_MemPool::Instance()->GetData(strData.size()<<1);
                iTmpLen=mysql_real_escape_string(mysql, pDataBuffer->m_uszData, strData.c_str(), strData.size());
                pDataBuffer->m_uszData[iTmpLen] = 0;
                pDataBuffer->m_sUsedSize = iTmpLen;
                szData = pDataBuffer->m_uszData;`)
			f.WriteString("\n            }\n")

			f.WriteString("            if (i==0)\n")
			f.WriteString("                iLen += snprintf(szSQLBegin+iLen, iTotalLen-iLen, \"`" + col.Name + "`=`%s`\", szData);\n")
			f.WriteString("            else\n")
			f.WriteString("                iLen += snprintf(szSQLBegin+iLen, iTotalLen-iLen, \" and `" + col.Name + "`=`%s`\", szData);\n")

			f.WriteString("            if (pDataBuffer != nullptr)\n")
			f.WriteString("                pDataBuffer->Release();\n")
		} else {
			f.WriteString("            if (i==0)\n")
			f.WriteString("                iLen += snprintf(szSQLBegin+iLen, iTotalLen-iLen, \"`" + col.Name + "`=" + CPPFieldPackRedis_COL_FORMAT(col.Type))
			f.WriteString("\", table_" + table.Name + "." + col.Name + "());\n")
			f.WriteString("            else\n")
			f.WriteString("                iLen += snprintf(szSQLBegin+iLen, iTotalLen-iLen, \"and `" + col.Name + "`=" + CPPFieldPackRedis_COL_FORMAT(col.Type))
			f.WriteString("\", table_" + table.Name + "." + col.Name + "());\n")
		}
		f.WriteString("            break;\n")
		f.WriteString("        }\n")
	}

	return 0
}

func CPPFieldsMapPackGetByNonPrimaryKeySQL_ForTables_One(table TableInfo, f *os.File) int {
	var BigTable string = strings.ToUpper(table.Name)
	var defineSQL string = CPPFieldsMapPackGetByNonPrimaryKeySQL_ForTables_One_GetDefineSql(table, false)
	f.WriteString((defineSQL))
	f.WriteString("int GORM_PackGet_By_Non_Primary_KeySQL" + BigTable + "_One(MYSQL* mysql, const GORM_PB_Table_" + table.Name + " &table_" + table.Name + ", const GORM_PB_REQ_HEADER &header, GORM_MemPoolData *&pReqData)\n")
	f.WriteString("{\n")
	f.WriteString(`
	string fieldMode = header.fieldmode();
    if (fieldMode == "")
        return GORM_REQ_NO_RECORDS;
    vector<int> vFields = GORM_FieldsOpt::GetFields(fieldMode.c_str(), fieldMode.size());
`)

	f.WriteString("    int iLen = strlen(GetByNonPrimaySQL_" + BigTable + ");\n")
	f.WriteString("    int iTotalLen = 64*vFields.size() + iLen + table_" + table.Name + ".ByteSizeLong();\n")
	f.WriteString("    pReqData = GORM_MemPool::Instance()->GetData(iTotalLen);\n")
	f.WriteString("    char *szSQLBegin = pReqData->m_uszData;\n")
	f.WriteString("    strncpy(szSQLBegin, " + "GetByNonPrimaySQL_" + BigTable + ", iLen);\n")
	f.WriteString("    for(int i=0; i<vFields.size(); i++)\n")
	f.WriteString("    {\n")
	f.WriteString("        switch (vFields[i])\n")
	f.WriteString("        {\n")
	if 0 != CPPFieldsMapPackGetByNonPrimaryKeySQL_ForTables_One_COL(table, f) {
		return -1
	}
	f.WriteString("        }\n")
	f.WriteString("    }\n")

	f.WriteString("    pReqData->m_sUsedSize = iLen;\n")
	f.WriteString("    return GORM_OK;\n")
	f.WriteString("}\n")
	return 0
}

func CPPFieldsMapPackGetByNonPrimaryKeySQL_ForTables_OneDEBUG(table TableInfo, f *os.File) int {
	var BigTable string = strings.ToUpper(table.Name)
	var defineSQL string = CPPFieldsMapPackGetByNonPrimaryKeySQL_ForTables_One_GetDefineSql(table, true)
	f.WriteString("#ifdef GORM_DEBUG\n")
	f.WriteString((defineSQL))
	f.WriteString("int GORM_PackGet_By_Non_Primary_KeySQL" + BigTable + "_One_DEBUG(GORM_MySQLEvent *pMySQLEvent, const GORM_PB_CUSTEM_COLUMNS &pbColumns, const GORM_PB_Table_" + table.Name + " &table_" + table.Name + ", const GORM_PB_REQ_HEADER &header, GORM_MemPoolData *&pReqData)\n")
	f.WriteString("{\n")
	f.WriteString(`
	MYSQL* mysql = pMySQLEvent->m_pMySQL;
	string fieldMode = header.fieldmode();
    if (fieldMode == "")
        return GORM_REQ_NO_RECORDS;
    vector<int> vFields = GORM_FieldsOpt::GetFields(fieldMode.c_str(), fieldMode.size());
`)

	f.WriteString("    int iLen = strlen(GetByNonPrimaySQL_" + BigTable + "_DEBUG);\n")
	f.WriteString("    int iTotalLen = 64*vFields.size() + pbColumns.ByteSizeLong() + iLen + table_" + table.Name + ".ByteSizeLong();\n")
	f.WriteString("    pReqData = GORM_MemPool::Instance()->GetData(iTotalLen);\n")
	f.WriteString("    char *szSQLBegin = pReqData->m_uszData;\n")
	f.WriteString("    strncpy(szSQLBegin, " + "GetByNonPrimaySQL_" + BigTable + "_DEBUG, iLen);\n")
	var columnLen string = strconv.FormatInt(int64(len(table.TableColumns)), 10)
	f.WriteString("    vector<string> &vColumns = pMySQLEvent->m_mapTablesColumnOrder[\"" + table.Name + "\"];\n")
	f.WriteString("    for (int i=" + columnLen + "; i<vColumns.size(); i++)\n")
	f.WriteString("    {\n")
	f.WriteString("        iLen += snprintf(szSQLBegin+iLen, iTotalLen-iLen, \", `%s`\", vColumns[i].c_str());\n")
	f.WriteString("    }\n")
	f.WriteString("    iLen += snprintf(szSQLBegin+iLen, iTotalLen-iLen, \" from " + table.Name + " where \");\n")
	f.WriteString("    for(int i=0; i<vFields.size(); i++)\n")
	f.WriteString("    {\n")
	f.WriteString("        switch (vFields[i])\n")
	f.WriteString("        {\n")
	if 0 != CPPFieldsMapPackGetByNonPrimaryKeySQL_ForTables_One_COL(table, f) {
		return -1
	}
	f.WriteString("        }\n")
	f.WriteString("    }\n")

	f.WriteString("    pReqData->m_sUsedSize = iLen;\n")
	f.WriteString("    return GORM_OK;\n")
	f.WriteString("}\n")
	f.WriteString("#endif\n")
	return 0
}

func CPPFieldsMapPackGetByNonPrimaryKeySQL_ForTables(games []XmlCfg, f *os.File) int {
	for _, game := range games {
		for _, table := range game.DB.TableList {
			if 0 != CPPFieldsMapPackGetByNonPrimaryKeySQL_ForTables_One(table, f) {
				return -1
			}
			if 0 != CPPFieldsMapPackGetByNonPrimaryKeySQL_ForTables_OneDEBUG(table, f) {
				return -1
			}
			var BigTable string = strings.ToUpper(table.Name)
			f.WriteString("int GORM_PackGet_By_Non_Primary_KeySQL")
			f.WriteString(BigTable)
			f.WriteString("(GORM_MySQLEvent *pMySQLEvent, MYSQL* mysql, const GORM_PB_GET_BY_NON_PRIMARY_KEY_REQ* pMsg, GORM_MemPoolData *&pReqData)\n")
			f.WriteString("{\n")
			f.WriteString(`
	if (!pMsg->has_header())
        return GORM_REQ_MSG_NO_HEADER;
    int iTableNum = pMsg->tables_size();
    if (iTableNum == 0)
        return GORM_REQ_NO_RECORDS;
    for (int i=0; i<iTableNum; i++)
    {
        const GORM_PB_TABLE &table = pMsg->tables(i);
`)
			f.WriteString("        if (!table.has_" + table.Name + "())\n")
			f.WriteString("            return GORM_REQ_NO_RECORDS;\n")
			f.WriteString("        const GORM_PB_Table_" + table.Name + " &table_" + table.Name + " = table." + table.Name + "();\n")
			f.WriteString("#ifdef GORM_DEBUG\n")
			f.WriteString("        GORM_MySQLUpdateTableSchema(pMySQLEvent, \"" + table.Name + "\", table.custom_columns());\n")
			f.WriteString("        return GORM_PackGet_By_Non_Primary_KeySQL" + BigTable + "_One_DEBUG(pMySQLEvent, table.custom_columns(), table_" + table.Name + ", pMsg->header(), pReqData);\n")
			f.WriteString("#endif\n")
			f.WriteString("        return GORM_PackGet_By_Non_Primary_KeySQL" + BigTable + "_One(mysql, table_" + table.Name + ", pMsg->header(), pReqData);\n")
			f.WriteString("    }\n")
			f.WriteString("    return GORM_OK;\n")
			f.WriteString("}\n")
		}
	}
	return 0
}

func CPPFieldsMapPackGetByNonPrimaryKeySQL(games []XmlCfg, f *os.File) int {
	CPPFieldsMapPackGetByNonPrimaryKeySQL_ForTables(games, f)
	CPPFields_GORM_PackSQL_TEMPLATE("Get_By_Non_Primary_Key", games, f)
	return 0
}
