#ifndef _GORM_WRAP_H__
#define _GORM_WRAP_H__
#include "gorm_client_table_opt_gorm-db.h"

class GORM_Wrap
{
public:
	int Init(char *cfgPath);

	static GORM_ServerWrap *Instance();
private:
	static GORM_ServerWrap *pServerWrap;
};

inline GORM_ServerWrap::Instance()
{
	return this->pServerWrap;
}


#endif