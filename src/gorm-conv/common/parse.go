package common

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

var (
	Type_Key_Name map[string]int = map[string]int{
		"int": 1,
	}
)

const (
	DBType_MYSQL = "mysql"
	DBType_MGO   = "mongo"
)

type TableIndex struct {
	Name    string `xml:"name,attr"`
	Columns string `xml:"columns,attr"`
	Unique  bool   `xml:"unique,attr"`
}

type SplitInfo struct {
	Columns   string `xml:"columns,attr"`
	Num       uint16 `xml:"num,attr"`
	SplitCols []string
}

type PrimaryKey struct {
	Column string
}

type TableColumn struct {
	GoName       string // golang使用的列名
	CPPName      string // cpp使用的列名
	SQLName      string
	Name         string `xml:"name,attr"`
	Type         string `xml:"type,attr"`
	NotNull      bool   `xml:"notnull,attr"`
	DefaultValue string `xml:"defaultvalue,attr"`
	Size         string `xml:"size,attr"`
	Version      uint64 `xml:"version,attr"`
}

type TableInfo struct {
	Name         string        `xml:"name,attr"`
	Version      uint64        `xml:"version,attr"`
	TableColumns []TableColumn `xml:"column"`
	TableIndex   []TableIndex  `xml:"index"`
	SplitInfo    SplitInfo     `xml:"splitinfo"`
	PrimaryKey   PrimaryKey
}

func (this *TableInfo) GetColumn(name string) TableColumn {
	for _, tc := range this.TableColumns {
		if name == tc.Name {
			return tc
		}
	}

	return TableColumn{}
}

type DataBaseInfo struct {
	Name     string `xml:"name,attr"`
	Host     string `xml:"host,attr"`
	Port     string `xml:"port,attr"`
	User     string `xml:"user,attr"`
	PassWD   string `xml:"password,attr"`
	Type     string `xml:"type,attr"`
	Database string `xml:"database,attr"`
}

type DataBase struct {
	DBList []DataBaseInfo `xml:"db"`
}

type RoutesTableDB struct {
	Name     string `xml:"name,attr"`
	SplitNum string `xml:"splittablenum,attr"`
}

type RoutesTable struct {
	Name              string          `xml:"name,attr"`
	RoutesTableDBList []RoutesTableDB `xml:"db"`
}

type Routes struct {
	TableList []RoutesTable `xml:"table"`
}

type GiantGame struct {
	DBList    DataBase    `xml:"databases"`
	TableList []TableInfo `xml:"table"`
	Routes    Routes      `xml:"routes"`
	Version   uint64      `xml:"version,attr"`
}

func (this *GiantGame) GetTableInfo(tableName string) (*TableInfo, bool) {
	for idx, table := range this.TableList {
		if table.Name == tableName {
			return &this.TableList[idx], true
		}
	}

	return nil, false
}

type XmlCfg struct {
	DB   GiantGame
	File string
}

func (this *XmlCfg) GetFileNameCharacter() (result string) {
	for i := 0; i < len(this.File); i++ {
		var c byte = this.File[i]
		if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9') {
			result += string(c)
		}
	}

	return
}

func GetXmlFiles(folder string) []string {
	var results []string
	files, _ := ioutil.ReadDir(folder) //specify the current dir
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		fileName := file.Name()
		vector := strings.Split(fileName, ".")
		if len(vector) > 1 && vector[len(vector)-1] == "xml" {
			results = append(results, file.Name())
		}
	}

	return results
}

func GetXmlCfg(infile string) (GiantGame, int) {
	var result GiantGame
	fmt.Println("begin to parse file:" + infile)
	file, err := os.Open(infile)
	if err != nil {
		fmt.Println("open file got error:", err.Error())
		return result, -1
	}

	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("read file got error:", err.Error())
		return result, -1
	}

	err = xml.Unmarshal(data, &result)
	if err != nil {
		fmt.Println("parse file got error:", err.Error())
		return result, -1
	}
	//fmt.Println(result)
	return result, 0
}

func ParseSplitInfo(table *TableInfo) int {
	table.SplitInfo.SplitCols = []string{}
	if table.SplitInfo.Columns == "" {
		return 0
	}
	table.SplitInfo.Columns = strings.ToLower(table.SplitInfo.Columns)
	var vcols []string = strings.Split(table.SplitInfo.Columns, ",")
	for _, c := range vcols {
		var bright bool = false
		for _, preC := range table.TableColumns {
			if c == preC.Name {
				bright = true
				break
			}
		}
		if !bright {
			fmt.Println("wrong splitinfo for table:" + table.Name)
			return -1
		}
		table.SplitInfo.SplitCols = append(table.SplitInfo.SplitCols, c)
	}
	return 0
}

func ParseXmls(strPath string, strFilePath string) ([]XmlCfg, int) {
	var results []XmlCfg
	var files []string
	if strPath != "" {
		files = GetXmlFiles(strPath)
	} else {
		files = append(files, strFilePath)
	}
	fmt.Println(files)
	for _, file := range files {
		var fileName string = file
		var cfg GiantGame
		var ret int
		// 如果参数是路径则带上路径
		if strPath != "" {
			cfg, ret = GetXmlCfg(strPath + "/" + file)
		} else { // 则直接取文件
			cfg, ret = GetXmlCfg(file)
			_, fileName = filepath.Split(fileName)
		}

		if ret != 0 {
			return nil, -1
		}
		results = append(results, XmlCfg{
			DB:   cfg,
			File: fileName,
		})
	}
	for _, result := range results {
		for idx, _ := range result.DB.DBList.DBList {
			db := &result.DB.DBList.DBList[idx]
			db.Database = strings.ToLower(db.Database)
			db.Name = strings.ToLower(db.Name)
		}
		for idx, _ := range result.DB.TableList {
			table := &result.DB.TableList[idx]
			if table.SplitInfo.Columns == "" {
				fmt.Println("table has no split info:" + table.Name)
				return nil, -1
			}
			table.PrimaryKey.Column = table.SplitInfo.Columns
			var TableColumns []TableColumn
			// 自动增加version字段
			TableColumns = append(TableColumns, TableColumn{
				Name:         "version",
				Type:         "uint64",
				DefaultValue: "0",
			})
			TableColumns = append(TableColumns, table.TableColumns...)
			table.TableColumns = TableColumns
			if table.Version == 0 {
				fmt.Println("table must has version attributer, and big than 0.")
				return nil, -1
			}
			table.Name = strings.ToLower(table.Name)
			table.PrimaryKey.Column = strings.ToLower(table.PrimaryKey.Column)
			for i, _ := range table.TableIndex {
				tIndex := &table.TableIndex[i]
				tIndex.Columns = strings.ToLower(tIndex.Columns)
				tIndex.Name = strings.ToLower(tIndex.Name)
			}
			var colName string
			for i, _ := range table.TableColumns {
				c := &table.TableColumns[i]
				size, _ := strconv.ParseInt(c.Size, 10, 64)
				// size没有则设置为默认值
				if size == 0 {
					c.Size = "1024"
				}
				colName = strings.ToLower(c.Name)
				if Type_Key_Name[colName] > 0 {
					fmt.Println("table column is class name, table:" + table.Name + ", column:" + colName)
					return nil, -1
				}
				c.Name = colName
				c.Type = strings.ToLower(c.Type)
				c.GoName = colName
				c.CPPName = colName
				c.SQLName = colName
				if colName == "int" {
					c.CPPName = "int_"
				}
			}
			if 0 != ParseSplitInfo(table) {
				return nil, -1
			}
		}
	}
	return results, 0
}
