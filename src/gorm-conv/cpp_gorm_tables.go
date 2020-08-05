package main

// 生成gorm_tables.h

import (
	"fmt"
	"os"
)

func CppGormTables_H(games []XmlCfg, outpath string) int {
	outfile := outpath + "gorm_tables.h"
	f, err := os.OpenFile(outfile, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err.Error())
		return -1
	}
	defer f.Close()
	fmt.Println("begin to general file:" + outfile)
	err = f.Truncate(0)

	// 1、输出固定的头/////////////////////////
	//#include "mysql.h"
	var header string = `#ifndef _GORM_TABLES_H__
#define _GORM_TABLES_H__

#include "gorm_table_field_map_define.h"
#include "gorm_pb_proto.pb.h"
#include "gorm_table_field_map.h"
#include "gorm_pb_tables_inc.pb.h"
#include "gorm_msg_helper.h"

`

	f.WriteString(header)

	for _, game := range games {
		_ = game
		file := game.File[0 : len(game.File)-4]
		f.WriteString(`#include "` + file + `.pb.h"`)
		f.WriteString("\n")
	}

	f.WriteString("\n#endif")
	return 0
}
