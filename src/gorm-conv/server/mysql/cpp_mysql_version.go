package mysql

import (
	"gorm-conv/common"
	"os"
)

func CPPFieldsMapPack_VERSION_SQL(games []common.XmlCfg, f *os.File) int {
	f.WriteString("int GORM_GETVERSION_SET(char *szSQLBegin, int iLen, GORM_CheckDataVerType vType, uint64 ulVersion)\n")
	f.WriteString("{\n")
	f.WriteString("    int iDataLen = 0;\n")
	f.WriteString("    switch (vType)\n")
	f.WriteString("    {\n")
	f.WriteString("    case CHECKDATAVERSION_AUTOINCREASE:\n")
	f.WriteString("    {\n")
	f.WriteString("        iDataLen = GORM_SafeSnprintf(szSQLBegin, iLen, \" `version`=`version`+1\");\n")
	f.WriteString("        break;\n")
	f.WriteString("    }\n")
	f.WriteString("    case NOCHECKDATAVERSION_OVERWRITE:\n")
	f.WriteString("    {\n")
	f.WriteString("        iDataLen = GORM_SafeSnprintf(szSQLBegin, iLen, \" `version`=%llu\", ulVersion);\n")
	f.WriteString("        break;\n")
	f.WriteString("    }\n")
	f.WriteString("    case NOCHECKDATAVERSION_AUTOINCREASE:\n")
	f.WriteString("    {\n")
	f.WriteString("        iDataLen = GORM_SafeSnprintf(szSQLBegin, iLen, \" `version`=`version`+1\", ulVersion);\n")
	f.WriteString("        break;\n")
	f.WriteString("    }\n")
	f.WriteString("    }\n")
	f.WriteString("    return iDataLen;\n")
	f.WriteString("}\n")

	f.WriteString("int GORM_GETVERSION_WHERE(char *szSQLBegin, int iLen, GORM_CheckDataVerType vType, uint64 ulVersion)\n")
	f.WriteString("{\n")
	f.WriteString("    int iDataLen = 0;\n")
	f.WriteString("    switch (vType)\n")
	f.WriteString("    {\n")
	f.WriteString("    case CHECKDATAVERSION_AUTOINCREASE:\n")
	f.WriteString("    {\n")
	f.WriteString("        //iDataLen = GORM_SafeSnprintf(szSQLBegin, iLen, \" and `version`=%llu\", ulVersion);\n")
	f.WriteString("        break;\n")
	f.WriteString("    }\n")
	f.WriteString("    case NOCHECKDATAVERSION_OVERWRITE:\n")
	f.WriteString("    {\n")
	f.WriteString("        break;\n")
	f.WriteString("    }\n")
	f.WriteString("    case NOCHECKDATAVERSION_AUTOINCREASE:\n")
	f.WriteString("    {\n")
	f.WriteString("        break;\n")
	f.WriteString("    }\n")
	f.WriteString("    }\n")
	f.WriteString("    return iDataLen;\n")
	f.WriteString("}\n")
	return 0
}
