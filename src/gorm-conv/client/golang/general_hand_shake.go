package golang

import (
	"fmt"
	"gorm-conv/common"
	"os"
	"strconv"
	"strings"
)

func GeneralGolang_HandShake(games []common.XmlCfg, outpath string) int {
	outfile := outpath + "/gorm_schema_info.go"
	f, err := os.OpenFile(outfile, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err.Error())
		return -1
	}

	fmt.Println("begin to general golang codes file:" + outfile)
	defer func() {
		f.Close()
		fmt.Println("end general golang codes file:" + outfile)
	}()
	f.Truncate(0)

	f.WriteString("package gorm\n\n")
	f.WriteString("var g_table_schema_info GORM_PB_HAND_SHAKE_REQ = GORM_PB_HAND_SHAKE_REQ{\n")
	f.WriteString("    Header:  &GORM_PB_REQ_HEADER{},\n")
	f.WriteString("    Version: 1,\n")
	f.WriteString("    Md5:     0,\n")
	f.WriteString("    Schemas: []*GORM_PB_TABLE_SCHEMA_INFO{\n")
	var tableIndex int64 = 0
	for _, game := range games {
		for _, table := range game.DB.TableList {
			tableIndex += 1
			var tableVersion string = strconv.FormatUint(table.Version, 10)
			var tableIndexStr string = strconv.FormatInt(tableIndex, 10)
			f.WriteString("        &GORM_PB_TABLE_SCHEMA_INFO{\n")
			f.WriteString("            Version:   " + tableVersion + ",\n")
			f.WriteString("            TableName: \"" + table.Name + "\",\n")
			f.WriteString("            TableIdx:  " + tableIndexStr + ",\n")
			f.WriteString("            Columns: []*GORM_PB_TABLE_SCHEMA_INFO_COLUMN{\n")
			for _, col := range table.TableColumns {
				var colVer string = strconv.FormatUint(col.Version, 10)
				f.WriteString("                &GORM_PB_TABLE_SCHEMA_INFO_COLUMN{\n")
				f.WriteString("                    Version:  " + colVer + ",\n")
				f.WriteString("                    Name:     \"" + col.GoName + "\",\n")
				f.WriteString("                    TypeDesc: \"" + col.Type + "\",\n")
				f.WriteString("                    Type:     GORM_PB_COLUMN_TYPE_GORM_PB_COLUMN_TYPE_" + strings.ToUpper(col.Type) + ",\n")
				f.WriteString("                },\n")
			}
			f.WriteString("            },\n")
			f.WriteString("        },\n")
		}
	}
	f.WriteString("    },\n")
	f.WriteString("}\n\n")

	f.WriteString(`

func GORM_HandShakeReq() *GORM_PB_HAND_SHAKE_REQ {
	return &g_table_schema_info
}

`)
	return 0
}
