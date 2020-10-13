#ifndef _GORM_CLIENT_TABLE_OPT_TEST_H__
#define _GORM_CLIENT_TABLE_OPT_TEST_H__
#include "test.pb.h"
#include "gorm_define.h"

/*
// 函数Get、Delete、SetPbMsg、SaveToDB、GetPbMsg为所有的表都会对应的操作函数
// 除了上面4个函数，其它函数都为更新表结构中的字段
class GORM_ClientTableAccount
{
public:
	// @desc 	以同步方式获取account表中的一条数据，此函数有一个对应的不带区服信息的函数
	// @param	region数据所在可用域（机房）
	// @param	logic_zone数据所属逻辑区
	// @param	physics_zone数组所属物理区
	// @param	id表的主键，不同的表主键类型与个数不同
	// @retval	获取到的数据
    static GORM_ClientTableAccount* Get(int region, int logic_zone, int physics_zone, int32 id);
    // 不带区服信息的函数
    static GORM_ClientTableAccount* Get(int32 id);
    
    // @desc 	以异步方式获取account表中的一条数据，此函数有一个对应的不带区服信息的函数
	// @param	region数据所在可用域（机房）
	// @param	logic_zone数据所属逻辑区
	// @param	physics_zone数组所属物理区
	// @param	cbId异步请求的回调id
	// @param	id表的主键，不同的表主键类型与个数不同
	// @param	cb异步请求的回调函数
	//			回调函数的参数解释:
	//			int64对应上面的cbId, GORM_ClientTableAccount*获取到的数据的
	// @retval	0  成功，等待回调
				<0 失败
    static int Get(int region, int logic_zone, int physics_zone, int32 id, int64 &cbId, int (*cb)(int64, GORM_ClientTableAccount*));
    // 不带区服信息的函数
    static int Get(int32 id, int64 &cbId, int (*cb)(int64, GORM_ClientTableAccount*));
    
    // @desc 	以异步方式删除account表中的一条数据，此函数有一个对应的不带区服信息的函数
	// @param	region数据所在可用域（机房）
	// @param	logic_zone数据所属逻辑区
	// @param	physics_zone数组所属物理区
	// @param	cbId异步请求的回调id
	// @param	cb异步请求的回调函数,可为nullptr，表示不关心结果
	//			回调函数的参数解释:
	//			int64对应上面的cbId
	// @param	id表的主键，不同的表主键类型与个数不同
	// @retval	0  成功，等待回调
				<0 失败
    static int Delete(int region, int logic_zone, int physics_zone, int32 id, int64 &cbId, int (*cb)(int64));
    // 不带区服信息的函数
    static int Delete(int32 id, int64 &cbId, int (*cb)(int64));
    
    // @desc 	全量更新数据，此函数有一个对应的不带区服信息的函数
    // @param	pbMsg需要覆盖更新的pb数据
    // @param	forceSave是否立即持久化，默认值为false不立即持久化
    // @retval	0  成功
				<0 失败
    static int SetPbMsg(int region, int logic_zone, int physics_zone, GORM_PB_Table_account *pbMsg, bool forceSave=false);
    // 不带区服信息的函数
    static int SetPbMsg(GORM_PB_Table_account *pbMsg, bool forceSave=false);
    
    // @desc 删除本条数据
    int Delete(int64 &cbId, int (*cb)(int64));
    
    // @desc 全量覆盖更新本条数据
    int SetPbMsg(GORM_PB_Table_account *pbMsg, bool forceSave=false);
    
    // @desc 立即将本条数据持久化保存
    int SaveToDB();
    
    // @desc 获取原始pb结构数据，长用于发送给后端逻辑节点使用
    GORM_PB_Table_account *GetPbMsg();
    
    // 以下为针对字段的Get，Set操作
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
    GORM_PB_Table_account *tablePbValue = nullptr;
};
*/
namespace gorm{

// 表account
class GORM_ClientTableAccount
{
public:
    // static带区服的接口，用于分区分服架构
    static GORM_ClientTableAccount* Get(int region, int logic_zone, int physics_zone, int32 id);
    static int Get(int region, int logic_zone, int physics_zone, int32 id, int64 &cbId, int (*cb)(int64, GORM_ClientTableAccount*));
    static int Delete(int region, int logic_zone, int physics_zone, int32 id, int64 &cbId, int (*cb)(int64));
    static int SetPbMsg(int region, int logic_zone, int physics_zone, GORM_PB_Table_account *pbMsg, bool forceSave=false);

    // static不带区服的接口，用于全区全服架构
    static GORM_ClientTableAccount* Get(int32 id);
    static int Get(int32 id, int64 &cbId, int (*cb)(int64, GORM_ClientTableAccount*));
    static int Delete(int32 id, int64 &cbId, int (*cb)(int64));
    static int SetPbMsg(GORM_PB_Table_account *pbMsg, bool forceSave=false);

    // 本地操作接口
    int Delete(int64 &cbId, int (*cb)(int64));
    int SetPbMsg(GORM_PB_Table_account *pbMsg, bool forceSave=false);
    int SaveToDB();
    GORM_PB_Table_account *GetPbMsg();

    /* 下面为针对表的每个字段的操作,Get为获取字段的原始数据,Set为更新字段的值 */
    // version
    uint64 GetVersion();
    int SetVersion(uint64 version, bool forceSave=false);
    // id
    int32 GetId();
    int SetId(int32 id, bool forceSave=false);
    // account
    const string &GetAccount() const;
    int SetAccount(const string &account, bool forceSave=false);
    int SetAccount(const char* account, size_t size, bool forceSave=false);
    // allbinary
    const string &GetAllbinary() const;
    int SetAllbinary(const string &allbinary, bool forceSave=false);
    int SetAllbinary(const char* allbinary, size_t size, bool forceSave=false);
private:
    GORM_PB_Table_account *tablePbValue = nullptr;
    GORM_FieldsOpt m_fieldOpt;
};

// 表bag
class GORM_ClientTableBag
{
public:
    // static带区服的接口，用于分区分服架构
    static GORM_ClientTableBag* Get(int region, int logic_zone, int physics_zone, int32 id);
    static int Get(int region, int logic_zone, int physics_zone, int32 id, int64 &cbId, int (*cb)(int64, GORM_ClientTableBag*));
    static int Delete(int region, int logic_zone, int physics_zone, int32 id, int64 &cbId, int (*cb)(int64));
    static int SetPbMsg(int region, int logic_zone, int physics_zone, GORM_PB_Table_bag *pbMsg, bool forceSave=false);

    // static不带区服的接口，用于全区全服架构
    static GORM_ClientTableBag* Get(int32 id);
    static int Get(int32 id, int64 &cbId, int (*cb)(int64, GORM_ClientTableBag*));
    static int Delete(int32 id, int64 &cbId, int (*cb)(int64));
    static int SetPbMsg(GORM_PB_Table_bag *pbMsg, bool forceSave=false);

    // 本地操作接口
    int Delete(int64 &cbId, int (*cb)(int64));
    int SetPbMsg(GORM_PB_Table_bag *pbMsg, bool forceSave=false);
    int SaveToDB();
    GORM_PB_Table_bag *GetPbMsg();

    /* 下面为针对表的每个字段的操作,Get为获取字段的原始数据,Set为更新字段的值 */
    // version
    uint64 GetVersion();
    int SetVersion(uint64 version, bool forceSave=false);
    // id
    int32 GetId();
    int SetId(int32 id, bool forceSave=false);
    // allbinary
    const string &GetAllbinary() const;
    int SetAllbinary(const string &allbinary, bool forceSave=false);
    int SetAllbinary(const char* allbinary, size_t size, bool forceSave=false);
private:
    GORM_PB_Table_bag *tablePbValue = nullptr;
    GORM_FieldsOpt m_fieldOpt;
};

inline GORM_PB_Table_account*GORM_ClientTableAccount::GetPbMsg()
{
    return this->tablePbValue;
}
inline uint64 GORM_ClientTableAccount::GetVersion()
{
    return this->tablePbValue->version();
}
inline int32 GORM_ClientTableAccount::GetId()
{
    return this->tablePbValue->id();
}
inline const string& GORM_ClientTableAccount::GetAccount() const
{
    return this->tablePbValue->account();
}
inline const string& GORM_ClientTableAccount::GetAllbinary() const
{
    return this->tablePbValue->allbinary();
}
inline int GORM_ClientTableAccount::SetPbMsg(GORM_PB_Table_account *pbMsg, bool forceSave)
{
    this->SetVersion(pbMsg->version());
    if (forceSave)
        return this->SaveToDB();
    this->SetId(pbMsg->id());
    if (forceSave)
        return this->SaveToDB();
    this->SetAccount(pbMsg->account());
    if (forceSave)
        return this->SaveToDB();
    this->SetAllbinary(pbMsg->allbinary());
    if (forceSave)
        return this->SaveToDB();
    return 0;
}
inline int GORM_ClientTableAccount::SetVersion(uint64 version, bool forceSave)
{
    this->tablePbValue->set_version(version);
    this->m_fieldOpt.AddField(0, 1);
    if (forceSave)
        return this->SaveToDB();
    return 0;
}
inline int GORM_ClientTableAccount::SetId(int32 id, bool forceSave)
{
    this->tablePbValue->set_id(id);
    this->m_fieldOpt.AddField(0, 2);
    if (forceSave)
        return this->SaveToDB();
    return 0;
}
inline int GORM_ClientTableAccount::SetAccount(const string &account, bool forceSave)
{
    this->tablePbValue->set_account(account);
    this->m_fieldOpt.AddField(0, 4);
    if (forceSave)
        return this->SaveToDB();
    return 0;
}
inline int GORM_ClientTableAccount::SetAccount(const char* account, size_t size, bool forceSave)
{
    this->tablePbValue->set_account(account, size);
    this->m_fieldOpt.AddField(0, 4);
    if (forceSave)
        return this->SaveToDB();
    return 0;
}
inline int GORM_ClientTableAccount::SetAllbinary(const string &allbinary, bool forceSave)
{
    this->tablePbValue->set_allbinary(allbinary);
    this->m_fieldOpt.AddField(0, 8);
    if (forceSave)
        return this->SaveToDB();
    return 0;
}
inline int GORM_ClientTableAccount::SetAllbinary(const char* allbinary, size_t size, bool forceSave)
{
    this->tablePbValue->set_allbinary(allbinary, size);
    this->m_fieldOpt.AddField(0, 8);
    if (forceSave)
        return this->SaveToDB();
    return 0;
}
inline GORM_PB_Table_bag*GORM_ClientTableBag::GetPbMsg()
{
    return this->tablePbValue;
}
inline uint64 GORM_ClientTableBag::GetVersion()
{
    return this->tablePbValue->version();
}
inline int32 GORM_ClientTableBag::GetId()
{
    return this->tablePbValue->id();
}
inline const string& GORM_ClientTableBag::GetAllbinary() const
{
    return this->tablePbValue->allbinary();
}
inline int GORM_ClientTableBag::SetPbMsg(GORM_PB_Table_bag *pbMsg, bool forceSave)
{
    this->SetVersion(pbMsg->version());
    if (forceSave)
        return this->SaveToDB();
    this->SetId(pbMsg->id());
    if (forceSave)
        return this->SaveToDB();
    this->SetAllbinary(pbMsg->allbinary());
    if (forceSave)
        return this->SaveToDB();
    return 0;
}
inline int GORM_ClientTableBag::SetVersion(uint64 version, bool forceSave)
{
    this->tablePbValue->set_version(version);
    this->m_fieldOpt.AddField(0, 1);
    if (forceSave)
        return this->SaveToDB();
    return 0;
}
inline int GORM_ClientTableBag::SetId(int32 id, bool forceSave)
{
    this->tablePbValue->set_id(id);
    this->m_fieldOpt.AddField(0, 2);
    if (forceSave)
        return this->SaveToDB();
    return 0;
}
inline int GORM_ClientTableBag::SetAllbinary(const string &allbinary, bool forceSave)
{
    this->tablePbValue->set_allbinary(allbinary);
    this->m_fieldOpt.AddField(0, 4);
    if (forceSave)
        return this->SaveToDB();
    return 0;
}
inline int GORM_ClientTableBag::SetAllbinary(const char* allbinary, size_t size, bool forceSave)
{
    this->tablePbValue->set_allbinary(allbinary, size);
    this->m_fieldOpt.AddField(0, 4);
    if (forceSave)
        return this->SaveToDB();
    return 0;
}

}



#endif