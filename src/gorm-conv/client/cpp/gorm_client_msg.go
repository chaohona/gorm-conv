package cpp

import (
	"fmt"
	"os"
)

func General_GormClientMsg(outpath string) int {
	outfile := outpath + "/gorm_client_msg.cc"
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

	f.WriteString(`
#include "gorm_client_msg.h"
#include "gorm_wrap.h"
#include "gorm_pb_proto.pb.h"
#include "gorm_pb_tables_inc.pb.h"
#include "gorm_msg_helper.h"


namespace gorm{

GORM_ClientMsg::~GORM_ClientMsg()
{
    this->Reset();
}

void GORM_ClientMsg::Reset()
{
    if (this->reqMemData != nullptr)
    {
        this->reqMemData->Release();
        this->reqMemData = nullptr;
    }
    if (this->pbRspMsg != nullptr)
    {
        delete this->pbRspMsg;
        this->pbRspMsg = nullptr;
    }
    if (this->pbReqMsg != nullptr)
    {
        delete this->pbReqMsg;
        this->pbReqMsg = nullptr;
    }
    sendChannelId = -1;   
    region = 0;
    logicZone = 0;
    physicZone = 0;
    reqCmd = GORM_CMD_INVALID;
    tableId = GORM_PB_TABLE_IDX_MIN__;
    verPolicy = NOCHECKDATAVERSION_AUTOINCREASE;
    reqFlag = GORM_ResultFlag_RETURN_CODE;
    cbId = 0;
    reqMemData = nullptr;
    pbReqMsg = nullptr;
    pbRspMsg = nullptr;
    rspCode.Reset();
    syncCall = 0;
    fieldOpt = nullptr;
}

int GORM_ClientMsg::ParseRsp(char *msgBeginPos, int msgLen)
{
    // 组装消息
    switch (reqCmd)
    {
    case GORM_CMD_GET:
    {
        return this->ParseRspGet(msgBeginPos, msgLen);
    }
    }
    
    return GORM_OK;
}

int GORM_ClientMsg::ParseRspGet(char *msgBeginPos, int msgLen)
{
    GORM_PB_GET_RSP *pbGetMsg = new GORM_PB_GET_RSP();
    this->pbRspMsg = pbGetMsg;
    if (!this->pbRspMsg->ParseFromArray(msgBeginPos, msgLen))
    {
        this->rspCode.code = GORM_RSP_UNPACK_FAILED;
        return GORM_ERROR;
    }
    if (pbGetMsg->has_retcode())
    {
        const GORM_PB_Ret_Code &pbRetCode = pbGetMsg->retcode();
        this->rspCode.code = pbRetCode.code();
        if (pbRetCode.dbcode() != 0)
        {
            this->rspCode.dbError = pbRetCode.dbcode();
            this->rspCode.dbErrorInfo[0] = '\n';
            const string &dbErrInfo = pbRetCode.dberrinfo();
            int errMsgLen = dbErrInfo.length();
            if (errMsgLen >= GORM_MAX_DB_ERR_INFO)
            {
                errMsgLen = GORM_MAX_DB_ERR_INFO - 1;
            }
            if (errMsgLen >0)
            {
                memcpy(this->rspCode.dbErrorInfo, dbErrInfo.c_str(), errMsgLen);
                this->rspCode.dbErrorInfo[errMsgLen] = '\n';
            }
        }
    }
    if (this->rspCode.code != GORM_OK)
    {
        return this->rspCode.code;
    }
    if (!pbGetMsg->has_table())
    {
        return GORM_OK;
    }

    return GORM_OK;
}

#define GORM_CLIENTREQUEST_SETHEADER()                                          \
if (fieldOpt->iUsedIdx<0) return GORM_ERROR;                                    \
GORM_PB_REQ_HEADER *header = pPbReq->mutable_header();                          \
header->set_reqflag(reqFlag);                                                   \
header->set_fieldmode(fieldOpt->szFieldCollections, fieldOpt->iUsedIdx+1);      \
header->set_tableid(tableId);                                                   \
header->set_verpolice(verPolicy);

int GORM_ClientMsg::PackReq()
{
    // 组装消息
    switch (reqCmd)
    {
    case GORM_CMD_INSERT:
    {
        return this->PackReqInsert();
    }
    case GORM_CMD_GET:
    {
        return this->PackReqGet();
    }
    case GORM_CMD_DELETE:
    {
        return this->PackReqDelete();
    }
    case GORM_CMD_UPDATE:
    {
        return this->PackReqUpdate();
    }
    case GORM_CMD_REPLACE:
    {
        return this->PackReplace();
    }
    }
    
    return GORM_OK;
}

int GORM_ClientMsg::PackReqInsert()
{
    GORM_PB_INSERT_REQ  *pPbReq = dynamic_cast<GORM_PB_INSERT_REQ*>(pbReqMsg);
    GORM_CLIENTREQUEST_SETHEADER();
    
    return this->MakeSendBuff();
}

int GORM_ClientMsg::PackReqDelete()
{
    GORM_PB_DELETE_REQ  *pPbReq = dynamic_cast<GORM_PB_DELETE_REQ*>(pbReqMsg);
    GORM_CLIENTREQUEST_SETHEADER();
    
    return this->MakeSendBuff();
}

int GORM_ClientMsg::PackReqGet()
{
    GORM_PB_GET_REQ  *pPbReq = dynamic_cast<GORM_PB_GET_REQ*>(pbReqMsg);
	GORM_CLIENTREQUEST_SETHEADER();
    
    return this->MakeSendBuff();
}

int GORM_ClientMsg::PackReqUpdate()
{
    GORM_PB_UPDATE_REQ  *pPbReq = dynamic_cast<GORM_PB_UPDATE_REQ*>(pbReqMsg);
    GORM_CLIENTREQUEST_SETHEADER();
    
    return this->MakeSendBuff();
}

int GORM_ClientMsg::PackReplace()
{
    GORM_PB_REPLACE_REQ  *pPbReq = dynamic_cast<GORM_PB_REPLACE_REQ*>(pbReqMsg);
    GORM_CLIENTREQUEST_SETHEADER();
    
    return this->MakeSendBuff();
}

int GORM_ClientMsg::MakeSendBuff()
{
    // 3、打包数据到buffer
    size_t sPbSize = pbReqMsg->ByteSizeLong() + GORM_REQ_MSG_HEADER_LEN;
    auto wrapInstance = GORM_Wrap::Instance();
    this->cbId = ++(wrapInstance->seqIdx);
    if (cbId == 0x7FFFFFFF)
    {
        wrapInstance->seqIdx = 0;
    }
    reqMemData = wrapInstance->memPool->GetData(sPbSize);
    if (reqMemData == nullptr)
    {
        GORM_CUSTOM_LOGE(logHandle, "pack request, get buffer failed, buffer size:%d", sPbSize);
        return GORM_ERROR;
    }
    reqMemData->m_pMemPool = wrapInstance->memPool;

    // 设置发送消息头
    GORM_SET_REQ_PRE_HEADER(reqMemData->m_uszData, sPbSize, reqCmd, cbId, 0);
    // 压缩pb数据到内存
    if (!pbReqMsg->SerializeToArray(reqMemData->m_uszData + GORM_REQ_MSG_HEADER_LEN, sPbSize-GORM_REQ_MSG_HEADER_LEN))
    {
        GORM_CUSTOM_LOGD(logHandle, "serialize msg to buffer failed.");
        return GORM_ERROR;
    }
    reqMemData->m_sUsedSize = sPbSize;

    return GORM_OK;
}


}

`)

	return 0
}
