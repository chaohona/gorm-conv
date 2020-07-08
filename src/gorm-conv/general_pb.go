package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func GetPBString(c TableColumn) string {
	switch c.Type {
	case "double":
		{
			return "double " + c.Name
		}
	case "int8", "uint8":
		{
			return "int32 " + c.Name
		}
	case "int", "int32":
		{
			//return "sfixed32 " + c.Name
			return "int32 " + c.Name
		}
	case "uint32":
		{
			//return "fixed32" + c.Name
			return "uint32" + c.Name
		}
	case "long", "int64":
		{
			//return "sfixed64 " + c.Name
			return "int64 " + c.Name
		}
	case "uint64":
		{
			//return "fixed64 " + c.Name
			return "uint64 " + c.Name
		}
	case "string", "char":
		{
			return "string " + c.Name
		}
	case "blob":
		{
			return "string " + c.Name
		}
	}
	return ""
}

func OutPutPBTable(table TableInfo, f *os.File) int {
	/*version := "    fixed64	version = 1;\n"
	_, err := f.WriteString(version)
	if err != nil {
		fmt.Println(err.Error())
		return -1
	}*/
	var index int64 = 0
	for _, c := range table.TableColumns {
		index += 1
		pbStr := GetPBString(c)
		if pbStr == "" {
			fmt.Println("general pb file failed, table:", table.Name, ", column:", c.Name)
			return -1
		}
		str := "    " + pbStr + " = "
		str += strconv.FormatInt(index, 10)
		str += ";\n"
		_, err := f.WriteString(str)
		if err != nil {
			fmt.Println(err.Error())
			return -1
		}
	}
	return 0
}

func GeneralPBFile(game XmlCfg, outpath string) int {
	fileLen := len(game.File)
	outfile := outpath + "/" + game.File[:fileLen-4] + ".proto"
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
	var header string = `syntax = "proto3";
	package gorm;

	`
	_, err = f.WriteString(header)
	if err != nil {
		fmt.Println(err.Error())
		return -1
	}

	// 输出表结构
	/*
		   message GORM_PB_Table_account
		   {
				uint64	version = 1;
				int id = 2;
				string account = 3;
				string allbinary = 4;
		   }
	*/

	for _, table := range game.DB.TableList {
		msgHeader := "message GORM_PB_Table_" + table.Name + "{\n"
		_, err = f.WriteString(msgHeader)
		if err != nil {
			fmt.Println(err.Error())
			return -1
		}
		if 0 != OutPutPBTable(table, f) {
			fmt.Println("genreral message for table failed:" + table.Name)
			return -1
		}
		_, err = f.WriteString("}\n\n")
		if err != nil {
			fmt.Println(err.Error())
			return -1
		}
	}

	return 0
}

func GeneralPBColumnIndex(games []XmlCfg, f *os.File) int {
	for _, game := range games {
		for _, table := range game.DB.TableList {
			var colIndex int64 = 0
			var enum string = "enum GORM_PB_"
			enum += strings.ToUpper(table.Name)
			enum += "_FIELD_INDEX\n{\n"
			_, err := f.WriteString(enum)
			if err != nil {
				fmt.Println(err.Error())
				return -1
			}

			outPre := "	GORM_PB_FIELD_" + strings.ToUpper(table.Name) + "_"
			out := outPre
			_, err = f.WriteString(out)
			if err != nil {
				fmt.Println(err.Error())
				return -1
			}
			for _, col := range table.TableColumns {
				strColIndex := strconv.FormatInt(int64(colIndex), 10)
				colIndex += 1
				out = outPre + strings.ToUpper(col.Name)
				out += " = "
				out += strColIndex
				out += ";\n"
				_, err := f.WriteString(out)
				if err != nil {
					fmt.Println(err.Error())
					return -1
				}
			}
			out = "}\n\n"
			_, err = f.WriteString(out)
			if err != nil {
				fmt.Println(err.Error())
				return -1
			}
		}
	}
	return 0
}

//
func GeneralTablesInc(games []XmlCfg, outpath string) int {
	outfile := outpath + "gorm_pb_tables_inc.proto"
	f, err := os.OpenFile(outfile, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err.Error())
		return -1
	}
	defer f.Close()
	fmt.Println("begin to general " + outfile)
	err = f.Truncate(0)
	if err != nil {
		fmt.Println(err.Error())
		return -1
	}

	var header string = `syntax = "proto3";
package gorm;

`
	_, err = f.WriteString(header)
	if err != nil {
		fmt.Println(err.Error())
		return -1
	}
	for _, game := range games {
		_, err = f.WriteString(string("import \""))
		if err != nil {
			fmt.Println(err.Error())
			return -1
		}
		fileLen := len(game.File)
		proto := game.File[:fileLen-4] + ".proto\";\n"
		_, err = f.WriteString(proto)
		if err != nil {
			fmt.Println(err.Error())
			return -1
		}
	}
	f.WriteString("\n\n")

	//输出
	/*
		enum GORM_PB_TABLE_INDEX
		{
			account = 1;
			bag = 2;
		}
	*/
	var enum string = `enum GORM_PB_TABLE_INDEX
{
    GORM_PB_TABLE_IDX_MIN__ = 0;
`
	_, err = f.WriteString(enum)
	if err != nil {
		fmt.Println(err.Error())
		return -1
	}

	var tableIndx int64 = 0
	for _, game := range games {
		for _, table := range game.DB.TableList {
			//GORM_PB_Table_account account = 1;
			tableIndx += 1
			out := "	GORM_PB_TABLE_IDX_" + strings.ToUpper(table.Name) + " = "
			out += strconv.FormatInt(int64(tableIndx), 10)
			_, err = f.WriteString(out + ";\n")
			if err != nil {
				fmt.Println(err.Error())
				return -1
			}
		}
	}
	tableIndx += 1
	out := "	GORM_PB_TABLE_IDX_MAX__ = "
	out += strconv.FormatInt(int64(tableIndx), 10)
	_, err = f.WriteString(out + ";\n")
	if err != nil {
		fmt.Println(err.Error())
		return -1
	}
	_, err = f.WriteString("}\n\n")
	if err != nil {
		fmt.Println(err.Error())
		return -1
	}

	// 输出所有列的宏
	if 0 != GeneralPBColumnIndex(games, f) {
		fmt.Println("general column index failed.")
		return -1
	}

	// 输出所有表的集合
	/*
		message GORM_PB_TABLE
		{
			GORM_PB_Table_account account = 1;
			GORM_PB_Table_bag bag = 2;
		}
	*/
	var table string = `message GORM_PB_TABLE
{
`
	_, err = f.WriteString(table)
	if err != nil {
		fmt.Println(err.Error())
		return -1
	}

	tableIndx = 0
	for _, game := range games {
		for _, table := range game.DB.TableList {
			//GORM_PB_Table_account account = 1;
			tableIndx += 1
			out := "	GORM_PB_Table_" + table.Name + " " + table.Name + " = "
			out += strconv.FormatInt(int64(tableIndx), 10)
			_, err = f.WriteString(out + ";\n")
			if err != nil {
				fmt.Println(err.Error())
				return -1
			}
		}
	}
	_, err = f.WriteString("}\n\n")
	if err != nil {
		fmt.Println(err.Error())
		return -1
	}

	// 输出
	/*
		message GORM_PB_TABLES
		{
			repeated GORM_PB_Table_account account = 1;
			repeated GORM_PB_Table_bag bag = 2;
		}
	*/
	tableIndx = 0
	table = `message GORM_PB_TABLES
{
`
	_, err = f.WriteString(table)
	if err != nil {
		fmt.Println(err.Error())
		return -1
	}
	for _, game := range games {
		for _, table := range game.DB.TableList {
			//GORM_PB_Table_account account = 1;
			tableIndx += 1
			out := "	repeated GORM_PB_Table_" + table.Name + " " + table.Name + " = "
			out += strconv.FormatInt(int64(tableIndx), 10)
			_, err = f.WriteString(out + ";\n")
			if err != nil {
				fmt.Println(err.Error())
				return -1
			}
		}
	}
	_, err = f.WriteString("}\n\n")
	if err != nil {
		fmt.Println(err.Error())
		return -1
	}

	return 0
}

func GeneralPBFiles(games []XmlCfg, outpath string) int {
	if 0 != GeneralTablesInc(games, outpath) {
		fmt.Println("general tables inc proto file failed.")
		return -1
	}
	for _, game := range games {
		if 0 != GeneralPBFile(game, outpath) {
			return -1
		}
	}

	return 0
}
