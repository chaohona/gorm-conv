package golang

import (
	"fmt"
	"gorm-conv/common"
	"os"
	"strconv"
)

func GeneralGolang_Table_Records_Table(table common.TableInfo, outpath string) int {
	outfile := outpath + "/gorm_table_" + table.Name + ".go"
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
	err = f.Truncate(0)

	var tablePbName string = getGolangPbFieldName(table.Name)
	var pbStructName string = getGolangPBStructName(table.Name)

	f.WriteString("package gorm\n")
	f.WriteString("type " + tablePbName + " struct{\n")
	f.WriteString("    Record\n")
	f.WriteString("}\n")

	// Init函数
	f.WriteString("func (this *" + tablePbName + ") Init() GORM_CODE {\n")
	f.WriteString("    this.msg = &" + pbStructName + "{}\n")
	f.WriteString("    this.fieldMode.Init()\n")
	f.WriteString("    this.modifyFiled = true\n")
	f.WriteString("    return GORM_CODE_OK\n")
	f.WriteString("}\n")

	// InitEx函数
	f.WriteString("func (this *" + tablePbName + ") InitEx() GORM_CODE {\n")
	f.WriteString("    this.msg = &" + pbStructName + "{}\n")
	f.WriteString("    this.modifyFiled = false\n")
	f.WriteString("    this.fieldMode.Init()\n")
	for i := 0; i < len(table.TableColumns); i++ {
		var shiftNum int64 = (int64(i) >> 3)
		var shift string = strconv.FormatInt(shiftNum, 10)
		var modeNum int = 1 << (i & 0x07)
		var mode string = strconv.FormatInt(int64(modeNum), 10)
		f.WriteString("    this.fieldCollections[" + shift + "] |= " + mode + "\n")
	}
	var shiftNum int64 = (int64(len(table.TableColumns)) >> 3)
	var shift string = strconv.FormatInt(shiftNum, 10)
	f.WriteString("    this.usedIdx = " + shift + "\n")
	f.WriteString("    return GORM_CODE_OK\n")
	f.WriteString("}\n")

	// GetReadTble函数
	f.WriteString("func (this *" + pbStructName + ") GetOnlyReadTbl() *" + pbStructName + " {\n")
	f.WriteString("    return this.msg.(*" + pbStructName + ")")
	f.WriteString("}\n")

	// 对各个字段的设置函数
	for idx, col := range table.TableColumns {
		var colPbName string = getGolangPbFieldName(col.Name)
		f.WriteString("func (this *" + tablePbName + ") Set" + colPbName + "(inArg " + CPPField_GolangType(col.Type) + ") {\n")
		f.WriteString("    var msg *" + pbStructName + "  = this.msg.(*" + pbStructName + ")\n")
		f.WriteString("    msg." + colPbName + " = inArg\n")

		var shiftNum int64 = (int64(idx) >> 3)
		var shift string = strconv.FormatInt(shiftNum, 10)
		var modeNum int = 1 << (idx & 0x07)
		var mode string = strconv.FormatInt(int64(modeNum), 10)
		f.WriteString("    this.fieldCollections[" + shift + "] |= " + mode + "\n")
		f.WriteString("    if this.usedIdx < " + shift + " {\n")
		f.WriteString("        this.usedIdx = " + shift + "\n")
		f.WriteString("    }\n")
		f.WriteString("    this.modifyFiled = true")
		f.WriteString("}\n")
	}
	return 0
}

func GeneralGolang_Table_Records(games []common.XmlCfg, outpath string) int {
	for _, game := range games {
		for _, table := range game.DB.TableList {
			if 0 != GeneralGolang_Table_Records_Table(table, outpath) {
				fmt.Println("general codes for table failed, table:", table.Name)
				return -1
			}
		}
	}
	return 0
}
