package main

import (
	"fmt"
	"os"
)

func GetFBString(c TableColumn) string {
	switch c.Type {
	case "int":
		return c.Name + ":int;"
	case "uint64":
		return c.Name + ":uint64;"
	case "string":
		return c.Name + ":[uint8];"
	case "blob":
		return c.Name + ":[uint8];"
	}
	return ""
}

func OutPutFB(table TableInfo, f *os.File) int {
	// 1、输出头
	/*
		table GORM_Table_account{
	*/
	_, err := f.WriteString(string("table GORM_FB_Table_"))
	if err != nil {
		fmt.Println(err.Error())
	}
	_, err = f.WriteString(table.Name)
	if err != nil {
		fmt.Println(err.Error())
	}
	_, err = f.WriteString(string(" {\n"))
	if err != nil {
		fmt.Println(err.Error())
	}
	// 2、输出结构体
	// 输出version
	_, err = f.WriteString(string("    version:ulong;\n"))
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, column := range table.TableColumns {
		/*if column.Name == "version" {
			fmt.Println("there should not version columns")
			return -1
		}*/
		_ = column
		c := GetFBString(column)
		_, err = f.WriteString(string("    "))
		if err != nil {
			fmt.Println(err.Error())
		}
		_, err = f.WriteString(c)
		if err != nil {
			fmt.Println(err.Error())
		}
		_, err = f.WriteString(string("\n"))
		if err != nil {
			fmt.Println(err.Error())
		}
	}
	// 3、输出结尾
	/*
		}
	*/
	_, err = f.WriteString(string("}\n\n"))
	if err != nil {
		fmt.Println(err.Error())
	}
	return 0
}

func PrintFBFileHeader(f *os.File) int {
	/*
		namespace gorm;

		table GORM_Text {
			Capacity:int;
			Size:int;
			Data:[uint8];
		}

		table GORM_Ret_Code {
			Code:ushort;
		}
	*/
	var header string = `namespace gorm;
`
	_, err := f.WriteString(header)
	if err != nil {
		fmt.Println(err.Error())
	}

	return 0
}

func PrintFBFileTail(f *os.File) int {
	var tail string = `
table GORM_FB_Text {
    Capacity:int;
    Size:int;
    Data:[uint8];
}

table GORM_FB_Ret_Code {
    Code:ushort;
}

table GORM_FB_REQ_HEADER {
    TableType:ushort;
    BusinessID:uint;
    VerPolice:uint8;
}

table GORM_FB_HGET_REQ {
    Header:GORM_FB_REQ_HEADER;
    Columns:[ushort];
    Where:[uint8];
}

table GORM_FB_HGET_RSP {
    RetCode:GORM_FB_Ret_Code;
    Column:GORM_FB_Tables;
}

table GORM_FB_HDEL_REQ {
    Header:GORM_FB_REQ_HEADER;
}

table GORM_FB_HDEL_RSP {
    RetCode:GORM_FB_Ret_Code;
    Count:int;
}

table GORM_FB_HSET_REQ {
    Header:GORM_FB_REQ_HEADER;
    Columns:GORM_FB_Tables;
}

table GORM_FB_HSET_RSP {
    RetCode:GORM_FB_Ret_Code;
    InsertID:ulong;
}

table GORM_FB_COUNT_REQ {
    Header:GORM_FB_REQ_HEADER;
    Where:[uint8];
}

table GORM_FB_HADD_REQ {
    Header:GORM_FB_REQ_HEADER;
    Columns:GORM_FB_Tables;
}

table GORM_FB_HADD_RSP {
    RetCode:GORM_FB_Ret_Code;
    InsertID:ulong;
}

table GORM_FB_COUNT_RSP {
    RetCode:GORM_FB_Ret_Code;
    Count:uint;
}

table GORM_FB_SELECT_REQ {
    Header:GORM_FB_REQ_HEADER;
    ColID:[ushort];
    Where:[uint8];
    Limit:ushort;
}

table GORM_FB_SELECT_RSP {
    RetCode:GORM_FB_Ret_Code;
    TotalCnt:uint;
    NowCnt:uint;
    Columns:[GORM_FB_Tables];
}

table GORM_FB_DELETE_REQ {
    Header:GORM_FB_REQ_HEADER;
    Where:[uint8];
}

table GORM_FB_DELETE_RSP {
    RetCode:GORM_FB_Ret_Code;
    Count:uint;
}

table GORM_FB_UPDATE_REQ {
    Header:GORM_FB_REQ_HEADER;
    Where:[uint8];
}

table GORM_FB_UPDATE_RSP {
    RetCode:GORM_FB_Ret_Code;
    Count:ulong;
}

table GORM_FB_INSERT_REQ {
    Header:GORM_FB_REQ_HEADER;
    Columns:[GORM_FB_Tables];
}

table GORM_FB_INSERT_RSP {
    RetCode:GORM_FB_Ret_Code;
    InsertID:[ulong];
}

union GORM_FB_REQ {
    Update:GORM_FB_UPDATE_REQ,
    Delete:GORM_FB_DELETE_REQ,
    Select:GORM_FB_SELECT_REQ,
    Insert:GORM_FB_INSERT_REQ,
    Count:GORM_FB_COUNT_REQ,
    HGet:GORM_FB_HGET_REQ,
    HSet:GORM_FB_HSET_REQ,
}

union GORM_FB_RSP {
    Update:GORM_FB_UPDATE_RSP,
    Delete:GORM_FB_DELETE_RSP,
    Select:GORM_FB_SELECT_RSP,
    Insert:GORM_FB_INSERT_RSP,
    Count:GORM_FB_COUNT_RSP,
    HGet:GORM_FB_HGET_RSP,
    HSet:GORM_FB_HSET_RSP,
}

table GORM_FB_MSG {
    REQ:GORM_FB_REQ;
    RSP:GORM_FB_RSP;
}

root_type GORM_FB_MSG;
	`
	_, err := f.WriteString(tail)
	if err != nil {
		fmt.Println(err.Error())
	}
	return 0
}

func GeneralFBFile(game XmlCfg, outpath string) int {
	fileLen := len(game.File)
	outfile := outpath + "/" + game.File[:fileLen-4] + ".fbs"
	f, err := os.OpenFile(outfile, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer f.Close()
	err = f.Truncate(0)
	if err != nil {
		fmt.Println(err.Error())
	}
	// 1、输出开头

	if 0 != PrintFBFileHeader(f) {
		return -1
	}
	var tableList []string
	// 2、输出中间表对应的结构体
	for _, table := range game.DB.TableList {
		if ret := OutPutFB(table, f); ret != 0 {
			fmt.Println("generate fb file failed:" + outfile)
			return ret
		}
		tableList = append(tableList, "GORM_FB_Table_"+table.Name)
	}
	// 3、输出表的列表
	/*
		union GORM_Tables {
		    GORM_Table_account,
		    GORM_Table_bag,
		}

	*/
	_, err = f.WriteString(string("union GORM_FB_Tables {\n"))
	if err != nil {
		fmt.Println(err.Error())
	}
	for idx, name := range tableList {
		_, err = f.WriteString(string("    "))
		if err != nil {
			fmt.Println(err.Error())
		}
		_, err = f.WriteString(game.DB.TableList[idx].Name)
		if err != nil {
			fmt.Println(err.Error())
		}
		_, err = f.WriteString(string(":"))
		if err != nil {
			fmt.Println(err.Error())
		}
		_, err = f.WriteString(name)
		if err != nil {
			fmt.Println(err.Error())
		}
		_, err = f.WriteString(string(",\n"))
		if err != nil {
			fmt.Println(err.Error())
		}
	}
	_, err = f.WriteString(string("}\n\n"))
	if err != nil {
		fmt.Println(err.Error())
	}
	// 4、输出结尾
	if 0 != PrintFBFileTail(f) {
		return -1
	}
	return 0
}

func GeneralFBFiles(games []XmlCfg, outpath string) int {
	for _, game := range games {
		if 0 != GeneralFBFile(game, outpath) {
			return -1
		}
	}
	return 0
}
