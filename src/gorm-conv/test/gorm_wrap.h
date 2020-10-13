#ifndef _GORM_WRAP_H__
#define _GORM_WRAP_H__
#include "gorm_define.h"
#include "gorm_client_table_opt_test.h"


namespace gorm{


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


#endif