package mysql

import (
	"gorm-conv/common"
	"os"
	"strings"
)

func CPPFields_GORM_PackSQL_TEMPLATE(opt string, games []common.XmlCfg, f *os.File) int {
	f.WriteString("int GORM_Pack" + opt + "SQL(shared_ptr<GORM_MemPool> &pMemPool, GORM_MySQLEvent *pMySQLEvent, MYSQL* mysql, int iTableId, uint32 uiHashValue, const GORM_PB_" + strings.ToUpper(opt) + "_REQ* pMsg, GORM_MemPoolData *&pReqData)\n")
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
			var hashValue string = "uiHashValue"
			f.WriteString("(pMemPool, pMySQLEvent, mysql, " + hashValue + ", pMsg, pReqData);\n")
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

func CPPFields_GORM_Table_PackSQL_TEMPLATE(opt string, games []common.XmlCfg, f *os.File) int {
	f.WriteString("int GORM_Pack" + opt + "SQL(shared_ptr<GORM_MemPool> &pMemPool, , MYSQL* mysql, int iTableId, const GORM_PB_REQ_HEADER &header, const GORM_PB_TABLE &table, GORM_MemPoolData *&pReqData)\n")
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
			f.WriteString("(mysql, header, table, pReqData);\n")
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
