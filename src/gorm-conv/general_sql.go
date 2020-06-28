package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func CheckTable(TableInfo TableInfo) int {
	return 0
}

func GeneralSQLColumn(c TableColumn) (string, int) {
	var strOut string = "`"
	strOut += strings.ToLower(c.Name)
	strOut += "` "
	switch c.Type {
	case "double":
		{
			strOut += "double "
			if c.NotNull {
				strOut += "NOT NULL "
			}
			if c.DefaultValue != "" {
				strOut += "DEFAULT " + c.DefaultValue
			}
		}
	case "int8", "uint8":
		{
			strOut += "tinyint "
			if c.NotNull {
				strOut += "NOT NULL "
			}
			if c.DefaultValue != "" {
				strOut += "DEFAULT " + c.DefaultValue
			}
		}
	case "int16", "uint16":
		{
			strOut += "smallint "
			if c.NotNull {
				strOut += "NOT NULL "
			}
			if c.DefaultValue != "" {
				strOut += "DEFAULT " + c.DefaultValue
			}
		}
	case "int", "int32", "uint32":
		{
			strOut += "int "
			if c.NotNull {
				strOut += "NOT NULL "
			}
			if c.DefaultValue != "" {
				strOut += "DEFAULT " + c.DefaultValue
			}
		}
	case "uint64", "int64":
		{
			strOut += "bigint "
			if c.NotNull {
				strOut += "NOT NULL "
			}
			if c.DefaultValue != "" {
				strOut += "DEFAULT " + c.DefaultValue
			}
		}
	case "string":
		{
			strOut += "varchar("
			strOut += c.Size
			strOut += ")"
			if c.NotNull {
				strOut += " NOT NULL"
			} else if c.DefaultValue != "" {
				strOut += " DEFAULT "
				strOut += c.DefaultValue
			} else {
				strOut += " DEFAULT "
				strOut += "NULL"
			}
		}
	case "char":
		{
			strSize, _ := strconv.ParseInt(c.Size, 10, 64)
			if strSize <= int64(255) {
				strOut += "char("
			} else {
				strOut += "varchar("
			}
			strOut += c.Size
			strOut += ")"
			if c.NotNull {
				strOut += " NOT NULL"
			} else if c.DefaultValue != "" {
				strOut += " DEFAULT "
				strOut += c.DefaultValue
			} else {
				strOut += " DEFAULT "
				strOut += "NULL"
			}
		}
	case "blob":
		{
			strOut += "mediumblob"
		}
	default:
		{
			fmt.Println("invalid column type:"+c.Type, ", name:"+c.Name)
			return "", -1
		}
	}
	return strOut, 0
}

func CreateTableSQL(table TableInfo) ([]byte, int) {
	var strOut string
	for idx, c := range table.TableColumns {
		column, ret := GeneralSQLColumn(c)
		column = "    " + column
		if ret != 0 {
			return nil, ret
		}
		if idx != 0 && idx != len(table.TableColumns) {
			strOut += ",\n"
		}
		strOut += column
	}
	if len(table.TableIndex) != 0 {
		for _, tIndex := range table.TableIndex {

			if tIndex.Unique {
				strOut += ",\n    UNIQUE INDEX " + tIndex.Name + "(" + tIndex.Columns + ")"
			} else {
				strOut += ",\n    INDEX " + tIndex.Name + "(" + tIndex.Columns + ")"
			}
		}
	}
	if table.PrimaryKey.Column != "" {
		strOut += ",\n    PRIMARY KEY (`" + table.PrimaryKey.Column + "`)"
	}
	strOut += "\n) ENGINE=InnoDB DEFAULT CHARSET=utf8;\n"
	return []byte(strOut), 0
}

func GeneralSQLBuff(table TableInfo, tableNum int) ([]byte, int) {
	var totalOut string
	if table.SplitInfo.Num == 0 {
		var strOut string = "CREATE TABLE `" + table.Name + "`(\n"
		totalOut += strOut
		tableOut, ret := CreateTableSQL(table)
		if ret != 0 {
			return nil, -1
		}
		totalOut += string(tableOut)
	} else {
		for index := uint16(1); index <= table.SplitInfo.Num; index++ {
			var strOut string = "CREATE TABLE `" + table.Name + "_" + strconv.FormatInt(int64(index), 10) + "`(\n"
			totalOut += strOut
			tableOut, ret := CreateTableSQL(table)
			if ret != 0 {
				return nil, -1
			}
			totalOut += string(tableOut)
		}
	}
	return []byte(totalOut), 0
}

func CreateSQLFile(db GiantGame, file string) int {
	// 1.校验所有的表
	var tableNames map[string]int = make(map[string]int)
	for idx, table := range db.TableList {
		if tableNames[table.Name] != 0 {
			fmt.Println("repeat table names:", table.Name)
			return -1
		}
		if ret := CheckTable(table); ret != 0 {
			fmt.Println("invalid table, index:" + strconv.FormatInt(int64(idx), 10) + ", name:" + table.Name)
			return ret
		}
	}
	// 2.获取每一种表的数据库列表
	// 3.为每一种表生成创建语句
	var outBuffer []byte
	for _, table := range db.TableList {
		data, ret := GeneralSQLBuff(table, 1)
		if ret != 0 {
			fmt.Println("general sql failed:" + table.Name)
			return ret
		}
		outBuffer = append(outBuffer, data...)
	}

	fmt.Println("sql out file:" + file)
	//fmt.Println(string(outBuffer))
	// 4.写SQL文件
	f, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer f.Close()
	_, err = f.Write(outBuffer)
	if err != nil {
		fmt.Println(err.Error())
	}
	return 0
}

func GeneralSQLFiles(games []XmlCfg, outpath string) int {
	for _, game := range games {
		fileLen := len(game.File)
		outfile := outpath + "/" + game.File[:fileLen-4] + ".sql"
		//fmt.Println(outfile)
		if ret := CreateSQLFile(game.DB, outfile); ret != 0 {
			fmt.Println("generate sql file failed:" + outfile)
			return ret
		}
		fmt.Println("generate sql file success:" + outfile)
	}
	return 0
}
