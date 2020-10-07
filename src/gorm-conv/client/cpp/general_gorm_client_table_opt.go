package cpp

import (
	"fmt"
	"gorm-conv/common"
	"os"
	"strings"
)

func GeneralClientCPPCodes_GeneralGormClientTableOpt_Table_SplitInfo(table common.TableInfo) (result string) {
	for idx, str := range table.SplitInfo.SplitCols {
		if idx != 0 {
			result += ", "
		}
		col := table.GetColumn(str)
		result += common.CPPField_CPPType(col.Type)
		result += " "
		result += str
	}
	return result
}

func GeneralClientCPPCodes_GeneralGormClientTableOpt_Table_Columns(table common.TableInfo, f *os.File) int {
	for _, col := range table.TableColumns {
		colType := common.CPPField_CPPType(col.Type)
		colStructName := common.CPP_TableColumnName(col.Name)
		getColFunc := "Get" + colStructName
		setColFunc := "Set" + colStructName
		if colType == "string" {
			f.WriteString("    string &" + getColFunc + "();\n")
			f.WriteString("    void " + setColFunc + "(string &" + col.Name + ");\n")
			f.WriteString("    void " + setColFunc + "(string &&" + col.Name + ");\n")
			f.WriteString("    void " + setColFunc + "(const char*" + col.Name + ");\n")
			f.WriteString("    void " + setColFunc + "(const char*" + col.Name + ", size_t size);\n")
		} else {
			f.WriteString("    " + colType + " " + getColFunc + "();\n")
			f.WriteString("    void " + setColFunc + "(" + colType + " " + col.Name + ");\n")
		}
	}

	return -1
}

func GeneralClientCPPCodes_GeneralGormClientTableOpt_Table(table common.TableInfo, f *os.File) int {
	structName := "GORM_ClientTable" + common.CPP_TableStruct(table.Name)
	pbStructName := "GORM_PB_Table_" + table.Name
	f.WriteString("class " + structName + "\n")
	f.WriteString("{\n")
	f.WriteString("public:\n")
	f.WriteString("    static " + structName + "* Get(" + GeneralClientCPPCodes_GeneralGormClientTableOpt_Table_SplitInfo(table) + ");\n")
	f.WriteString("static int Get(" + GeneralClientCPPCodes_GeneralGormClientTableOpt_Table_SplitInfo(table) + ", int64 &cbId, int (*cb)(int64));\n")
	f.WriteString("int Delete(int (*cb)(int64));\n")
	f.WriteString("void RemoveFromLocal();\n")
	f.WriteString(pbStructName + " *GetPbMsg();\n")

	// 设置每个字段的存取方法
	if 0 != GeneralClientCPPCodes_GeneralGormClientTableOpt_Table_Columns(table, f) {
		return -1
	}
	f.WriteString("private:\n")
	f.WriteString("    " + pbStructName + " *pTablePbValue;\n")
	f.WriteString("};\n")

	return 0
}

func GeneralClientCPPCodes_GeneralGormClientTableOpt(games []common.XmlCfg, outpath string) int {
	for _, game := range games {
		fileName := game.File[:len(game.File)-3]
		outfile := outpath + "/gorm_client_table_opt_" + fileName + ".h"
		f, err := os.OpenFile(outfile, os.O_RDWR|os.O_CREATE, 0666)
		_ = f
		if err != nil {
			fmt.Println(err.Error())
			return -1
		}

		bigFile := strings.ToUpper(fileName)
		f.WriteString("#ifndef _GORM_CLIENT_TABLE_OPT_" + bigFile + "_H__\n")
		f.WriteString("#define _GORM_CLIENT_TABLE_OPT_" + bigFile + "_H__\n")
		f.WriteString("#include \"" + fileName + ".pb.h\n")
		f.WriteString("\n")

		for _, table := range game.DB.TableList {
			if 0 != GeneralClientCPPCodes_GeneralGormClientTableOpt_Table(table, f) {
				fmt.Println("GeneralClientCPPCodes_GeneralGormClientTableOpt_Table failed, for table:", table.Name)
				return -1
			}
		}

	}
	return -1
}
