package cpp

import (
	"fmt"
	"gorm-conv/common"
	gorm_cpp "gorm-conv/cpp"
	"os"
)

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
#include "gorm_client_msg.h"

#include <mutex>

namespace gorm{

`
	_, err = f.WriteString(header)
	if err != nil {
		fmt.Println(err.Error())
		return -1
	}

	if 0 != gorm_cpp.GORM_InitTableSchemaInfo(games, f) {
		fmt.Println("GORM_InitTableSchemaInfo failed.")
		return -1
	}

	f.WriteString(`

GORM_ClientMsg *GORM_GetHandShakeMessage(int &iResult)
{
    static GORM_ClientMsg clientMsg;
    static std::once_flag flag;
    static int packResult = GORM_OK;
    call_once(flag, [&](){
        unique_lock<mutex> lck(clientMsg.mtx);
        clientMsg.reqCmd = GORM_CMD_HAND_SHAKE;
        GORM_PB_HAND_SHAKE_REQ handShakeReq;
        clientMsg.pbReqMsg = &handShakeReq;
        if (GORM_OK != GORM_InitTableSchemaInfo(&handShakeReq))
        {
            packResult = GORM_ERROR;
        }
        if (GORM_OK != clientMsg.PackReq())
        {
            packResult = GORM_ERROR;
        }
    });

    iResult = packResult;
    return &clientMsg;
}


`)

	f.WriteString("\n}\n")
	return 0
}
