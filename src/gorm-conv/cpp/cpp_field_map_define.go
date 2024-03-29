package cpp

// 生成文件gorm_fields_map_define.cc

import (
	"fmt"
	"gorm-conv/common"
	"os"
	"strconv"
	"strings"
)

func CppFieldsMapDefine1(games []common.XmlCfg, f *os.File) int {
	var id2name string = `int GORM_SetTableFieldId2Name(int iTableType, OUT FieldId2Name &mapId2Name)
{
    switch (iTableType)
    {
`
	f.WriteString(id2name)
	// 输出各种case
	for _, game := range games {
		for _, table := range game.DB.TableList {
			/*
				case GORM_PB_TABLE_IDX_ACCOUNT:
				{
					return GORM_SetTableACCOUNTId2Name(mapId2Name)
				}
			*/
			var caseStr string = "    case GORM_PB_TABLE_IDX_" + strings.ToUpper(table.Name) + ":\n"
			caseStr += "    {\n"
			caseStr += "        return GORM_SetTable" + strings.ToUpper(table.Name) + "Id2Name(mapId2Name);\n"
			caseStr += "    }\n"
			f.WriteString(caseStr)
		}
	}
	// 输出default
	/*
			default:
		        return GORM_ERROR;
		    }

		    return GORM_OK
		}
	*/
	var def string = `    default:
        return GORM_ERROR;
    }

    return GORM_OK;
}
`
	f.WriteString(def)
	return 0
}

func CppFieldsMapDefine2(games []common.XmlCfg, f *os.File) int {
	var name2id string = `int GORM_SetTableFieldName2Id(int iTableType, OUT FieldName2Id &mapName2Id)
{
    switch (iTableType)
    {
`
	f.WriteString(name2id)
	// 输出各种case
	for _, game := range games {
		for _, table := range game.DB.TableList {
			/*
				case GORM_PB_TABLE_IDX_ACCOUNT:
				{
					return GORM_SetTableFieldName2Id(mapName2Id)
				}
			*/
			var caseStr string = "    case GORM_PB_TABLE_IDX_" + strings.ToUpper(table.Name) + ":\n"
			caseStr += "    {\n"
			caseStr += "        return GORM_SetTable" + strings.ToUpper(table.Name) + "Name2Id(mapName2Id);\n"
			caseStr += "    }\n"
			f.WriteString(caseStr)
		}
	}
	// 输出default
	/*
			default:
		        return GORM_ERROR;
		    }

		    return GORM_OK
		}
	*/
	var def string = `    default:
        return GORM_ERROR;
    }

    return GORM_OK;
}
`
	f.WriteString(def)
	return 0
}

func CppFieldsMapDefine3(games []common.XmlCfg, f *os.File) int {
	/*
		int GORM_SetTableACCOUNTId2Name(OUT FieldId2Name &mapId2Name)
		{
		    return GORM_OK;
		}
	*/
	for _, game := range games {
		for _, table := range game.DB.TableList {
			tableName := strings.ToUpper(table.Name)
			var header = "int GORM_SetTable" + tableName + "Id2Name(OUT FieldId2Name &mapId2Name)\n"
			header += "{\n"
			//header += "    mapId2Name[GORM_PB_FIELD_" + tableName + "_VERSION] = \"version\";\n"
			f.WriteString(header)
			for _, col := range table.TableColumns {
				colName := strings.ToUpper(col.Name)
				colStr := "    mapId2Name[GORM_PB_FIELD_" + tableName + "_" + colName + "] = \"" + col.Name + "\";\n"
				f.WriteString(colStr)
			}
			end := "    return GORM_OK;\n}\n"
			f.WriteString(end)
		}
	}
	return 0
}

func CppFieldsMapDefine4(games []common.XmlCfg, f *os.File) int {
	/*
		int GORM_SetTableACCOUNTName2Id(OUT FieldName2Id &mapName2Id)
		{
		    return GORM_OK;
		}
	*/
	for _, game := range games {
		for _, table := range game.DB.TableList {
			tableName := strings.ToUpper(table.Name)
			var header = "int GORM_SetTable" + tableName + "Name2Id(OUT FieldName2Id &mapName2Id)\n"
			header += "{\n"
			//header += "    mapName2Id[\"version\"] = GORM_PB_FIELD_" + tableName + "_VERSION;\n"
			f.WriteString(header)
			for _, col := range table.TableColumns {
				colName := strings.ToUpper(col.Name)
				colStr := "    mapName2Id[\"" + col.Name + "\"] = GORM_PB_FIELD_" + tableName + "_" + colName + ";\n"
				f.WriteString(colStr)
			}
			end := "    return GORM_OK;\n}\n"
			f.WriteString(end)
		}
	}
	return 0
}

func CppFieldsMapDefine5(games []common.XmlCfg, f *os.File) int {
	/*
		int GORM_SetTableACCOUNTId2Name(OUT FieldId2Name &mapId2Name);
		int GORM_SetTableBAGId2Name(OUT FieldId2Name &mapId2Name);
		int GORM_SetTableACCOUNTName2Id(OUT FieldId2Name &mapName2Id);
		int GORM_SetTableBAGName2Id(OUT FieldId2Name &mapName2Id);
	*/
	for _, game := range games {
		for _, table := range game.DB.TableList {
			var id2name string = "int GORM_SetTable" + strings.ToUpper(table.Name) + "Id2Name(OUT FieldId2Name &mapId2Name);\n"
			f.WriteString(id2name)
			var name2id string = "int GORM_SetTable" + strings.ToUpper(table.Name) + "Name2Id(OUT FieldName2Id &mapName2Id);\n"
			f.WriteString(name2id)
		}
	}
	f.WriteString("\n\n")
	return 0
}

func CppFieldsMapDefine6(games []common.XmlCfg, f *os.File) int {
	/*
		int GORM_SetTableName2Id(OUT TableName2Id &mapName2Id)
		{
		    mapName2Id["account"] = GORM_PB_TABLE_IDX_ACCOUNT;
		    mapName2Id["bag"] = GORM_PB_TABLE_IDX_BAG;
		    return GORM_OK;
		}
	*/
	f.WriteString("int GORM_SetTableName2Id(OUT TableName2Id &mapName2Id)\n{\n")
	for _, game := range games {
		for _, table := range game.DB.TableList {
			var name2id = "    mapName2Id[\"" + table.Name + "\"] = GORM_PB_TABLE_IDX_" + strings.ToUpper(table.Name) + ";\n"
			f.WriteString(name2id)
		}
	}
	f.WriteString("    return GORM_OK;\n}\n\n")
	return 0
}

func CppFieldsMapDefine7(games []common.XmlCfg, f *os.File) int {
	/*
		int GORM_SetTableId2Name(OUT TableId2Name &mapId2Name)
		{
		    mapId2Name[GORM_PB_TABLE_IDX_ACCOUNT] = "account";
		    mapId2Name[GORM_PB_TABLE_IDX_BAG] = "bag";
		    return GORM_OK;
		}
	*/
	f.WriteString("int GORM_SetTableId2Name(OUT TableId2Name &mapId2Name)\n{\n")
	for _, game := range games {
		for _, table := range game.DB.TableList {
			var name2id = "    mapId2Name[GORM_PB_TABLE_IDX_" + strings.ToUpper(table.Name) + "] = \"" + table.Name + "\";\n"
			f.WriteString(name2id)
		}
	}
	f.WriteString("    return GORM_OK;\n}\n\n")
	return 0
}

func CppFieldsMapDefine8(games []common.XmlCfg, f *os.File) int {
	/*
		int GORM_SetTableVersion(OUT TableVersionMap& mapTableVersion)
		{
		    mapTableVersion[GORM_PB_TABLE_IDX_ACCOUNT] = 1;
		    return GORM_OK;
		}
	*/
	f.WriteString("int GORM_SetTableVersion(OUT TableVersionMap& mapTableVersion)\n{\n")
	for _, game := range games {
		for _, table := range game.DB.TableList {
			var tableVersion = "    mapTableVersion[GORM_PB_TABLE_IDX_" + strings.ToUpper(table.Name) + "] = "
			tableVersion += strconv.FormatInt(int64(table.Version), 10) + ";\n"
			f.WriteString(tableVersion)
		}
	}
	f.WriteString("    return GORM_OK;\n}\n\n")
	return 0
}

func CPPFieldsMapSetTableFieldValueSub(games []common.XmlCfg, f *os.File) int {
	totalIdx := 0
	for _, game := range games {
		for _, table := range game.DB.TableList {
			colTotal := 0
			for _, col := range table.TableColumns {
				if col.Type != "blob" && col.Type != "char" && col.Type != "string" && col.Type != "bytes" {
					totalIdx += 1
					colTotal += 1
					if totalIdx == 1 {
						f.WriteString("    switch (iTableId)\n    {\n")
					}
					if colTotal == 1 {
						// 输出表名 case GORM_PB_TABLE_IDX_ACCOUNT:
						f.WriteString("    case GORM_PB_TABLE_IDX_")
						f.WriteString(strings.ToUpper(table.Name))
						f.WriteString(":\n")
						f.WriteString("    {\n")
						f.WriteString("        switch (iFieldId)\n        {\n")
					}
					// 输出列名 case GORM_PB_FIELD_ACCOUNT_VERSION
					f.WriteString("        case GORM_PB_FIELD_")
					f.WriteString(strings.ToUpper(table.Name))
					f.WriteString("_")
					f.WriteString(strings.ToUpper(col.Name))
					f.WriteString(":\n")
					f.WriteString("        {\n")
					// 输出return语句 return (GORM_PB_Table_account*)(pMsg)->set_version(value);
					var PBTableType string = "GORM_PB_Table_" + table.Name
					f.WriteString("            " + PBTableType + "* pPbReal = dynamic_cast<" + PBTableType + "*>(pMsg);\n")
					f.WriteString("            return pPbReal->set_")
					f.WriteString(col.Name)
					f.WriteString("(value);\n")
					f.WriteString("        }\n")
				}
			}
			if colTotal > 0 {
				f.WriteString("        }\n")
				f.WriteString("    }\n")
			}
		}
		/*if tableMatch {
			f.WriteString("        }\n")
			f.WriteString("    }\n")
		}*/
	}
	if totalIdx > 0 {
		f.WriteString("    }\n")
	}
	f.WriteString("}\n\n")

	return 0
}

func CPPFieldsMapSetTableFieldValue(games []common.XmlCfg, f *os.File) int {
	// 1.void GORM_SetTableFieldValue(PB_MSG_PTR pMsg, int iTableId, int iFieldId, const char * value, const size_t size)
	f.WriteString("void GORM_SetTableFieldValue(PB_MSG_PTR pMsg, int iTableId, int iFieldId, const char * value, const size_t size)\n{\n")
	totalIdx := 0
	for _, game := range games {
		for _, table := range game.DB.TableList {
			colTotal := 0
			for _, col := range table.TableColumns {
				if col.Type == "blob" || col.Type == "char" || col.Type == "string" || col.Type == "bytes" {
					totalIdx += 1
					colTotal += 1
					if totalIdx == 1 {
						f.WriteString("    switch (iTableId)\n    {\n")
					}
					if colTotal == 1 {
						// 输出表名 case GORM_PB_TABLE_IDX_ACCOUNT:
						f.WriteString("    case GORM_PB_TABLE_IDX_")
						f.WriteString(strings.ToUpper(table.Name))
						f.WriteString(":\n")
						f.WriteString("    {\n")
						f.WriteString("        switch (iFieldId)\n        {\n")
					}
					// 输出列名 case GORM_PB_FIELD_ACCOUNT_VERSION
					f.WriteString("        case GORM_PB_FIELD_")
					f.WriteString(strings.ToUpper(table.Name))
					f.WriteString("_")
					f.WriteString(strings.ToUpper(col.Name))
					f.WriteString(":\n")
					f.WriteString("        {\n")
					// 输出return语句 return (GORM_PB_Table_account*)(pMsg)->set_version(value);
					var PBTableType string = "GORM_PB_Table_" + table.Name
					f.WriteString("            " + PBTableType + "* pPbReal = dynamic_cast<" + PBTableType + "*>(pMsg);\n")
					f.WriteString("            return pPbReal->set_")
					f.WriteString(col.Name)
					f.WriteString("((const char*)value, size);\n")
					f.WriteString("        }\n")
				}
			}
			if colTotal > 0 {
				f.WriteString("        }\n")
				f.WriteString("    }\n")
			}
		}
		/*if tableMatch {
			f.WriteString("        }\n")
			f.WriteString("    }\n")
		}*/
	}
	if totalIdx > 0 {
		f.WriteString("    }\n")
	}
	f.WriteString("}\n\n")
	////////////////////////////////////////////////////////////////

	// 2.void GORM_SetTableFieldValue(PB_MSG_PTR pMsg, int iTableId, int iFieldId, const char * value)
	f.WriteString("void GORM_SetTableFieldValue(PB_MSG_PTR pMsg, int iTableId, int iFieldId, const char * value)\n{\n")
	totalIdx = 0
	for _, game := range games {
		for _, table := range game.DB.TableList {
			colTotal := 0
			for _, col := range table.TableColumns {
				if col.Type == "blob" || col.Type == "char" || col.Type == "string" || col.Type == "bytes" {
					totalIdx += 1
					colTotal += 1
					if totalIdx == 1 {
						f.WriteString("    switch (iTableId)\n    {\n")
					}
					if colTotal == 1 {
						// 输出表名 case GORM_PB_TABLE_IDX_ACCOUNT:
						f.WriteString("    case GORM_PB_TABLE_IDX_")
						f.WriteString(strings.ToUpper(table.Name))
						f.WriteString(":\n")
						f.WriteString("    {\n")
						f.WriteString("        switch (iFieldId)\n        {\n")
					}
					// 输出列名 case GORM_PB_FIELD_ACCOUNT_VERSION
					f.WriteString("        case GORM_PB_FIELD_")
					f.WriteString(strings.ToUpper(table.Name))
					f.WriteString("_")
					f.WriteString(strings.ToUpper(col.Name))
					f.WriteString(":\n")
					f.WriteString("        {\n")
					// 输出return语句 return (GORM_PB_Table_account*)(pMsg)->set_version(value);
					var PBTableType string = "GORM_PB_Table_" + table.Name
					f.WriteString("            " + PBTableType + "* pPbReal = dynamic_cast<" + PBTableType + "*>(pMsg);\n")
					f.WriteString("            return pPbReal->set_")
					f.WriteString(col.Name)
					f.WriteString("((const char*)value);\n")
					f.WriteString("        }\n")
				}
			}
			if colTotal > 0 {
				f.WriteString("        }\n")
				f.WriteString("    }\n")
			}
		}
		/*if tableMatch {
			f.WriteString("        }\n")
			f.WriteString("    }\n")
		}*/
	}
	if totalIdx > 0 {
		f.WriteString("    }\n")
	}
	f.WriteString("}\n\n")
	////////////////////////////////////////////////////////////////

	var setFuncs []string = []string{
		"void GORM_SetTableFieldValue(PB_MSG_PTR pMsg, int iTableId, int iFieldId, int8 value)",
		"void GORM_SetTableFieldValue(PB_MSG_PTR pMsg, int iTableId, int iFieldId, int16 value)",
		"void GORM_SetTableFieldValue(PB_MSG_PTR pMsg, int iTableId, int iFieldId, int32 value)",
		"void GORM_SetTableFieldValue(PB_MSG_PTR pMsg, int iTableId, int iFieldId, int64 value)",
		"void GORM_SetTableFieldValue(PB_MSG_PTR pMsg, int iTableId, int iFieldId, double value)",
		"void GORM_SetTableFieldValue(PB_MSG_PTR pMsg, int iTableId, int iFieldId, uint8 value)",
		"void GORM_SetTableFieldValue(PB_MSG_PTR pMsg, int iTableId, int iFieldId, uint16 value)",
		"void GORM_SetTableFieldValue(PB_MSG_PTR pMsg, int iTableId, int iFieldId, uint32 value)",
		"void GORM_SetTableFieldValue(PB_MSG_PTR pMsg, int iTableId, int iFieldId, uint64 value)",
	}
	for _, str := range setFuncs {
		f.WriteString(str)
		f.WriteString("\n{\n")
		CPPFieldsMapSetTableFieldValueSub(games, f)
	}

	return 0
}

func CPPFieldsMapGetTableFieldValueSub(games []common.XmlCfg, f *os.File) int {
	totalIdx := 0
	for _, game := range games {
		for _, table := range game.DB.TableList {
			colTotal := 0
			for _, col := range table.TableColumns {
				if col.Type != "blob" && col.Type != "char" && col.Type != "string" && col.Type != "bytes" {
					totalIdx += 1
					colTotal += 1
					if totalIdx == 1 {
						f.WriteString("    switch (iTableId)\n    {\n")
					}
					if colTotal == 1 {
						// 输出表名 case GORM_PB_TABLE_IDX_ACCOUNT:
						f.WriteString("    case GORM_PB_TABLE_IDX_")
						f.WriteString(strings.ToUpper(table.Name))
						f.WriteString(":\n")
						f.WriteString("    {\n")
						f.WriteString("        switch (iFieldId)\n        {\n")
					}
					// 输出列名 case GORM_PB_FIELD_ACCOUNT_VERSION
					f.WriteString("        case GORM_PB_FIELD_")
					f.WriteString(strings.ToUpper(table.Name))
					f.WriteString("_")
					f.WriteString(strings.ToUpper(col.Name))
					f.WriteString(":\n")
					f.WriteString("        {\n")
					// 输出return语句 return (GORM_PB_Table_account*)(pMsg)->set_version(value);
					var PBTableType string = "GORM_PB_Table_" + table.Name
					f.WriteString("            " + PBTableType + "* pPbReal = dynamic_cast<" + PBTableType + "*>(pMsg);\n")
					f.WriteString("            value = pPbReal->")
					f.WriteString(col.Name)
					f.WriteString("();\n")
					f.WriteString("            return GORM_OK;\n")
					f.WriteString("        }\n")
				}
			}
			if colTotal > 0 {
				f.WriteString("        }\n")
				f.WriteString("    }\n")
			}
		}
		/*if tableMatch {
			f.WriteString("        }\n")
			f.WriteString("    }\n")
		}*/
	}
	if totalIdx > 0 {
		f.WriteString("    }\n")
		f.WriteString("\n    return GORM_ERROR;\n")
	} else {
		f.WriteString("    return GORM_OK;\n")
	}
	f.WriteString("}\n\n")

	return 0
}

func CPPFieldsMapGetTableFieldValue(games []common.XmlCfg, f *os.File) int {
	// 1.int GORM_GetTableFieldValue(PB_MSG_PTR pMsg, int iTableId, int iFieldId, string &strValue);
	f.WriteString("int GORM_GetTableFieldValue(PB_MSG_PTR pMsg, int iTableId, int iFieldId, string &value)\n{\n")
	totalIdx := 0
	for _, game := range games {
		for _, table := range game.DB.TableList {
			colTotal := 0
			for _, col := range table.TableColumns {
				if col.Type == "blob" || col.Type == "char" || col.Type == "string" || col.Type == "bytes" {
					totalIdx += 1
					colTotal += 1
					if totalIdx == 1 {
						f.WriteString("    switch (iTableId)\n    {\n")
					}
					if colTotal == 1 {
						// 输出表名 case GORM_PB_TABLE_IDX_ACCOUNT:
						f.WriteString("    case GORM_PB_TABLE_IDX_")
						f.WriteString(strings.ToUpper(table.Name))
						f.WriteString(":\n")
						f.WriteString("    {\n")
						f.WriteString("        switch (iFieldId)\n        {\n")
					}
					// 输出列名 case GORM_PB_FIELD_ACCOUNT_VERSION
					f.WriteString("        case GORM_PB_FIELD_")
					f.WriteString(strings.ToUpper(table.Name))
					f.WriteString("_")
					f.WriteString(strings.ToUpper(col.Name))
					f.WriteString(":\n")
					f.WriteString("        {\n")
					// 输出return语句 strValue = (GORM_PB_Table_account*)pMsg->account();
					var PBTableType string = "GORM_PB_Table_" + table.Name
					f.WriteString("            " + PBTableType + "* pPbReal = dynamic_cast<" + PBTableType + "*>(pMsg);\n")
					f.WriteString("            value = pPbReal->")
					f.WriteString(col.Name)
					f.WriteString("();\n")
					f.WriteString("            return GORM_OK;\n")
					f.WriteString("        }\n")
				}
			}
			if colTotal > 0 {
				f.WriteString("        }\n")
				f.WriteString("    }\n")
			}
		}
		/*if tableMatch {
			f.WriteString("        }\n")
			f.WriteString("    }\n")
		}*/
	}
	if totalIdx > 0 {
		f.WriteString("    }\n")
		f.WriteString("\n    return GORM_ERROR;\n")
	} else {
		f.WriteString("    return GORM_OK;\n")
	}
	f.WriteString("}\n\n")
	////////////////////////////////////////////////////////////////

	// 2.int GORM_GetTableFieldValue(PB_MSG_PTR pMsg, int iTableId, int iFieldId, uint8 *&value, size_t &size);
	f.WriteString("int GORM_GetTableFieldValue(PB_MSG_PTR pMsg, int iTableId, int iFieldId, uint8 *&value, size_t &size)\n{\n")
	totalIdx = 0
	for _, game := range games {
		for _, table := range game.DB.TableList {
			colTotal := 0
			for _, col := range table.TableColumns {
				if col.Type == "blob" || col.Type == "char" || col.Type == "string" || col.Type == "bytes" {
					totalIdx += 1
					colTotal += 1
					if totalIdx == 1 {
						f.WriteString("    switch (iTableId)\n    {\n")
					}
					if colTotal == 1 {
						// 输出表名 case GORM_PB_TABLE_IDX_ACCOUNT:
						f.WriteString("    case GORM_PB_TABLE_IDX_")
						f.WriteString(strings.ToUpper(table.Name))
						f.WriteString(":\n")
						f.WriteString("    {\n")
						f.WriteString("        switch (iFieldId)\n        {\n")
					}
					// 输出列名 case GORM_PB_FIELD_ACCOUNT_VERSION
					f.WriteString("        case GORM_PB_FIELD_")
					f.WriteString(strings.ToUpper(table.Name))
					f.WriteString("_")
					f.WriteString(strings.ToUpper(col.Name))
					f.WriteString(":\n")
					f.WriteString("        {\n")
					f.WriteString("")
					// 输出return语句 return (GORM_PB_Table_account*)(pMsg)->set_version(value);
					var PBTableType string = "GORM_PB_Table_" + table.Name
					f.WriteString("            " + PBTableType + "* pPbReal = dynamic_cast<" + PBTableType + "*>(pMsg);\n")
					f.WriteString("            string strValue = pPbReal->")
					f.WriteString(col.Name)
					f.WriteString("();\n")
					f.WriteString("            value=(uint8*)strValue.c_str();\n")
					f.WriteString("            size=strValue.size();\n")
					f.WriteString("            return GORM_OK;\n")
					f.WriteString("        }\n")
				}
			}
			if colTotal > 0 {
				f.WriteString("        }\n")
				f.WriteString("    }\n")
			}
		}
		/*if tableMatch {
			f.WriteString("        }\n")
			f.WriteString("    }\n")
		}*/
	}
	if totalIdx > 0 {
		f.WriteString("    }\n")
		f.WriteString("\n    return GORM_ERROR;\n")
	} else {
		f.WriteString("    return GORM_OK;\n")
	}
	f.WriteString("}\n\n")
	////////////////////////////////////////////////////////////////

	var getFuncs []string = []string{
		"int GORM_GetTableFieldValue(PB_MSG_PTR pMsg, int iTableId, int iFieldId, int8 &value)",
		"int GORM_GetTableFieldValue(PB_MSG_PTR pMsg, int iTableId, int iFieldId, int16 &value)",
		"int GORM_GetTableFieldValue(PB_MSG_PTR pMsg, int iTableId, int iFieldId, int32 &value)",
		"int GORM_GetTableFieldValue(PB_MSG_PTR pMsg, int iTableId, int iFieldId, int64 &value)",
		"int GORM_GetTableFieldValue(PB_MSG_PTR pMsg, int iTableId, int iFieldId, double &value)",
		"int GORM_GetTableFieldValue(PB_MSG_PTR pMsg, int iTableId, int iFieldId, uint8 &value)",
		"int GORM_GetTableFieldValue(PB_MSG_PTR pMsg, int iTableId, int iFieldId, uint16 &value)",
		"int GORM_GetTableFieldValue(PB_MSG_PTR pMsg, int iTableId, int iFieldId, uint32 &value)",
		"int GORM_GetTableFieldValue(PB_MSG_PTR pMsg, int iTableId, int iFieldId, uint64 &value)",
	}
	_ = getFuncs
	for _, str := range getFuncs {
		f.WriteString(str)
		f.WriteString("\n{\n")
		CPPFieldsMapGetTableFieldValueSub(games, f)
	}

	return 0
}

func GetCustomPbMsg(games []common.XmlCfg, f *os.File) int {
	var customFunc string = `
int GORM_GetCustomPbMsg(PB_MSG_PTR &pMsgPtr)
{
    pMsgPtr = new GORM_PB_CUSTEM_COLUMNS();
    return GORM_OK;
}
`
	f.WriteString(customFunc)

	return 0
}

func GetTablePbMsgDefine(games []common.XmlCfg, f *os.File) int {
	f.WriteString("int GetTablePbMsgDefine(int iTableId, PB_MSG_PTR &pMsgPtr)\n")
	f.WriteString("{\n")

	totalIdx := 0
	for _, game := range games {
		for _, table := range game.DB.TableList {
			totalIdx += 1
			if totalIdx == 1 {
				f.WriteString("    switch (iTableId)\n")
				f.WriteString("    {\n")
			}
			f.WriteString("    case GORM_PB_TABLE_IDX_")
			f.WriteString(strings.ToUpper(table.Name))
			f.WriteString(":\n")
			f.WriteString("    {\n")
			var tableName string = "GORM_PB_Table_" + table.Name
			f.WriteString("        pMsgPtr = new ")
			f.WriteString(tableName + "();\n")
			f.WriteString("        return GORM_OK;\n")
			f.WriteString("    }\n")
		}
	}

	if totalIdx > 0 {
		f.WriteString("    }\n")
	}
	f.WriteString("    return GORM_INVALID_TABLE;\n")
	f.WriteString("}\n")
	return 0
}

func GORM_AddRecordToReqPbMsgDefine(games []common.XmlCfg, f *os.File) int {
	f.WriteString("int GORM_AddRecordToReqPbMsgDefine(int iTableId, GORM_PB_TABLE *pPbTable, PB_MSG_PTR pPbMsg)\n")
	f.WriteString("{\n")

	totalIdx := 0
	for _, game := range games {
		for _, table := range game.DB.TableList {
			totalIdx += 1
			if totalIdx == 1 {
				f.WriteString("    switch (iTableId)\n")
				f.WriteString("    {\n")
			}
			f.WriteString("    case GORM_PB_TABLE_IDX_")
			f.WriteString(strings.ToUpper(table.Name))
			f.WriteString(":\n")
			f.WriteString("    {\n")
			f.WriteString("        GORM_PB_Table_")
			f.WriteString(table.Name)
			f.WriteString(" *pTableMsg = dynamic_cast<GORM_PB_Table_")
			f.WriteString(table.Name)
			f.WriteString("*>(pPbMsg);\n")
			f.WriteString("        pPbTable->set_allocated_")
			f.WriteString(table.Name)
			f.WriteString("(pTableMsg);\n")
			f.WriteString("        return GORM_OK;\n")
			f.WriteString("    }\n")
		}
	}

	if totalIdx > 0 {
		f.WriteString("    }\n")
	}
	f.WriteString("    return GORM_INVALID_TABLE;\n")
	f.WriteString("}\n")
	return 0
}

func GORM_TableHasData(games []common.XmlCfg, f *os.File) int {
	f.WriteString("bool GORM_TableHasData(GORM_PB_TABLE *pTable, int iTableId)\n")
	f.WriteString("{\n")
	f.WriteString("    switch (iTableId)\n")
	f.WriteString("    {\n")
	for _, game := range games {
		for _, table := range game.DB.TableList {
			f.WriteString("    case GORM_PB_TABLE_IDX_" + strings.ToUpper(table.Name) + ":\n")
			f.WriteString("        return pTable->has_" + table.Name + "();\n")
		}
	}
	f.WriteString("    }\n\n")
	f.WriteString("    return false;\n")
	f.WriteString("}\n")

	return 0
}

func GORM_GetTableSrcPbMsg(games []common.XmlCfg, f *os.File) int {
	f.WriteString("int GORM_GetTableSrcPbMsg(int iTableId, GORM_PB_TABLE *pTable, PB_MSG_PTR &pMsgPtr)\n")
	f.WriteString("{\n")
	f.WriteString("    switch (iTableId)\n")
	f.WriteString("    {\n")
	for _, game := range games {
		for _, table := range game.DB.TableList {
			f.WriteString("    case GORM_PB_TABLE_IDX_" + strings.ToUpper(table.Name) + ":\n")
			f.WriteString("    {\n")
			f.WriteString("        pMsgPtr = pTable->mutable_" + table.Name + "();\n")
			f.WriteString("        return GORM_OK;\n")
			f.WriteString("    }\n")
		}
	}
	f.WriteString("    }\n\n")
	f.WriteString("    return false;\n")
	f.WriteString("}\n")

	return 0
}

func CPPFieldInitTableColumnInfo_ForTable(games []common.XmlCfg, f *os.File) int {
	for _, game := range games {
		for _, table := range game.DB.TableList {
			f.WriteString("int GORM_InitTableColumnInfo_" + table.Name + "(unordered_map<string, vector<string>> &mapTablesColumnOrder, unordered_map<string, unordered_map<string, GORM_PB_COLUMN_TYPE>> &mapTablesColumnInfo)\n")
			f.WriteString("{\n")
			f.WriteString("    vector<string> vColumns = {")
			for idx, col := range table.TableColumns {
				if idx != 0 {
					f.WriteString(",")
				}
				f.WriteString("\"" + col.Name + "\"")
			}
			f.WriteString("};\n")
			f.WriteString("    mapTablesColumnOrder[\"" + table.Name + "\"] = vColumns;\n")
			f.WriteString("    unordered_map<string, GORM_PB_COLUMN_TYPE> mapColumnType = {\n")
			for _, col := range table.TableColumns {
				f.WriteString("        {\"" + col.Name + "\", " + common.CPPFieldInitTableColumnInfo_ForTable_COLTYPE(col.Type) + "},\n")
			}
			f.WriteString("    };\n")
			f.WriteString("    mapTablesColumnInfo[\"" + table.Name + "\"] = mapColumnType;\n")
			f.WriteString("    return GORM_OK;\n")
			f.WriteString("}\n")
		}
	}
	return 0
}

func GORM_InitTableSchemaInfo(games []common.XmlCfg, f *os.File) int {
	f.WriteString("int GORM_InitTableSchemaInfo(PB_MSG_PTR pMsgPtr)\n")
	f.WriteString("{\n")
	f.WriteString("    GORM_PB_HAND_SHAKE_REQ *pHandShake = dynamic_cast<GORM_PB_HAND_SHAKE_REQ*>(pMsgPtr);\n")
	f.WriteString("    if (pHandShake == nullptr)\n")
	f.WriteString("        return GORM_ERROR;\n")
	f.WriteString("    GORM_PB_TABLE_SCHEMA_INFO *pInfo;\n")
	f.WriteString("    GORM_PB_TABLE_SCHEMA_INFO_COLUMN *pColumn;\n")
	f.WriteString("    pHandShake->mutable_header();\n")
	f.WriteString("    pHandShake->set_version(1);\n")
	f.WriteString("    pHandShake->set_md5(1);\n")
	var tableIdx int64 = 0
	for _, game := range games {
		for _, table := range game.DB.TableList {
			f.WriteString("    // for table " + table.Name + "\n")
			f.WriteString("    pInfo = pHandShake->add_schemas();\n")
			f.WriteString("    if (pInfo == nullptr)\n")
			f.WriteString("        return GORM_ERROR;\n")
			var tableVer string = strconv.FormatUint(table.Version, 10)
			f.WriteString("    pInfo->set_version(" + tableVer + ");\n")
			f.WriteString("    pInfo->set_tablename(\"" + table.Name + "\");\n")
			tableIdx += 1
			var tableIdxStr string = strconv.FormatInt(tableIdx, 10)
			f.WriteString("    pInfo->set_tableidx(" + tableIdxStr + ");\n")
			for _, col := range table.TableColumns {
				f.WriteString("    pColumn = pInfo->add_columns();\n")
				f.WriteString("    if (pColumn == nullptr)\n")
				f.WriteString("        return GORM_ERROR;\n")
				var colVer string = strconv.FormatUint(col.Version, 10)
				f.WriteString("    pColumn->set_version(" + colVer + ");\n")
				f.WriteString("    pColumn->set_name(\"" + col.Name + "\");\n")
				f.WriteString("    pColumn->set_typedesc(\"" + col.Type + "\");\n")
				f.WriteString("    pColumn->set_type(GORM_PB_COLUMN_TYPE_" + strings.ToUpper(col.Type) + ");\n")
			}
		}
	}

	f.WriteString("    return GORM_OK;\n")
	f.WriteString("}\n")
	return 0
}

func CPPFieldInitTableColumnInfo(games []common.XmlCfg, f *os.File) int {
	if 0 != CPPFieldInitTableColumnInfo_ForTable(games, f) {
		return -1
	}
	f.WriteString("int GORM_InitTableColumnInfo(unordered_map<string, vector<string>> &mapTablesColumnOrder, unordered_map<string, unordered_map<string, GORM_PB_COLUMN_TYPE>> &mapTablesColumnInfo)\n")
	f.WriteString("{\n")
	for _, game := range games {
		for _, table := range game.DB.TableList {
			f.WriteString("    if (GORM_InitTableColumnInfo_" + table.Name + "(mapTablesColumnOrder, mapTablesColumnInfo))\n")
			f.WriteString("        return GORM_ERROR;\n")
		}
	}
	f.WriteString("    return GORM_OK;\n")
	f.WriteString("}\n")
	return 0
}

func CppFieldsMapDefine(games []common.XmlCfg, outpath string) int {
	outfile := outpath + "gorm_table_field_map_define.cc"
	f, err := os.OpenFile(outfile, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err.Error())
		return -1
	}
	defer f.Close()
	fmt.Println("begin to general file:" + outfile)
	err = f.Truncate(0)

	f.WriteString(common.NOT_EDIT_WARNING)

	// 1、输出固定的头/////////////////////////
	//#include "mysql.h"
	var header string = `#include "gorm_table_field_map_define.h"
#include "gorm_pb_tables_inc.pb.h"
#include "gorm_pb_proto.pb.h"
#include "gorm_mempool.h"
#include "gorm_hash.h"

namespace gorm{

`
	_, err = f.WriteString(header)
	if err != nil {
		fmt.Println(err.Error())
		return -1
	}

	// >> 输出函数声明
	if 0 != CppFieldsMapDefine5(games, f) {
		fmt.Println("CppFieldsMapDefine5 failed")
		return -1
	}
	if 0 != GORM_InitTableSchemaInfo(games, f) {
		fmt.Println("GORM_InitTableSchemaInfo failed.")
		return -1
	}
	if 0 != GORM_TableHash(games, f) {
		fmt.Println("GORM_TableHash failed.")
		return -1
	}
	if 0 != GetCustomPbMsg(games, f) {
		fmt.Println("GetCustomPbMsg failed")
		return -1
	}
	if 0 != GetTablePbMsgDefine(games, f) {
		fmt.Println("GetCustomerPbMsgDefine failed")
		return -1
	}
	if 0 != GORM_TableHasData(games, f) {
		fmt.Println("GORM_TableHasData failed")
		return -1
	}
	if 0 != GORM_GetTableSrcPbMsg(games, f) {
		fmt.Println("GORM_GetTableSrcPbMsg failed")
		return -1
	}
	if 0 != GORM_AddRecordToReqPbMsgDefine(games, f) {
		fmt.Println("GORM_AddRecordToReqPbMsgDefine failed")
		return -1
	}
	// >> 设置表的版本号
	if 0 != CppFieldsMapDefine8(games, f) {
		fmt.Println("CppFieldsMapDefine8 failed")
		return -1
	}
	// >> 输出表名与宏的映射
	if 0 != CppFieldsMapDefine6(games, f) {
		fmt.Println("CppFieldsMapDefine6 failed")
		return -1
	}
	// >> 输出宏与表名的映射
	if 0 != CppFieldsMapDefine7(games, f) {
		fmt.Println("CppFieldsMapDefine7 failed")
		return -1
	}

	// 2、输出id2name入口函数///////////////////////
	if 0 != CppFieldsMapDefine1(games, f) {
		fmt.Println("CppFieldsMapDefine1 failed")
		return -1
	}
	// 3、输出name2id入口函数
	if 0 != CppFieldsMapDefine2(games, f) {
		fmt.Println("CppFieldsMapDefine2 failed")
		return -1
	}
	// 4、输出id2name转换函数
	if 0 != CppFieldsMapDefine3(games, f) {
		fmt.Println("CppFieldsMapDefine3 failed")
		return -1
	}
	// 5、输出name2id转换函数
	if 0 != CppFieldsMapDefine4(games, f) {
		fmt.Println("CppFieldsMapDefine4 failed")
		return -1
	}
	if 0 != CPPFieldInitTableColumnInfo(games, f) {
		fmt.Println("CPPFieldInitTableColumnInfo faild.")
		return -1
	}
	if 0 != CPPFieldsMapSetTableFieldValue(games, f) {
		fmt.Println("CPPFieldsMapSetTableFieldValue failed")
		return -1
	}
	if 0 != CPPFieldsMapGetTableFieldValue(games, f) {
		fmt.Println("CPPFieldsMapGetTableFieldValue failed")
		return -1
	}

	// 6、输出结尾
	var end string = "\n}"
	_, err = f.WriteString(end)
	if err != nil {
		fmt.Println(err)
		return -1
	}
	return 0
}
