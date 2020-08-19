package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getGolangPBStructName(inName string) string {
	return "GORM_PB_Table" + getGolangPbFieldName(inName)
}

func getGolangPbFieldName(srcField string) string {
	var strList []string = strings.Split(srcField, "_")
	var result string
	for idx, s := range strList {
		if byte(s[0]) >= 'A' && byte(s[0]) <= 'Z' {
			if idx != 0 {
				result += "_"
			}
			result += s
		} else {
			result += strings.ToUpper(string(s[0])) + s[1:]
		}
	}
	return result
}

func CPPField_GolangType(colType string) string {
	switch colType {
	case "int8", "int16", "int32", "uint8", "uint16", "uint32", "int64", "uint64":
		{
			return colType
		}
	case "int":
		return "int32"
	case "double":
		return "floag64"
	default:
		{
			return "string"
		}
	}
	return "string"
}

func GeneralGolangCodes(games []XmlCfg, outpath string) int {
	outfile := outpath + "/gorm_table_field_map.go"
	f, err := os.OpenFile(outfile, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err.Error())
		return -1
	}

	fmt.Println("begin to general golang codes file:" + outfile)
	defer func() {
		f.Close()
		fmt.Println("end general golang codes file:" + outfile)
	}()
	err = f.Truncate(0)

	f.WriteString(`package gorm

import (
	"strings"

	"github.com/golang/protobuf/proto"
)

`)

	// 1.表名与ID映射
	f.WriteString("var g_tablename_to_tableid_map map[string]int32 = map[string]int32{\n")
	for _, game := range games {
		for _, table := range game.DB.TableList {
			f.WriteString("    \"" + table.Name + "\": int32(GORM_PB_TABLE_INDEX_GORM_PB_TABLE_IDX_" + strings.ToUpper(table.Name) + "),\n")
		}
	}
	f.WriteString("}\n")

	f.WriteString("var g_tableid_to_tablename_map map[int32]string = map[int32]string{\n")
	for _, game := range games {
		for _, table := range game.DB.TableList {
			f.WriteString("    int32(GORM_PB_TABLE_INDEX_GORM_PB_TABLE_IDX_" + strings.ToUpper(table.Name) + "): \"" + table.Name + "\",\n")
		}
	}
	f.WriteString("}\n")

	f.WriteString(`func GORM_GetTableId(strTable string) int32 {
	return g_tablename_to_tableid_map[strings.ToLower(strTable)]
}

func GORM_GetTableName(tableId int32) string {
	return g_tableid_to_tablename_map[tableId]
}
`)

	// 2.表对应的PB结构体
	f.WriteString("func GORM_GetTablePbMsgDefind(tableId int32) (msg proto.Message) {\n")
	f.WriteString("    switch GORM_PB_TABLE_INDEX(tableId) {\n")
	for _, game := range games {
		for _, table := range game.DB.TableList {
			f.WriteString("        case GORM_PB_TABLE_INDEX_GORM_PB_TABLE_IDX_" + strings.ToUpper(table.Name) + ":\n")
			f.WriteString("            return &" + getGolangPBStructName(table.Name) + "{}\n")
		}
	}
	f.WriteString("    }\n")
	f.WriteString("    return nil\n")
	f.WriteString("}\n")

	// 3.表中字段与ID的映射关系
	f.WriteString("var GORM_Table_FieldName_To_Id_Map []map[string]int32 = []map[string]int32{\n")
	var index int = -1
	for _, game := range games {
		for _, table := range game.DB.TableList {
			index += 1
			var bigTable string = strings.ToUpper(table.Name)
			f.WriteString("    " + strconv.FormatInt(int64(index), 10) + ": map[string]int32{\n")
			for _, col := range table.TableColumns {
				var bigColName string = strings.ToUpper(col.Name)
				var colValue string = "GORM_PB_" + bigTable + "_FIELD_INDEX_GORM_PB_FIELD_" + bigTable + "_" + bigColName
				f.WriteString("        \"" + col.Name + "\": int32(" + colValue + "),\n")
			}
			f.WriteString("    },\n")
		}
	}
	f.WriteString("}\n\n")
	f.WriteString(`func GORM_GetTableFieldId(tableId int32, fieldName string) (int32, GORM_CODE) {
	if tableId <= int32(GORM_PB_TABLE_INDEX_GORM_PB_TABLE_IDX_MIN__) || tableId >= int32(GORM_PB_TABLE_INDEX_GORM_PB_TABLE_IDX_MAX__) {
		return 0, GORM_CODE_INVALID_FIELD
	}
	fieldId, ok := GORM_Table_FieldName_To_Id_Map[tableId-1][strings.ToLower(fieldName)]
	if !ok {
		return 0, GORM_CODE_INVALID_FIELD
	}

	return fieldId, GORM_CODE_OK
}
`)

	// 4.GORM_SetTableFieldIntValue
	f.WriteString("func GORM_SetTableFieldIntValue(msg proto.Message, tableId int32, fieldId int32, value int64) GORM_CODE {\n")
	GORM_SetTableFieldValue(games, f, false)
	f.WriteString("}\n")
	// 5.GORM_SetTableFieldUintValue
	f.WriteString("func GORM_SetTableFieldUintValue(msg proto.Message, tableId int32, fieldId int32, value uint64) GORM_CODE {\n")
	GORM_SetTableFieldValue(games, f, false)
	f.WriteString("}\n")
	// 6.GORM_SetTableFieldDoubleValue
	f.WriteString("func GORM_SetTableFieldDoubleValue(msg proto.Message, tableId int32, fieldId int32, value float64) GORM_CODE {\n")
	GORM_SetTableFieldValue(games, f, false)
	f.WriteString("}\n")
	// 7.GORM_SetTableFieldStrValue
	f.WriteString("func GORM_SetTableFieldStrValue(msg proto.Message, tableId int32, fieldId int32, value string) GORM_CODE {\n")
	GORM_SetTableFieldValue(games, f, true)
	f.WriteString("}\n")

	// 8.GORM_GetFieldIntValueByID
	f.WriteString("func GORM_GetFieldIntValueByID(msg proto.Message, tableId int32, fieldId int32) (int64, GORM_CODE) {\n")
	GORM_GetTableFieldValue(games, f, false, "int64")
	f.WriteString("}\n")
	// 9.GORM_GetFieldUintValueByID
	f.WriteString("func GORM_GetFieldUintValueByID(msg proto.Message, tableId int32, fieldId int32) (uint64, GORM_CODE) {\n")
	GORM_GetTableFieldValue(games, f, false, "uint64")
	f.WriteString("}\n")
	// 10.GORM_GetFieldDoubleValueByID
	f.WriteString("func GORM_GetFieldDoubleValueByID(msg proto.Message, tableId int32, fieldId int32) (float64, GORM_CODE) {\n")
	GORM_GetTableFieldValue(games, f, false, "float64")
	f.WriteString("}\n")
	// 11.GORM_GetFieldStrValueByID
	f.WriteString("func GORM_GetFieldStrValueByID(msg proto.Message, tableId int32, fieldId int32) (string, GORM_CODE) {\n")
	GORM_GetTableFieldValue(games, f, true, "string")
	f.WriteString("}\n")

	// 12.GORM_AddRecordToReqPbMsgDefine
	f.WriteString("func GORM_AddRecordToReqPbMsgDefine(pbTable *GORM_PB_TABLE, recordMsg proto.Message, tableId int32) GORM_CODE {\n")
	f.WriteString("    switch GORM_PB_TABLE_INDEX(tableId) {\n")
	for _, game := range games {
		for _, table := range game.DB.TableList {
			f.WriteString("    case GORM_PB_TABLE_INDEX_GORM_PB_TABLE_IDX_" + strings.ToUpper(table.Name) + ":\n")
			f.WriteString("        pbTable." + getGolangPbFieldName(table.Name) + " = recordMsg.(*" + getGolangPBStructName(table.Name) + ")\n")
		}
	}
	f.WriteString("    }\n")
	f.WriteString("    return GORM_CODE_OK\n")
	f.WriteString("}\n")

	// 13.GORM_GetTableSrcPbMsg
	f.WriteString("func GORM_GetTableSrcPbMsg(tableId int32, pbTable *GORM_PB_TABLE) proto.Message {\n")
	f.WriteString("    switch GORM_PB_TABLE_INDEX(tableId) {\n")
	for _, game := range games {
		for _, table := range game.DB.TableList {
			f.WriteString("    case GORM_PB_TABLE_INDEX_GORM_PB_TABLE_IDX_" + strings.ToUpper(table.Name) + ":\n")
			f.WriteString("        return pbTable." + getGolangPbFieldName(table.Name) + "\n")
		}
	}
	f.WriteString("    }\n")
	f.WriteString("    return nil\n")
	f.WriteString("}\n")

	return 0
}

func GORM_SetTableFieldValue(games []XmlCfg, f *os.File, bStr bool) {
	f.WriteString("    switch GORM_PB_TABLE_INDEX(tableId) {\n")
	for _, game := range games {
		for _, table := range game.DB.TableList {
			var tablePbName string = getGolangPBStructName(table.Name)
			var bigTable string = strings.ToUpper(table.Name)
			f.WriteString("    case GORM_PB_TABLE_INDEX_GORM_PB_TABLE_IDX_" + bigTable + ":\n")
			f.WriteString("        var " + table.Name + " *" + tablePbName + " =  msg.(*" + tablePbName + ")\n")
			f.WriteString("        switch GORM_PB_" + bigTable + "_FIELD_INDEX(fieldId) {\n")
			for _, col := range table.TableColumns {
				var colType string = CPPField_GolangType(col.Type)
				if bStr {
					if colType != "string" {
						continue
					}
				} else {
					if colType == "string" {
						continue
					}
				}

				f.WriteString("        case GORM_PB_" + bigTable + "_FIELD_INDEX_GORM_PB_FIELD_" + bigTable + "_" + strings.ToUpper(col.Name) + ":\n")
				f.WriteString("            " + table.Name + "." + getGolangPbFieldName(col.Name) + " = " + colType + "(value)\n")
				f.WriteString("            return GORM_CODE_OK\n")
			}
			f.WriteString("        }\n")
		}
		f.WriteString("    }\n")
	}
	f.WriteString("    return GORM_CODE_INVALID_FIELD\n")
}

func GORM_GetTableFieldValue(games []XmlCfg, f *os.File, bStr bool, inType string) {
	f.WriteString("    switch GORM_PB_TABLE_INDEX(tableId) {\n")
	for _, game := range games {
		for _, table := range game.DB.TableList {
			var tablePbName string = getGolangPBStructName(table.Name)
			var bigTable string = strings.ToUpper(table.Name)
			f.WriteString("    case GORM_PB_TABLE_INDEX_GORM_PB_TABLE_IDX_" + bigTable + ":\n")
			f.WriteString("        var " + table.Name + " *" + tablePbName + " =  msg.(*" + tablePbName + ")\n")
			f.WriteString("        switch GORM_PB_" + bigTable + "_FIELD_INDEX(fieldId) {\n")
			for _, col := range table.TableColumns {
				var colType string = CPPField_GolangType(col.Type)
				if bStr {
					if colType != "string" {
						continue
					}
				} else {
					if colType == "string" {
						continue
					}
				}

				f.WriteString("        case GORM_PB_" + bigTable + "_FIELD_INDEX_GORM_PB_FIELD_" + bigTable + "_" + strings.ToUpper(col.Name) + ":\n")
				f.WriteString("            return " + inType + "(" + table.Name + "." + getGolangPbFieldName(col.Name) + "), GORM_CODE_OK\n")
			}
			f.WriteString("        }\n")
		}
		f.WriteString("    }\n")
	}
	if !bStr {
		f.WriteString("    return  0, GORM_CODE_INVALID_FIELD\n")
	} else {
		f.WriteString("    return \"\", GORM_CODE_INVALID_FIELD\n")
	}
}
