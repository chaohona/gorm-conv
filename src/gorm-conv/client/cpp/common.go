//test

package cpp

import (
	"fmt"
	"gorm-conv/common"
)

func GeneralClientCPPCodes(games []common.XmlCfg, outpath string) int {
	if 0 != GeneralClientCPPCodes_GeneralGormClientTableOpt(games, outpath) {
		fmt.Println("GeneralClientCPPCodes_GeneralGormClientTableOpt failed")
		return -1
	}

	if 0 != GeneralClientCPPCodes_GenralGormH(games, outpath) {
		fmt.Println("GeneralClientCPPCodes_GenralGormH failed")
		return -1
	}

	if 0 != General_GormClientMsg(outpath) {
		fmt.Println("General_GormClientMsg failed")
		return -1
	}
	return 0
}
