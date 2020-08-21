package cpp

import (
	"gorm-conv/common"
	"os"
	"strconv"
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
				for _, c := range table.TableColumns {
					if s == c.Name {
						var cType string = CPPField_CPPType(c.Type)
						if cType == "string" {
							f.WriteString("        const string &str" + strings.ToUpper(s) + " = " + table.Name + "." + c.Name + "();\n")
						} else {
							f.WriteString("        " + c.Type + " num" + strings.ToUpper(s) + " = " + table.Name + "." + c.Name + "();\n")
						}
						break
					}
				}
			}
			var iSplitNum int = len(table.SplitInfo.SplitCols)
			if iSplitNum > 4 {
				iSplitNum = 4
			}
			f.WriteString("        return GORM_Hash::Crc32_" + strconv.FormatInt(int64(iSplitNum), 10) + "(")
			for idx, s := range table.SplitInfo.SplitCols {
				for _, c := range table.TableColumns {
					var cBIG string = strings.ToUpper(c.Name)
					if s == c.Name {
						if idx != 0 {
							f.WriteString(", ")
						}
						var cType string = CPPField_CPPType(c.Type)
						if cType == "string" {
							f.WriteString("str" + cBIG + ".c_str(), str" + cBIG + ".size()")
						} else {
							var iLen int = CPPTypeLen(c.Type)
							f.WriteString("(const char*)&num" + cBIG + ", " + strconv.FormatInt(int64(iLen), 10))
						}
						break
					}
				}
				if idx == 3 {
					break
				}
			}
			f.WriteString(");\n")
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
