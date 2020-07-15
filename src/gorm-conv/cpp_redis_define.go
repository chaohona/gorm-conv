package main

import (
	"fmt"
	"os"
	"strings"
)

func CppRedisDefine_Opt(opt string, games []XmlCfg, f *os.File) int {
	opt = strings.ToUpper(opt)
	f.WriteString("\n")
	f.WriteString("int GORM_REDIS_" + opt + "(int iTableId, redisContext *pRedisConn, GORM_PB_TABLE *pPbTable)\n")
	f.WriteString("{\n")
	f.WriteString("    switch (iTableId)\n")
	f.WriteString("    {\n")
	for _, game := range games {
		for _, table := range game.DB.TableList {
			f.WriteString("    case GORM_PB_TABLE_IDX_" + strings.ToUpper(table.Name) + ":\n")
			f.WriteString("    {\n")
			f.WriteString("        if (!pPbTable->has_" + table.Name + "())\n")
			f.WriteString("            return GORM_REQ_NO_RECORDS;\n")
			f.WriteString("        return GORM_REDIS_" + opt + "_" + strings.ToUpper(table.Name) + "(pPbTable->" + table.Name + "(), pRedisConn);\n")
			f.WriteString("    }\n")
		}
	}
	f.WriteString("    default:\n")
	f.WriteString("        return GORM_OK;\n")
	f.WriteString("    }\n")
	f.WriteString("}\n")
	return 0
}

func CppRedisDefine(games []XmlCfg, outpath string) int {
	outfile := outpath + "gorm_redis_define.cc"
	f, err := os.OpenFile(outfile, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err.Error())
		return -1
	}
	defer f.Close()
	fmt.Println("begin to general file:" + outfile)
	err = f.Truncate(0)

	var header string = `#include "gorm_redis_define.h"

using namespace gorm;
`

	f.WriteString(header)

	if 0 != CppRedisDefine_Hmget(games, f) {
		fmt.Println("CppRedisDefine_Hmget failed.")
		return -1
	}
	if 0 != CppRedisDefine_Insert(games, f) {
		fmt.Println("CppRedisDefine_Insert failed.")
		return -1
	}
	if 0 != CppRedisDefine_Delete(games, f) {
		fmt.Println("CppRedisDefine_Delete failed.")
		return -1
	}
	return 0
}
