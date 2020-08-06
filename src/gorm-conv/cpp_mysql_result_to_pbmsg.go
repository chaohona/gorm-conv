package main

import (
	"fmt"
	"strconv"
	"strings"

	//"fmt"
	"os"
	//"strconv"
)

func GORM_MySQLResult2PbMSG_SwitchTable(games []XmlCfg, f *os.File) int {
	f.WriteString("int GORM_MySQLResult2PbMSG(GORM_MySQLEvent *pMySQLEvent, int iTableId, GORM_PB_TABLE *pPbTable, MYSQL_ROW row, unsigned long *lengths)\n")
	f.WriteString("{\n")
	f.WriteString("    switch (iTableId)\n")
	f.WriteString("    {\n")
	for _, game := range games {
		for _, table := range game.DB.TableList {
			f.WriteString("    case GORM_PB_TABLE_IDX_" + strings.ToUpper(table.Name) + ":\n")
			f.WriteString("    {\n")
			f.WriteString("    #ifdef GORM_DEBUG\n")
			f.WriteString("        GORM_PB_Table_" + table.Name + " *pSrcTable = pPbTable->mutable_" + table.Name + "();\n")
			f.WriteString("        GORM_PB_CUSTEM_COLUMNS *pColumns = pPbTable->mutable_custom_columns();\n")
			f.WriteString("        return GORM_MySQLResult2PbMSG_" + strings.ToUpper(table.Name) + "_DEBUG(pMySQLEvent, pSrcTable, pColumns, row, lengths);\n")
			f.WriteString("    #else\n")
			f.WriteString("        GORM_PB_Table_" + table.Name + " *pSrcTable = pPbTable->mutable_" + table.Name + "();\n")
			f.WriteString("        return GORM_MySQLResult2PbMSG_" + strings.ToUpper(table.Name) + "(pSrcTable, row, lengths);\n")
			f.WriteString("    #endif\n")
			f.WriteString("    }\n")
		}
	}
	f.WriteString("    default:\n")
	f.WriteString("        return GORM_INVALID_TABLE;\n")
	f.WriteString("    }\n\n")

	f.WriteString("    return GORM_OK;\n")
	f.WriteString("}\n")
	return 0
}

func GORM_MySQLResult2PbMSG_Tables(games []XmlCfg, f *os.File) int {
	for _, game := range games {
		for _, table := range game.DB.TableList {
			f.WriteString("int GORM_MySQLResult2PbMSG_" + strings.ToUpper(table.Name) + "(GORM_PB_Table_" + table.Name + " *pPbTable, MYSQL_ROW row, unsigned long *lengths)\n")
			f.WriteString("{\n")
			if 0 != GORM_MySQLResult2PbMSG_Table(table, f) {
				fmt.Println("GORM_MySQLResult2PbMSG_Tables failed, table:", table.Name)
				return -1
			}

			f.WriteString("    return GORM_OK;\n")
			f.WriteString("}\n")
		}
	}
	return 0
}

func GORM_MySQLResult2PbMSG_Tables_DEBUG(games []XmlCfg, f *os.File) int {
	for _, game := range games {
		for _, table := range game.DB.TableList {
			f.WriteString("#ifdef GORM_DEBUG\n")
			f.WriteString("int GORM_MySQLResult2PbMSG_" + strings.ToUpper(table.Name) + "_DEBUG(GORM_MySQLEvent *pMySQLEvent, GORM_PB_Table_" + table.Name + " *pPbTable, GORM_PB_CUSTEM_COLUMNS *pColumns, MYSQL_ROW row, unsigned long *lengths)\n")
			f.WriteString("{\n")
			f.WriteString("    if (GORM_OK != GORM_MySQLResult2PbMSG_" + strings.ToUpper(table.Name) + "(pPbTable, row, lengths))\n")
			f.WriteString("        return GORM_ERROR;\n")
			f.WriteString("    return GORM_OK;\n")
			f.WriteString("    vector<string> &vTablesColumnOrder = pMySQLEvent->m_mapTablesColumnOrder[\"" + table.Name + "\"];\n")
			f.WriteString("    unordered_map<string, GORM_PB_COLUMN_TYPE> &mapTablesColumnType = pMySQLEvent->m_mapTablesColumnInfo[\"" + table.Name + "\"];\n")

			var strLen string = strconv.FormatInt(int64(len(table.TableColumns)), 10)
			f.WriteString("    for (int i=" + strLen + "; i<vTablesColumnOrder.size(); i++)\n")
			f.WriteString(`    {
        GORM_PB_COLUMN *pColumn = pColumns->add_columns();
        string &strColumn = vTablesColumnOrder[i];
        pColumn->set_name(strColumn);
        GORM_PB_COLUMN_VALUE *pColumnValue = pColumn->mutable_value();
        GORM_PB_COLUMN_TYPE iType = mapTablesColumnType[strColumn];
        pColumnValue->set_type(iType);
        switch (iType)
        {
        case GORM_PB_COLUMN_TYPE_INT:
        {
            // 转成int
            if (row[i] == nullptr || lengths[i] == 0)
                pColumnValue->set_intvalue(0);
            else
                pColumnValue->set_intvalue(strtol(row[i], (char **)NULL,10));
            break;
        }
        case GORM_PB_COLUMN_TYPE_UINT:
        {
            if (row[i] == nullptr || lengths[i] == 0)
                pColumnValue->set_uintvalue(0);
            else
                pColumnValue->set_uintvalue(strtoull(row[i], (char **)NULL,10));
            break;
        }
        case GORM_PB_COLUMN_TYPE_DOUBLE:
        {
            char* pEnd;
            if (row[i] == nullptr || lengths[i] == 0)
                pColumnValue->set_uintvalue(0);
            else
                pColumnValue->set_uintvalue(strtod(row[i], &pEnd));
            break;
        }
        case GORM_PB_COLUMN_TYPE_STRING:
        {
            if (row[i] == nullptr || lengths[i] == 0)
                pColumnValue->set_stringvalue("");
            else
                pColumnValue->set_stringvalue(row[i], lengths[i]);
            break;
        }
        }
    }
`)
			f.WriteString("    return GORM_OK;\n")
			f.WriteString("}\n")
			f.WriteString("#endif\n")
		}
	}
	return 0
}

func GORM_MySQLResult2PbMSG_Table(table TableInfo, f *os.File) int {
	for idx, col := range table.TableColumns {
		var idxStr string = strconv.FormatInt(int64(idx), 10)
		valStr, emptyStr := GORM_MySQLResult2PbMSG_Table_Col(idxStr, col)
		f.WriteString("    if (nullptr != row[" + idxStr + "] && lengths[" + idxStr + "] > 0)\n")
		f.WriteString("        pPbTable->set_" + col.Name + "(" + valStr + ");\n")
		f.WriteString("    else\n")
		f.WriteString("        pPbTable->set_" + col.Name + "(" + emptyStr + ");\n")
		f.WriteString("\n")
	}
	return 0
}

func GORM_MySQLResult2PbMSG_Table_Col(idxStr string, col TableColumn) (string, string) {
	var valStr, emptyStr string
	switch col.Type {
	case "int8", "int16", "int32", "int":
		{
			valStr = "strtol(row[" + idxStr + "], (char **)NULL,10)"
			emptyStr = "0"
		}
	case "uint8", "uint16", "uint32":
		{
			valStr = "strtoul(row[" + idxStr + "], (char **)NULL,10)"
			emptyStr = "0"
		}
	case "int64":
		{
			valStr = "strtoll(row[" + idxStr + "], (char **)NULL,10)"
			emptyStr = "0"
		}
	case "uint64":
		{
			valStr = "strtoull(row[" + idxStr + "], (char **)NULL,10)"
			emptyStr = "0"
		}
	case "double":
		{
			valStr = "strtoull(row[" + idxStr + "], (char **)NULL,10)"
			emptyStr = "0"
		}
	default:
		{
			valStr = "row[" + idxStr + "], lengths[" + idxStr + "]"
			emptyStr = "\"\""
		}
	}
	return valStr, emptyStr
}

func GORM_MySQLResult2PbMSG(games []XmlCfg, f *os.File) int {
	GORM_MySQLResult2PbMSG_Tables(games, f)
	GORM_MySQLResult2PbMSG_Tables_DEBUG(games, f)
	GORM_MySQLResult2PbMSG_SwitchTable(games, f)
	return 0
}
