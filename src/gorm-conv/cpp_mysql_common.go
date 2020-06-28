package main

import (
	"os"
	"strings"
)

func CPPFieldPackSQL_COL_FORMAT(colType string) string {
	var DefineSQL string
	switch colType {
	case "int8", "int16", "int32", "int":
		{
			DefineSQL += "%d"
		}
	case "uint8", "uint16", "uint32":
		{
			DefineSQL += "%u"
		}
	case "int64":
		{
			DefineSQL += "%ll"
		}
	case "uint64":
		{
			DefineSQL += "%llu"
		}
	default:
		{
			DefineSQL += "\\\"%s\\\""
		}
	}
	return DefineSQL
}

func CPPFields_GORM_PackSQL_TEMPLATE(opt string, games []XmlCfg, f *os.File) int {
	f.WriteString("int GORM_Pack" + opt + "SQL(MYSQL* mysql, int iTableId, const GORM_PB_" + strings.ToUpper(opt) + "_REQ* pMsg, GORM_MemPoolData *&pReqData)\n")
	f.WriteString("{\n")
	totalIndex := 0
	for _, game := range games {
		for _, table := range game.DB.TableList {
			totalIndex += 1
			if totalIndex == 1 {
				f.WriteString("    switch (iTableId)\n")
				f.WriteString("    {\n")
			}
			var BigTable string = strings.ToUpper(table.Name)
			f.WriteString("    case GORM_PB_TABLE_IDX_")
			f.WriteString(BigTable)
			f.WriteString(":\n")
			f.WriteString("        return GORM_Pack" + opt + "SQL")
			f.WriteString(BigTable)
			f.WriteString("(mysql, pMsg, pReqData);\n")
			f.WriteString("    \n")
		}
	}
	if totalIndex > 0 {
		f.WriteString("    }\n")
	}

	f.WriteString("    return GORM_INVALID_TABLE;\n")
	f.WriteString("}\n\n")
	return 0
}
