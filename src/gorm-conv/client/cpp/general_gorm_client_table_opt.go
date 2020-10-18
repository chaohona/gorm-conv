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
		f.WriteString("    // " + col.Name + "\n")
		colType := common.CPPField_CPPType(col.Type)
		colStructName := common.CPP_TableColumnName(col.Name)
		getColFunc := "Get" + colStructName
		setColFunc := "Set" + colStructName
		if colType == "string" {
			f.WriteString("    const string &" + getColFunc + "() const;\n")
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

// 头文件中的inline函数实现
func GeneralClientCPPCodes_GeneralGormClientTableOpt_Table_H_Columns_Define(table common.TableInfo, f *os.File) int {
	pbStructName := "GORM_PB_Table_" + table.Name
	structName := "GORM_ClientTable" + common.CPP_TableStruct(table.Name)
	f.WriteString("inline " + pbStructName + "*" + structName + "::GetPbMsg()\n")
	f.WriteString("{\n")
	f.WriteString("    return this->tablePbValue;\n")
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
		f.WriteString("    return this->tablePbValue->" + col.Name + "();\n")
		f.WriteString("}\n")
	}
	// 生成更新所有字段函数
	f.WriteString("inline int " + structName + "::SetPbMsg(" + pbStructName + " *pbMsg, bool forceSave)\n")
	f.WriteString("{\n")
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
			f.WriteString("    this->tablePbValue->set_" + col.Name + "(" + col.Name + ");\n")
			f.WriteString(setFiledMode(idx))
			f.WriteString("    if (this->dirtyFlag!=0) this->dirtyFlag = 2;\n")
			f.WriteString("    if (forceSave)\n")
			f.WriteString("        return this->SaveToDB();\n")
			f.WriteString("    return 0;\n")
			f.WriteString("}\n")

			//Set(const char*, size_t)
			f.WriteString("inline int " + structName + "::" + setColFunc + "(const char* " + col.Name + ", size_t size, bool forceSave)\n")
			f.WriteString("{\n")
			f.WriteString("    this->tablePbValue->set_" + col.Name + "(" + col.Name + ", size);\n")
			f.WriteString(setFiledMode(idx))
			f.WriteString("    if (this->dirtyFlag!=0) this->dirtyFlag = 2;\n")
			f.WriteString("    if (forceSave)\n")
			f.WriteString("        return this->SaveToDB();\n")
			f.WriteString("    return 0;\n")
			f.WriteString("}\n")
		} else {
			f.WriteString("inline int " + structName + "::" + setColFunc + "(" + colType + " " + col.Name + ", bool forceSave)\n")
			f.WriteString("{\n")
			f.WriteString("    this->tablePbValue->set_" + col.Name + "(" + col.Name + ");\n")
			f.WriteString(setFiledMode(idx))
			f.WriteString("    if (this->dirtyFlag!=0) this->dirtyFlag = 2;\n")
			f.WriteString("    if (forceSave)\n")
			f.WriteString("        return this->SaveToDB();\n")
			f.WriteString("    return 0;\n")
			f.WriteString("}\n")
		}
	}
	return 0
}

func GeneralClientCPPCodes_GeneralGormClientTableOpt_Table_H(table common.TableInfo, f *os.File) int {
	// 表对应的类的声明
	structName := "GORM_ClientTable" + common.CPP_TableStruct(table.Name)
	pbStructName := "GORM_PB_Table_" + table.Name

	f.WriteString("// 表" + table.Name + "\n")
	f.WriteString("class " + structName + "\n")
	f.WriteString("{\n")
	f.WriteString("public:\n")
	f.WriteString("    // static带区服的接口，用于分区分服架构\n")
	f.WriteString("    static " + structName + "* Get(int region, int logic_zone, int physics_zone, " + GeneralClientCPPCodes_GeneralGormClientTableOpt_Table_SplitInfo(table) + ");\n")
	f.WriteString("    static int Get(int region, int logic_zone, int physics_zone, " + GeneralClientCPPCodes_GeneralGormClientTableOpt_Table_SplitInfo(table) + ", int64 &cbId, int (*cb)(int64, " + structName + "*));\n")
	f.WriteString("    static int Delete(int region, int logic_zone, int physics_zone, " + GeneralClientCPPCodes_GeneralGormClientTableOpt_Table_SplitInfo(table) + ", int64 &cbId, int (*cb)(int64));\n")
	f.WriteString("    static int SetPbMsg(int region, int logic_zone, int physics_zone, " + pbStructName + " *pbMsg, bool forceSave=false);\n")
	f.WriteString("\n")
	f.WriteString("    // static不带区服的接口，用于全区全服架构\n")
	f.WriteString("    static " + structName + "* Get(" + GeneralClientCPPCodes_GeneralGormClientTableOpt_Table_SplitInfo(table) + ");\n")
	f.WriteString("    static int Get(" + GeneralClientCPPCodes_GeneralGormClientTableOpt_Table_SplitInfo(table) + ", int64 &cbId, int (*cb)(int64, " + structName + "*));\n")
	f.WriteString("    static int Delete(" + GeneralClientCPPCodes_GeneralGormClientTableOpt_Table_SplitInfo(table) + ", int64 &cbId, int (*cb)(int64));\n")
	//f.WriteString("    static int SetPbMsg(" + pbStructName + " *pbMsg, bool forceSave=false);\n")

	f.WriteString("\n")
	f.WriteString("    // 本地操作接口\n")
	f.WriteString("    int Delete(int64 &cbId, int (*cb)(int64));\n")
	f.WriteString("    int SetPbMsg(" + pbStructName + " *pbMsg, bool forceSave=false);\n")
	//f.WriteString("    int RemoveFromLocal();\n")
	f.WriteString("    int SaveToDB();\n")
	f.WriteString("    " + pbStructName + " *GetPbMsg();\n")

	f.WriteString("\n")
	// 声明每个字段的存取方法
	if 0 != GeneralClientCPPCodes_GeneralGormClientTableOpt_Table_H_Columns_Desc(table, f) {
		return -1
	}
	f.WriteString("private:\n")
	f.WriteString("    int DoGet();\n")
	f.WriteString("    int DoDelete();\n")
	f.WriteString("    int DoUpdate();\n")
	f.WriteString("    int DoInsert();\n")
	f.WriteString("public:\n")
	f.WriteString("    mutex mtx;\n")
	// 其它变量
	f.WriteString("private:\n")
	f.WriteString("    " + pbStructName + " *tablePbValue = nullptr;\n")
	f.WriteString("    GORM_FieldsOpt fieldOpt;\n")
	f.WriteString("    int dirtyFlag = 0;\n")
	f.WriteString("};\n\n")

	return 0
}

func GeneralClientCPPCodes_GeneralGormClientTableOpt_CPP_Table_DoFunc(table common.TableInfo, f *os.File, opt string) int {
	var bigOpt string = strings.ToUpper(opt)
	var bigTableName string = strings.ToUpper(table.Name)

	f.WriteString("    GORM_ClientMsg *clientMsg = new GORM_ClientMsg();\n")
	f.WriteString("    clientMsg->tableId = GORM_PB_TABLE_IDX_" + bigTableName + ";\n")
	f.WriteString("    clientMsg->reqCmd = GORM_CMD_" + bigOpt + ";\n")
	f.WriteString("    GORM_PB_" + bigOpt + "_REQ *getReq = new GORM_PB_" + bigOpt + "_REQ();\n")
	f.WriteString("    clientMsg->pbReqMsg = getReq;\n")
	if bigOpt == "DELETE" {
		f.WriteString("    GORM_PB_TABLE *pbTableAll = getReq->mutable_table();\n")
	} else {
		f.WriteString("    GORM_PB_TABLE *pbTableAll = getReq->add_tables();\n")
	}

	f.WriteString("    pbTableAll->set_allocated_" + table.Name + "(this->tablePbValue);\n\n")
	f.WriteString(`
    // 打包
    if (GORM_OK != clientMsg->PackReq())
    {
        delete clientMsg;
        return GORM_ERROR;
    }
    // 发送Get请求
    if (GORM_OK != GORM_ClientThreadPool::Instance()->SendRequest(clientMsg, clientMsg->cbId))
    {
        delete clientMsg;
        return GORM_ERROR;
    }
    // 获取结果
    clientMsg = nullptr;
    for (;;)
    {
        if (GORM_OK != GORM_ClientThreadPool::Instance()->GetResponse(clientMsg))
        {
            return GORM_ERROR;
            break;
        }
        if (clientMsg != nullptr)
        {
            break;
        }
        
    }

    clientMsg->mtx.lock();
    int code = clientMsg->rspCode.code;
    clientMsg->mtx.unlock();
    delete clientMsg;
    return code;
`)
	return 0
}

func GeneralClientCPPCodes_GeneralGormClientTableOpt_CPP_Table_DoGet(table common.TableInfo, f *os.File, opt string) int {
	var bigOpt string = strings.ToUpper(opt)
	var bigTableName string = strings.ToUpper(table.Name)
	pbStructName := "GORM_PB_Table_" + table.Name

	f.WriteString("    GORM_ClientMsg *clientMsg = new GORM_ClientMsg();\n")
	f.WriteString("    clientMsg->tableId = GORM_PB_TABLE_IDX_" + bigTableName + ";\n")
	f.WriteString("    clientMsg->reqCmd = GORM_CMD_" + bigOpt + ";\n")
	f.WriteString("    GORM_PB_GET_REQ *getReq = new GORM_PB_" + bigOpt + "_REQ();\n")
	f.WriteString("    clientMsg->pbReqMsg = getReq;\n")
	f.WriteString("    GORM_PB_TABLE *pbTableAll = getReq->mutable_table();\n")
	f.WriteString("    pbTableAll->set_allocated_" + table.Name + "(this->tablePbValue);\n\n")
	f.WriteString(`
    // 打包
    if (GORM_OK != clientMsg->PackReq())
    {
        delete clientMsg;
        return GORM_ERROR;
    }
    // 发送Get请求
    if (GORM_OK != GORM_ClientThreadPool::Instance()->SendRequest(clientMsg, clientMsg->cbId))
    {
        delete clientMsg;
        return GORM_ERROR;
    }
    // 获取结果
    clientMsg = nullptr;
    for (;;)
    {
        if (GORM_OK != GORM_ClientThreadPool::Instance()->GetResponse(clientMsg))
        {
            return GORM_ERROR;
            break;
        }
        if (clientMsg != nullptr)
        {
            break;
        }
        
    }

    clientMsg->mtx.lock();
    if (GORM_OK != clientMsg->rspCode.code)
    {
        clientMsg->mtx.unlock();
        delete clientMsg;
        return GORM_ERROR;
    }`)
	f.WriteString("    this->tablePbValue = dynamic_cast<" + pbStructName + "*>(clientMsg->pbRspMsg);\n")
	f.WriteString(`
    clientMsg->mtx.unlock();
    delete clientMsg;
    return GORM_OK;
`)
	return 0
}

func GeneralClientCPPCodes_GeneralGormClientTableOpt_CPP_Table(table common.TableInfo, f *os.File) int {
	structName := "GORM_ClientTable" + common.CPP_TableStruct(table.Name)
	pbStructName := "GORM_PB_Table_" + table.Name
	//////////////////////////////////////// 带区服的接口
	// static 同步Get函数
	f.WriteString(structName + "* " + structName + "::Get(int region, int logic_zone, int physics_zone," + GeneralClientCPPCodes_GeneralGormClientTableOpt_Table_SplitInfo(table) + ")\n")
	f.WriteString("{\n")
	f.WriteString("    unique_lock<mutex> lck(GORM_Wrap::Instance()->mtx);\n")

	f.WriteString("    return nullptr;\n")
	f.WriteString("}\n")

	// static 异步Get函数
	f.WriteString("int " + structName + "::Get(int region, int logic_zone, int physics_zone, " + GeneralClientCPPCodes_GeneralGormClientTableOpt_Table_SplitInfo(table) + ", int64 &cbId, int (*cb)(int64, " + structName + "*))\n")
	f.WriteString("{\n")
	f.WriteString("    unique_lock<mutex> lck(GORM_Wrap::Instance()->mtx);\n")
	f.WriteString("    return 0;\n")
	f.WriteString("}\n")

	// static Delete函数
	f.WriteString("int " + structName + "::Delete(int region, int logic_zone, int physics_zone, " + GeneralClientCPPCodes_GeneralGormClientTableOpt_Table_SplitInfo(table) + ", int64 &cbId, int (*cb)(int64))\n")
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
	f.WriteString("    return table->Delete(cbId, cb);\n")
	f.WriteString("}\n")

	// static SetPbMsg函数
	f.WriteString("int " + structName + "::SetPbMsg(int region, int logic_zone, int physics_zone, " + pbStructName + " *pbMsg, bool forceSave)\n")
	f.WriteString("{\n")
	f.WriteString("    shared_ptr<" + structName + "> table = make_shared<" + structName + ">();\n")
	f.WriteString("    return table->SetPbMsg(pbMsg, true);\n")
	f.WriteString("}\n")
	//////////////////////////////////////// 带区服的接口

	/////////////////////////////////////////// 不带区服的接口
	// static 不带区服的同步Get函数
	f.WriteString(structName + "* " + structName + "::Get(" + GeneralClientCPPCodes_GeneralGormClientTableOpt_Table_SplitInfo(table) + ")\n")
	f.WriteString("{\n")
	f.WriteString("    return " + structName + "::Get(0,0,0," + GeneralClientCPPCodes_GeneralGormClientTableOpt_Table_SplitInfo_Param(table) + ");\n")
	f.WriteString("}\n")

	// static 不带区服，异步Get函数
	f.WriteString("int " + structName + "::Get(" + GeneralClientCPPCodes_GeneralGormClientTableOpt_Table_SplitInfo(table) + ", int64 &cbId, int (*cb)(int64, " + structName + "*))\n")
	f.WriteString("{\n")
	f.WriteString("    return " + structName + "::Get(0,0,0," + GeneralClientCPPCodes_GeneralGormClientTableOpt_Table_SplitInfo_Param(table) + ", cbId, cb);\n")
	f.WriteString("}\n")

	// static Delete函数
	f.WriteString("int " + structName + "::Delete(" + GeneralClientCPPCodes_GeneralGormClientTableOpt_Table_SplitInfo(table) + ", int64 &cbId, int (*cb)(int64))\n")
	f.WriteString("{\n")
	f.WriteString("    return " + structName + "::Delete(0,0,0," + GeneralClientCPPCodes_GeneralGormClientTableOpt_Table_SplitInfo_Param(table) + ", cbId, cb);\n")
	f.WriteString("}\n")
	/////////////////////////////////////////// 不带区服的接口

	// Delete函数
	f.WriteString("int " + structName + "::Delete(int64 &cbId, int (*cb)(int64))\n")
	f.WriteString("{\n")
	f.WriteString("    unique_lock<mutex> lck(GORM_Wrap::Instance()->mtx);\n")
	f.WriteString("    return 0;\n")
	f.WriteString("}\n")

	// SaveToDB函数
	f.WriteString("int " + structName + "::SaveToDB()\n")
	f.WriteString("{\n")
	f.WriteString("    int flag = this->dirtyFlag;\n")
	f.WriteString("    this->dirtyFlag = 1;\n")
	f.WriteString("    if (flag == 2)\n")
	f.WriteString("        return this->DoUpdate();\n")
	f.WriteString("    else if (flag == 0)\n")
	f.WriteString("        return this->DoInsert();\n")
	f.WriteString("    return 0;\n")
	f.WriteString("}\n")

	// DoGet函数
	f.WriteString("int " + structName + "::DoGet()\n")
	f.WriteString("{\n")
	f.WriteString("    unique_lock<mutex> lck(this->mtx);\n")
	GeneralClientCPPCodes_GeneralGormClientTableOpt_CPP_Table_DoGet(table, f, "get")
	f.WriteString("}\n")
	// DoDelete函数
	f.WriteString("int " + structName + "::DoDelete()\n")
	f.WriteString("{\n")
	f.WriteString("    unique_lock<mutex> lck(this->mtx);\n")
	GeneralClientCPPCodes_GeneralGormClientTableOpt_CPP_Table_DoFunc(table, f, "delete")
	f.WriteString("}\n")
	// DoUpdate函数
	f.WriteString("int " + structName + "::DoUpdate()\n")
	f.WriteString("{\n")
	f.WriteString("    unique_lock<mutex> lck(this->mtx);\n")
	GeneralClientCPPCodes_GeneralGormClientTableOpt_CPP_Table_DoFunc(table, f, "update")
	f.WriteString("}\n")
	// DoInsert函数
	f.WriteString("int " + structName + "::DoInsert()\n")
	f.WriteString("{\n")
	f.WriteString("    unique_lock<mutex> lck(this->mtx);\n")
	GeneralClientCPPCodes_GeneralGormClientTableOpt_CPP_Table_DoFunc(table, f, "insert")
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
    static GORM_ClientTableAccount* Get(int region, int logic_zone, int physics_zone, int32 id);
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
    static int Get(int region, int logic_zone, int physics_zone, int32 id, int64 &cbId, int (*cb)(int64, GORM_ClientTableAccount*));
    // 不带区服信息的函数
    static int Get(int32 id, int64 &cbId, int (*cb)(int64, GORM_ClientTableAccount*));
    
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
    static int Delete(int region, int logic_zone, int physics_zone, int32 id, int64 &cbId, int (*cb)(int64));
    // 不带区服信息的函数
    static int Delete(int32 id, int64 &cbId, int (*cb)(int64));
    
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
    
    // @desc 立即将本条数据持久化保存
    int SaveToDB();
    
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
    GORM_PB_Table_account *tablePbValue = nullptr;
    GORM_FieldsOpt fieldOpt;
    // 0:新的没有持久化的数据，1:从持久化存储拉到内存的数据，2:已经持久化了的并且有更新的数据
    int dirtyFlag = 0;  
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
