package mysql

import (
	"gorm-conv/common"
	"os"
	"strings"
)

func CPPFieldsMapPackBatchGetSQL(games []common.XmlCfg, f *os.File) int {
	f.WriteString("int GORM_PackGetSQLTable(shared_ptr<GORM_MemPool> &pMemPool, GORM_MySQLEvent *pMySQLEvent, MYSQL* mysql, int iTableId, uint32 uiHashValue, const GORM_PB_TABLE& table, GORM_MemPoolData *&pReqData)\n")
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

			var tableIndex string = "uiHashValue"
			f.WriteString("        return GORM_PackGetSQL" + bigTable + "_ONE(pMemPool, mysql, " + tableIndex + ", table." + table.Name + "(), pReqData);\n")
			f.WriteString("    }\n")
		}
	}
	f.WriteString("    }\n")
	f.WriteString("    return GORM_INVALID_TABLE;\n")
	f.WriteString("}\n")
	return 0
}

func CPPFieldsMapPackInsertTableSQL(games []common.XmlCfg, f *os.File) int {
	f.WriteString("int GORM_PackInsertSQLTable(shared_ptr<GORM_MemPool> &pMemPool, GORM_MySQLEvent *pMySQLEvent, MYSQL* mysql, int iTableId, uint32 uiHashValue, const GORM_PB_TABLE& table, GORM_MemPoolData *&pReqData)\n")
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

			var tableIndex string = "uiHashValue"
			f.WriteString("        return GORM_PackInsertSQL" + bigTable + "_ONE(pMemPool, pMySQLEvent, pMySQLEvent->m_pMySQL, " + tableIndex + ", table." + table.Name + "(), pReqData);\n")
			f.WriteString("    }\n")
		}
	}
	f.WriteString("    }\n")
	f.WriteString("    return GORM_INVALID_TABLE;\n")
	f.WriteString("}\n")
	return 0
}
