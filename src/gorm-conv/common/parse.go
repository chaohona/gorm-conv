package common

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
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

type XmlCfg struct {
	DB   GiantGame
	File string
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

func ParseXmls(strPath string) ([]XmlCfg, int) {
	var results []XmlCfg
	files := GetXmlFiles(strPath)
	for _, file := range files {
		cfg, ret := GetXmlCfg(strPath + "/" + file)
		if ret != 0 {
			return nil, -1
		}
		results = append(results, XmlCfg{
			DB:   cfg,
			File: file,
		})
	}
	for _, result := range results {
		for idx, db := range result.DB.DBList.DBList {
			result.DB.DBList.DBList[idx].Database = strings.ToLower(db.Database)
			result.DB.DBList.DBList[idx].Name = strings.ToLower(db.Name)
		}
		for idx, _ := range result.DB.TableList {
			table := &result.DB.TableList[idx]
			table.PrimaryKey.Column = table.SplitInfo.Columns
			var TableColumns []TableColumn
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
			for i, c := range table.TableColumns {
				table.TableColumns[i].Name = strings.ToLower(c.Name)
				table.TableColumns[i].Type = strings.ToLower(c.Type)
			}
			if 0 != ParseSplitInfo(table) {
				return nil, -1
			}
		}
	}
	return results, 0
}
