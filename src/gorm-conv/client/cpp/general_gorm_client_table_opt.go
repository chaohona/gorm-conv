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

func GeneralClientCPPCodes_GeneralGormClientTableOpt_Table_H_Columns_Desc(table common.TableInfo, f *os.File) int {
	for _, col := range table.TableColumns {
		colType := common.CPPField_CPPType(col.Type)
		colStructName := common.CPP_TableColumnName(col.Name)
		getColFunc := "Get" + colStructName
		setColFunc := "Set" + colStructName
		if colType == "string" {
			f.WriteString("    const string &" + getColFunc + "() const;\n")
			f.WriteString("    void " + setColFunc + "(const string &" + col.Name + ", bool forceSave=false);\n")
			f.WriteString("    void " + setColFunc + "(const char* " + col.Name + ", size_t size, bool forceSave=false);\n")
		} else {
			f.WriteString("    " + colType + " " + getColFunc + "();\n")
			f.WriteString("    void " + setColFunc + "(" + colType + " " + col.Name + ", bool forceSave=false);\n")
		}
	}

	return 0
}

// 头文件中的inline函数实现
func GeneralClientCPPCodes_GeneralGormClientTableOpt_Table_H_Columns_Define(table common.TableInfo, f *os.File) int {
	pbStructName := "GORM_PB_Table_" + table.Name
	structName := "GORM_ClientTable" + common.CPP_TableStruct(table.Name)
	f.WriteString("inline " + pbStructName + "*" + structName + "::GetPbMsg()\n")
	f.WriteString("{\n")
	f.WriteString("    return this->pTablePbValue;\n")
	f.WriteString("}\n")
	// 生成Get函数
	for _, col := range table.TableColumns {
		colType := common.CPPField_CPPType(col.Type)
		colStructName := common.CPP_TableColumnName(col.Name)
		getColFunc := "Get" + colStructName
		if colType == "string" {
			f.WriteString("inline const string& " + structName + "::" + getColFunc + "() const\n")
		} else {
			f.WriteString("inline " + colType + " " + structName + "::" + getColFunc + "()\n")
		}
		f.WriteString("{\n")
		f.WriteString("    return this->pTablePbValue->" + col.Name + "();\n")
		f.WriteString("}\n")
	}
	// 生成更新所有字段函数
	f.WriteString("inline void " + structName + "::SetPbMsg(" + pbStructName + " *pbMsg, bool forceSave)\n")
	f.WriteString("{\n")
	for _, col := range table.TableColumns {
		colStructName := common.CPP_TableColumnName(col.Name)
		setColFunc := "Set" + colStructName
		f.WriteString("    this->" + setColFunc + "(pbMsg->" + col.Name + "(), forceSave);\n")
	}
	f.WriteString("    return;\n")
	f.WriteString("}\n")
	// 生成set函数
	for _, col := range table.TableColumns {
		colType := common.CPPField_CPPType(col.Type)
		colStructName := common.CPP_TableColumnName(col.Name)
		setColFunc := "Set" + colStructName
		if colType == "string" {
			// Set(string&)
			f.WriteString("inline void " + structName + "::" + setColFunc + "(const string &" + col.Name + ", bool forceSave)\n")
			f.WriteString("{\n")
			f.WriteString("    this->pTablePbValue->set_" + col.Name + "(" + col.Name + ");\n")
			f.WriteString("    return;\n")
			f.WriteString("}\n")

			//Set(const char*, size_t)
			f.WriteString("inline void " + structName + "::" + setColFunc + "(const char* " + col.Name + ", size_t size, bool forceSave)\n")
			f.WriteString("{\n")
			f.WriteString("    this->pTablePbValue->set_" + col.Name + "(" + col.Name + ", size);\n")
			f.WriteString("    return;\n")
			f.WriteString("}\n")
		} else {
			f.WriteString("inline void " + structName + "::" + setColFunc + "(" + colType + " " + col.Name + ", bool forceSave)\n")
			f.WriteString("{\n")
			f.WriteString("    this->pTablePbValue->set_" + col.Name + "(" + col.Name + ");\n")
			f.WriteString("    return;\n")
			f.WriteString("}\n")
		}
	}
	return 0
}

func GeneralClientCPPCodes_GeneralGormClientTableOpt_Table_H(table common.TableInfo, f *os.File) int {
	// 表对应的类的声明
	structName := "GORM_ClientTable" + common.CPP_TableStruct(table.Name)
	pbStructName := "GORM_PB_Table_" + table.Name
	f.WriteString("class " + structName + "\n")
	f.WriteString("{\n")
	f.WriteString("public:\n")
	f.WriteString("    static " + structName + "* Get(int region, int logic_zone, int physics_zone, " + GeneralClientCPPCodes_GeneralGormClientTableOpt_Table_SplitInfo(table) + ");\n")
	f.WriteString("    static int Get(int region, int logic_zone, int physics_zone, int64 &cbId, " + GeneralClientCPPCodes_GeneralGormClientTableOpt_Table_SplitInfo(table) + ", int (*cb)(int64, " + structName + "*));\n")
	f.WriteString("    static int Delete(int region, int logic_zone, int physics_zone, int64 &cbId, int (*cb)(int64), " + GeneralClientCPPCodes_GeneralGormClientTableOpt_Table_SplitInfo(table) + ");\n")
	f.WriteString("    static void SetPbMsg(int region, int logic_zone, int physics_zone, " + pbStructName + " *pbMsg, bool forceSave=false);\n")
	f.WriteString("    int Delete(int (*cb)(int64));\n")
	f.WriteString("    void SetPbMsg(" + pbStructName + " *pbMsg, bool forceSave=false);\n")
	//f.WriteString("    void RemoveFromLocal();\n")
	f.WriteString("    int SaveToDB();\n")
	f.WriteString("    " + pbStructName + " *GetPbMsg();\n")

	// 声明每个字段的存取方法
	if 0 != GeneralClientCPPCodes_GeneralGormClientTableOpt_Table_H_Columns_Desc(table, f) {
		return -1
	}
	// 其它变量
	f.WriteString("private:\n")
	f.WriteString("    " + pbStructName + " *pTablePbValue = nullptr;\n")
	f.WriteString("};\n\n")

	return 0
}

func GeneralClientCPPCodes_GeneralGormClientTableOpt_CPP_Table(table common.TableInfo, f *os.File) int {
	structName := "GORM_ClientTable" + common.CPP_TableStruct(table.Name)
	pbStructName := "GORM_PB_Table_" + table.Name
	// 同步Get函数
	f.WriteString(structName + "* " + structName + "::Get(int region, int logic_zone, int physics_zone," + GeneralClientCPPCodes_GeneralGormClientTableOpt_Table_SplitInfo(table) + ")\n")
	f.WriteString("{\n")
	f.WriteString("    return nullptr;\n")
	f.WriteString("}\n")

	// 异步Get函数
	f.WriteString(structName + "* " + structName + "::Get(int region, int logic_zone, int physics_zone, int64 &cbId, " + GeneralClientCPPCodes_GeneralGormClientTableOpt_Table_SplitInfo(table) + ", int (*cb)(int64, " + structName + "*))\n")
	f.WriteString("{\n")
	f.WriteString("    return nullptr;\n")
	f.WriteString("}\n")

	// static Delete函数
	f.WriteString("int " + structName + "::Delete(int region, int logic_zone, int physics_zone, int64 &cbId, int (*cb)(int64), " + GeneralClientCPPCodes_GeneralGormClientTableOpt_Table_SplitInfo(table) + ")\n")
	f.WriteString("{\n")
	f.WriteString("    return 0;\n")
	f.WriteString("}\n")

	// SetPbMsg函数
	f.WriteString("int " + structName + "::SetPbMsg(int region, int logic_zone, int physics_zone, " + pbStructName + " *pbMsg, bool forceSave)\n")
	f.WriteString("{\n")
	f.WriteString("    return 0;\n")
	f.WriteString("}\n")

	// Delete函数
	f.WriteString("int " + structName + "::Delete(" + GeneralClientCPPCodes_GeneralGormClientTableOpt_Table_SplitInfo(table) + ")\n")
	f.WriteString("{\n")
	f.WriteString("    return 0;\n")
	f.WriteString("}\n")
	return 0
}

// 源文件实现
func GeneralClientCPPCodes_GeneralGormClientTableOpt_CPP(game common.XmlCfg, outpath string) int {
	fileName := game.File[:len(game.File)-4]
	outfile := outpath + "/gorm_client_table_opt_" + fileName + ".cc"
	f, err := os.OpenFile(outfile, os.O_RDWR|os.O_CREATE, 0666)
	_ = f
	if err != nil {
		fmt.Println(err.Error())
		return -1
	}

	headerFile := "/gorm_client_table_opt_" + fileName + ".h"
	f.WriteString("#include " + headerFile + "\n\n")
	f.WriteString("namespace gorm{\n\n")

	for _, table := range game.DB.TableList {
		// 生成响应的表的实现函数
		if 0 != GeneralClientCPPCodes_GeneralGormClientTableOpt_CPP_Table(table, f) {
			return -1
		}
	}
	f.WriteString("\n\n}\n")

	return 0
}

// 头文件实现
func GeneralClientCPPCodes_GeneralGormClientTableOpt_H(game common.XmlCfg, outpath string) int {
	fileName := game.File[:len(game.File)-4]
	outfile := outpath + "/gorm_client_table_opt_" + fileName + ".h"
	fmt.Println(outfile)
	f, err := os.OpenFile(outfile, os.O_RDWR|os.O_CREATE, 0666)
	_ = f
	if err != nil {
		fmt.Println("create file failed:", outfile)
		fmt.Println(err.Error())
		return -1
	}
	f.Truncate(0)

	bigName := strings.ToUpper(fileName)
	_, err = f.WriteString("#ifndef _GORM_CLIENT_TABLE_OPT_" + bigName + "_H__\n")
	fmt.Println(err)
	f.WriteString("#define _GORM_CLIENT_TABLE_OPT_" + bigName + "_H__\n")
	f.WriteString("#include \"" + fileName + ".pb.h\"\n")
	f.WriteString("\n")
	f.WriteString("namespace gorm{\n\n")

	for _, table := range game.DB.TableList {
		if 0 != GeneralClientCPPCodes_GeneralGormClientTableOpt_Table_H(table, f) {
			fmt.Println("GeneralClientCPPCodes_GeneralGormClientTableOpt_Table_H failed, for table:", table.Name)
			return -1
		}
	}
	for _, table := range game.DB.TableList {
		// 头文件中实现的inline函数
		if 0 != GeneralClientCPPCodes_GeneralGormClientTableOpt_Table_H_Columns_Define(table, f) {
			return -1
		}
	}

	// 结尾
	// namespace gorm 结尾
	f.WriteString("\n}\n\n")
	// define 结尾
	f.WriteString("\n\n#endif")
	f.Close()
	return 0
}

func GeneralClientCPPCodes_GeneralGormClientTableOpt(games []common.XmlCfg, outpath string) int {
	for _, game := range games {
		// 头文件
		if 0 != GeneralClientCPPCodes_GeneralGormClientTableOpt_H(game, outpath) {
			return -1
		}

		// 源文件
		if 0 != GeneralClientCPPCodes_GeneralGormClientTableOpt_CPP(game, outpath) {
			return -1
		}
	}
	return 0
}
