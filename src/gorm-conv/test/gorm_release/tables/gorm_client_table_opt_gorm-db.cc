#include /gorm_client_table_opt_gorm-db.h
GORM_ClientTableAccount* GORM_ClientTableAccount::Get(int region, int logic_zone, int physics_zone,int32 id)
{
    return nullptr;
}
GORM_ClientTableAccount* GORM_ClientTableAccount::Get(int region, int logic_zone, int physics_zone, int64 &cbId, int32 id, int (*cb)(int64, GORM_ClientTableAccount*))
{
    return nullptr;
}
int GORM_ClientTableAccount::Delete(int region, int logic_zone, int physics_zone, int64 &cbId, int (*cb)(int64), int32 id)
{
    return 0;
}
int GORM_ClientTableAccount::SetPbMsg(int region, int logic_zone, int physics_zone, GORM_PB_Table_account *pbMsg, bool forceSave)
{
    return 0;
}
int GORM_ClientTableAccount::Delete(int32 id)
{
    return 0;
}
GORM_ClientTableBag* GORM_ClientTableBag::Get(int region, int logic_zone, int physics_zone,int32 id)
{
    return nullptr;
}
GORM_ClientTableBag* GORM_ClientTableBag::Get(int region, int logic_zone, int physics_zone, int64 &cbId, int32 id, int (*cb)(int64, GORM_ClientTableBag*))
{
    return nullptr;
}
int GORM_ClientTableBag::Delete(int region, int logic_zone, int physics_zone, int64 &cbId, int (*cb)(int64), int32 id)
{
    return 0;
}
int GORM_ClientTableBag::SetPbMsg(int region, int logic_zone, int physics_zone, GORM_PB_Table_bag *pbMsg, bool forceSave)
{
    return 0;
}
int GORM_ClientTableBag::Delete(int32 id)
{
    return 0;
}
