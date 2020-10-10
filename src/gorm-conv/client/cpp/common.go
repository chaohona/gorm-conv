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

	if 0 != GeneralClientCPPCodes_GeneralGormServerWrap(games, outpath) {
		fmt.Println("GeneralClientCPPCodes_GeneralGormServerWrap failed")
		return -1
	}
	return 0
}
