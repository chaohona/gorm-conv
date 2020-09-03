package mysql

import (
	"os"
)

func PrintGetBuffFromMemPool(f *os.File, buffStr string, sizeStr string) {
	f.WriteString("    GORM_MallocFromSharedPool(pMemPool, " + buffStr + ", " + sizeStr + ");\n")
}
