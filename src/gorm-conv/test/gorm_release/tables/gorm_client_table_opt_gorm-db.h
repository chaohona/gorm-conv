#ifndef _GORM_CLIENT_TABLE_OPT_GORM-DB_H__
#define _GORM_CLIENT_TABLE_OPT_GORM-DB_H__
#include "gorm-db.pb.h

class GORM_ClientTableAccount
{
public:
    static GORM_ClientTableAccount* Get(int region, int logic_zone, int physics_zone, int32 id);
    static int Get(int region, int logic_zone, int physics_zone, int64 &cbId, int32 id, int (*cb)(int64, GORM_ClientTableAccount*));
    static int Delete(int region, int logic_zone, int physics_zone, int64 &cbId, int (*cb)(int64), int32 id);
    static void SetPbMsg(int region, int logic_zone, int physics_zone, GORM_PB_Table_account *pbMsg, bool forceSave=false);
    int Delete(int (*cb)(int64));
    void SetPbMsg(GORM_PB_Table_account *pbMsg, bool forceSave=false);
    int SaveToDB();
    GORM_PB_Table_account *GetPbMsg();
    uint64 GetVersion();
    void SetVersion(uint64 version, bool forceSave=false);
    int32 GetId();
    void SetId(int32 id, bool forceSave=false);
    string &GetAccount();
    void SetAccount(string &account, bool forceSave=false);
    void SetAccount(const char* account, size_t size, bool forceSave=false);
    string &GetAllbinary();
    void SetAllbinary(string &allbinary, bool forceSave=false);
    void SetAllbinary(const char* allbinary, size_t size, bool forceSave=false);
private:
    GORM_PB_Table_account *pTablePbValue = nullptr;
};

class GORM_ClientTableBag
{
public:
    static GORM_ClientTableBag* Get(int region, int logic_zone, int physics_zone, int32 id);
    static int Get(int region, int logic_zone, int physics_zone, int64 &cbId, int32 id, int (*cb)(int64, GORM_ClientTableBag*));
    static int Delete(int region, int logic_zone, int physics_zone, int64 &cbId, int (*cb)(int64), int32 id);
    static void SetPbMsg(int region, int logic_zone, int physics_zone, GORM_PB_Table_bag *pbMsg, bool forceSave=false);
    int Delete(int (*cb)(int64));
    void SetPbMsg(GORM_PB_Table_bag *pbMsg, bool forceSave=false);
    int SaveToDB();
    GORM_PB_Table_bag *GetPbMsg();
    uint64 GetVersion();
    void SetVersion(uint64 version, bool forceSave=false);
    int32 GetId();
    void SetId(int32 id, bool forceSave=false);
    string &GetAllbinary();
    void SetAllbinary(string &allbinary, bool forceSave=false);
    void SetAllbinary(const char* allbinary, size_t size, bool forceSave=false);
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
inline string& GORM_ClientTableAccount::GetAccount()
{
    return this->pTablePbValue->account();
}
inline string& GORM_ClientTableAccount::GetAllbinary()
{
    return this->pTablePbValue->allbinary();
}
inline void GORM_ClientTableAccount::SetPbMsg(GORM_PB_Table_account *pbMsg, bool forceSave=false)
{
    this->SetVersion(pbMsg->version(), forceSave);
    this->SetId(pbMsg->id(), forceSave);
    this->SetAccount(pbMsg->account(), forceSave);
    this->SetAllbinary(pbMsg->allbinary(), forceSave);
    return;
}
inline void GORM_ClientTableAccount::SetVersion(uint64 version, bool forceSave=false)
{
    this->pTablePbValue->set_version(version);
    return;
}
inline void GORM_ClientTableAccount::SetId(int32 id, bool forceSave=false)
{
    this->pTablePbValue->set_id(id);
    return;
}
inline void GORM_ClientTableAccount::SetAccount(string &account, bool forceSave=false)
{
    this->pTablePbValue->set_account(account);
    return;
}
inline void GORM_ClientTableAccount::SetAccount(const char* account, size_t size, bool forceSave=false)
{
    this->pTablePbValue->set_account(account, size);
    return;
}
inline void GORM_ClientTableAccount::SetAllbinary(string &allbinary, bool forceSave=false)
{
    this->pTablePbValue->set_allbinary(allbinary);
    return;
}
inline void GORM_ClientTableAccount::SetAllbinary(const char* allbinary, size_t size, bool forceSave=false)
{
    this->pTablePbValue->set_allbinary(allbinary, size);
    return;
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
inline string& GORM_ClientTableBag::GetAllbinary()
{
    return this->pTablePbValue->allbinary();
}
inline void GORM_ClientTableBag::SetPbMsg(GORM_PB_Table_bag *pbMsg, bool forceSave=false)
{
    this->SetVersion(pbMsg->version(), forceSave);
    this->SetId(pbMsg->id(), forceSave);
    this->SetAllbinary(pbMsg->allbinary(), forceSave);
    return;
}
inline void GORM_ClientTableBag::SetVersion(uint64 version, bool forceSave=false)
{
    this->pTablePbValue->set_version(version);
    return;
}
inline void GORM_ClientTableBag::SetId(int32 id, bool forceSave=false)
{
    this->pTablePbValue->set_id(id);
    return;
}
inline void GORM_ClientTableBag::SetAllbinary(string &allbinary, bool forceSave=false)
{
    this->pTablePbValue->set_allbinary(allbinary);
    return;
}
inline void GORM_ClientTableBag::SetAllbinary(const char* allbinary, size_t size, bool forceSave=false)
{
    this->pTablePbValue->set_allbinary(allbinary, size);
    return;
}


#endif