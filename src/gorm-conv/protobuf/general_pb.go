package protobuf

import (
	"fmt"
	"gorm-conv/common"
	"os"
	"strconv"
	"strings"
)

func GetPBString(c common.TableColumn) string {
	switch c.Type {
	case "double":
		{
			return "double " + c.Name
		}
	case "int8", "int", "int32":
		{
			return "sfixed32 " + c.Name
		}
	case "uint8", "uint32":
		{
			return "fixed32" + c.Name
		}
	case "long", "int64":
		{
			return "sfixed64 " + c.Name
		}
	case "uint64":
		{
			return "fixed64 " + c.Name
		}
	case "string", "char":
		{
			return "string " + c.Name
		}
	case "blob", "bytes":
		{
			return "bytes " + c.Name
		}
	}
	return ""
}

func OutPutPBTable(table common.TableInfo, f *os.File) int {
	/*version := "    fixed64	version = 1;\n"
	_, err := f.WriteString(version)
	if err != nil {
		fmt.Println(err.Error())
		return -1
	}*/
	var index int64 = 0
	for _, c := range table.TableColumns {
		index += 1
		pbStr := GetPBString(c)
		if pbStr == "" {
			fmt.Println("general pb file failed, table:", table.Name, ", column:", c.Name)
			return -1
		}
		str := "    " + common.Proto_optional + pbStr + " = "
		str += strconv.FormatInt(index, 10)
		str += ";\n"
		_, err := f.WriteString(str)
		if err != nil {
			fmt.Println(err.Error())
			return -1
		}
	}
	return 0
}

func GeneralPBFile(game common.XmlCfg, outpath string) int {
	fileLen := len(game.File)
	outfile := outpath + "/" + game.File[:fileLen-4] + ".proto"
	f, err := os.OpenFile(outfile, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err.Error())
		return -1
	}
	defer f.Close()
	fmt.Println("begin to generate pb file:%s", outfile)
	err = f.Truncate(0)
	if err != nil {
		fmt.Println(err.Error())
	}
	// 1、输出开头
	f.WriteString("syntax = \"proto" + *common.Protoversion + "\";\n")
	f.WriteString("package gorm;\n")
	f.WriteString("option go_package = \"" + *common.Gopackage + "\";\n")
	if err != nil {
		fmt.Println(err.Error())
		return -1
	}

	// 输出表结构
	/*
		   message GORM_PB_Table_account
		   {
				fixed64	version = 1;
				int id = 2;
				string account = 3;
				string allbinary = 4;
		   }
	*/

	for _, table := range game.DB.TableList {
		msgHeader := "message GORM_PB_Table_" + table.Name + "{\n"
		_, err = f.WriteString(msgHeader)
		if err != nil {
			fmt.Println(err.Error())
			return -1
		}
		if 0 != OutPutPBTable(table, f) {
			fmt.Println("genreral message for table failed:" + table.Name)
			return -1
		}
		_, err = f.WriteString("}\n\n")
		if err != nil {
			fmt.Println(err.Error())
			return -1
		}
	}

	return 0
}

func GeneralPBColumnIndex(games []common.XmlCfg, f *os.File) int {
	for _, game := range games {
		for _, table := range game.DB.TableList {
			var colIndex int64 = 0
			var enum string = "enum GORM_PB_"
			enum += strings.ToUpper(table.Name)
			enum += "_FIELD_INDEX\n{\n"
			_, err := f.WriteString(enum)
			if err != nil {
				fmt.Println(err.Error())
				return -1
			}

			outPre := "	GORM_PB_FIELD_" + strings.ToUpper(table.Name) + "_"
			out := ""
			_, err = f.WriteString(out)
			if err != nil {
				fmt.Println(err.Error())
				return -1
			}
			for _, col := range table.TableColumns {
				strColIndex := strconv.FormatInt(int64(colIndex), 10)
				colIndex += 1
				out = outPre + strings.ToUpper(col.Name)
				out += " = "
				out += strColIndex
				out += ";\n"
				_, err := f.WriteString(out)
				if err != nil {
					fmt.Println(err.Error())
					return -1
				}
			}
			out = "}\n\n"
			_, err = f.WriteString(out)
			if err != nil {
				fmt.Println(err.Error())
				return -1
			}
		}
	}
	return 0
}

//
func GeneralTablesInc(games []common.XmlCfg, outpath string) int {
	outfile := outpath + "gorm_pb_tables_inc.proto"
	f, err := os.OpenFile(outfile, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err.Error())
		return -1
	}
	defer f.Close()
	fmt.Println("begin to general " + outfile)
	err = f.Truncate(0)
	if err != nil {
		fmt.Println(err.Error())
		return -1
	}
	f.WriteString("syntax = \"proto" + *common.Protoversion + "\";\n")
	f.WriteString("package gorm;\n")
	f.WriteString("option go_package = \"" + *common.Gopackage + "\";\n")
	if err != nil {
		fmt.Println(err.Error())
		return -1
	}
	for _, game := range games {
		_, err = f.WriteString(string("import \""))
		if err != nil {
			fmt.Println(err.Error())
			return -1
		}
		fileLen := len(game.File)
		proto := game.File[:fileLen-4] + ".proto\";\n"
		_, err = f.WriteString(proto)
		if err != nil {
			fmt.Println(err.Error())
			return -1
		}
	}
	f.WriteString("\n\n")

	//输出
	/*
		enum GORM_PB_TABLE_INDEX
		{
			account = 1;
			bag = 2;
		}
	*/
	var enum string = `enum GORM_PB_TABLE_INDEX
{
    GORM_PB_TABLE_IDX_MIN__ = 0;
`
	_, err = f.WriteString(enum)
	if err != nil {
		fmt.Println(err.Error())
		return -1
	}

	var tableIndx int64 = 0
	for _, game := range games {
		for _, table := range game.DB.TableList {
			//GORM_PB_Table_account account = 1;
			tableIndx += 1
			out := "	GORM_PB_TABLE_IDX_" + strings.ToUpper(table.Name) + " = "
			out += strconv.FormatInt(int64(tableIndx), 10)
			_, err = f.WriteString(out + ";\n")
			if err != nil {
				fmt.Println(err.Error())
				return -1
			}
		}
	}
	tableIndx += 1
	out := "	GORM_PB_TABLE_IDX_MAX__ = "
	out += strconv.FormatInt(int64(tableIndx), 10)
	_, err = f.WriteString(out + ";\n")
	if err != nil {
		fmt.Println(err.Error())
		return -1
	}
	_, err = f.WriteString("}\n\n")
	if err != nil {
		fmt.Println(err.Error())
		return -1
	}

	// 输出所有列的宏
	if 0 != GeneralPBColumnIndex(games, f) {
		fmt.Println("general column index failed.")
		return -1
	}

	// 输出列的类型
	var columns string = `
enum GORM_PB_COLUMN_TYPE
{
	GORM_PB_COLUMN_TYPE_INVALID = 0;
	GORM_PB_COLUMN_TYPE_INT		= 1;
	GORM_PB_COLUMN_TYPE_UINT 	= 2;
	GORM_PB_COLUMN_TYPE_DOUBLE	= 3;
	GORM_PB_COLUMN_TYPE_STRING	= 4;
	GORM_PB_COLUMN_TYPE_INT8	= 5;
	GORM_PB_COLUMN_TYPE_INT16	= 6;
	GORM_PB_COLUMN_TYPE_INT32	= 7;
	GORM_PB_COLUMN_TYPE_INT64	= 8;
	GORM_PB_COLUMN_TYPE_UINT8	= 9;
	GORM_PB_COLUMN_TYPE_UINT16	= 10;
	GORM_PB_COLUMN_TYPE_UINT32	= 11;
	GORM_PB_COLUMN_TYPE_UINT64	= 12;
	GORM_PB_COLUMN_TYPE_BLOB	= 13;
	GORM_PB_COLUMN_TYPE_CHAR	= 14;
}

message GORM_PB_COLUMN_VALUE
{
	optional GORM_PB_COLUMN_TYPE	type 		= 1;
	optional fixed64				uintvalue 	= 2;
	optional sfixed64				intvalue 	= 3;
	optional double					doublevalue = 4;
	optional string					stringvalue	= 5;
}

message GORM_PB_COLUMN
{
	optional string  				name 	= 1;				// 列名
	optional GORM_PB_COLUMN_VALUE 	value 	= 2;	// 存储的数据
}

`
	if *common.Protoversion == "3" {
		columns = strings.Replace(columns, "optional ", "", -1)
	}
	f.WriteString(columns)

	// 输出所有表的集合
	/*
		message GORM_PB_TABLE
		{
			GORM_PB_Table_account account = 1;
			GORM_PB_Table_bag bag = 2;
		}
	*/
	var custom_columns string = `
message GORM_PB_CUSTEM_COLUMNS
{
	repeated	GORM_PB_COLUMN	columns 	= 1;
}
`
	f.WriteString(custom_columns)
	var table string = `message GORM_PB_TABLE
{
	optional sfixed32 TableId = 1;
	optional GORM_PB_CUSTEM_COLUMNS	custom_columns = 2;
`
	if *common.Protoversion == "3" {
		table = strings.Replace(table, "optional ", "", -1)
	}
	_, err = f.WriteString(table)
	if err != nil {
		fmt.Println(err.Error())
		return -1
	}

	tableIndx = 2
	for _, game := range games {
		for _, table := range game.DB.TableList {
			tableIndx += 1
			out := "	" + common.Proto_optional + " GORM_PB_Table_" + table.Name + " " + table.Name + " = "
			out += strconv.FormatInt(int64(tableIndx), 10)
			_, err = f.WriteString(out + ";\n")
			if err != nil {
				fmt.Println(err.Error())
				return -1
			}
		}
	}
	_, err = f.WriteString("}\n\n")
	if err != nil {
		fmt.Println(err.Error())
		return -1
	}

	// 输出
	/*
		message GORM_PB_TABLES
		{
			repeated GORM_PB_Table_account account = 1;
			repeated GORM_PB_Table_bag bag = 2;
		}
	*/
	tableIndx = 0
	table = `message GORM_PB_TABLES
{
`
	_, err = f.WriteString(table)
	if err != nil {
		fmt.Println(err.Error())
		return -1
	}
	for _, game := range games {
		for _, table := range game.DB.TableList {
			tableIndx += 1
			out := "	repeated GORM_PB_Table_" + table.Name + " " + table.Name + " = "
			out += strconv.FormatInt(int64(tableIndx), 10)
			_, err = f.WriteString(out + ";\n")
			if err != nil {
				fmt.Println(err.Error())
				return -1
			}
		}
	}
	_, err = f.WriteString("}\n\n")
	if err != nil {
		fmt.Println(err.Error())
		return -1
	}

	return 0
}

func GeneralPBFiles(games []common.XmlCfg, outpath string) int {
	if 0 != GeneralTablesInc(games, outpath) {
		fmt.Println("general tables inc proto file failed.")
		return -1
	}
	for _, game := range games {
		if 0 != GeneralPBFile(game, outpath) {
			return -1
		}
	}

	return GeneragePbProtoFile(outpath)
}

func GeneragePbProtoFile(outpath string) int {
	outfile := outpath + "gorm_pb_proto.proto"
	f, err := os.OpenFile(outfile, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err.Error())
		return -1
	}
	defer f.Close()

	f.Truncate(0)
	f.WriteString("syntax = \"proto" + *common.Protoversion + "\";\n")
	f.WriteString(`package gorm;

option go_package = "`)
	f.WriteString(*common.Gopackage)
	f.WriteString("\";\n")
	var proto_string string = `

// 此协议文件不需要变更，里面保存了GORM客户端与服务器通信的所需要的协议
//gorm_tables_inc.proto包含所有表的定义
import "gorm_pb_tables_inc.proto";
`
	f.WriteString(proto_string)
	if common.Gooutpath != nil && *common.Gooutpath != "" {
		proto_string = `

enum GORM_CODE
{
	OK = 0;
	
	ERROR              = -1;
	EAGAIN             = -2;
	INVALID_CLIENT     = -3;
	PART_FAILED        = -4; // 请求部分失败
	INVALID_TABLE      = -5; // 无效的表
	RESET              = -6; // 重复设置
	TOO_MUCH_RECORD    = -7; // 太多record
	INIT_RECORD        = -8;
	INVALID_FIELD      = -9;  // 无效的field
	PACK_REQ_ERROR     = -10; // 打包请求失败
	REQ_NO_RECORDS     = -11; // 没有往请求中放入record
	MULTI_TABLES       = -12; // 有多张表
	RSP_UNPACK_FAILED  = -13; // 响应解包失败
	CONN_CLOSED        = -14; // 连接已经关闭了
	CONN_FAILED        = -15; // 连接RetErr服务器失败
	DB_ERROR           = -16; // 数据库发生错误,此时需要根据db信息获取进一步错误信息
	NO_DB              = -17; // 没有找到db
	REQ_MSG_ERROR      = -18; // 请求信息错误
	NOT_SUPPORT_CMD    = -19; // 不支持的命令
	UNPACK_REQ         = -20; // 解压缩请求信息出错
	PACK_RSP_ERROR     = -21; // 压缩响应消息出错
	REQ_MSG_NO_HEADER  = -22; // 请求没有设置消息头
	REQ_NEED_SPLIT     = -23; // split信息没有带全
	REQ_TOO_LARGE      = -24; // 请求数据太大
	DB_2_STRUCT_ERROR  = -25; // db结果转换到struct出错，一般都是版本对不上导致
	NO_MORE_RECORD     = -26; // 没有更多record
	VERSION_NOT_SET    = -27; // 没有设置版本号
	CACHE_ERROR        = -28; // 操作缓存错误
	NO_VALUE           = -29; // 没有对应的值
	INVALID_VALUE_TYPE = -30; // 无效的类型
	NEED_HAND_SHAKE    = -31; // 客户端没有握手直接发送消息
}
`
	} else {
		proto_string = `

enum GORM_CODE
{
	GORM_CODE_OK = 0;
	
	GORM_CODE_ERROR              = -1;
	GORM_CODE_EAGAIN             = -2;
	GORM_CODE_INVALID_CLIENT     = -3;
	GORM_CODE_PART_FAILED        = -4; // 请求部分失败
	GORM_CODE_INVALID_TABLE      = -5; // 无效的表
	GORM_CODE_RESET              = -6; // 重复设置
	GORM_CODE_TOO_MUCH_RECORD    = -7; // 太多record
	GORM_CODE_INIT_RECORD        = -8;
	GORM_CODE_INVALID_FIELD      = -9;  // 无效的field
	GORM_CODE_PACK_REQ_ERROR     = -10; // 打包请求失败
	GORM_CODE_REQ_NO_RECORDS     = -11; // 没有往请求中放入record
	GORM_CODE_MULTI_TABLES       = -12; // 有多张表
	GORM_CODE_RSP_UNPACK_FAILED  = -13; // 响应解包失败
	GORM_CODE_CONN_CLOSED        = -14; // 连接已经关闭了
	GORM_CODE_CONN_FAILED        = -15; // 连接RetErr服务器失败
	GORM_CODE_DB_ERROR           = -16; // 数据库发生错误,此时需要根据db信息获取进一步错误信息
	GORM_CODE_NO_DB              = -17; // 没有找到db
	GORM_CODE_REQ_MSG_ERROR      = -18; // 请求信息错误
	GORM_CODE_NOT_SUPPORT_CMD    = -19; // 不支持的命令
	GORM_CODE_UNPACK_REQ         = -20; // 解压缩请求信息出错
	GORM_CODE_PACK_RSP_ERROR     = -21; // 压缩响应消息出错
	GORM_CODE_REQ_MSG_NO_HEADER  = -22; // 请求没有设置消息头
	GORM_CODE_REQ_NEED_SPLIT     = -23; // split信息没有带全
	GORM_CODE_REQ_TOO_LARGE      = -24; // 请求数据太大
	GORM_CODE_DB_2_STRUCT_ERROR  = -25; // db结果转换到struct出错，一般都是版本对不上导致
	GORM_CODE_NO_MORE_RECORD     = -26; // 没有更多record
	GORM_CODE_VERSION_NOT_SET    = -27; // 没有设置版本号
	GORM_CODE_CACHE_ERROR        = -28; // 操作缓存错误
	GORM_CODE_NO_VALUE           = -29; // 没有对应的值
	GORM_CODE_INVALID_VALUE_TYPE = -30; // 无效的类型
	GORM_CODE_NEED_HAND_SHAKE    = -31; // 客户端没有握手直接发送消息
}
`
	}

	f.WriteString(proto_string)
	proto_string = `
enum GORM_CMD
{
	GORM_CMD_INVALID 			= 0;
	GORM_CMD_HEART              = 1;    // 心跳，内部使用的协议
    GORM_CMD_HAND_SHAKE         = 2;    // 握手，获取客户端id过程
    GORM_CMD_INSERT             = 3;    // 增加记录
    GORM_CMD_REPLACE            = 4;    // 有则替换，没有则插入
    GORM_CMD_INCREASE           = 5;    // 增量更新请求
    GORM_CMD_GET                = 6;    // 单条查询请求
    GORM_CMD_DELETE             = 7;    // 删除请求
    GORM_CMD_BATCH_GET          = 8;    // 批量查询请求
    GORM_CMD_GET_BY_PARTKEY     = 9;    // 部分key查询请求
    GORM_CMD_UPDATE             = 10;    // 更新请求
    GORM_CMD_GET_BY_NON_PRIMARY_KEY = 11;   // 通过非主键批量获取 
}

message GORM_PB_SPLIT_INFO
{
	message GORM_COLUMN_VALUE
	{
		optional sfixed32 	Column = 1;
		optional bytes 	Value = 2;
	}
	repeated GORM_COLUMN_VALUE SplitInfo = 1;
}

message GORM_PB_Ret_Code 
{
	optional sfixed32 	Code = 1;			// 0 为成功，-1为失败，-4为部分失败
	optional sfixed32 	DBCode = 2;			// 后端db的错误码
	optional string 	DBErrInfo = 3;		// 后端db错误的详细信息
	optional sfixed32	SucessNum	 = 4;	// 成功的请求个数
	optional sfixed32 	TotalNum	 = 5;	// 所有的请求
}

message GORM_PB_RELOAD_TABLE_REQ
{
	optional GORM_PB_REQ_HEADER 		Header = 1; 
	optional fixed64 					TableVersion = 2;
}

message GORM_PB_RELOAD_TABLE_RSP
{
	optional GORM_PB_Ret_Code 		RetCode = 1;
}

message GORM_PB_REQ_HEADER 
{
	optional sfixed32   			TableId	 = 1;	// 表的类型
	optional sfixed32    			BusinessID = 2;	// 串行化ID
	optional sfixed32 				VerPolice = 3;	// 版本校验规则
	optional fixed32 				ReqFlag  = 4;	// 参见GORM_CMD_FLAG_XXX
	optional bytes					FieldMode	= 5;
	optional GORM_PB_SPLIT_INFO 	SplitTableInfo = 6;	// 分库分表信息
}

message GORM_PB_HEART_REQ
{
	optional GORM_PB_REQ_HEADER 		Header = 1; 
}

message GORM_PB_HEART_RSP
{
	optional GORM_PB_Ret_Code 		RetCode = 1;
} 

message GORM_PB_TABLE_SCHEMA_INFO_COLUMN
{
	fixed64	Version = 1;
	string 	Name = 2;
	string	TypeDesc = 3;
	GORM_PB_COLUMN_TYPE Type = 4;
}

message GORM_PB_TABLE_SCHEMA_INFO
{
	fixed64		Version = 1;
	string 		TableName = 2;
	sfixed32  	TableIdx = 3;
	repeated 	GORM_PB_TABLE_SCHEMA_INFO_COLUMN Columns = 4;
}

message GORM_PB_HAND_SHAKE_REQ
{
	GORM_PB_REQ_HEADER 		Header = 1;
	fixed64		Version = 2;
	fixed32		Md5     = 3;
	repeated 	GORM_PB_TABLE_SCHEMA_INFO Schemas = 4;
}

message GORM_PB_HAND_SHAKE_RSP
{
	optional GORM_PB_Ret_Code 		RetCode = 1;
	fixed64	 ClientId	= 2;	// 客户端ID
}

// 插入暂时只能支持单挑记录插入
message GORM_PB_INSERT_REQ
{
	optional GORM_PB_REQ_HEADER 	Header = 1;
	repeated GORM_PB_TABLE 			Tables = 2;
}

message GORM_PB_INSERT_RSP
{
	optional GORM_PB_Ret_Code 		RetCode = 1;
	repeated GORM_PB_TABLE 			Tables = 2;	// 如果需要回带数据，则将数据带回
	optional fixed64				LastInsertId = 3;
}

message GORM_PB_UPDATE_REQ
{
	optional GORM_PB_REQ_HEADER 	Header 	= 1;
	repeated GORM_PB_TABLE 			Tables 	= 2;
}

message GORM_PB_UPDATE_RSP
{
	optional GORM_PB_Ret_Code 		RetCode = 1;
	repeated GORM_PB_TABLE 			Tables 	= 2;	// 如果需要回带数据，则将数据带回
	optional sfixed32				AffectedNum = 3;
}

message GORM_PB_REPLACE_REQ
{
	optional GORM_PB_REQ_HEADER 	Header 	= 1;
	repeated GORM_PB_TABLE 			Tables 	= 2;
}

message GORM_PB_REPLACE_RSP
{
	optional GORM_PB_Ret_Code 		RetCode = 1;
	repeated GORM_PB_TABLE 			Tables 	= 2;	// 如果需要回带数据，则将数据带回
	optional sfixed32				AffectedNum = 3;
}

message GORM_PB_GET_REQ 
{
	optional GORM_PB_REQ_HEADER 	Header 	= 1;
	optional GORM_PB_TABLE			Table	= 2;
	optional sfixed32				GetFlag = 3;
}

message GORM_PB_GET_RSP
{
	optional GORM_PB_Ret_Code 		RetCode = 1;
	optional GORM_PB_TABLE 			Table 	= 2;
	optional sfixed32				NewInsert = 3;	// 结果是否是新插入的
	optional fixed64				LastInsertId = 4;	// 结果是否是新插入的
}

message GORM_PB_BATCH_GET_REQ 
{
	optional GORM_PB_REQ_HEADER 		Header 	= 1;
	repeated GORM_PB_TABLE 				Tables 	= 2;	// 如果需要回带数据，则将数据带回
}

message GORM_PB_BATCH_GET_RSP
{
	optional GORM_PB_Ret_Code 		RetCode = 1;
	repeated GORM_PB_TABLE 			Tables 	= 2;
}

message GORM_PB_INCREASE_REQ
{
	optional GORM_PB_REQ_HEADER 			Header 	 = 1;
	repeated GORM_PB_TABLE 					Tables 	= 2;
	optional string 						PlusColumns = 3;	// 增加的字段
	optional string 						MinusColumns = 4;	// 减少的字段
}

message GORM_PB_INCREASE_RSP
{
	optional GORM_PB_Ret_Code 			RetCode = 1;
	repeated GORM_PB_TABLE 				Tables 	= 2;
	optional sfixed32					AffectedNum = 3;
}

message GORM_PB_DELETE_REQ
{
	optional GORM_PB_REQ_HEADER 		Header = 1;
	optional GORM_PB_TABLE				Table	= 2;
}

message GORM_PB_DELETE_RSP
{
	optional GORM_PB_Ret_Code 			RetCode = 1;
	optional sfixed32					AffectedNum = 2;
}

message GORM_PB_GET_BY_PARTKEY_REQ
{
	optional GORM_PB_REQ_HEADER 		Header = 1;
	repeated GORM_PB_TABLE 				Tables 	= 2;	// 如果需要回带数据，则将数据带回
}

message GORM_PB_GET_BY_PARTKEY_RSP
{
	optional GORM_PB_Ret_Code 		RetCode = 1;
	repeated GORM_PB_TABLE 			Tables 	= 2;
}

message GORM_PB_GET_BY_NON_PRIMARY_KEY_REQ
{
	optional GORM_PB_REQ_HEADER 		Header = 1;
	repeated GORM_PB_TABLE 				Tables 	= 2;	// 如果需要回带数据，则将数据带回
}

message GORM_PB_GET_BY_NON_PRIMARY_KEY_RSP
{
	optional GORM_PB_Ret_Code 		RetCode = 1;
	repeated GORM_PB_TABLE 			Tables 	= 2;
}
`
	if *common.Protoversion == "3" {
		proto_string = strings.Replace(proto_string, "optional ", "", -1)
	}
	f.WriteString(proto_string)

	return 0
}
