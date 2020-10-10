#include "gorm_wrap.h"

using namespace gorm;

GORM_Wrap *global_gorm_wrap = GORM_Wrap::Instance();

int main()
{
	// 初始化gorm驱动
	if (0 != global_gorm_wrap->Init())
	{
		return -1;
	}
	
	// 获取主键为1的account表数据
	GORM_ClientTableAccount *pAccount = GORM_ClientTableAccount::Get(0,0,0,1);
	
	// 获取pb数据，发送给逻辑进程使用
	GORM_PB_Table_account *pPbAcc = pAccount->GetPbMsg();
	
	// 更新并持久化保存数据
	char *newAcc = "new account";
	pAccount->SetAccount(newAcc, strlen(newAcc));
	pAccount->SaveToDB();
	return 0;
}