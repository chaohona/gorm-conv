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
	Name         string `xml:"name,attr"`
	Columns      string `xml:"columns,attr"`
	Unique       bool   `xml:"unique,attr"`
	IndexColumns []string
}

type SplitInfo struct {
	Columns   string
	SplitCols []string
	Num       uint16
}

type TableColumn struct {
	PrimaryKey   bool   // 是否是主键
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
	Split        string        `xml:"splitinfo,attr"`
	PrimaryKey   string        `xml:"primarykey,attr"`
	SplitInfo    SplitInfo
}

func (this *TableInfo) GetColumn(name string) *TableColumn {
	for idx, tc := range this.TableColumns {
		if name == tc.Name {
			return &this.TableColumns[idx]
		}
	}

	return &TableColumn{}
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

func (this *Routes) GetRouteNum(tableName string) uint16 {
	var result uint16 = 0
	for _, r := range this.TableList {
		if r.Name != tableName {
			continue
		}
		for _, tdb := range r.RoutesTableDBList {
			num, _ := strconv.ParseInt(tdb.SplitNum, 10, 64)
			result += uint16(num)
		}
	}

	return result
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

func ParseSplitInfo(table *TableInfo, router Routes) int {
	table.SplitInfo.SplitCols = []string{}
	table.SplitInfo.Columns = table.Split
	table.SplitInfo.Num = router.GetRouteNum(table.Name)
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
			fmt.Println(vcols)
			fmt.Println(table.TableColumns)
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
			if table.Split == "" {
				fmt.Println("table has no split info:" + table.Name)
				return nil, -1
			}
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

				// 判断此字段是否为主键
				/*for _, sc := range table.SplitInfo.SplitCols {
					if sc == c.Name {
						c.PrimaryKey = true
					}
				}*/
			}
			if 0 != ParseSplitInfo(table, result.DB.Routes) {
				return nil, -1
			}
			// 主键判断
			for _, colName := range table.SplitInfo.SplitCols {
				col := table.GetColumn(colName)
				col.PrimaryKey = true
			}
			// split信息判断
			for idx, _ := range table.TableIndex {
				tableIndex := &table.TableIndex[idx]
				if tableIndex.Columns == "" {
					fmt.Println("table index config failed, table name:" + table.Name + ", index name:" + tableIndex.Name)
					return nil, -1
				}
				tableIndex.Columns = strings.ToLower(tableIndex.Columns)
				tableIndex.IndexColumns = strings.Split(tableIndex.Columns, ",")
				// 判断index是否正确
				for _, cname := range tableIndex.IndexColumns {
					col := table.GetColumn(cname)
					if col == nil {
						fmt.Println("table index config failed, table name:" + table.Name + ", index name:" + tableIndex.Name)
						return nil, -1
					}
				}
				// 判断index名字是否重复
				for nowIdx, nowTableIndex := range table.TableIndex {
					if nowIdx == idx {
						continue
					}
					if nowTableIndex.Name == table.TableIndex[idx].Name {
						fmt.Println("table index config failed, same index name, table name:" + table.Name + ", index name:" + tableIndex.Name)
						return nil, -1
					}
				}
				// 判断index是否包含splitinfo
				for _, sname := range table.SplitInfo.SplitCols {
					var match bool = false
					for _, cname := range tableIndex.IndexColumns {
						if sname == cname {
							match = true
							break
						}
					}
					if !match {
						fmt.Println("table index config failed, table name:" + table.Name + ", index name:" + tableIndex.Name)
						return nil, -1
					}
				}
			}
		}
	}
	return results, 0
}
