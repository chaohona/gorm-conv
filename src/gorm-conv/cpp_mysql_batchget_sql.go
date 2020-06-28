package main

import (
	"os"
	"strings"
)

func CPPFieldsMapPackBatchGetSQL(games []XmlCfg, f *os.File) int {
	f.WriteString("int GORM_PackGetSQL(MYSQL* mysql, int iTableId, const GORM_PB_TABLE& table, GORM_MemPoolData *&pReqData)\n")
	f.WriteString("{\n")
	f.WriteString("    switch (iTableId)\n")
	f.WriteString("    {\n")
	for _, game := range games {
		for _, table := range game.DB.TableList {
			f.WriteString("    case GORM_PB_TABLE_IDX_" + strings.ToUpper(table.Name) + ":\n")
			f.WriteString("    {\n")
			f.WriteString("        if (!table.has_" + table.Name + "())\n")
			f.WriteString("            return GORM_REQ_NO_RECORDS;\n")
			f.WriteString("        return GORM_PackGetSQL" + strings.ToUpper(table.Name) + "_ONE(mysql, table." + table.Name + "(), pReqData);\n")
			f.WriteString("    }\n")
		}
	}
	f.WriteString("    }\n")
	f.WriteString("    return GORM_INVALID_TABLE;\n")
	f.WriteString("}\n")
	return 0
}
