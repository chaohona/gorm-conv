package golang

import (
	"gorm-conv/common"
)

func GeneralClientGolangCodes(games []common.XmlCfg, outpath string) int {
	if 0 != GeneralGolangCodes_Common(games, outpath) {
		return -1
	}

	if 0 != GeneralGolang_Table_Records(games, outpath) {
		return -1
	}

	if 0 != GeneralGolang_HandShake(games, outpath) {
		return -1
	}

	return 0
}
