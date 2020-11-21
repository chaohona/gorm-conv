package cpp

import (
	"fmt"
	"gorm-conv/common"
	"os"
	"strconv"
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

func GeneralClientCPPCodes_GeneralGormClientTableOpt_Table_TableIndex(table common.TableInfo, tableIndex common.TableIndex) (result string) {
	for idx, str := range tableIndex.IndexColumns {
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

func GeneralClientCPPCodes_GeneralGormClientTableOpt_Table_SplitInfoName(table common.TableInfo) (result string) {
	for idx, str := range table.SplitInfo.SplitCols {
		if idx != 0 {
			result += ", "
		}
		result += str
	}
	return result
}

func GeneralClientCPPCodes_GeneralGormClientTableOpt_Table_SplitInfo_Param(table common.TableInfo) (result string) {
	for idx, str := range table.SplitInfo.SplitCols {
		if idx != 0 {
			result += ", "
		}
		result += str
	}
	return result
}

func GeneralClientCPPCodes_GeneralGormClientTableOpt_Table_TableIndex_Param(tableIndex common.TableIndex) (result string) {
	for idx, str := range tableIndex.IndexColumns {
		if idx != 0 {
			result += ", "
		}
		result += str
	}
	return result
}

func GeneralClientCPPCodes_GeneralGormClientTableOpt_Table_H_Columns_Desc(table common.TableInfo, f *os.File) int {
	f.WriteString("    /* 下面为针对表的每个字段的操作,Get为获取字段的原始数据,Set为更新字段的值 */\n")
	for _, col := range table.TableColumns {
		if col.PrimaryKey {
			f.WriteString("    // 此字段为路由项，设置之后不能更改\n")
		}
		f.WriteString("    // " + col.Name + "\n")
		colType := common.CPPField_CPPType(col.Type)
		colStructName := common.CPP_TableColumnName(col.Name)
		getColFunc := "Get" + colStructName
		setColFunc := "Set" + colStructName
		if colType == "string" {
			f.WriteString("    const string &" + getColFunc + "();\n")
			f.WriteString("    int " + setColFunc + "(const string &" + col.Name + ", bool forceSave=false);\n")
			f.WriteString("    int " + setColFunc + "(const char* " + col.Name + ", size_t size, bool forceSave=false);\n")
		} else {
			f.WriteString("    " + colType + " " + getColFunc + "();\n")
			f.WriteString("    int " + setColFunc + "(" + colType + " " + col.Name + ", bool forceSave=false);\n")
		}
	}

	return 0
}

func setFiledMode(index int) string {
	var idx int = index / 8
	var strIdx string = strconv.FormatInt(int64(idx), 10)
	_ = strIdx
	var mode int = 1 << (index - idx*8)
	var strMode string = strconv.FormatInt(int64(mode), 10)
	_ = strMode
	return "    this->fieldOpt.AddField(" + strIdx + ", " + strMode + ");\n"
}

func CheckAndNewPb(pbStructName string) string {
	return string("    if (this->tablePbValue == nullptr) this->tablePbValue = new " + pbStructName + "();\n")
}

// 头文件中的inline函数实现
func GeneralClientCPPCodes_GeneralGormClientTableOpt_Table_H_Columns_Define(table common.TableInfo, f *os.File) int {
	pbStructName := "GORM_PB_Table_" + table.Name
	structName := "GORM_ClientTable" + common.CPP_TableStruct(table.Name)

	//////////////////////////////////////// 带区服的接口
	for _, tableIndex := range table.TableIndex {
		f.WriteString("inline shared_ptr<" + structName + "> " + structName + "::GetBy" + tableIndex.Name + "(int region, int logic_zone, int physics_zone, int &retCode, " + GeneralClientCPPCodes_GeneralGormClientTableOpt_Table_TableIndex(table, tableIndex) + ")\n")
		f.WriteString("{\n")
		f.WriteString("    shared_ptr<" + structName + "> table = make_shared<" + structName + ">();\n")
		f.WriteString("    uint32 cbId = 0;\n")
		f.WriteString("    retCode = table->DoGetBy" + tableIndex.Name + "(cbId, " + GeneralClientCPPCodes_GeneralGormClientTableOpt_Table_TableIndex_Param(tableIndex) + ");\n")
		f.WriteString("    if (GORM_OK != retCode)\n")
		f.WriteString("    {\n")
		f.WriteString("        if (retCode == GORM_NO_MORE_RECORD) retCode = GORM_OK;")
		f.WriteString("        return nullptr;\n")
		f.WriteString("    }\n")
		f.WriteString("    return table;\n")
		f.WriteString("}\n")
	}
	// static 同步Get函数
	f.WriteString("inline shared_ptr<" + structName + "> " + structName + "::Get(int region, int logic_zone, int physics_zone, int &retCode, " + GeneralClientCPPCodes_GeneralGormClientTableOpt_Table_SplitInfo(table) + ")\n")
	f.WriteString("{\n")
	f.WriteString("    shared_ptr<" + structName + "> table = make_shared<" + structName + ">();\n")
	f.WriteString("    uint32 cbId = 0;\n")
	f.WriteString("    retCode = table->DoGet(cbId, " + GeneralClientCPPCodes_GeneralGormClientTableOpt_Table_SplitInfo_Param(table) + ");\n")
	f.WriteString("    if (GORM_OK != retCode)\n")
	f.WriteString("    {\n")
	f.WriteString("        if (retCode == GORM_NO_MORE_RECORD) retCode = GORM_OK;")
	f.WriteString("        return nullptr;\n")
	f.WriteString("    }\n")
	f.WriteString("    return table;\n")
	f.WriteString("}\n")

	// static 异步Get函数
	f.WriteString("inline int " + structName + "::Get(int region, int logic_zone, int physics_zone, " + GeneralClientCPPCodes_GeneralGormClientTableOpt_Table_SplitInfo(table) + ", uint32 &cbId, GORM_CbFun cbFunc)\n")
	f.WriteString("{\n")
	f.WriteString("    shared_ptr<" + structName + "> table = make_shared<" + structName + ">();\n")
	f.WriteString("    int retCode = table->DoGet(cbId, " + GeneralClientCPPCodes_GeneralGormClientTableOpt_Table_SplitInfo_Param(table) + ");\n")
	f.WriteString("    if (GORM_OK != retCode)\n")
	f.WriteString("    {\n")
	f.WriteString("        if (retCode == GORM_NO_MORE_RECORD) retCode = GORM_OK;")
	f.WriteString("        return retCode;\n")
	f.WriteString("    }\n")
	f.WriteString("    if (cbFunc != nullptr) cbFunc(cbId, retCode, table);\n")
	f.WriteString("    return GORM_OK;\n")
	f.WriteString("}\n")

	// static Delete函数
	f.WriteString("inline int " + structName + "::Delete(int region, int logic_zone, int physics_zone, " + GeneralClientCPPCodes_GeneralGormClientTableOpt_Table_SplitInfo(table) + ", uint32 &cbId, GORM_CbFun cbFunc)\n")
	f.WriteString("{\n")
	f.WriteString("    shared_ptr<" + structName + "> table = make_shared<" + structName + ">();\n")
	f.WriteString("    shared_ptr<" + pbStructName + "> pbTable = make_shared<" + pbStructName + ">();\n")
	f.WriteString("    table->tablePbValue = pbTable.get();\n")
	for _, c := range table.SplitInfo.SplitCols {
		col := table.GetColumn(c)
		colStructName := common.CPP_TableColumnName(col.Name)
		setColFunc := "Set" + colStructName
		f.WriteString("    table->" + setColFunc + "(" + col.Name + ");\n")
	}
	f.WriteString("    return table->Delete(cbId, cbFunc);\n")
	f.WriteString("}\n")

	// static SetPbMsg函数
	f.WriteString("inline int " + structName + "::SetPbMsg(int region, int logic_zone, int physics_zone, " + pbStructName + " *pbMsg, bool forceSave)\n")
	f.WriteString("{\n")
	f.WriteString("    " + structName + " table;\n")
	f.WriteString("    return table.SetPbMsg(pbMsg, forceSave);\n")
	f.WriteString("}\n")
	//////////////////////////////////////// 带区服的接口

	/////////////////////////////////////////// 不带区服的接口
	// static 不带区服的同步Get函数
	for _, tableIndex := range table.TableIndex {
		f.WriteString("inline int " + structName + "::GetVectorBy" + tableIndex.Name + "(vector<shared_ptr<" + structName + ">> &outResult, " + GeneralClientCPPCodes_GeneralGormClientTableOpt_Table_TableIndex(table, tableIndex) + ")\n")
		f.WriteString("{\n")
		f.WriteString("    return " + structName + "::GetVectorBy" + tableIndex.Name + "(0,0,0, outResult, " + GeneralClientCPPCodes_GeneralGormClientTableOpt_Table_TableIndex_Param(tableIndex) + ");\n")
		f.WriteString("}\n")

		f.WriteString("inline shared_ptr<" + structName + "> " + structName + "::GetBy" + tableIndex.Name + "(int &retCode, " + GeneralClientCPPCodes_GeneralGormClientTableOpt_Table_TableIndex(table, tableIndex) + ")\n")
		f.WriteString("{\n")
		f.WriteString("    return " + structName + "::GetBy" + tableIndex.Name + "(0,0,0, retCode, " + GeneralClientCPPCodes_GeneralGormClientTableOpt_Table_TableIndex_Param(tableIndex) + ");\n")
		f.WriteString("}\n")
	}
	f.WriteString("inline int " + structName + "::GetVector(vector<shared_ptr<" + structName + ">> &outResult, " + GeneralClientCPPCodes_GeneralGormClientTableOpt_Table_SplitInfo(table) + ")\n")
	f.WriteString("{\n")
	f.WriteString("    return " + structName + "::GetVector(0,0,0, outResult, " + GeneralClientCPPCodes_GeneralGormClientTableOpt_Table_SplitInfo_Param(table) + ");\n")
	f.WriteString("}\n")
	f.WriteString("inline shared_ptr<" + structName + "> " + structName + "::Get(int &retCode, " + GeneralClientCPPCodes_GeneralGormClientTableOpt_Table_SplitInfo(table) + ")\n")
	f.WriteString("{\n")
	f.WriteString("    return " + structName + "::Get(0,0,0, retCode, " + GeneralClientCPPCodes_GeneralGormClientTableOpt_Table_SplitInfo_Param(table) + ");\n")
	f.WriteString("}\n")

	// static 不带区服，异步Get函数
	f.WriteString("inline int " + structName + "::Get(" + GeneralClientCPPCodes_GeneralGormClientTableOpt_Table_SplitInfo(table) + ", uint32 &cbId, GORM_CbFun cbFunc)\n")
	f.WriteString("{\n")
	f.WriteString("    return " + structName + "::Get(0,0,0, " + GeneralClientCPPCodes_GeneralGormClientTableOpt_Table_SplitInfo_Param(table) + ", cbId, cbFunc);\n")
	f.WriteString("}\n")

	// static Delete函数
	f.WriteString("inline int " + structName + "::Delete(" + GeneralClientCPPCodes_GeneralGormClientTableOpt_Table_SplitInfo(table) + ", uint32 &cbId, GORM_CbFun cbFunc)\n")
	f.WriteString("{\n")
	f.WriteString("    return " + structName + "::Delete(0,0,0," + GeneralClientCPPCodes_GeneralGormClientTableOpt_Table_SplitInfo_Param(table) + ", cbId, cbFunc);\n")
	f.WriteString("}\n")

	/////////////////////////////////////////// 不带区服的接口

	/////////////////////////////////// 操作对象的接口
	// Delete函数
	f.WriteString("inline int " + structName + "::Delete(uint32 &cbId, GORM_CbFun cbFunc)\n")
	f.WriteString("{\n")
	f.WriteString("    int retCode = this->DoDelete(cbId);\n")
	f.WriteString(`
    if (cbFunc != nullptr)
    {
        cbFunc(cbId, retCode, nullptr);
    }
    
    return GORM_OK;
`)
	f.WriteString("}\n")

	// SaveToDB函数
	f.WriteString("inline int " + structName + "::SaveToDB()\n")
	f.WriteString("{\n")
	f.WriteString("    int flag = this->dirtyFlag;\n")
	f.WriteString("    this->dirtyFlag = 1;\n")
	f.WriteString("    uint32 cbId;\n")
	f.WriteString("    if (flag == 2)\n")
	f.WriteString("        return this->DoUpdate(cbId);\n")
	f.WriteString("    else if (flag == 0)\n")
	f.WriteString("        return this->DoUpSert(cbId);\n")
	f.WriteString("    return 0;\n")
	f.WriteString("}\n")

	f.WriteString("inline int " + structName + "::Add()\n")
	f.WriteString("{\n")
	f.WriteString("    uint32 cbId;\n")
	f.WriteString("    return this->DoInsert(cbId);\n")
	f.WriteString("}\n")

	f.WriteString("inline int " + structName + "::Update()\n")
	f.WriteString("{\n")
	f.WriteString("    uint32 cbId;\n")
	f.WriteString("    return this->DoUpdate(cbId);\n")
	f.WriteString("}\n")

	f.WriteString("inline " + pbStructName + "* " + structName + "::GetPbMsg()\n")
	f.WriteString("{\n")
	f.WriteString("    if (this->tablePbValue == nullptr) this->tablePbValue = new " + pbStructName + "();\n")
	f.WriteString("    return this->tablePbValue;\n")
	f.WriteString("}\n")
	f.WriteString("inline " + pbStructName + "* " + structName + "::ReleasePbMsg() // 将pb消息返回，并将本地数据置为空(转移数据属主)\n")
	f.WriteString("{\n")
	f.WriteString("    auto *pbMsg = this->tablePbValue;\n")
	f.WriteString("    this->tablePbValue = nullptr;\n")
	f.WriteString("    this->fieldOpt.Reset();\n")
	f.WriteString("    return pbMsg;\n")
	f.WriteString("}\n")
	// 生成Get函数
	for _, col := range table.TableColumns {
		colType := common.CPPField_CPPType(col.Type)
		colStructName := common.CPP_TableColumnName(col.Name)
		getColFunc := "Get" + colStructName
		if colType == "string" {
			f.WriteString("inline const string& " + structName + "::" + getColFunc + "()\n")
			f.WriteString("{\n")
			f.WriteString("    if (this->tablePbValue == nullptr) new " + pbStructName + "();\n")
			f.WriteString("    return this->tablePbValue->" + col.Name + "();\n")
			f.WriteString("}\n")
		} else {
			f.WriteString("inline " + colType + " " + structName + "::" + getColFunc + "()\n")
			f.WriteString("{\n")
			f.WriteString("    if (this->tablePbValue == nullptr) return 0;\n")
			f.WriteString("    return this->tablePbValue->" + col.Name + "();\n")
			f.WriteString("}\n")
		}
	}
	// 生成更新所有字段函数
	f.WriteString("inline int " + structName + "::SetPbMsg(" + pbStructName + " *pbMsg, bool forceSave)\n")
	f.WriteString("{\n")
	f.WriteString(CheckAndNewPb(pbStructName))
	for _, col := range table.TableColumns {
		colStructName := common.CPP_TableColumnName(col.Name)
		setColFunc := "Set" + colStructName
		f.WriteString("    this->" + setColFunc + "(pbMsg->" + col.Name + "());\n")
	}
	f.WriteString("    if (forceSave)\n")
	f.WriteString("        return this->SaveToDB();\n")
	f.WriteString("    return 0;\n")
	f.WriteString("}\n")
	// 生成set函数
	for idx, col := range table.TableColumns {
		colType := common.CPPField_CPPType(col.Type)
		colStructName := common.CPP_TableColumnName(col.Name)
		setColFunc := "Set" + colStructName
		if colType == "string" {
			// Set(string&)
			f.WriteString("inline int " + structName + "::" + setColFunc + "(const string &" + col.Name + ", bool forceSave)\n")
			f.WriteString("{\n")
			f.WriteString(CheckAndNewPb(pbStructName))
			f.WriteString("    this->tablePbValue->set_" + col.Name + "(" + col.Name + ");\n")
			f.WriteString(setFiledMode(idx))
			f.WriteString("    if (this->dirtyFlag != 0) this->dirtyFlag = 2;\n")
			f.WriteString("    if (forceSave)\n")
			f.WriteString("        return this->SaveToDB();\n")
			f.WriteString("    return 0;\n")
			f.WriteString("}\n")

			//Set(const char*, size_t)
			f.WriteString("inline int " + structName + "::" + setColFunc + "(const char* " + col.Name + ", size_t size, bool forceSave)\n")
			f.WriteString("{\n")
			f.WriteString(CheckAndNewPb(pbStructName))
			f.WriteString("    this->tablePbValue->set_" + col.Name + "(" + col.Name + ", size);\n")
			f.WriteString(setFiledMode(idx))
			f.WriteString("    if (this->dirtyFlag != 0) this->dirtyFlag = 2;\n")
			f.WriteString("    if (forceSave)\n")
			f.WriteString("        return this->SaveToDB();\n")
			f.WriteString("    return 0;\n")
			f.WriteString("}\n")
		} else {
			f.WriteString("inline int " + structName + "::" + setColFunc + "(" + colType + " " + col.Name + ", bool forceSave)\n")
			f.WriteString("{\n")
			f.WriteString(CheckAndNewPb(pbStructName))
			f.WriteString("    this->tablePbValue->set_" + col.Name + "(" + col.Name + ");\n")
			f.WriteString(setFiledMode(idx))
			f.WriteString("    if (this->dirtyFlag != 0) this->dirtyFlag = 2;\n")
			f.WriteString("    if (forceSave)\n")
			f.WriteString("        return this->SaveToDB();\n")
			f.WriteString("    return 0;\n")
			f.WriteString("}\n")
		}
	}

	// HasSetPrimaryKey
	f.WriteString("inline bool " + structName + "::HasSetPrimaryKey()\n")
	f.WriteString("{\n")
	// return ((fieldOpt.szFieldCollections[0] & 1) >0) && ((fieldOpt.szFieldCollections[1] & 1) >0);
	f.WriteString("    return ")
	for nameIdx, cname := range table.SplitInfo.SplitCols {
		for idx, c := range table.TableColumns {
			if cname != c.Name {
				continue
			}
			var index int = idx >> 3
			var value int = 1 << (idx & 0x07)
			var strIndex string = strconv.FormatInt(int64(index), 10)
			var strValue string = strconv.FormatInt(int64(value), 10)
			if nameIdx != 0 {
				f.WriteString(" && ")
			}
			f.WriteString("((fieldOpt.szFieldCollections[" + strIndex + "]&" + strValue + ")>0)")
		}
	}
	f.WriteString(";\n")
	f.WriteString("}\n")

	// 函数GetCallBack
	f.WriteString("inline void " + structName + "::GetCallBack(GORM_ClientMsg *clientMsg)\n")
	f.WriteString("{\n")
	f.WriteString("    if (clientMsg == nullptr || clientMsg->cbFunc == nullptr) return;\n")
	f.WriteString("    GORM_PB_GET_RSP *pbRspMsg = dynamic_cast<GORM_PB_GET_RSP*>(clientMsg->pbRspMsg);\n")
	f.WriteString("    if (pbRspMsg == nullptr || !pbRspMsg->table().has_" + table.Name + "())\n")
	f.WriteString("    {\n")
	f.WriteString("        clientMsg->cbFunc(clientMsg->cbId, clientMsg->rspCode.code, nullptr);\n")
	f.WriteString("        return;\n")
	f.WriteString("    }\n")
	f.WriteString("    shared_ptr<" + structName + "> table = make_shared<" + structName + ">(pbRspMsg->mutable_table()->release_" + table.Name + "());\n")
	f.WriteString("    clientMsg->cbFunc(clientMsg->cbId, clientMsg->rspCode.code, table);\n")
	f.WriteString("    return;\n")
	f.WriteString("}\n")
	return 0
}

func GeneralClientCPPCodes_GeneralGormClientTableOpt_Table_H(table common.TableInfo, f *os.File) int {
	// 表对应的类的声明
	structName := "GORM_ClientTable" + common.CPP_TableStruct(table.Name)
	pbStructName := "GORM_PB_Table_" + table.Name

	f.WriteString("// 表" + table.Name + "\n")
	f.WriteString("class " + structName + ": public GORM_ClientTable\n")
	f.WriteString("{\n")
	f.WriteString("public:\n")
	f.WriteString("    " + structName + "(){}\n")
	f.WriteString("    " + structName + "(int region, int logic_zone, int physics_zone){}\n")
	f.WriteString("    " + structName + "(" + pbStructName + " *tablePbValue){this->SetPbMsg(tablePbValue);}\n")
	f.WriteString("    " + structName + "(int region, int logic_zone, int physics_zone, " + pbStructName + " *tablePbValue){this->SetPbMsg(tablePbValue);}\n")
	f.WriteString("    ~" + structName + "()\n")
	f.WriteString("    {\n")
	f.WriteString("        if (this->tablePbValue != nullptr)\n")
	f.WriteString("        {\n")
	f.WriteString("            delete this->tablePbValue;\n")
	f.WriteString("            this->tablePbValue = nullptr;\n")
	f.WriteString("        }\n")
	f.WriteString("    }\n")
	f.WriteString("public:\n")
	f.WriteString("    // 取出分表中的所有数据，暂时最多只支持取出GORM_MAX_LIMIT_NUM条数据\n")
	f.WriteString("    static int GetAllRows(vector<shared_ptr<" + structName + ">> &outResult, int tableIndex=0);\n")
	f.WriteString("    // static带区服的接口，用于分区分服架构\n")
	// 根据index取数据的接口
	for _, tableIndex := range table.TableIndex {
		f.WriteString("    static shared_ptr<" + structName + "> GetBy" + tableIndex.Name + "(int region, int logic_zone, int physics_zone, int &retCode, " + GeneralClientCPPCodes_GeneralGormClientTableOpt_Table_TableIndex(table, tableIndex) + ");\n")
		f.WriteString("    static int GetVectorBy" + tableIndex.Name + "(int region, int logic_zone, int physics_zone, vector<shared_ptr<" + structName + ">> &outResult, " + GeneralClientCPPCodes_GeneralGormClientTableOpt_Table_TableIndex(table, tableIndex) + ");\n")
	}
	f.WriteString("    static shared_ptr<" + structName + "> Get(int region, int logic_zone, int physics_zone, int &retCode, " + GeneralClientCPPCodes_GeneralGormClientTableOpt_Table_SplitInfo(table) + ");\n")
	f.WriteString("    static int GetVector(int region, int logic_zone, int physics_zone, vector<shared_ptr<" + structName + ">> &outResult, " + GeneralClientCPPCodes_GeneralGormClientTableOpt_Table_SplitInfo(table) + ");\n")
	f.WriteString("    static int Get(int region, int logic_zone, int physics_zone, " + GeneralClientCPPCodes_GeneralGormClientTableOpt_Table_SplitInfo(table) + ", uint32 &cbId, GORM_CbFun cbFunc);\n")
	f.WriteString("    static int Delete(int region, int logic_zone, int physics_zone, " + GeneralClientCPPCodes_GeneralGormClientTableOpt_Table_SplitInfo(table) + ", uint32 &cbId, GORM_CbFun cbFunc);\n")
	f.WriteString("    static int SetPbMsg(int region, int logic_zone, int physics_zone, " + pbStructName + " *pbMsg, bool forceSave=false);\n")
	f.WriteString("\n")
	f.WriteString("    // static不带区服的接口，用于全区全服架构\n")
	for _, tableIndex := range table.TableIndex {
		f.WriteString("    static shared_ptr<" + structName + "> GetBy" + tableIndex.Name + "(int &retCode, " + GeneralClientCPPCodes_GeneralGormClientTableOpt_Table_TableIndex(table, tableIndex) + ");\n")
		f.WriteString("    static int GetVectorBy" + tableIndex.Name + "(vector<shared_ptr<" + structName + ">> &outResult, " + GeneralClientCPPCodes_GeneralGormClientTableOpt_Table_TableIndex(table, tableIndex) + ");\n")
	}
	f.WriteString("    static shared_ptr<" + structName + "> Get(int &retCode, " + GeneralClientCPPCodes_GeneralGormClientTableOpt_Table_SplitInfo(table) + ");\n")
	f.WriteString("    static int GetVector(vector<shared_ptr<" + structName + ">> &outResult, " + GeneralClientCPPCodes_GeneralGormClientTableOpt_Table_SplitInfo(table) + ");\n")
	f.WriteString("    static int Get(" + GeneralClientCPPCodes_GeneralGormClientTableOpt_Table_SplitInfo(table) + ", uint32 &cbId, GORM_CbFun cbFunc);\n")
	f.WriteString("    static int Delete(" + GeneralClientCPPCodes_GeneralGormClientTableOpt_Table_SplitInfo(table) + ", uint32 &cbId, GORM_CbFun cbFunc);\n")

	f.WriteString("\n")
	f.WriteString("    // 本地操作接口\n")
	f.WriteString("    int Delete(uint32 &cbId, GORM_CbFun cbFunc);\n")
	f.WriteString("    int SetPbMsg(" + pbStructName + " *pbMsg, bool forceSave=false);\n")
	f.WriteString("    // 更新数据到数据库，有则更新，没有则插入。如果业务侧确认是新插入数据则建议使用Update接口，确认是新增数据则建议使用Add接口\n")
	f.WriteString("    int SaveToDB();\n")
	f.WriteString("    int Add();\n")
	f.WriteString("    int Update();\n")
	f.WriteString("    " + pbStructName + " *GetPbMsg();\n")
	f.WriteString("    " + pbStructName + " *ReleasePbMsg(); // 将pb消息返回，并将本地数据置为空(转移数据属主)\n")

	f.WriteString("\n")
	// 声明每个字段的存取方法
	if 0 != GeneralClientCPPCodes_GeneralGormClientTableOpt_Table_H_Columns_Desc(table, f) {
		return -1
	}
	f.WriteString("private:\n")
	for _, tableIndex := range table.TableIndex {
		f.WriteString("    int DoGetBy" + tableIndex.Name + "(uint32 &cbId, " + GeneralClientCPPCodes_GeneralGormClientTableOpt_Table_TableIndex(table, tableIndex) + ");\n")
		//f.WriteString("    int DoGetVectorBy" + tableIndex.Name + "(uint32 &cbId, " + GeneralClientCPPCodes_GeneralGormClientTableOpt_Table_TableIndex(table, tableIndex) + ");\n")
	}
	//f.WriteString("    int DoGetVector(uint32 &cbId, " + GeneralClientCPPCodes_GeneralGormClientTableOpt_Table_SplitInfo(table) + ");\n")
	f.WriteString("    int DoGet(uint32 &cbId, " + GeneralClientCPPCodes_GeneralGormClientTableOpt_Table_SplitInfo(table) + ");\n")
	f.WriteString("    int DoDelete(uint32 &cbId);\n")
	f.WriteString("    int DoUpdate(uint32 &cbId);\n")
	f.WriteString("    int DoInsert(uint32 &cbId);\n")
	f.WriteString("    int DoUpSert(uint32 &cbId);\n")
	f.WriteString("    bool HasSetPrimaryKey();\n")
	f.WriteString("\n")
	f.WriteString("    friend class GORM_ClientMsg;\n")
	f.WriteString("    static void GetCallBack(GORM_ClientMsg *clientMsg);\n")
	f.WriteString("public:\n")
	f.WriteString("    mutex mtx;\n")
	// 其它变量
	f.WriteString("private:\n")
	f.WriteString("    char dirtyFlag = 0;\n")
	f.WriteString("    " + pbStructName + " *tablePbValue = nullptr;\n")
	f.WriteString("    GORM_FieldsOpt fieldOpt;\n")
	f.WriteString("};\n\n")

	return 0
}

func GeneralClientCPPCodes_GeneralGormClientTableOpt_CPP_Table(table common.TableInfo, f *os.File) int {
	structName := "GORM_ClientTable" + common.CPP_TableStruct(table.Name)

	f.WriteString("int " + structName + "::GetAllRows(vector<shared_ptr<" + structName + ">> &outResult, int tableIndex)\n")
	f.WriteString("{\n")
	GeneralClientCPPCodes_GeneralGormClientTableOpt_CPP_Table_DoGetAllRows(table, f)
	f.WriteString("}\n\n")

	// DoGet函数
	f.WriteString("int " + structName + "::DoGet(uint32 &cbId, " + GeneralClientCPPCodes_GeneralGormClientTableOpt_Table_SplitInfo(table) + ")\n")
	f.WriteString("{\n")
	GeneralClientCPPCodes_GeneralGormClientTableOpt_CPP_Table_DoGet(table, f, "get")
	f.WriteString("}\n")
	// DoDelete函数
	f.WriteString("int " + structName + "::DoDelete(uint32 &cbId)\n")
	f.WriteString("{\n")
	f.WriteString("    if (!this->HasSetPrimaryKey()) return GORM_NO_PRIMARY_KEY;\n")
	GeneralClientCPPCodes_GeneralGormClientTableOpt_CPP_Table_DoFunc(table, f, "delete")
	f.WriteString("}\n")
	// DoUpdate函数
	f.WriteString("int " + structName + "::DoUpdate(uint32 &cbId)\n")
	f.WriteString("{\n")
	f.WriteString("    if (!this->HasSetPrimaryKey()) return GORM_NO_PRIMARY_KEY;\n")
	GeneralClientCPPCodes_GeneralGormClientTableOpt_CPP_Table_DoFunc(table, f, "update")
	f.WriteString("}\n")
	// DoInsert函数
	f.WriteString("int " + structName + "::DoInsert(uint32 &cbId)\n")
	f.WriteString("{\n")
	f.WriteString("    if (!this->HasSetPrimaryKey()) return GORM_NO_PRIMARY_KEY;\n")
	GeneralClientCPPCodes_GeneralGormClientTableOpt_CPP_Table_DoFunc(table, f, "insert")
	f.WriteString("}\n")
	// DoUpSert函数
	f.WriteString("int " + structName + "::DoUpSert(uint32 &cbId)\n")
	f.WriteString("{\n")
	f.WriteString("    if (!this->HasSetPrimaryKey()) return GORM_NO_PRIMARY_KEY;\n")
	GeneralClientCPPCodes_GeneralGormClientTableOpt_CPP_Table_DoFunc(table, f, "replace")
	f.WriteString("}\n")

	// 为索引生成获取函数
	for _, tableIndex := range table.TableIndex {
		f.WriteString("int " + structName + "::DoGetBy" + tableIndex.Name + "(uint32 &cbId, " + GeneralClientCPPCodes_GeneralGormClientTableOpt_Table_TableIndex(table, tableIndex) + ")\n")
		f.WriteString("{\n")
		GeneralClientCPPCodes_GeneralGormClientTableOpt_CPP_Table_DoGetByIndex(table, tableIndex, f)
		f.WriteString("}\n")
	}

	// 生成所有的返回数组的函数GetVector
	f.WriteString("int " + structName + "::GetVector(int region, int logic_zone, int physics_zone, vector<shared_ptr<" + structName + ">> &outResult, " + GeneralClientCPPCodes_GeneralGormClientTableOpt_Table_SplitInfo(table) + ")\n")
	f.WriteString("{\n")
	GeneralClientCPPCodes_GeneralGormClientTableOpt_CPP_Table_DoGetVector(table, 1, 0, f)
	f.WriteString("}\n\n")
	for i, tableIndex := range table.TableIndex {
		f.WriteString("int " + structName + "::GetVectorBy" + tableIndex.Name + "(int region, int logic_zone, int physics_zone, vector<shared_ptr<" + structName + ">> &outResult, " + GeneralClientCPPCodes_GeneralGormClientTableOpt_Table_TableIndex(table, tableIndex) + ")\n")
		f.WriteString("{\n")
		GeneralClientCPPCodes_GeneralGormClientTableOpt_CPP_Table_DoGetVector(table, 2, i, f)
		f.WriteString("}\n\n")
	}
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
	defer f.Close()
	f.Truncate(0)

	f.WriteString(common.NOT_EDIT_WARNING)

	headerFile := "gorm_client_table_opt_" + fileName + ".h"
	f.WriteString("#include \"" + headerFile + "\"\n")
	f.WriteString("#include \"gorm_wrap.h\"\n")
	f.WriteString("#include \"gorm_sys_inc.h\"\n")
	f.WriteString("#include \"gorm_client_thread.h\"\n")
	f.WriteString("#include \"gorm_utils.h\"\n")
	f.WriteString("#include \"gorm_pb_proto.pb.h\"\n")
	if SupportCppCoroutine {
		f.WriteString("#include \"gamesh/framework/framework.h\"\n")
		f.WriteString("#include \"gamesh/coroutine/coroutine_mgr.h\"\n")
	}
	f.WriteString("\n")
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
	defer f.Close()
	f.Truncate(0)

	f.WriteString(common.NOT_EDIT_WARNING)

	bigName := strings.ToUpper(game.GetFileNameCharacter())
	f.WriteString("#ifndef _GORM_CLIENT_TABLE_OPT_" + bigName + "_H__\n")
	f.WriteString("#define _GORM_CLIENT_TABLE_OPT_" + bigName + "_H__\n")
	f.WriteString("#include \"" + fileName + ".pb.h\"\n")
	f.WriteString("#include \"gorm_sys_inc.h\"\n")
	f.WriteString("#include \"gorm_define.h\"\n")
	f.WriteString("#include \"gorm_utils.h\"\n")
	f.WriteString("#include \"gorm_error.h\"\n")
	f.WriteString("#include \"gorm_table.h\"\n")
	f.WriteString("#include \"gorm_client_msg.h\"\n")
	f.WriteString("#include \"gorm_pb_proto.pb.h\"\n")
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
