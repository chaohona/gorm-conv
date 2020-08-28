package golang

import (
	"fmt"
	"gorm-conv/common"
	"os"
	"strconv"
)

func GeneralRecods(outpath string) int {
	outfile := outpath + "/gorm_record.go"
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

	f.WriteString(`package gorm

var (
	FIELDS_OPT_COLLECTIONS_MAX_LEN int = 128
)

type Record interface {
	Init() int
}

`)
	return 0
}

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

	f.WriteString("package gorm\n")
	f.WriteString("type " + tablePbName + " struct{\n")
	f.WriteString("    Record\n")
	f.WriteString("    " + getGolangPBStructName(table.Name) + "\n")
	f.WriteString("    fieldCollections []byte\n")
	f.WriteString("    usedIdx          int\n")
	f.WriteString("}\n")

	f.WriteString("func (this *" + tablePbName + ") Init() {\n")
	f.WriteString("    this.fieldCollections = make([]byte, FIELDS_OPT_COLLECTIONS_MAX_LEN)\n")
	f.WriteString("}\n")

	for idx, col := range table.TableColumns {
		var colPbName string = getGolangPbFieldName(col.Name)
		f.WriteString("func (this *" + tablePbName + ") Set" + colPbName + "(inArg " + CPPField_GolangType(col.Type) + ") {\n")
		f.WriteString("    this." + colPbName + " = inArg\n")

		var shiftNum int64 = (int64(idx) >> 3)
		var shift string = strconv.FormatInt(shiftNum, 10)
		var modeNum int = 1 << (idx & 0x07)
		var mode string = strconv.FormatInt(int64(modeNum), 10)
		f.WriteString("    this.fieldCollections[" + shift + "] |= " + mode + "\n")
		f.WriteString("    if this.usedIdx < " + shift + " {\n")
		f.WriteString("        this.usedIdx = " + shift + "\n")
		f.WriteString("    }\n")
		f.WriteString("}\n")
	}
	return 0
}

func GeneralGolang_Table_Records(games []common.XmlCfg, outpath string) int {
	if 0 != GeneralRecods(outpath) {
		return -1
	}
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
