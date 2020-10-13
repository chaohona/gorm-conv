package cpp

/*
import (
	"fmt"
	"gorm-conv/common"
	"os"
)

func GeneralClientCPPCodes_GeneralGormServerWrap(games []common.XmlCfg, outpath string) int {
	outfile := outpath + "/gorm_wrap.h"
	fmt.Println(outfile)
	f, err := os.OpenFile(outfile, os.O_RDWR|os.O_CREATE, 0666)
	_ = f
	if err != nil {
		fmt.Println("create file failed:", outfile)
		fmt.Println(err.Error())
		return -1
	}
	f.Truncate(0)

	f.WriteString(`#ifndef _GORM_WRAP_H__
#define _GORM_WRAP_H__
#include "gorm_define.h"
`)

	// 包含所有的文件
	for _, game := range games {
		fileName := game.File[:len(game.File)-4]
		headerName := "gorm_client_table_opt_" + fileName + ".h"
		f.WriteString("#include \"" + headerName + "\"\n")
	}

	f.WriteString("\n\nnamespace gorm{\n\n")

	f.WriteString(`
class GORM_Wrap
{
public:
	int Init(char *cfgPath);

	static GORM_Wrap *Instance();
private:
	static GORM_Wrap *gormWrap;
};

inline GORM_Wrap* GORM_Wrap::Instance()
{
	return GORM_Wrap::gormWrap;
}

// namespace gorm end
}
`)

	f.WriteString("\n\n#endif")
	return 0
}
*/
