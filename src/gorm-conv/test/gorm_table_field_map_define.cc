#include "gorm_table_field_map_define.h"
#include "gorm_pb_tables_inc.pb.h"
#include "gorm_pb_proto.pb.h"
#include "gorm_mempool.h"
#include "gorm_hash.h"

namespace gorm{

int GORM_SetTableACCOUNTId2Name(OUT FieldId2Name &mapId2Name);
int GORM_SetTableACCOUNTName2Id(OUT FieldName2Id &mapName2Id);
int GORM_SetTableBAGId2Name(OUT FieldId2Name &mapId2Name);
int GORM_SetTableBAGName2Id(OUT FieldName2Id &mapName2Id);


int GORM_InitTableSchemaInfo(PB_MSG_PTR pMsgPtr)
{
    GORM_PB_HAND_SHAKE_REQ *pHandShake = dynamic_cast<GORM_PB_HAND_SHAKE_REQ*>(pMsgPtr);
    if (pHandShake == nullptr)
        return GORM_ERROR;
    GORM_PB_TABLE_SCHEMA_INFO *pInfo;
    GORM_PB_TABLE_SCHEMA_INFO_COLUMN *pColumn;
    pHandShake->mutable_header();
    pHandShake->set_version(1);
    pHandShake->set_md5(1);
    // for table account
    pInfo = pHandShake->add_schemas();
    if (pInfo == nullptr)
        return GORM_ERROR;
    pInfo->set_version(1);
    pInfo->set_tablename("account");
    pInfo->set_tableidx(1);
    pColumn = pInfo->add_columns();
    if (pColumn == nullptr)
        return GORM_ERROR;
    pColumn->set_version(0);
    pColumn->set_name("version");
    pColumn->set_typedesc("uint64");
    pColumn->set_type(GORM_PB_COLUMN_TYPE_UINT64);
    pColumn = pInfo->add_columns();
    if (pColumn == nullptr)
        return GORM_ERROR;
    pColumn->set_version(0);
    pColumn->set_name("id");
    pColumn->set_typedesc("int");
    pColumn->set_type(GORM_PB_COLUMN_TYPE_INT);
    pColumn = pInfo->add_columns();
    if (pColumn == nullptr)
        return GORM_ERROR;
    pColumn->set_version(0);
    pColumn->set_name("account");
    pColumn->set_typedesc("string");
    pColumn->set_type(GORM_PB_COLUMN_TYPE_STRING);
    pColumn = pInfo->add_columns();
    if (pColumn == nullptr)
        return GORM_ERROR;
    pColumn->set_version(0);
    pColumn->set_name("allbinary");
    pColumn->set_typedesc("blob");
    pColumn->set_type(GORM_PB_COLUMN_TYPE_BLOB);
    // for table bag
    pInfo = pHandShake->add_schemas();
    if (pInfo == nullptr)
        return GORM_ERROR;
    pInfo->set_version(1);
    pInfo->set_tablename("bag");
    pInfo->set_tableidx(2);
    pColumn = pInfo->add_columns();
    if (pColumn == nullptr)
        return GORM_ERROR;
    pColumn->set_version(0);
    pColumn->set_name("version");
    pColumn->set_typedesc("uint64");
    pColumn->set_type(GORM_PB_COLUMN_TYPE_UINT64);
    pColumn = pInfo->add_columns();
    if (pColumn == nullptr)
        return GORM_ERROR;
    pColumn->set_version(0);
    pColumn->set_name("id");
    pColumn->set_typedesc("int");
    pColumn->set_type(GORM_PB_COLUMN_TYPE_INT);
    pColumn = pInfo->add_columns();
    if (pColumn == nullptr)
        return GORM_ERROR;
    pColumn->set_version(0);
    pColumn->set_name("allbinary");
    pColumn->set_typedesc("blob");
    pColumn->set_type(GORM_PB_COLUMN_TYPE_BLOB);
    return GORM_OK;
}
uint32 GORM_TableHash(int iTableId, const GORM_PB_TABLE &pbTable)
{
    switch (iTableId)
    {
    case GORM_PB_TABLE_IDX_ACCOUNT:
    {
        if (!pbTable.has_account())
            return 0;
        const GORM_PB_Table_account& account = pbTable.account();
        int tmp_ID = account.id();
        char szSrcHash[1024];
        int iTotalLen = GORM_SafeSnprintf(szSrcHash, 1024, "%d_" , tmp_ID);
        if (iTotalLen > 1024)
            iTotalLen = 1024;
        return GORM_Hash::Crc32_1((const char*)szSrcHash, iTotalLen);
    }
    case GORM_PB_TABLE_IDX_BAG:
    {
        if (!pbTable.has_bag())
            return 0;
        const GORM_PB_Table_bag& bag = pbTable.bag();
        int tmp_ID = bag.id();
        char szSrcHash[1024];
        int iTotalLen = GORM_SafeSnprintf(szSrcHash, 1024, "%d_" , tmp_ID);
        if (iTotalLen > 1024)
            iTotalLen = 1024;
        return GORM_Hash::Crc32_1((const char*)szSrcHash, iTotalLen);
    }
    default:
        return 0;
    }
    return 0;
}

int GORM_GetCustomPbMsg(PB_MSG_PTR &pMsgPtr)
{
    pMsgPtr = new GORM_PB_CUSTEM_COLUMNS();
    return GORM_OK;
}
int GetTablePbMsgDefine(int iTableId, PB_MSG_PTR &pMsgPtr)
{
    switch (iTableId)
    {
    case GORM_PB_TABLE_IDX_ACCOUNT:
    {
        pMsgPtr = new GORM_PB_Table_account();
        return GORM_OK;
    }
    case GORM_PB_TABLE_IDX_BAG:
    {
        pMsgPtr = new GORM_PB_Table_bag();
        return GORM_OK;
    }
    }
    return GORM_INVALID_TABLE;
}
bool GORM_TableHasData(GORM_PB_TABLE *pTable, int iTableId)
{
    switch (iTableId)
    {
    case GORM_PB_TABLE_IDX_ACCOUNT:
        return pTable->has_account();
    case GORM_PB_TABLE_IDX_BAG:
        return pTable->has_bag();
    }

    return false;
}
int GORM_GetTableSrcPbMsg(int iTableId, GORM_PB_TABLE *pTable, PB_MSG_PTR &pMsgPtr)
{
    switch (iTableId)
    {
    case GORM_PB_TABLE_IDX_ACCOUNT:
    {
        pMsgPtr = pTable->mutable_account();
        return GORM_OK;
    }
    case GORM_PB_TABLE_IDX_BAG:
    {
        pMsgPtr = pTable->mutable_bag();
        return GORM_OK;
    }
    }

    return false;
}
int GORM_AddRecordToReqPbMsgDefine(int iTableId, GORM_PB_TABLE *pPbTable, PB_MSG_PTR pPbMsg)
{
    switch (iTableId)
    {
    case GORM_PB_TABLE_IDX_ACCOUNT:
    {
        GORM_PB_Table_account *pTableMsg = dynamic_cast<GORM_PB_Table_account*>(pPbMsg);
        pPbTable->set_allocated_account(pTableMsg);
        return GORM_OK;
    }
    case GORM_PB_TABLE_IDX_BAG:
    {
        GORM_PB_Table_bag *pTableMsg = dynamic_cast<GORM_PB_Table_bag*>(pPbMsg);
        pPbTable->set_allocated_bag(pTableMsg);
        return GORM_OK;
    }
    }
    return GORM_INVALID_TABLE;
}
int GORM_SetTableVersion(OUT TableVersionMap& mapTableVersion)
{
    mapTableVersion[GORM_PB_TABLE_IDX_ACCOUNT] = 1;
    mapTableVersion[GORM_PB_TABLE_IDX_BAG] = 1;
    return GORM_OK;
}

int GORM_SetTableName2Id(OUT TableName2Id &mapName2Id)
{
    mapName2Id["account"] = GORM_PB_TABLE_IDX_ACCOUNT;
    mapName2Id["bag"] = GORM_PB_TABLE_IDX_BAG;
    return GORM_OK;
}

int GORM_SetTableId2Name(OUT TableId2Name &mapId2Name)
{
    mapId2Name[GORM_PB_TABLE_IDX_ACCOUNT] = "account";
    mapId2Name[GORM_PB_TABLE_IDX_BAG] = "bag";
    return GORM_OK;
}

int GORM_SetTableFieldId2Name(int iTableType, OUT FieldId2Name &mapId2Name)
{
    switch (iTableType)
    {
    case GORM_PB_TABLE_IDX_ACCOUNT:
    {
        return GORM_SetTableACCOUNTId2Name(mapId2Name);
    }
    case GORM_PB_TABLE_IDX_BAG:
    {
        return GORM_SetTableBAGId2Name(mapId2Name);
    }
    default:
        return GORM_ERROR;
    }

    return GORM_OK;
}
int GORM_SetTableFieldName2Id(int iTableType, OUT FieldName2Id &mapName2Id)
{
    switch (iTableType)
    {
    case GORM_PB_TABLE_IDX_ACCOUNT:
    {
        return GORM_SetTableACCOUNTName2Id(mapName2Id);
    }
    case GORM_PB_TABLE_IDX_BAG:
    {
        return GORM_SetTableBAGName2Id(mapName2Id);
    }
    default:
        return GORM_ERROR;
    }

    return GORM_OK;
}
int GORM_SetTableACCOUNTId2Name(OUT FieldId2Name &mapId2Name)
{
    mapId2Name[GORM_PB_FIELD_ACCOUNT_VERSION] = "version";
    mapId2Name[GORM_PB_FIELD_ACCOUNT_ID] = "id";
    mapId2Name[GORM_PB_FIELD_ACCOUNT_ACCOUNT] = "account";
    mapId2Name[GORM_PB_FIELD_ACCOUNT_ALLBINARY] = "allbinary";
    return GORM_OK;
}
int GORM_SetTableBAGId2Name(OUT FieldId2Name &mapId2Name)
{
    mapId2Name[GORM_PB_FIELD_BAG_VERSION] = "version";
    mapId2Name[GORM_PB_FIELD_BAG_ID] = "id";
    mapId2Name[GORM_PB_FIELD_BAG_ALLBINARY] = "allbinary";
    return GORM_OK;
}
int GORM_SetTableACCOUNTName2Id(OUT FieldName2Id &mapName2Id)
{
    mapName2Id["version"] = GORM_PB_FIELD_ACCOUNT_VERSION;
    mapName2Id["id"] = GORM_PB_FIELD_ACCOUNT_ID;
    mapName2Id["account"] = GORM_PB_FIELD_ACCOUNT_ACCOUNT;
    mapName2Id["allbinary"] = GORM_PB_FIELD_ACCOUNT_ALLBINARY;
    return GORM_OK;
}
int GORM_SetTableBAGName2Id(OUT FieldName2Id &mapName2Id)
{
    mapName2Id["version"] = GORM_PB_FIELD_BAG_VERSION;
    mapName2Id["id"] = GORM_PB_FIELD_BAG_ID;
    mapName2Id["allbinary"] = GORM_PB_FIELD_BAG_ALLBINARY;
    return GORM_OK;
}
int GORM_InitTableColumnInfo_account(unordered_map<string, vector<string>> &mapTablesColumnOrder, unordered_map<string, unordered_map<string, GORM_PB_COLUMN_TYPE>> &mapTablesColumnInfo)
{
    vector<string> vColumns = {"version","id","account","allbinary"};
    mapTablesColumnOrder["account"] = vColumns;
    unordered_map<string, GORM_PB_COLUMN_TYPE> mapColumnType = {
        {"version", GORM_PB_COLUMN_TYPE_UINT},
        {"id", GORM_PB_COLUMN_TYPE_INT},
        {"account", GORM_PB_COLUMN_TYPE_STRING},
        {"allbinary", GORM_PB_COLUMN_TYPE_STRING},
    };
    mapTablesColumnInfo["account"] = mapColumnType;
    return GORM_OK;
}
int GORM_InitTableColumnInfo_bag(unordered_map<string, vector<string>> &mapTablesColumnOrder, unordered_map<string, unordered_map<string, GORM_PB_COLUMN_TYPE>> &mapTablesColumnInfo)
{
    vector<string> vColumns = {"version","id","allbinary"};
    mapTablesColumnOrder["bag"] = vColumns;
    unordered_map<string, GORM_PB_COLUMN_TYPE> mapColumnType = {
        {"version", GORM_PB_COLUMN_TYPE_UINT},
        {"id", GORM_PB_COLUMN_TYPE_INT},
        {"allbinary", GORM_PB_COLUMN_TYPE_STRING},
    };
    mapTablesColumnInfo["bag"] = mapColumnType;
    return GORM_OK;
}
int GORM_InitTableColumnInfo(unordered_map<string, vector<string>> &mapTablesColumnOrder, unordered_map<string, unordered_map<string, GORM_PB_COLUMN_TYPE>> &mapTablesColumnInfo)
{
    if (GORM_InitTableColumnInfo_account(mapTablesColumnOrder, mapTablesColumnInfo))
        return GORM_ERROR;
    if (GORM_InitTableColumnInfo_bag(mapTablesColumnOrder, mapTablesColumnInfo))
        return GORM_ERROR;
    return GORM_OK;
}
void GORM_SetTableFieldValue(PB_MSG_PTR pMsg, int iTableId, int iFieldId, const char * value, const size_t size)
{
    switch (iTableId)
    {
    case GORM_PB_TABLE_IDX_ACCOUNT:
    {
        switch (iFieldId)
        {
        case GORM_PB_FIELD_ACCOUNT_ACCOUNT:
        {
            GORM_PB_Table_account* pPbReal = dynamic_cast<GORM_PB_Table_account*>(pMsg);
            return pPbReal->set_account((const char*)value, size);
        }
        case GORM_PB_FIELD_ACCOUNT_ALLBINARY:
        {
            GORM_PB_Table_account* pPbReal = dynamic_cast<GORM_PB_Table_account*>(pMsg);
            return pPbReal->set_allbinary((const char*)value, size);
        }
        }
    }
    case GORM_PB_TABLE_IDX_BAG:
    {
        switch (iFieldId)
        {
        case GORM_PB_FIELD_BAG_ALLBINARY:
        {
            GORM_PB_Table_bag* pPbReal = dynamic_cast<GORM_PB_Table_bag*>(pMsg);
            return pPbReal->set_allbinary((const char*)value, size);
        }
        }
    }
    }
}

void GORM_SetTableFieldValue(PB_MSG_PTR pMsg, int iTableId, int iFieldId, const char * value)
{
    switch (iTableId)
    {
    case GORM_PB_TABLE_IDX_ACCOUNT:
    {
        switch (iFieldId)
        {
        case GORM_PB_FIELD_ACCOUNT_ACCOUNT:
        {
            GORM_PB_Table_account* pPbReal = dynamic_cast<GORM_PB_Table_account*>(pMsg);
            return pPbReal->set_account((const char*)value);
        }
        case GORM_PB_FIELD_ACCOUNT_ALLBINARY:
        {
            GORM_PB_Table_account* pPbReal = dynamic_cast<GORM_PB_Table_account*>(pMsg);
            return pPbReal->set_allbinary((const char*)value);
        }
        }
    }
    case GORM_PB_TABLE_IDX_BAG:
    {
        switch (iFieldId)
        {
        case GORM_PB_FIELD_BAG_ALLBINARY:
        {
            GORM_PB_Table_bag* pPbReal = dynamic_cast<GORM_PB_Table_bag*>(pMsg);
            return pPbReal->set_allbinary((const char*)value);
        }
        }
    }
    }
}

void GORM_SetTableFieldValue(PB_MSG_PTR pMsg, int iTableId, int iFieldId, int8 value)
{
    switch (iTableId)
    {
    case GORM_PB_TABLE_IDX_ACCOUNT:
    {
        switch (iFieldId)
        {
        case GORM_PB_FIELD_ACCOUNT_VERSION:
        {
            GORM_PB_Table_account* pPbReal = dynamic_cast<GORM_PB_Table_account*>(pMsg);
            return pPbReal->set_version(value);
        }
        case GORM_PB_FIELD_ACCOUNT_ID:
        {
            GORM_PB_Table_account* pPbReal = dynamic_cast<GORM_PB_Table_account*>(pMsg);
            return pPbReal->set_id(value);
        }
        }
    }
    case GORM_PB_TABLE_IDX_BAG:
    {
        switch (iFieldId)
        {
        case GORM_PB_FIELD_BAG_VERSION:
        {
            GORM_PB_Table_bag* pPbReal = dynamic_cast<GORM_PB_Table_bag*>(pMsg);
            return pPbReal->set_version(value);
        }
        case GORM_PB_FIELD_BAG_ID:
        {
            GORM_PB_Table_bag* pPbReal = dynamic_cast<GORM_PB_Table_bag*>(pMsg);
            return pPbReal->set_id(value);
        }
        }
    }
    }
}

void GORM_SetTableFieldValue(PB_MSG_PTR pMsg, int iTableId, int iFieldId, int16 value)
{
    switch (iTableId)
    {
    case GORM_PB_TABLE_IDX_ACCOUNT:
    {
        switch (iFieldId)
        {
        case GORM_PB_FIELD_ACCOUNT_VERSION:
        {
            GORM_PB_Table_account* pPbReal = dynamic_cast<GORM_PB_Table_account*>(pMsg);
            return pPbReal->set_version(value);
        }
        case GORM_PB_FIELD_ACCOUNT_ID:
        {
            GORM_PB_Table_account* pPbReal = dynamic_cast<GORM_PB_Table_account*>(pMsg);
            return pPbReal->set_id(value);
        }
        }
    }
    case GORM_PB_TABLE_IDX_BAG:
    {
        switch (iFieldId)
        {
        case GORM_PB_FIELD_BAG_VERSION:
        {
            GORM_PB_Table_bag* pPbReal = dynamic_cast<GORM_PB_Table_bag*>(pMsg);
            return pPbReal->set_version(value);
        }
        case GORM_PB_FIELD_BAG_ID:
        {
            GORM_PB_Table_bag* pPbReal = dynamic_cast<GORM_PB_Table_bag*>(pMsg);
            return pPbReal->set_id(value);
        }
        }
    }
    }
}

void GORM_SetTableFieldValue(PB_MSG_PTR pMsg, int iTableId, int iFieldId, int32 value)
{
    switch (iTableId)
    {
    case GORM_PB_TABLE_IDX_ACCOUNT:
    {
        switch (iFieldId)
        {
        case GORM_PB_FIELD_ACCOUNT_VERSION:
        {
            GORM_PB_Table_account* pPbReal = dynamic_cast<GORM_PB_Table_account*>(pMsg);
            return pPbReal->set_version(value);
        }
        case GORM_PB_FIELD_ACCOUNT_ID:
        {
            GORM_PB_Table_account* pPbReal = dynamic_cast<GORM_PB_Table_account*>(pMsg);
            return pPbReal->set_id(value);
        }
        }
    }
    case GORM_PB_TABLE_IDX_BAG:
    {
        switch (iFieldId)
        {
        case GORM_PB_FIELD_BAG_VERSION:
        {
            GORM_PB_Table_bag* pPbReal = dynamic_cast<GORM_PB_Table_bag*>(pMsg);
            return pPbReal->set_version(value);
        }
        case GORM_PB_FIELD_BAG_ID:
        {
            GORM_PB_Table_bag* pPbReal = dynamic_cast<GORM_PB_Table_bag*>(pMsg);
            return pPbReal->set_id(value);
        }
        }
    }
    }
}

void GORM_SetTableFieldValue(PB_MSG_PTR pMsg, int iTableId, int iFieldId, int64 value)
{
    switch (iTableId)
    {
    case GORM_PB_TABLE_IDX_ACCOUNT:
    {
        switch (iFieldId)
        {
        case GORM_PB_FIELD_ACCOUNT_VERSION:
        {
            GORM_PB_Table_account* pPbReal = dynamic_cast<GORM_PB_Table_account*>(pMsg);
            return pPbReal->set_version(value);
        }
        case GORM_PB_FIELD_ACCOUNT_ID:
        {
            GORM_PB_Table_account* pPbReal = dynamic_cast<GORM_PB_Table_account*>(pMsg);
            return pPbReal->set_id(value);
        }
        }
    }
    case GORM_PB_TABLE_IDX_BAG:
    {
        switch (iFieldId)
        {
        case GORM_PB_FIELD_BAG_VERSION:
        {
            GORM_PB_Table_bag* pPbReal = dynamic_cast<GORM_PB_Table_bag*>(pMsg);
            return pPbReal->set_version(value);
        }
        case GORM_PB_FIELD_BAG_ID:
        {
            GORM_PB_Table_bag* pPbReal = dynamic_cast<GORM_PB_Table_bag*>(pMsg);
            return pPbReal->set_id(value);
        }
        }
    }
    }
}

void GORM_SetTableFieldValue(PB_MSG_PTR pMsg, int iTableId, int iFieldId, double value)
{
    switch (iTableId)
    {
    case GORM_PB_TABLE_IDX_ACCOUNT:
    {
        switch (iFieldId)
        {
        case GORM_PB_FIELD_ACCOUNT_VERSION:
        {
            GORM_PB_Table_account* pPbReal = dynamic_cast<GORM_PB_Table_account*>(pMsg);
            return pPbReal->set_version(value);
        }
        case GORM_PB_FIELD_ACCOUNT_ID:
        {
            GORM_PB_Table_account* pPbReal = dynamic_cast<GORM_PB_Table_account*>(pMsg);
            return pPbReal->set_id(value);
        }
        }
    }
    case GORM_PB_TABLE_IDX_BAG:
    {
        switch (iFieldId)
        {
        case GORM_PB_FIELD_BAG_VERSION:
        {
            GORM_PB_Table_bag* pPbReal = dynamic_cast<GORM_PB_Table_bag*>(pMsg);
            return pPbReal->set_version(value);
        }
        case GORM_PB_FIELD_BAG_ID:
        {
            GORM_PB_Table_bag* pPbReal = dynamic_cast<GORM_PB_Table_bag*>(pMsg);
            return pPbReal->set_id(value);
        }
        }
    }
    }
}

void GORM_SetTableFieldValue(PB_MSG_PTR pMsg, int iTableId, int iFieldId, uint8 value)
{
    switch (iTableId)
    {
    case GORM_PB_TABLE_IDX_ACCOUNT:
    {
        switch (iFieldId)
        {
        case GORM_PB_FIELD_ACCOUNT_VERSION:
        {
            GORM_PB_Table_account* pPbReal = dynamic_cast<GORM_PB_Table_account*>(pMsg);
            return pPbReal->set_version(value);
        }
        case GORM_PB_FIELD_ACCOUNT_ID:
        {
            GORM_PB_Table_account* pPbReal = dynamic_cast<GORM_PB_Table_account*>(pMsg);
            return pPbReal->set_id(value);
        }
        }
    }
    case GORM_PB_TABLE_IDX_BAG:
    {
        switch (iFieldId)
        {
        case GORM_PB_FIELD_BAG_VERSION:
        {
            GORM_PB_Table_bag* pPbReal = dynamic_cast<GORM_PB_Table_bag*>(pMsg);
            return pPbReal->set_version(value);
        }
        case GORM_PB_FIELD_BAG_ID:
        {
            GORM_PB_Table_bag* pPbReal = dynamic_cast<GORM_PB_Table_bag*>(pMsg);
            return pPbReal->set_id(value);
        }
        }
    }
    }
}

void GORM_SetTableFieldValue(PB_MSG_PTR pMsg, int iTableId, int iFieldId, uint16 value)
{
    switch (iTableId)
    {
    case GORM_PB_TABLE_IDX_ACCOUNT:
    {
        switch (iFieldId)
        {
        case GORM_PB_FIELD_ACCOUNT_VERSION:
        {
            GORM_PB_Table_account* pPbReal = dynamic_cast<GORM_PB_Table_account*>(pMsg);
            return pPbReal->set_version(value);
        }
        case GORM_PB_FIELD_ACCOUNT_ID:
        {
            GORM_PB_Table_account* pPbReal = dynamic_cast<GORM_PB_Table_account*>(pMsg);
            return pPbReal->set_id(value);
        }
        }
    }
    case GORM_PB_TABLE_IDX_BAG:
    {
        switch (iFieldId)
        {
        case GORM_PB_FIELD_BAG_VERSION:
        {
            GORM_PB_Table_bag* pPbReal = dynamic_cast<GORM_PB_Table_bag*>(pMsg);
            return pPbReal->set_version(value);
        }
        case GORM_PB_FIELD_BAG_ID:
        {
            GORM_PB_Table_bag* pPbReal = dynamic_cast<GORM_PB_Table_bag*>(pMsg);
            return pPbReal->set_id(value);
        }
        }
    }
    }
}

void GORM_SetTableFieldValue(PB_MSG_PTR pMsg, int iTableId, int iFieldId, uint32 value)
{
    switch (iTableId)
    {
    case GORM_PB_TABLE_IDX_ACCOUNT:
    {
        switch (iFieldId)
        {
        case GORM_PB_FIELD_ACCOUNT_VERSION:
        {
            GORM_PB_Table_account* pPbReal = dynamic_cast<GORM_PB_Table_account*>(pMsg);
            return pPbReal->set_version(value);
        }
        case GORM_PB_FIELD_ACCOUNT_ID:
        {
            GORM_PB_Table_account* pPbReal = dynamic_cast<GORM_PB_Table_account*>(pMsg);
            return pPbReal->set_id(value);
        }
        }
    }
    case GORM_PB_TABLE_IDX_BAG:
    {
        switch (iFieldId)
        {
        case GORM_PB_FIELD_BAG_VERSION:
        {
            GORM_PB_Table_bag* pPbReal = dynamic_cast<GORM_PB_Table_bag*>(pMsg);
            return pPbReal->set_version(value);
        }
        case GORM_PB_FIELD_BAG_ID:
        {
            GORM_PB_Table_bag* pPbReal = dynamic_cast<GORM_PB_Table_bag*>(pMsg);
            return pPbReal->set_id(value);
        }
        }
    }
    }
}

void GORM_SetTableFieldValue(PB_MSG_PTR pMsg, int iTableId, int iFieldId, uint64 value)
{
    switch (iTableId)
    {
    case GORM_PB_TABLE_IDX_ACCOUNT:
    {
        switch (iFieldId)
        {
        case GORM_PB_FIELD_ACCOUNT_VERSION:
        {
            GORM_PB_Table_account* pPbReal = dynamic_cast<GORM_PB_Table_account*>(pMsg);
            return pPbReal->set_version(value);
        }
        case GORM_PB_FIELD_ACCOUNT_ID:
        {
            GORM_PB_Table_account* pPbReal = dynamic_cast<GORM_PB_Table_account*>(pMsg);
            return pPbReal->set_id(value);
        }
        }
    }
    case GORM_PB_TABLE_IDX_BAG:
    {
        switch (iFieldId)
        {
        case GORM_PB_FIELD_BAG_VERSION:
        {
            GORM_PB_Table_bag* pPbReal = dynamic_cast<GORM_PB_Table_bag*>(pMsg);
            return pPbReal->set_version(value);
        }
        case GORM_PB_FIELD_BAG_ID:
        {
            GORM_PB_Table_bag* pPbReal = dynamic_cast<GORM_PB_Table_bag*>(pMsg);
            return pPbReal->set_id(value);
        }
        }
    }
    }
}

int GORM_GetTableFieldValue(PB_MSG_PTR pMsg, int iTableId, int iFieldId, string &value)
{
    switch (iTableId)
    {
    case GORM_PB_TABLE_IDX_ACCOUNT:
    {
        switch (iFieldId)
        {
        case GORM_PB_FIELD_ACCOUNT_ACCOUNT:
        {
            GORM_PB_Table_account* pPbReal = dynamic_cast<GORM_PB_Table_account*>(pMsg);
            value = pPbReal->account();
            return GORM_OK;
        }
        case GORM_PB_FIELD_ACCOUNT_ALLBINARY:
        {
            GORM_PB_Table_account* pPbReal = dynamic_cast<GORM_PB_Table_account*>(pMsg);
            value = pPbReal->allbinary();
            return GORM_OK;
        }
        }
    }
    case GORM_PB_TABLE_IDX_BAG:
    {
        switch (iFieldId)
        {
        case GORM_PB_FIELD_BAG_ALLBINARY:
        {
            GORM_PB_Table_bag* pPbReal = dynamic_cast<GORM_PB_Table_bag*>(pMsg);
            value = pPbReal->allbinary();
            return GORM_OK;
        }
        }
    }
    }

    return GORM_ERROR;
}

int GORM_GetTableFieldValue(PB_MSG_PTR pMsg, int iTableId, int iFieldId, uint8 *&value, size_t &size)
{
    switch (iTableId)
    {
    case GORM_PB_TABLE_IDX_ACCOUNT:
    {
        switch (iFieldId)
        {
        case GORM_PB_FIELD_ACCOUNT_ACCOUNT:
        {
            GORM_PB_Table_account* pPbReal = dynamic_cast<GORM_PB_Table_account*>(pMsg);
            string strValue = pPbReal->account();
            value=(uint8*)strValue.c_str();
            size=strValue.size();
            return GORM_OK;
        }
        case GORM_PB_FIELD_ACCOUNT_ALLBINARY:
        {
            GORM_PB_Table_account* pPbReal = dynamic_cast<GORM_PB_Table_account*>(pMsg);
            string strValue = pPbReal->allbinary();
            value=(uint8*)strValue.c_str();
            size=strValue.size();
            return GORM_OK;
        }
        }
    }
    case GORM_PB_TABLE_IDX_BAG:
    {
        switch (iFieldId)
        {
        case GORM_PB_FIELD_BAG_ALLBINARY:
        {
            GORM_PB_Table_bag* pPbReal = dynamic_cast<GORM_PB_Table_bag*>(pMsg);
            string strValue = pPbReal->allbinary();
            value=(uint8*)strValue.c_str();
            size=strValue.size();
            return GORM_OK;
        }
        }
    }
    }

    return GORM_ERROR;
}

int GORM_GetTableFieldValue(PB_MSG_PTR pMsg, int iTableId, int iFieldId, int8 &value)
{
    switch (iTableId)
    {
    case GORM_PB_TABLE_IDX_ACCOUNT:
    {
        switch (iFieldId)
        {
        case GORM_PB_FIELD_ACCOUNT_VERSION:
        {
            GORM_PB_Table_account* pPbReal = dynamic_cast<GORM_PB_Table_account*>(pMsg);
            value = pPbReal->version();
            return GORM_OK;
        }
        case GORM_PB_FIELD_ACCOUNT_ID:
        {
            GORM_PB_Table_account* pPbReal = dynamic_cast<GORM_PB_Table_account*>(pMsg);
            value = pPbReal->id();
            return GORM_OK;
        }
        }
    }
    case GORM_PB_TABLE_IDX_BAG:
    {
        switch (iFieldId)
        {
        case GORM_PB_FIELD_BAG_VERSION:
        {
            GORM_PB_Table_bag* pPbReal = dynamic_cast<GORM_PB_Table_bag*>(pMsg);
            value = pPbReal->version();
            return GORM_OK;
        }
        case GORM_PB_FIELD_BAG_ID:
        {
            GORM_PB_Table_bag* pPbReal = dynamic_cast<GORM_PB_Table_bag*>(pMsg);
            value = pPbReal->id();
            return GORM_OK;
        }
        }
    }
    }

    return GORM_ERROR;
}

int GORM_GetTableFieldValue(PB_MSG_PTR pMsg, int iTableId, int iFieldId, int16 &value)
{
    switch (iTableId)
    {
    case GORM_PB_TABLE_IDX_ACCOUNT:
    {
        switch (iFieldId)
        {
        case GORM_PB_FIELD_ACCOUNT_VERSION:
        {
            GORM_PB_Table_account* pPbReal = dynamic_cast<GORM_PB_Table_account*>(pMsg);
            value = pPbReal->version();
            return GORM_OK;
        }
        case GORM_PB_FIELD_ACCOUNT_ID:
        {
            GORM_PB_Table_account* pPbReal = dynamic_cast<GORM_PB_Table_account*>(pMsg);
            value = pPbReal->id();
            return GORM_OK;
        }
        }
    }
    case GORM_PB_TABLE_IDX_BAG:
    {
        switch (iFieldId)
        {
        case GORM_PB_FIELD_BAG_VERSION:
        {
            GORM_PB_Table_bag* pPbReal = dynamic_cast<GORM_PB_Table_bag*>(pMsg);
            value = pPbReal->version();
            return GORM_OK;
        }
        case GORM_PB_FIELD_BAG_ID:
        {
            GORM_PB_Table_bag* pPbReal = dynamic_cast<GORM_PB_Table_bag*>(pMsg);
            value = pPbReal->id();
            return GORM_OK;
        }
        }
    }
    }

    return GORM_ERROR;
}

int GORM_GetTableFieldValue(PB_MSG_PTR pMsg, int iTableId, int iFieldId, int32 &value)
{
    switch (iTableId)
    {
    case GORM_PB_TABLE_IDX_ACCOUNT:
    {
        switch (iFieldId)
        {
        case GORM_PB_FIELD_ACCOUNT_VERSION:
        {
            GORM_PB_Table_account* pPbReal = dynamic_cast<GORM_PB_Table_account*>(pMsg);
            value = pPbReal->version();
            return GORM_OK;
        }
        case GORM_PB_FIELD_ACCOUNT_ID:
        {
            GORM_PB_Table_account* pPbReal = dynamic_cast<GORM_PB_Table_account*>(pMsg);
            value = pPbReal->id();
            return GORM_OK;
        }
        }
    }
    case GORM_PB_TABLE_IDX_BAG:
    {
        switch (iFieldId)
        {
        case GORM_PB_FIELD_BAG_VERSION:
        {
            GORM_PB_Table_bag* pPbReal = dynamic_cast<GORM_PB_Table_bag*>(pMsg);
            value = pPbReal->version();
            return GORM_OK;
        }
        case GORM_PB_FIELD_BAG_ID:
        {
            GORM_PB_Table_bag* pPbReal = dynamic_cast<GORM_PB_Table_bag*>(pMsg);
            value = pPbReal->id();
            return GORM_OK;
        }
        }
    }
    }

    return GORM_ERROR;
}

int GORM_GetTableFieldValue(PB_MSG_PTR pMsg, int iTableId, int iFieldId, int64 &value)
{
    switch (iTableId)
    {
    case GORM_PB_TABLE_IDX_ACCOUNT:
    {
        switch (iFieldId)
        {
        case GORM_PB_FIELD_ACCOUNT_VERSION:
        {
            GORM_PB_Table_account* pPbReal = dynamic_cast<GORM_PB_Table_account*>(pMsg);
            value = pPbReal->version();
            return GORM_OK;
        }
        case GORM_PB_FIELD_ACCOUNT_ID:
        {
            GORM_PB_Table_account* pPbReal = dynamic_cast<GORM_PB_Table_account*>(pMsg);
            value = pPbReal->id();
            return GORM_OK;
        }
        }
    }
    case GORM_PB_TABLE_IDX_BAG:
    {
        switch (iFieldId)
        {
        case GORM_PB_FIELD_BAG_VERSION:
        {
            GORM_PB_Table_bag* pPbReal = dynamic_cast<GORM_PB_Table_bag*>(pMsg);
            value = pPbReal->version();
            return GORM_OK;
        }
        case GORM_PB_FIELD_BAG_ID:
        {
            GORM_PB_Table_bag* pPbReal = dynamic_cast<GORM_PB_Table_bag*>(pMsg);
            value = pPbReal->id();
            return GORM_OK;
        }
        }
    }
    }

    return GORM_ERROR;
}

int GORM_GetTableFieldValue(PB_MSG_PTR pMsg, int iTableId, int iFieldId, double &value)
{
    switch (iTableId)
    {
    case GORM_PB_TABLE_IDX_ACCOUNT:
    {
        switch (iFieldId)
        {
        case GORM_PB_FIELD_ACCOUNT_VERSION:
        {
            GORM_PB_Table_account* pPbReal = dynamic_cast<GORM_PB_Table_account*>(pMsg);
            value = pPbReal->version();
            return GORM_OK;
        }
        case GORM_PB_FIELD_ACCOUNT_ID:
        {
            GORM_PB_Table_account* pPbReal = dynamic_cast<GORM_PB_Table_account*>(pMsg);
            value = pPbReal->id();
            return GORM_OK;
        }
        }
    }
    case GORM_PB_TABLE_IDX_BAG:
    {
        switch (iFieldId)
        {
        case GORM_PB_FIELD_BAG_VERSION:
        {
            GORM_PB_Table_bag* pPbReal = dynamic_cast<GORM_PB_Table_bag*>(pMsg);
            value = pPbReal->version();
            return GORM_OK;
        }
        case GORM_PB_FIELD_BAG_ID:
        {
            GORM_PB_Table_bag* pPbReal = dynamic_cast<GORM_PB_Table_bag*>(pMsg);
            value = pPbReal->id();
            return GORM_OK;
        }
        }
    }
    }

    return GORM_ERROR;
}

int GORM_GetTableFieldValue(PB_MSG_PTR pMsg, int iTableId, int iFieldId, uint8 &value)
{
    switch (iTableId)
    {
    case GORM_PB_TABLE_IDX_ACCOUNT:
    {
        switch (iFieldId)
        {
        case GORM_PB_FIELD_ACCOUNT_VERSION:
        {
            GORM_PB_Table_account* pPbReal = dynamic_cast<GORM_PB_Table_account*>(pMsg);
            value = pPbReal->version();
            return GORM_OK;
        }
        case GORM_PB_FIELD_ACCOUNT_ID:
        {
            GORM_PB_Table_account* pPbReal = dynamic_cast<GORM_PB_Table_account*>(pMsg);
            value = pPbReal->id();
            return GORM_OK;
        }
        }
    }
    case GORM_PB_TABLE_IDX_BAG:
    {
        switch (iFieldId)
        {
        case GORM_PB_FIELD_BAG_VERSION:
        {
            GORM_PB_Table_bag* pPbReal = dynamic_cast<GORM_PB_Table_bag*>(pMsg);
            value = pPbReal->version();
            return GORM_OK;
        }
        case GORM_PB_FIELD_BAG_ID:
        {
            GORM_PB_Table_bag* pPbReal = dynamic_cast<GORM_PB_Table_bag*>(pMsg);
            value = pPbReal->id();
            return GORM_OK;
        }
        }
    }
    }

    return GORM_ERROR;
}

int GORM_GetTableFieldValue(PB_MSG_PTR pMsg, int iTableId, int iFieldId, uint16 &value)
{
    switch (iTableId)
    {
    case GORM_PB_TABLE_IDX_ACCOUNT:
    {
        switch (iFieldId)
        {
        case GORM_PB_FIELD_ACCOUNT_VERSION:
        {
            GORM_PB_Table_account* pPbReal = dynamic_cast<GORM_PB_Table_account*>(pMsg);
            value = pPbReal->version();
            return GORM_OK;
        }
        case GORM_PB_FIELD_ACCOUNT_ID:
        {
            GORM_PB_Table_account* pPbReal = dynamic_cast<GORM_PB_Table_account*>(pMsg);
            value = pPbReal->id();
            return GORM_OK;
        }
        }
    }
    case GORM_PB_TABLE_IDX_BAG:
    {
        switch (iFieldId)
        {
        case GORM_PB_FIELD_BAG_VERSION:
        {
            GORM_PB_Table_bag* pPbReal = dynamic_cast<GORM_PB_Table_bag*>(pMsg);
            value = pPbReal->version();
            return GORM_OK;
        }
        case GORM_PB_FIELD_BAG_ID:
        {
            GORM_PB_Table_bag* pPbReal = dynamic_cast<GORM_PB_Table_bag*>(pMsg);
            value = pPbReal->id();
            return GORM_OK;
        }
        }
    }
    }

    return GORM_ERROR;
}

int GORM_GetTableFieldValue(PB_MSG_PTR pMsg, int iTableId, int iFieldId, uint32 &value)
{
    switch (iTableId)
    {
    case GORM_PB_TABLE_IDX_ACCOUNT:
    {
        switch (iFieldId)
        {
        case GORM_PB_FIELD_ACCOUNT_VERSION:
        {
            GORM_PB_Table_account* pPbReal = dynamic_cast<GORM_PB_Table_account*>(pMsg);
            value = pPbReal->version();
            return GORM_OK;
        }
        case GORM_PB_FIELD_ACCOUNT_ID:
        {
            GORM_PB_Table_account* pPbReal = dynamic_cast<GORM_PB_Table_account*>(pMsg);
            value = pPbReal->id();
            return GORM_OK;
        }
        }
    }
    case GORM_PB_TABLE_IDX_BAG:
    {
        switch (iFieldId)
        {
        case GORM_PB_FIELD_BAG_VERSION:
        {
            GORM_PB_Table_bag* pPbReal = dynamic_cast<GORM_PB_Table_bag*>(pMsg);
            value = pPbReal->version();
            return GORM_OK;
        }
        case GORM_PB_FIELD_BAG_ID:
        {
            GORM_PB_Table_bag* pPbReal = dynamic_cast<GORM_PB_Table_bag*>(pMsg);
            value = pPbReal->id();
            return GORM_OK;
        }
        }
    }
    }

    return GORM_ERROR;
}

int GORM_GetTableFieldValue(PB_MSG_PTR pMsg, int iTableId, int iFieldId, uint64 &value)
{
    switch (iTableId)
    {
    case GORM_PB_TABLE_IDX_ACCOUNT:
    {
        switch (iFieldId)
        {
        case GORM_PB_FIELD_ACCOUNT_VERSION:
        {
            GORM_PB_Table_account* pPbReal = dynamic_cast<GORM_PB_Table_account*>(pMsg);
            value = pPbReal->version();
            return GORM_OK;
        }
        case GORM_PB_FIELD_ACCOUNT_ID:
        {
            GORM_PB_Table_account* pPbReal = dynamic_cast<GORM_PB_Table_account*>(pMsg);
            value = pPbReal->id();
            return GORM_OK;
        }
        }
    }
    case GORM_PB_TABLE_IDX_BAG:
    {
        switch (iFieldId)
        {
        case GORM_PB_FIELD_BAG_VERSION:
        {
            GORM_PB_Table_bag* pPbReal = dynamic_cast<GORM_PB_Table_bag*>(pMsg);
            value = pPbReal->version();
            return GORM_OK;
        }
        case GORM_PB_FIELD_BAG_ID:
        {
            GORM_PB_Table_bag* pPbReal = dynamic_cast<GORM_PB_Table_bag*>(pMsg);
            value = pPbReal->id();
            return GORM_OK;
        }
        }
    }
    }

    return GORM_ERROR;
}


}