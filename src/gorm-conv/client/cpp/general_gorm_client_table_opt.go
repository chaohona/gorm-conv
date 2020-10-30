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

func GeneralClientCPPCodes_GeneralGormClientTableOpt_Table_SplitInfo_Param(table common.TableInfo) (result string) {
	for idx, str := range table.SplitInfo.SplitCols {
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
			f.WriteString("    // 主键，设置之后不能更改")
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
	// static 同步Get函数
	f.WriteString("inline " + structName + "* " + structName + "::Get(int region, int logic_zone, int physics_zone, int &retCode, " + GeneralClientCPPCodes_GeneralGormClientTableOpt_Table_SplitInfo(table) + ")\n")
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
	f.WriteString("    uint32 cbId = 0;\n")
	f.WriteString("    retCode = table->DoGet(cbId);\n")
	f.WriteString("    if (GORM_OK != retCode)\n")
	f.WriteString("    {\n")
	f.WriteString("        return nullptr;\n")
	f.WriteString("    }\n")
	f.WriteString("    return table;\n")
	f.WriteString("}\n")

	// static 异步Get函数
	f.WriteString("inline int " + structName + "::Get(int region, int logic_zone, int physics_zone, " + GeneralClientCPPCodes_GeneralGormClientTableOpt_Table_SplitInfo(table) + ", uint32 &cbId, GORM_CbFun cbFunc)\n")
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
	f.WriteString("    int retCode = table->DoGet(cbId);\n")
	f.WriteString("    if (GORM_OK != retCode)\n")
	f.WriteString("    {\n")
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
	f.WriteString("inline " + structName + "* " + structName + "::Get(int &retCode, " + GeneralClientCPPCodes_GeneralGormClientTableOpt_Table_SplitInfo(table) + ")\n")
	f.WriteString("{\n")
	f.WriteString("    return " + structName + "::Get(0,0,0, retCode, " + GeneralClientCPPCodes_GeneralGormClientTableOpt_Table_SplitInfo_Param(table) + ");\n")
	f.WriteString("}\n")

	// static 不带区服，异步Get函数
	f.WriteString("inline int " + structName + "::Get(" + GeneralClientCPPCodes_GeneralGormClientTableOpt_Table_SplitInfo(table) + ", uint32 &cbId, GORM_CbFun cbFunc)\n")
	f.WriteString("{\n")
	f.WriteString("    return " + structName + "::Get(0,0,0, " + GeneralClientCPPCodes_GeneralGormClientTableOpt_Table_SplitInfo_Param(table) + ", cbId, cb);\n")
	f.WriteString("}\n")

	// static Delete函数
	f.WriteString("inline int " + structName + "::Delete(" + GeneralClientCPPCodes_GeneralGormClientTableOpt_Table_SplitInfo(table) + ", uint32 &cbId, GORM_CbFun cbFunc)\n")
	f.WriteString("{\n")
	f.WriteString("    return " + structName + "::Delete(0,0,0," + GeneralClientCPPCodes_GeneralGormClientTableOpt_Table_SplitInfo_Param(table) + ", cbId, cb);\n")
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
	f.WriteString("    if (clientMsg == nullptr) return;\n")
	f.WriteString("    GORM_PB_GET_RSP *pbRspMsg = dynamic_cast<GORM_PB_GET_RSP*>(clientMsg->pbRspMsg);\n")
	f.WriteString("    if (pbRspMsg == nullptr || !pbRspMsg->table().has_" + table.Name + "())\n")
	f.WriteString("    {\n")
	f.WriteString("        clientMsg->cbFunc(clientMsg->cbId, clientMsg->rspCode.code, nullptr);\n")
	f.WriteString("        return;\n")
	f.WriteString("    }\n")
	f.WriteString("    " + structName + " *table = new " + structName + "(pbRspMsg->table().release_" + table.Name + "));\n")
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
	f.WriteString("    // static带区服的接口，用于分区分服架构\n")
	f.WriteString("    static shared_ptr<" + structName + "> Get(int region, int logic_zone, int physics_zone, int &retCode, " + GeneralClientCPPCodes_GeneralGormClientTableOpt_Table_SplitInfo(table) + ");\n")
	f.WriteString("    static int Get(int region, int logic_zone, int physics_zone, " + GeneralClientCPPCodes_GeneralGormClientTableOpt_Table_SplitInfo(table) + ", uint32 &cbId, GORM_CbFun cbFunc);\n")
	f.WriteString("    static int Delete(int region, int logic_zone, int physics_zone, " + GeneralClientCPPCodes_GeneralGormClientTableOpt_Table_SplitInfo(table) + ", uint32 &cbId, GORM_CbFun cbFunc);\n")
	f.WriteString("    static int SetPbMsg(int region, int logic_zone, int physics_zone, " + pbStructName + " *pbMsg, bool forceSave=false);\n")
	f.WriteString("\n")
	f.WriteString("    // static不带区服的接口，用于全区全服架构\n")
	f.WriteString("    static shared_ptr<" + structName + "> Get(int &retCode, " + GeneralClientCPPCodes_GeneralGormClientTableOpt_Table_SplitInfo(table) + ");\n")
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
	f.WriteString("    int DoGet(uint32 &cbId);\n")
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

func GeneralClientCPPCodes_GeneralGormClientTableOpt_CPP_Table_DoFunc(table common.TableInfo, f *os.File, opt string) int {
	var bigOpt string = strings.ToUpper(opt)
	var bigTableName string = strings.ToUpper(table.Name)

	f.WriteString("    GORM_ClientMsg *clientMsg = new GORM_ClientMsg();\n")
	f.WriteString("    clientMsg->tableId = GORM_PB_TABLE_IDX_" + bigTableName + ";\n")
	f.WriteString("    clientMsg->reqCmd = GORM_CMD_" + bigOpt + ";\n")
	f.WriteString("    clientMsg->fieldOpt = &this->fieldOpt;\n")
	f.WriteString("    GORM_PB_" + bigOpt + "_REQ *getReq = new GORM_PB_" + bigOpt + "_REQ();\n")
	f.WriteString("    clientMsg->pbReqMsg = getReq;\n")
	if bigOpt == "DELETE" {
		f.WriteString("    GORM_PB_TABLE *pbTableAll = getReq->mutable_table();\n")
	} else {
		f.WriteString("    GORM_PB_TABLE *pbTableAll = getReq->add_tables();\n")
	}

	f.WriteString("    pbTableAll->set_allocated_" + table.Name + "(this->tablePbValue);\n\n")
	f.WriteString("    if (GORM_OK != clientMsg->PackReq())\n")
	f.WriteString("    {\n")
	f.WriteString("        pbTableAll->release_" + table.Name + "();\n")
	f.WriteString("        delete clientMsg;\n")
	f.WriteString("        return GORM_ERROR;\n")
	f.WriteString("    }\n")

	f.WriteString("    // 使用完，交还外部数据\n")
	f.WriteString("    pbTableAll->release_" + table.Name + "();\n")
	f.WriteString(`
    // 发送Get请求
    if (GORM_OK != GORM_ClientThreadPool::Instance()->SendRequest(clientMsg, clientMsg->cbId))
    {
        delete clientMsg;
        return GORM_ERROR;
    }
    cbId = clientMsg->cbId;
    clientMsg->Wait();	// 等待响应
    clientMsg = nullptr;
    // 获取结果
    for (int i=0; i<1000;i++)
    {
        if (GORM_OK != GORM_ClientThreadPool::Instance()->GetResponse(clientMsg))
        {
            return GORM_ERROR;
        }
        if (clientMsg != nullptr)
        {
            break;
        }
        ThreadSleepMilliSeconds(1); 
    }

    if (clientMsg == nullptr)
        return GORM_ERROR;

    shared_ptr<GORM_ClientMsg> sharedClientMsg(clientMsg);
    unique_lock<mutex> lck(sharedClientMsg->mtx);
    int code = sharedClientMsg->rspCode.code;
    
    return code;
`)
	return 0
}

func GeneralClientCPPCodes_GeneralGormClientTableOpt_CPP_Table_DoGet(table common.TableInfo, f *os.File, opt string) int {
	var bigOpt string = strings.ToUpper(opt)
	var bigTableName string = strings.ToUpper(table.Name)
	structName := "GORM_ClientTable" + common.CPP_TableStruct(table.Name)
	///	pbStructName := "GORM_PB_Table_" + table.Name

	f.WriteString("    GORM_ClientMsg *clientMsg = new GORM_ClientMsg();\n")
	f.WriteString("    clientMsg->tableId = GORM_PB_TABLE_IDX_" + bigTableName + ";\n")
	f.WriteString("    clientMsg->reqCmd = GORM_CMD_" + bigOpt + ";\n")
	f.WriteString("    clientMsg->fieldOpt = &this->fieldOpt;\n")
	f.WriteString("    GORM_PB_GET_REQ *getReq = new GORM_PB_" + bigOpt + "_REQ();\n")
	f.WriteString("    clientMsg->pbReqMsg = getReq;\n")
	f.WriteString("    clientMsg->getCBFunc = " + structName + "::GetCallBack;\n")
	f.WriteString("    GORM_PB_TABLE *pbTableAll = getReq->mutable_table();\n")
	f.WriteString("    pbTableAll->set_allocated_" + table.Name + "(this->tablePbValue);\n\n")

	f.WriteString("    if (GORM_OK != clientMsg->PackReq())\n")
	f.WriteString("    {\n")
	f.WriteString("        pbTableAll->release_" + table.Name + "();\n")
	f.WriteString("        delete clientMsg;\n")
	f.WriteString("        return GORM_ERROR;\n")
	f.WriteString("    }\n")

	f.WriteString("    // 使用完，交还外部数据\n")
	f.WriteString("    pbTableAll->release_" + table.Name + "();\n")
	f.WriteString(`
    // 发送Get请求
    if (GORM_OK != GORM_ClientThreadPool::Instance()->SendRequest(clientMsg, clientMsg->cbId))
    {
        delete clientMsg;
        return GORM_ERROR;
    }
    cbId = clientMsg->cbId;
    clientMsg->Wait();	// 等待响应
    // 获取结果
    clientMsg = nullptr;
    for (int i=0;i<1000;i++)
    {
        if (GORM_OK != GORM_ClientThreadPool::Instance()->GetResponse(clientMsg))
        {
            return GORM_ERROR;
        }
        if (clientMsg != nullptr)
        {
            break;
        }
        ThreadSleepMilliSeconds(1); 
    }
    if (clientMsg == nullptr)
        return GORM_ERROR;

    shared_ptr<GORM_ClientMsg> sharedClientMsg(clientMsg);
    unique_lock<mutex> lck(sharedClientMsg->mtx);
    if (GORM_OK != sharedClientMsg->rspCode.code)
    {
        return sharedClientMsg->rspCode.code;
    }
    GORM_PB_GET_RSP *pbRspMsg = dynamic_cast<GORM_PB_GET_RSP*>(sharedClientMsg->pbRspMsg);
`)
	f.WriteString("    if (pbRspMsg == nullptr || !pbRspMsg->table().has_" + table.Name + "())\n")
	f.WriteString("    {\n")
	f.WriteString("        return GORM_NO_MORE_RECORD;\n")
	f.WriteString("    }\n")
	f.WriteString("    this->tablePbValue = pbRspMsg->mutable_table()->release_" + table.Name + "();\n")
	f.WriteString("    this->dirtyFlag = 1;\n")
	f.WriteString("\n")
	f.WriteString("    return GORM_OK;\n")
	return 0
}

func GeneralClientCPPCodes_GeneralGormClientTableOpt_CPP_Table(table common.TableInfo, f *os.File) int {
	structName := "GORM_ClientTable" + common.CPP_TableStruct(table.Name)

	// DoGet函数
	f.WriteString("int " + structName + "::DoGet(uint32 &cbId)\n")
	f.WriteString("{\n")
	f.WriteString("    if (!this->HasSetPrimaryKey()) return GORM_NO_PRIMARY_KEY;\n")
	f.WriteString("    unique_lock<mutex> lck(this->mtx);\n")
	GeneralClientCPPCodes_GeneralGormClientTableOpt_CPP_Table_DoGet(table, f, "get")
	f.WriteString("}\n")
	// DoDelete函数
	f.WriteString("int " + structName + "::DoDelete(uint32 &cbId)\n")
	f.WriteString("{\n")
	f.WriteString("    if (!this->HasSetPrimaryKey()) return GORM_NO_PRIMARY_KEY;\n")
	f.WriteString("    unique_lock<mutex> lck(this->mtx);\n")
	GeneralClientCPPCodes_GeneralGormClientTableOpt_CPP_Table_DoFunc(table, f, "delete")
	f.WriteString("}\n")
	// DoUpdate函数
	f.WriteString("int " + structName + "::DoUpdate(uint32 &cbId)\n")
	f.WriteString("{\n")
	f.WriteString("    if (!this->HasSetPrimaryKey()) return GORM_NO_PRIMARY_KEY;\n")
	f.WriteString("    if (this->dirtyFlag == 1) return GORM_NOT_DIRTY_DATA;\n")
	f.WriteString("    unique_lock<mutex> lck(this->mtx);\n")
	GeneralClientCPPCodes_GeneralGormClientTableOpt_CPP_Table_DoFunc(table, f, "update")
	f.WriteString("}\n")
	// DoInsert函数
	f.WriteString("int " + structName + "::DoInsert(uint32 &cbId)\n")
	f.WriteString("{\n")
	f.WriteString("    if (!this->HasSetPrimaryKey()) return GORM_NO_PRIMARY_KEY;\n")
	f.WriteString("    if (this->dirtyFlag == 1) return GORM_NOT_DIRTY_DATA;\n")
	f.WriteString("    unique_lock<mutex> lck(this->mtx);\n")
	GeneralClientCPPCodes_GeneralGormClientTableOpt_CPP_Table_DoFunc(table, f, "insert")
	f.WriteString("}\n")
	// DoUpSert函数
	f.WriteString("int " + structName + "::DoUpSert(uint32 &cbId)\n")
	f.WriteString("{\n")
	f.WriteString("    if (!this->HasSetPrimaryKey()) return GORM_NO_PRIMARY_KEY;\n")
	f.WriteString("    if (this->dirtyFlag == 1) return GORM_NOT_DIRTY_DATA;\n")
	f.WriteString("    unique_lock<mutex> lck(this->mtx);\n")
	GeneralClientCPPCodes_GeneralGormClientTableOpt_CPP_Table_DoFunc(table, f, "replace")
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

var exampleCodes string = `/*
// 函数Get、Delete、SetPbMsg、SaveToDB、GetPbMsg为所有的表都会对应的操作函数
// 除了上面4个函数，其它函数都为更新表结构中的字段
class GORM_ClientTableAccount
{
public:
	// @desc 	以同步方式获取account表中的一条数据，此函数有一个对应的不带区服信息的函数
	// @param	region数据所在可用域（机房）
	// @param	logic_zone数据所属逻辑区
	// @param	physics_zone数组所属物理区
	// @param	id表的主键，不同的表主键类型与个数不同
	// @retval	获取到的数据
    static GORM_ClientTableAccount* Get(int region, int logic_zone, int physics_zone, int &retCode, int32 id);
    // 不带区服信息的函数
    static GORM_ClientTableAccount* Get(int32 id);
    
    // @desc 	以异步方式获取account表中的一条数据，此函数有一个对应的不带区服信息的函数
	// @param	region数据所在可用域（机房）
	// @param	logic_zone数据所属逻辑区
	// @param	physics_zone数组所属物理区
	// @param	cbId异步请求的回调id
	// @param	id表的主键，不同的表主键类型与个数不同
	// @param	cb异步请求的回调函数
	//			回调函数的参数解释:
	//			int64对应上面的cbId, GORM_ClientTableAccount*获取到的数据的
	// @retval	0  成功，等待回调
				<0 失败
    static int Get(int region, int logic_zone, int physics_zone, int32 id, int64 &cbId, int (*cb)(int64, int&, GORM_ClientTableAccount*));
    // 不带区服信息的函数
    static int Get(int32 id, int64 &cbId, int (*cb)(int64, int&, GORM_ClientTableAccount*));
    
    // @desc 	以异步方式删除account表中的一条数据，此函数有一个对应的不带区服信息的函数
	// @param	region数据所在可用域（机房）
	// @param	logic_zone数据所属逻辑区
	// @param	physics_zone数组所属物理区
	// @param	cbId异步请求的回调id
	// @param	cb异步请求的回调函数,可为nullptr，表示不关心结果
	//			回调函数的参数解释:
	//			int64对应上面的cbId
	// @param	id表的主键，不同的表主键类型与个数不同
	// @retval	0  成功，等待回调
				<0 失败
    static int Delete(int region, int logic_zone, int physics_zone, int32 id, int64 &cbId, int (*cb)(int64, int&));
    // 不带区服信息的函数
    static int Delete(int32 id, int64 &cbId, int (*cb)(int64, int&));
    
    // @desc 	全量更新数据，此函数有一个对应的不带区服信息的函数
    // @param	pbMsg需要覆盖更新的pb数据
    // @param	forceSave是否立即持久化，默认值为false不立即持久化
    // @retval	0  成功
				<0 失败
    static int SetPbMsg(int region, int logic_zone, int physics_zone, GORM_PB_Table_account *pbMsg, bool forceSave=false);
    // 不带区服信息的函数
    static int SetPbMsg(GORM_PB_Table_account *pbMsg, bool forceSave=false);
    
    // @desc 删除本条数据
    int Delete(int64 &cbId, int (*cb)(int64));
    
    // @desc 全量覆盖更新本条数据
    int SetPbMsg(GORM_PB_Table_account *pbMsg, bool forceSave=false);
    
    // @desc 立即将本条数据持久化保存,已经存在的数据则更新，没有的数据则插入一条新数据
    int SaveToDB();
    
    // @desc 增加新的记录到数据库,插入失败则直接返回
    int Add();
    // @desc 更新一条记录，更新失败则返回
    int Update();
    
    // @desc 获取原始pb结构数据，长用于发送给后端逻辑节点使用
    GORM_PB_Table_account *GetPbMsg();
    
    // 以下为针对字段的Get，Set操作
    uint64 GetVersion();
    int SetVersion(uint64 version, bool forceSave=false);
    int32 GetId();
    int SetId(int32 id, bool forceSave=false);
    const string &GetAccount() const;
    int SetAccount(const string &account, bool forceSave=false);
    int SetAccount(const char* account, size_t size, bool forceSave=false);
    const string &GetAllbinary() const;
    int SetAllbinary(const string &allbinary, bool forceSave=false);
    int SetAllbinary(const char* allbinary, size_t size, bool forceSave=false);
private:
	// 0:新的没有持久化的数据，1:从持久化存储拉到内存的数据，2:已经持久化了的并且有更新的数据
	char dirtyFlag = 0;
    GORM_PB_Table_account *tablePbValue = nullptr;
    GORM_FieldsOpt fieldOpt;
};
*/

`

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
	f.WriteString("#include \"gorm_define.h\"\n")
	f.WriteString("#include \"gorm_utils.h\"\n")
	f.WriteString("#include \"gorm_error.h\"\n")
	f.WriteString("#include \"gorm_table.h\"\n")
	f.WriteString("#include \"gorm_client_msg.h\"")
	f.WriteString("\n")
	f.WriteString(exampleCodes)
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
