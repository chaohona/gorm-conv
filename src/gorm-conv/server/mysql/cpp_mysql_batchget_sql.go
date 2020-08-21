package mysql

import (
	"gorm-conv/common"
	"os"
	"strconv"
	"strings"
)

func CPPFieldsMapPackBatchGetSQL(games []common.XmlCfg, f *os.File) int {
	f.WriteString("int GORM_PackGetSQLTable(GORM_MySQLEvent *pMySQLEvent, MYSQL* mysql, int iTableId, uint32 uiHashValue, const GORM_PB_TABLE& table, GORM_MemPoolData *&pReqData)\n")
	f.WriteString("{\n")

	f.WriteString("    switch (iTableId)\n")
	f.WriteString("    {\n")
	for _, game := range games {
		for _, table := range game.DB.TableList {
			var bigTable string = strings.ToUpper(table.Name)
			f.WriteString("    case GORM_PB_TABLE_IDX_" + bigTable + ":\n")
			f.WriteString("    {\n")
			f.WriteString("        if (!table.has_" + table.Name + "())\n")
			f.WriteString("            return GORM_REQ_NO_RECORDS;\n")

			f.WriteString("#ifdef GORM_DEBUG\n")
			f.WriteString("        GORM_MySQLUpdateTableSchema(pMySQLEvent, \"" + table.Name + "\", table.custom_columns());\n")
			f.WriteString("#endif\n")
			var tableIndex string = "0"
			if table.SplitInfo.Num > 1 {
				tableIndex = "uiHashValue%"
				tableIndex += strconv.FormatInt(int64(table.SplitInfo.Num), 10)
			}
			f.WriteString("        return GORM_PackGetSQL" + bigTable + "_ONE(mysql, " + tableIndex + ", table." + table.Name + "(), pReqData);\n")
			f.WriteString("    }\n")
		}
	}
	f.WriteString("    }\n")
	f.WriteString("    return GORM_INVALID_TABLE;\n")
	f.WriteString("}\n")
	return 0
}

func CPPFieldsMapPackInsertTableSQL(games []common.XmlCfg, f *os.File) int {
	f.WriteString("int GORM_PackInsertSQLTable(GORM_MySQLEvent *pMySQLEvent, MYSQL* mysql, int iTableId, uint32 uiHashValue, const GORM_PB_TABLE& table, GORM_MemPoolData *&pReqData)\n")
	f.WriteString("{\n")

	f.WriteString("    switch (iTableId)\n")
	f.WriteString("    {\n")
	for _, game := range games {
		for _, table := range game.DB.TableList {
			var bigTable string = strings.ToUpper(table.Name)
			f.WriteString("    case GORM_PB_TABLE_IDX_" + bigTable + ":\n")
			f.WriteString("    {\n")
			f.WriteString("        if (!table.has_" + table.Name + "())\n")
			f.WriteString("            return GORM_REQ_NO_RECORDS;\n")

			f.WriteString("#ifdef GORM_DEBUG\n")
			f.WriteString("        GORM_MySQLUpdateTableSchema(pMySQLEvent, \"" + table.Name + "\", table.custom_columns());\n")
			f.WriteString("#endif\n")
			var tableIndex string = "0"
			if table.SplitInfo.Num > 1 {
				tableIndex = "uiHashValue%"
				tableIndex += strconv.FormatInt(int64(table.SplitInfo.Num), 10)
			}
			f.WriteString("        return GORM_PackInsertSQL" + bigTable + "_ONE(pMySQLEvent, pMySQLEvent->m_pMySQL, " + tableIndex + ", table." + table.Name + "(), pReqData);\n")
			f.WriteString("    }\n")
		}
	}
	f.WriteString("    }\n")
	f.WriteString("    return GORM_INVALID_TABLE;\n")
	f.WriteString("}\n")
	return 0
}
