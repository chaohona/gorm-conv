package cpp

import (
	"fmt"
	"gorm-conv/common"
	"os"
)

func GeneralClientCPPCodes_GenralGormH(games []common.XmlCfg, outpath string) int {
	outfile := outpath + "/gorm.h"
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

	f.WriteString(`#ifndef _GORM_H__
#define _GORM_H__

#include "gorm_define.h"
#include "gorm_wrap.h"

`)

	// 包含所有的文件
	for _, game := range games {
		fileName := game.File[:len(game.File)-4]
		headerName := "gorm_client_table_opt_" + fileName + ".h"
		f.WriteString("#include \"" + headerName + "\"\n")
	}

	f.WriteString("\n")
	f.WriteString("#endif")

	return 0
}
