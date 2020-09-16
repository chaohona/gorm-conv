package cpp

import (
	"gorm-conv/common"
	"os"
	"strings"
)

func GORM_TableHash(games []common.XmlCfg, f *os.File) int {
	f.WriteString("uint32 GORM_TableHash(int iTableId, const GORM_PB_TABLE &pbTable)\n")
	f.WriteString("{\n")
	f.WriteString("    switch (iTableId)\n")
	f.WriteString("    {\n")
	for _, game := range games {
		for _, table := range game.DB.TableList {
			f.WriteString("    case GORM_PB_TABLE_IDX_" + strings.ToUpper(table.Name) + ":\n")
			f.WriteString("    {\n")
			f.WriteString("        if (!pbTable.has_" + table.Name + "())\n")
			f.WriteString("            return 0;\n")
			f.WriteString("        const GORM_PB_Table_" + table.Name + "& " + table.Name + " = pbTable." + table.Name + "();\n")
			for _, s := range table.SplitInfo.SplitCols {
				c := table.GetColumn(s)
				var cType string = CPPField_CPPType(c.Type)
				if cType == "string" {
					f.WriteString("        const string &tmp_" + strings.ToUpper(s) + " = " + table.Name + "." + c.Name + "();\n")
				} else {
					f.WriteString("        " + c.Type + " tmp_" + strings.ToUpper(s) + " = " + table.Name + "." + c.Name + "();\n")
				}
			}
			f.WriteString("        char szSrcHash[1024];\n")
			f.WriteString("        int iTotalLen = GORM_SafeSnprintf(szSrcHash, 1024, \"")
			for _, splitCol := range table.SplitInfo.SplitCols {
				col := table.GetColumn(splitCol)
				f.WriteString(CPPFieldPackRedis_COL_FORMAT(col.Type))
				f.WriteString("_")
			}
			f.WriteString("\" ")
			for _, splitCol := range table.SplitInfo.SplitCols {
				f.WriteString(", ")
				f.WriteString("tmp_" + strings.ToUpper(splitCol))
			}
			f.WriteString(");\n")
			f.WriteString("        if (iTotalLen > 1024)\n")
			f.WriteString("            iTotalLen = 1024;\n")
			f.WriteString("        return GORM_Hash::Crc32_1((const char*)szSrcHash, iTotalLen);\n")
			f.WriteString("    }\n")
		}
	}
	f.WriteString("    default:\n")
	f.WriteString("        return 0;\n")
	f.WriteString("    }\n")
	f.WriteString("    return 0;\n")
	f.WriteString("}\n")
	return 0
}
