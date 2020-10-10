#ifndef _GORM_SERVER_WRAP_H__
#define _GORM_SERVER_WRAP_H__
#include "gorm_client_table_opt_test.h"

class GORM_ServerWrap
{
public:
	int Init(char *dbPath);

	static GORM_ServerWrap *Instance();
private:
	static GORM_ServerWrap *pServerWrap;
};

inline GORM_ServerWrap::Instance()
{
	return this->pServerWrap;
}

#endif