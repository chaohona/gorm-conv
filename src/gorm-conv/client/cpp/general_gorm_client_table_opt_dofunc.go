package cpp

import (
	"gorm-conv/common"
	"os"
	"strings"
)

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
    unique_lock<mutex> msgLk(sharedClientMsg->mtx);
    int code = sharedClientMsg->rspCode.code;
    
    return code;
`)
	return 0
}

func GeneralClientCPPCodes_GeneralGormClientTableOpt_CPP_Table_DoGet(table common.TableInfo, f *os.File, opt string) int {
	var bigOpt string = strings.ToUpper(opt)
	var bigTableName string = strings.ToUpper(table.Name)
	structName := "GORM_ClientTable" + common.CPP_TableStruct(table.Name)
	pbStructName := "GORM_PB_Table_" + table.Name

	f.WriteString("    GORM_ClientMsg *clientMsg = new GORM_ClientMsg();\n")
	f.WriteString("    clientMsg->tableId = GORM_PB_TABLE_IDX_" + bigTableName + ";\n")
	f.WriteString("    clientMsg->reqCmd = GORM_CMD_" + bigOpt + ";\n")
	f.WriteString("    clientMsg->fieldOpt = &this->fieldOpt;\n")
	f.WriteString("    GORM_PB_GET_REQ *getReq = new GORM_PB_" + bigOpt + "_REQ();\n")
	f.WriteString("    clientMsg->pbReqMsg = getReq;\n")
	f.WriteString("    clientMsg->getCBFunc = " + structName + "::GetCallBack;\n")
	f.WriteString("    GORM_PB_TABLE *pbTableAll = getReq->mutable_table();\n")

	f.WriteString("    shared_ptr<" + pbStructName + "> sharedPbValue = nullptr;\n")
	f.WriteString("    " + pbStructName + " *tmpPbValue = nullptr;\n")
	f.WriteString("    if (this->tablePbValue != nullptr)\n")
	f.WriteString("        tmpPbValue = this->tablePbValue;\n")
	f.WriteString("    else\n")
	f.WriteString("    {\n")
	f.WriteString("        sharedPbValue = make_shared<" + pbStructName + ">();\n")
	f.WriteString("        this->tablePbValue = sharedPbValue.get();\n")
	for _, colName := range table.SplitInfo.SplitCols {
		//f.WriteString("        sharedPbValue->set_" + colName + "(" + colName + ");\n")
		colStructName := common.CPP_TableColumnName(colName)
		f.WriteString("        this->Set" + colStructName + "(" + colName + ");\n")
	}
	f.WriteString("        this->tablePbValue = nullptr;\n")
	f.WriteString("        tmpPbValue = sharedPbValue.get();\n")
	f.WriteString("    }\n")

	f.WriteString("    pbTableAll->set_allocated_" + table.Name + "(tmpPbValue);\n\n")

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
    unique_lock<mutex> msgLk(sharedClientMsg->mtx);
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
	f.WriteString("    if (this->tablePbValue != nullptr) delete this->tablePbValue;\n")
	f.WriteString("    this->tablePbValue = pbRspMsg->mutable_table()->release_" + table.Name + "();\n")
	f.WriteString("    this->dirtyFlag = 1;\n")
	f.WriteString("\n")
	f.WriteString("    return GORM_OK;\n")
	return 0
}

func GeneralClientCPPCodes_GeneralGormClientTableOpt_CPP_Table_DoGetByIndex(table common.TableInfo, tableIndex common.TableIndex, f *os.File) int {
	var bigTableName string = strings.ToUpper(table.Name)
	structName := "GORM_ClientTable" + common.CPP_TableStruct(table.Name)
	pbStructName := "GORM_PB_Table_" + table.Name

	f.WriteString("    GORM_ClientMsg *clientMsg = new GORM_ClientMsg();\n")
	f.WriteString("    clientMsg->tableId = GORM_PB_TABLE_IDX_" + bigTableName + ";\n")
	f.WriteString("    clientMsg->reqCmd = GORM_CMD_GET_BY_NON_PRIMARY_KEY;\n")
	f.WriteString("    clientMsg->fieldOpt = &this->fieldOpt;\n")
	f.WriteString("    GORM_PB_GET_BY_NON_PRIMARY_KEY_REQ *getReq = new GORM_PB_GET_BY_NON_PRIMARY_KEY_REQ();\n")
	f.WriteString("    clientMsg->pbReqMsg = getReq;\n")
	f.WriteString("    clientMsg->getCBFunc = " + structName + "::GetCallBack;\n")
	f.WriteString("    GORM_PB_TABLE *pbTableAll = getReq->mutable_table();\n")

	f.WriteString("    shared_ptr<" + pbStructName + "> sharedPbValue = nullptr;\n")
	f.WriteString("    " + pbStructName + " *tmpPbValue = nullptr;\n")
	f.WriteString("    if (this->tablePbValue != nullptr)\n")
	f.WriteString("        tmpPbValue = this->tablePbValue;\n")
	f.WriteString("    else\n")
	f.WriteString("    {\n")
	f.WriteString("        sharedPbValue = make_shared<" + pbStructName + ">();\n")
	f.WriteString("        this->tablePbValue = sharedPbValue.get();\n")
	for _, colName := range tableIndex.IndexColumns {
		colStructName := common.CPP_TableColumnName(colName)
		f.WriteString("        this->Set" + colStructName + "(" + colName + ");\n")
	}
	f.WriteString("        this->tablePbValue = nullptr;\n")
	f.WriteString("        tmpPbValue = sharedPbValue.get();\n")
	f.WriteString("    }\n")

	f.WriteString("    pbTableAll->set_allocated_" + table.Name + "(tmpPbValue);\n\n")

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
    unique_lock<mutex> msgLk(sharedClientMsg->mtx);
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
	f.WriteString("    if (this->tablePbValue != nullptr) delete this->tablePbValue;\n")
	f.WriteString("    this->tablePbValue = pbRspMsg->mutable_table()->release_" + table.Name + "();\n")
	f.WriteString("    this->dirtyFlag = 1;\n")
	f.WriteString("\n")
	f.WriteString("    return GORM_OK;\n")
	return 0
}

// forIndexOrSplit 为1则是split，为2则是index
func GeneralClientCPPCodes_GeneralGormClientTableOpt_CPP_Table_DoGetVector(table common.TableInfo, forIndexOrSplit int, indexIndex int, f *os.File) {
	var bigTableName string = strings.ToUpper(table.Name)
	structName := "GORM_ClientTable" + common.CPP_TableStruct(table.Name)
	pbStructName := "GORM_PB_Table_" + table.Name

	f.WriteString("    GORM_ClientMsg *clientMsg = new GORM_ClientMsg();\n")
	f.WriteString("    vector<" + structName + "*> result;\n")
	f.WriteString("    " + structName + " forRequest;\n")
	f.WriteString("    clientMsg->tableId = GORM_PB_TABLE_IDX_" + bigTableName + ";\n")
	f.WriteString("    clientMsg->reqCmd = GORM_CMD_GET_BY_NON_PRIMARY_KEY;\n")
	f.WriteString("    clientMsg->fieldOpt = &forRequest->fieldOpt;\n")
	f.WriteString("    GORM_PB_GET_BY_NON_PRIMARY_KEY_REQ *getReq = new GORM_PB_GET_BY_NON_PRIMARY_KEY_REQ();\n")
	f.WriteString("    clientMsg->pbReqMsg = getReq;\n")
	f.WriteString("    clientMsg->getCBFunc = " + structName + "::GetCallBack;\n")
	f.WriteString("    GORM_PB_TABLE *pbTableAll = getReq->mutable_table();\n")

	f.WriteString("    shared_ptr<" + pbStructName + "> sharedPbValue = nullptr;\n")
	f.WriteString("    " + pbStructName + " *tmpPbValue = nullptr;\n")
	f.WriteString("    sharedPbValue = make_shared<" + pbStructName + ">();\n")
	f.WriteString("    forRequest->tablePbValue = sharedPbValue.get();\n")
	if forIndexOrSplit == 1 {
		for _, colName := range table.TableIndex[indexIndex].IndexColumns {
			colStructName := common.CPP_TableColumnName(colName)
			f.WriteString("    forRequest->Set" + colStructName + "(" + colName + ");\n")
		}
	} else {
		for _, colName := range table.SplitInfo.SplitCols {
			colStructName := common.CPP_TableColumnName(colName)
			f.WriteString("    forRequest->Set" + colStructName + "(" + colName + ");\n")
		}
	}

	f.WriteString("    forRequest->tablePbValue = nullptr;\n")
	f.WriteString("    tmpPbValue = sharedPbValue.get();\n")

	f.WriteString("    pbTableAll->set_allocated_" + table.Name + "(tmpPbValue);\n\n")

	f.WriteString("    if (GORM_OK != clientMsg->PackReq())\n")
	f.WriteString("    {\n")
	f.WriteString("        pbTableAll->release_" + table.Name + "();\n")
	f.WriteString("        delete clientMsg;\n")
	f.WriteString("        retCode = GORM_ERROR;\n")
	f.WriteString("        return result;\n")
	f.WriteString("    }\n")

	f.WriteString("    // 使用完，交还外部数据\n")
	f.WriteString("    pbTableAll->release_" + table.Name + "();\n")
	f.WriteString(`
    // 发送Get请求
    if (GORM_OK != GORM_ClientThreadPool::Instance()->SendRequest(clientMsg, clientMsg->cbId))
    {
        delete clientMsg;
        retCode = GORM_ERROR;
        return result;
    }
    clientMsg->Wait();	// 等待响应
    // 获取结果
    clientMsg = nullptr;
    for (int i=0;i<1000;i++)
    {
        if (GORM_OK != GORM_ClientThreadPool::Instance()->GetResponse(clientMsg))
        {
            retCode = GORM_ERROR;
			return result; 
        }
        if (clientMsg != nullptr)
        {
            break;
        }
        ThreadSleepMilliSeconds(1); 
    }
    if (clientMsg == nullptr)
    {
		retCode = GORM_ERROR;
		return result;  	
    }

    shared_ptr<GORM_ClientMsg> sharedClientMsg(clientMsg);
    unique_lock<mutex> msgLk(sharedClientMsg->mtx);
    if (GORM_OK != sharedClientMsg->rspCode.code)
    {
        return sharedClientMsg->rspCode.code;
    }
    GORM_PB_GET_BY_NON_PRIMARY_KEY_RSP *pbRspMsg = dynamic_cast<GORM_PB_GET_BY_NON_PRIMARY_KEY_RSP*>(sharedClientMsg->pbRspMsg);
    for(int i=0; i<pbRspMsg->tables_size(); i++)
    {
        auto pbTables = pbRspMsg->mutable_tables(i);
        if (!pbTables->has_currency())
        {
            continue;
        }
`)
	f.WriteString("        " + structName + " *nowTable = new " + structName + "();\n")
	f.WriteString("        nowTable->tablePbValue = pbTables->release_" + table.Name + "();\n")
	f.WriteString("        nowTable->dirtyFlag = 1;\n")
	f.WriteString("        result.push_back(nowTable);\n")
	f.WriteString("    }\n")
	f.WriteString("    return result;\n")
}
