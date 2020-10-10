#include "gorm_wrap.h"

namespace gorm{

GORM_Wrap *GORM_Wrap::pServerWrap = new GORM_Wrap();

int GORM_Wrap::Init(char *cfgPath)
{
	return 0;
}

}