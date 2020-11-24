package cpp

import (
	"gorm-conv/common"
	"os"
	"strings"
)

func GeneralClientCPPCodes_GeneralGormClientTableOpt_CPP_Table_DoFunc(table common.TableInfo, f *os.File, opt string) int {
	var bigOpt string = strings.ToUpper(opt)
	var bigTableName string = strings.ToUpper(table.Name)

	f.WriteString("    GORM_ClientMsg *clientMsg = nullptr;\n")
	f.WriteString("    clientMsg = new GORM_ClientMsg();\n")
	f.WriteString("{\n")
	f.WriteString("    clientMsg->mtx.lock();\n")
	if SupportCppCoroutine {
		f.WriteString("    clientMsg->container_ = GORM_Wrap::Instance()->GetContainer();\n")
	}
	f.WriteString("    clientMsg->tableId = GORM_PB_TABLE_IDX_" + bigTableName + ";\n")
	f.WriteString("    clientMsg->reqCmd = GORM_CMD_" + bigOpt + ";\n")
	f.WriteString("    clientMsg->fieldOpt = &this->fieldOpt;\n")
	f.WriteString("    GORM_PB_" + bigOpt + "_REQ getReq;\n")
	f.WriteString("    clientMsg->pbReqMsg = &getReq;\n")
	if bigOpt == "DELETE" {
		f.WriteString("    GORM_PB_TABLE *pbTableAll = getReq.mutable_table();\n")
	} else {
		f.WriteString("    GORM_PB_TABLE *pbTableAll = getReq.add_tables();\n")
	}

	f.WriteString("    pbTableAll->set_allocated_" + table.Name + "(this->tablePbValue);\n\n")
	f.WriteString("    if (GORM_OK != clientMsg->PackReq())\n")
	f.WriteString("    {\n")
	f.WriteString("        pbTableAll->release_" + table.Name + "();\n")
	f.WriteString("        clientMsg->mtx.unlock();\n")
	f.WriteString("        delete clientMsg;\n")
	f.WriteString("        return GORM_ERROR;\n")
	f.WriteString("    }\n")
	f.WriteString("    clientMsg->mtx.unlock();\n")

	f.WriteString("    // 使用完，交还外部数据\n")
	f.WriteString("    pbTableAll->release_" + table.Name + "();\n")
	f.WriteString("}// end of unique_lock<mutex> msgLk(clientMsg->mtx)\n")
	f.WriteString(`
    // 发送Get请求
	int sendResult = GORM_ClientThreadPool::Instance()->SendRequest(clientMsg);
    if (GORM_OK != sendResult)
    {
        delete clientMsg;
        return sendResult;
    }
`)

	// 支持协程
	if SupportCppCoroutine {
		f.WriteString("    clientMsg->YieldCo();\n")
	} else {
		f.WriteString(`
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
`)
	}

	f.WriteString("    shared_ptr<GORM_ClientMsg> sharedClientMsg(clientMsg);\n")
	f.WriteString("    unique_lock<mutex> msgLk(sharedClientMsg->mtx);\n")
	f.WriteString("    int code = sharedClientMsg->rspCode.code;\n")
	f.WriteString("    return code;\n")

	return 0
}

func GeneralClientCPPCodes_GeneralGormClientTableOpt_CPP_Table_DoGet(table common.TableInfo, f *os.File, opt string) int {
	var bigOpt string = strings.ToUpper(opt)
	var bigTableName string = strings.ToUpper(table.Name)
	structName := "GORM_ClientTable" + common.CPP_TableStruct(table.Name)
	pbStructName := "GORM_PB_Table_" + table.Name

	f.WriteString("    GORM_ClientMsg *clientMsg = nullptr;\n")
	f.WriteString("    clientMsg = new GORM_ClientMsg();\n")
	f.WriteString("{\n")
	f.WriteString("    clientMsg->mtx.lock();\n")
	if SupportCppCoroutine {
		f.WriteString("    clientMsg->container_ = GORM_Wrap::Instance()->GetContainer();\n")
	}
	f.WriteString("    clientMsg->tableId = GORM_PB_TABLE_IDX_" + bigTableName + ";\n")
	f.WriteString("    clientMsg->reqCmd = GORM_CMD_" + bigOpt + ";\n")
	f.WriteString("    clientMsg->fieldOpt = &this->fieldOpt;\n")
	f.WriteString("    GORM_PB_GET_REQ getReq;\n")
	f.WriteString("    clientMsg->pbReqMsg = &getReq;\n")
	f.WriteString("    clientMsg->getCBFunc = " + structName + "::GetCallBack;\n")
	f.WriteString("    GORM_PB_TABLE *pbTableAll = getReq.mutable_table();\n")

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
	f.WriteString("        clientMsg->mtx.unlock();\n")
	f.WriteString("        delete clientMsg;\n")
	f.WriteString("        return GORM_ERROR;\n")
	f.WriteString("    }\n")
	f.WriteString("    clientMsg->mtx.unlock();\n")

	f.WriteString("    // 使用完，交还外部数据\n")
	f.WriteString("    pbTableAll->release_" + table.Name + "();\n")
	f.WriteString("}// end of unique_lock<mutex> msgLk(clientMsg->mtx)")
	f.WriteString(`
    // 发送Get请求
    int sendResult = GORM_ClientThreadPool::Instance()->SendRequest(clientMsg);
    if (GORM_OK != sendResult)
    {
        delete clientMsg;
        return sendResult;
    }
    cbId = clientMsg->cbId;

`)
	if SupportCppCoroutine {
		f.WriteString("    clientMsg->YieldCo();\n")
	} else {
		f.WriteString(`
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
`)
	}

	f.WriteString(`
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

	f.WriteString("    GORM_ClientMsg *clientMsg = nullptr;\n")
	f.WriteString("    clientMsg = new GORM_ClientMsg();\n")
	f.WriteString("{\n")
	f.WriteString("    clientMsg->mtx.lock();\n")
	if SupportCppCoroutine {
		f.WriteString("    clientMsg->container_ = GORM_Wrap::Instance()->GetContainer();\n")
	}
	f.WriteString("    clientMsg->tableId = GORM_PB_TABLE_IDX_" + bigTableName + ";\n")
	f.WriteString("    clientMsg->reqCmd = GORM_CMD_GET_BY_NON_PRIMARY_KEY;\n")
	f.WriteString("    clientMsg->fieldOpt = &this->fieldOpt;\n")
	f.WriteString("    GORM_PB_GET_BY_NON_PRIMARY_KEY_REQ getReq;\n")
	f.WriteString("    clientMsg->pbReqMsg = &getReq;\n")
	f.WriteString("    clientMsg->getCBFunc = " + structName + "::GetCallBack;\n")
	f.WriteString("    GORM_PB_TABLE *pbTableAll = getReq.add_tables();\n")

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
	f.WriteString("        clientMsg->mtx.unlock();\n")
	f.WriteString("        delete clientMsg;\n")
	f.WriteString("        return GORM_ERROR;\n")
	f.WriteString("    }\n")
	f.WriteString("    clientMsg->mtx.unlock();\n")

	f.WriteString("    // 使用完，交还外部数据\n")
	f.WriteString("    pbTableAll->release_" + table.Name + "();\n")
	f.WriteString("} // end of unique_lock<mutex> msgLk(clientMsg->mtx)\n")
	f.WriteString(`
    // 发送Get请求
    int sendResult = GORM_ClientThreadPool::Instance()->SendRequest(clientMsg);
    if (GORM_OK != sendResult)
    {
        delete clientMsg;
        return sendResult;
    }

`)
	if SupportCppCoroutine {
		f.WriteString("    clientMsg->YieldCo();\n")
	} else {
		f.WriteString(`
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
`)
	}

	f.WriteString(`
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

	f.WriteString("    GORM_ClientMsg *clientMsg = nullptr;\n")
	f.WriteString("    clientMsg = new GORM_ClientMsg();\n")
	f.WriteString("{\n")
	f.WriteString("    clientMsg->mtx.lock();\n")
	f.WriteString("    " + structName + " forRequest;\n")
	if SupportCppCoroutine {
		f.WriteString("    clientMsg->container_ = GORM_Wrap::Instance()->GetContainer();\n")
	}
	f.WriteString("    clientMsg->tableId = GORM_PB_TABLE_IDX_" + bigTableName + ";\n")
	f.WriteString("    clientMsg->reqCmd = GORM_CMD_GET_BY_NON_PRIMARY_KEY;\n")
	f.WriteString("    clientMsg->fieldOpt = &forRequest.fieldOpt;\n")
	f.WriteString("    GORM_PB_GET_BY_NON_PRIMARY_KEY_REQ getReq;\n")
	f.WriteString("    clientMsg->pbReqMsg = &getReq;\n")
	f.WriteString("    clientMsg->limitNum = GORM_MAX_LIMIT_NUM;\n")
	f.WriteString("    clientMsg->getCBFunc = " + structName + "::GetCallBack;\n")
	f.WriteString("    GORM_PB_TABLE *pbTableAll = getReq.add_tables();\n")

	f.WriteString("    shared_ptr<" + pbStructName + "> sharedPbValue = nullptr;\n")
	f.WriteString("    " + pbStructName + " *tmpPbValue = nullptr;\n")
	f.WriteString("    sharedPbValue = make_shared<" + pbStructName + ">();\n")
	f.WriteString("    forRequest.tablePbValue = sharedPbValue.get();\n")
	if forIndexOrSplit == 2 {
		for _, colName := range table.TableIndex[indexIndex].IndexColumns {
			colStructName := common.CPP_TableColumnName(colName)
			f.WriteString("    forRequest.Set" + colStructName + "(" + colName + ");\n")
		}
	} else {
		for _, colName := range table.SplitInfo.SplitCols {
			colStructName := common.CPP_TableColumnName(colName)
			f.WriteString("    forRequest.Set" + colStructName + "(" + colName + ");\n")
		}
	}

	f.WriteString("    forRequest.tablePbValue = nullptr;\n")
	f.WriteString("    tmpPbValue = sharedPbValue.get();\n")

	f.WriteString("    pbTableAll->set_allocated_" + table.Name + "(tmpPbValue);\n\n")

	f.WriteString("    if (GORM_OK != clientMsg->PackReq())\n")
	f.WriteString("    {\n")
	f.WriteString("        pbTableAll->release_" + table.Name + "();\n")
	f.WriteString("        clientMsg->mtx.unlock();\n")
	f.WriteString("        delete clientMsg;\n")
	f.WriteString("        return GORM_PACK_REQ_ERROR;\n")
	f.WriteString("    }\n")
	f.WriteString("    clientMsg->mtx.unlock();\n")

	f.WriteString("    // 使用完，交还外部数据\n")
	f.WriteString("    pbTableAll->release_" + table.Name + "();\n")
	f.WriteString("}// end of unique_lock<mutex> msgLk(clientMsg->mtx)\n")
	f.WriteString(`
    // 发送Get请求
    int sendResult = GORM_ClientThreadPool::Instance()->SendRequest(clientMsg);
    if (GORM_OK != sendResult)
    {
        delete clientMsg;
        return sendResult;
    }

`)
	if SupportCppCoroutine {
		f.WriteString("    clientMsg->YieldCo();\n")
	} else {
		f.WriteString(`
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
    {
		return GORM_ERROR;  	
    }
`)
	}

	f.WriteString(`
    shared_ptr<GORM_ClientMsg> sharedClientMsg(clientMsg);
    unique_lock<mutex> msgLk(sharedClientMsg->mtx);
    if (GORM_OK != sharedClientMsg->rspCode.code)
    {
		return sharedClientMsg->rspCode.code;  
    }
    GORM_PB_GET_BY_NON_PRIMARY_KEY_RSP *pbRspMsg = dynamic_cast<GORM_PB_GET_BY_NON_PRIMARY_KEY_RSP*>(sharedClientMsg->pbRspMsg);
    if (pbRspMsg == nullptr)
    	return GORM_OK;
    for(int i=0; i<pbRspMsg->tables_size(); i++)
    {
        auto pbTables = pbRspMsg->mutable_tables(i);
        if (!pbTables->has_currency())
        {
            continue;
        }
`)

	f.WriteString("        shared_ptr<" + structName + "> nowTable = make_shared<" + structName + ">();\n")
	f.WriteString("        nowTable->tablePbValue = pbTables->release_" + table.Name + "();\n")
	f.WriteString("        nowTable->dirtyFlag = 1;")
	f.WriteString("        outResult.push_back(nowTable);\n")
	f.WriteString("    }\n")
	f.WriteString("    return GORM_OK;\n")
}

// forIndexOrSplit 为1则是split，为2则是index
func GeneralClientCPPCodes_GeneralGormClientTableOpt_CPP_Table_DoGetAllRows(table common.TableInfo, f *os.File) {
	var bigTableName string = strings.ToUpper(table.Name)
	structName := "GORM_ClientTable" + common.CPP_TableStruct(table.Name)

	f.WriteString("    GORM_ClientMsg *clientMsg = nullptr;\n")
	f.WriteString("    clientMsg = new GORM_ClientMsg();\n")
	f.WriteString("{\n")
	f.WriteString("    clientMsg->mtx.lock();\n")
	f.WriteString("    " + structName + " forRequest;\n")
	if SupportCppCoroutine {
		f.WriteString("    clientMsg->container_ = GORM_Wrap::Instance()->GetContainer();\n")
	}
	f.WriteString("    clientMsg->reqFlag |= GORM_REQ_REF_TO_TABLE_INDEX;\n")
	f.WriteString("    clientMsg->refTableIndex = tableIndex;\n")
	f.WriteString("    clientMsg->tableId = GORM_PB_TABLE_IDX_" + bigTableName + ";\n")
	f.WriteString("    clientMsg->reqCmd = GORM_CMD_GET_BY_NON_PRIMARY_KEY;\n")
	f.WriteString("    clientMsg->fieldOpt = &forRequest.fieldOpt;\n")
	f.WriteString("    GORM_PB_GET_BY_NON_PRIMARY_KEY_REQ getReq;\n")
	f.WriteString("    clientMsg->pbReqMsg = &getReq;\n")
	f.WriteString("    clientMsg->limitNum = GORM_MAX_LIMIT_NUM;\n")
	f.WriteString("    clientMsg->getCBFunc = " + structName + "::GetCallBack;\n")
	f.WriteString("    if (GORM_OK != clientMsg->PackReq())\n")
	f.WriteString("    {\n")
	f.WriteString("        clientMsg->mtx.unlock();\n")
	f.WriteString("        delete clientMsg;\n")
	f.WriteString("        return GORM_PACK_REQ_ERROR;\n")
	f.WriteString("    }\n")
	f.WriteString("    clientMsg->mtx.unlock();\n")

	f.WriteString("}// end of unique_lock<mutex> msgLk(clientMsg->mtx)\n")
	f.WriteString(`
    // 发送Get请求
    int sendResult = GORM_ClientThreadPool::Instance()->SendRequest(clientMsg);
    if (GORM_OK != sendResult)
    {
        delete clientMsg;
        return sendResult;
    }

`)
	if SupportCppCoroutine {
		f.WriteString("    clientMsg->YieldCo();\n")
	} else {
		f.WriteString(`
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
    {
		return GORM_ERROR;  	
    }
`)
	}
	f.WriteString(`
    shared_ptr<GORM_ClientMsg> sharedClientMsg(clientMsg);
    unique_lock<mutex> msgLk(sharedClientMsg->mtx);
    if (GORM_OK != sharedClientMsg->rspCode.code)
    {
		return sharedClientMsg->rspCode.code;
    }
    GORM_PB_GET_BY_NON_PRIMARY_KEY_RSP *pbRspMsg = dynamic_cast<GORM_PB_GET_BY_NON_PRIMARY_KEY_RSP*>(sharedClientMsg->pbRspMsg);
    if (pbRspMsg == nullptr)
    	return GORM_OK;
    for(int i=0; i<pbRspMsg->tables_size(); i++)
    {
        auto pbTables = pbRspMsg->mutable_tables(i);
        if (!pbTables->has_currency())
        {
            continue;
        }
`)

	f.WriteString("        shared_ptr<" + structName + "> nowTable = make_shared<" + structName + ">();\n")
	f.WriteString("        nowTable->tablePbValue = pbTables->release_" + table.Name + "();\n")
	f.WriteString("        nowTable->dirtyFlag = 1;")
	f.WriteString("        outResult.push_back(nowTable);\n")
	f.WriteString("    }\n")
	f.WriteString("    return GORM_OK;\n")
}
