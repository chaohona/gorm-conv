#ifndef _GORM_CLIENT_TABLE_OPT_GORM-DB_H__
#define _GORM_CLIENT_TABLE_OPT_GORM-DB_H__
#include "gorm-db.pb.h"
#include "gorm_define.h"

namespace gorm{

class GORM_ClientTableAccount
{
public:
    static GORM_ClientTableAccount* Get(int region, int logic_zone, int physics_zone, int32 id);
    static int Get(int region, int logic_zone, int physics_zone, int64 &cbId, int32 id, int (*cb)(int64, GORM_ClientTableAccount*));
    static int Delete(int region, int logic_zone, int physics_zone, int64 &cbId, int (*cb)(int64), int32 id);
    static int SetPbMsg(int region, int logic_zone, int physics_zone, GORM_PB_Table_account *pbMsg, bool forceSave=false);
    int Delete(int64 &cbId, int (*cb)(int64));
    int SetPbMsg(GORM_PB_Table_account *pbMsg, bool forceSave=false);
    int SaveToDB();
    GORM_PB_Table_account *GetPbMsg();
    uint64 GetVersion();
    int SetVersion(uint64 version, bool forceSave=false);
    int32 GetId();
    int SetId(int32 id, bool forceSave=false);
    const string &GetAccount() const;
    int SetAccount(const string &account, bool forceSave=false);
    int SetAccount(const char* account, size_t size, bool forceSave=false);
    const string &GetAllbinary() const;
    int SetAllbinary(const string &allbinary, bool forceSave=false);
    int SetAllbinary(const char* allbinary, size_t size, bool forceSave=false);
private:
    GORM_PB_Table_account *pTablePbValue = nullptr;
};

class GORM_ClientTableBag
{
public:
    static GORM_ClientTableBag* Get(int region, int logic_zone, int physics_zone, int32 id);
    static int Get(int region, int logic_zone, int physics_zone, int64 &cbId, int32 id, int (*cb)(int64, GORM_ClientTableBag*));
    static int Delete(int region, int logic_zone, int physics_zone, int64 &cbId, int (*cb)(int64), int32 id);
    static int SetPbMsg(int region, int logic_zone, int physics_zone, GORM_PB_Table_bag *pbMsg, bool forceSave=false);
    int Delete(int64 &cbId, int (*cb)(int64));
    int SetPbMsg(GORM_PB_Table_bag *pbMsg, bool forceSave=false);
    int SaveToDB();
    GORM_PB_Table_bag *GetPbMsg();
    uint64 GetVersion();
    int SetVersion(uint64 version, bool forceSave=false);
    int32 GetId();
    int SetId(int32 id, bool forceSave=false);
    const string &GetAllbinary() const;
    int SetAllbinary(const string &allbinary, bool forceSave=false);
    int SetAllbinary(const char* allbinary, size_t size, bool forceSave=false);
private:
    GORM_PB_Table_bag *pTablePbValue = nullptr;
};

inline GORM_PB_Table_account*GORM_ClientTableAccount::GetPbMsg()
{
    return this->pTablePbValue;
}
inline uint64 GORM_ClientTableAccount::GetVersion()
{
    return this->pTablePbValue->version();
}
inline int32 GORM_ClientTableAccount::GetId()
{
    return this->pTablePbValue->id();
}
inline const string& GORM_ClientTableAccount::GetAccount() const
{
    return this->pTablePbValue->account();
}
inline const string& GORM_ClientTableAccount::GetAllbinary() const
{
    return this->pTablePbValue->allbinary();
}
inline int GORM_ClientTableAccount::SetPbMsg(GORM_PB_Table_account *pbMsg, bool forceSave)
{
    this->SetVersion(pbMsg->version(), forceSave);
    this->SetId(pbMsg->id(), forceSave);
    this->SetAccount(pbMsg->account(), forceSave);
    this->SetAllbinary(pbMsg->allbinary(), forceSave);
    return 0;
}
inline int GORM_ClientTableAccount::SetVersion(uint64 version, bool forceSave)
{
    this->pTablePbValue->set_version(version);
    return 0;
}
inline int GORM_ClientTableAccount::SetId(int32 id, bool forceSave)
{
    this->pTablePbValue->set_id(id);
    return 0;
}
inline int GORM_ClientTableAccount::SetAccount(const string &account, bool forceSave)
{
    this->pTablePbValue->set_account(account);
    return 0;
}
inline int GORM_ClientTableAccount::SetAccount(const char* account, size_t size, bool forceSave)
{
    this->pTablePbValue->set_account(account, size);
    return 0;
}
inline int GORM_ClientTableAccount::SetAllbinary(const string &allbinary, bool forceSave)
{
    this->pTablePbValue->set_allbinary(allbinary);
    return 0;
}
inline int GORM_ClientTableAccount::SetAllbinary(const char* allbinary, size_t size, bool forceSave)
{
    this->pTablePbValue->set_allbinary(allbinary, size);
    return 0;
}
inline GORM_PB_Table_bag*GORM_ClientTableBag::GetPbMsg()
{
    return this->pTablePbValue;
}
inline uint64 GORM_ClientTableBag::GetVersion()
{
    return this->pTablePbValue->version();
}
inline int32 GORM_ClientTableBag::GetId()
{
    return this->pTablePbValue->id();
}
inline const string& GORM_ClientTableBag::GetAllbinary() const
{
    return this->pTablePbValue->allbinary();
}
inline int GORM_ClientTableBag::SetPbMsg(GORM_PB_Table_bag *pbMsg, bool forceSave)
{
    this->SetVersion(pbMsg->version(), forceSave);
    this->SetId(pbMsg->id(), forceSave);
    this->SetAllbinary(pbMsg->allbinary(), forceSave);
    return 0;
}
inline int GORM_ClientTableBag::SetVersion(uint64 version, bool forceSave)
{
    this->pTablePbValue->set_version(version);
    return 0;
}
inline int GORM_ClientTableBag::SetId(int32 id, bool forceSave)
{
    this->pTablePbValue->set_id(id);
    return 0;
}
inline int GORM_ClientTableBag::SetAllbinary(const string &allbinary, bool forceSave)
{
    this->pTablePbValue->set_allbinary(allbinary);
    return 0;
}
inline int GORM_ClientTableBag::SetAllbinary(const char* allbinary, size_t size, bool forceSave)
{
    this->pTablePbValue->set_allbinary(allbinary, size);
    return 0;
}

}



#endif