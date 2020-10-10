package common

func CPP_TableColumnName(col string) string {
	return CPP_TableStruct(col)
}

// 第一个字符大写
func CPP_TableStruct(table string) string {
	vv := []rune(table)
	if vv[0] >= 'a' && vv[0] <= 'z' {
		vv[0] -= 32
	}

	return string(vv)
}

func CPPFieldInitTableColumnInfo_ForTable_COLTYPE(colType string) string {
	switch colType {
	case "int8", "int16", "int32", "int":
		{
			return "GORM_PB_COLUMN_TYPE_INT"
		}
	case "uint8", "uint16", "uint32":
		{
			return "GORM_PB_COLUMN_TYPE_UINT"
		}
	case "int64":
		{
			return "GORM_PB_COLUMN_TYPE_INT"
		}
	case "uint64":
		{
			return "GORM_PB_COLUMN_TYPE_UINT"
		}
	case "double":
		{
			return "GORM_PB_COLUMN_TYPE_DOUBLE"
		}
	}

	return "GORM_PB_COLUMN_TYPE_STRING"
}

func CPPFieldsMapPackSQL_ForTables_COL2SQL_GET_LONG_ULONG(colType string) string {
	switch colType {
	case "int8", "int16", "int32", "int":
		{
			return "long"
		}
	case "uint8", "uint16", "uint32":
		{
			return "ulong"
		}
	case "int64":
		{
			return "long"
		}
	case "uint64":
		{
			return "ulong"
		}
	case "double":
		{
			return "long"
		}
	}

	return "ulong"
}

func CPPFieldsMapPackSQL_ForTables_COL2SQL_GetCPPType(colType string) string {
	switch colType {
	case "int8", "int16", "int32", "int":
		{
			return "int32"
		}
	case "uint8", "uint16", "uint32":
		{
			return "uint32"
		}
	case "int64":
		{
			return "int64"
		}
	case "uint64":
		{
			return "uint64"
		}
	case "double":
		{
			return "double"
		}
	default:
		{
			return "string"
		}
	}

	return "string"
}

func CPPTypeLen(colType string) int {
	switch colType {
	case "int8", "uint8":
		{
			return 1
		}
	case "int16", "uint16":
		{
			return 2
		}
	case "int32", "int", "uint32":
		{
			return 4
		}
	case "int64", "uint64", "double":
		{
			return 8
		}
	default:
		{
			return 0
		}
	}
	return 0
}

func CPPField_CPPType(colType string) string {
	switch colType {
	case "int8", "int16", "int32", "uint8", "uint16", "uint32", "int64", "uint64", "double":
		{
			return colType
		}
	case "int":
		return "int32"
	default:
		{
			return "string"
		}
	}
	return "string"
}

func CPPFieldPackRedis_COL_FORMAT(colType string) string {
	var DefineSQL string
	switch colType {
	case "int8", "int16", "int32", "int":
		{
			DefineSQL = "%d"
		}
	case "uint8", "uint16", "uint32":
		{
			DefineSQL = "%u"
		}
	case "int64":
		{
			DefineSQL = "%lld"
		}
	case "uint64":
		{
			DefineSQL = "%llu"
		}
	case "double":
		{
			DefineSQL = "%g"
		}
	default:
		{
			DefineSQL = "%s"
		}
	}
	return DefineSQL
}

func CPPFieldPackSQL_COL_FORMAT(colType string) string {
	var DefineSQL string
	switch colType {
	case "int8", "int16", "int32", "int":
		{
			DefineSQL += "%d"
		}
	case "uint8", "uint16", "uint32":
		{
			DefineSQL += "%u"
		}
	case "int64":
		{
			DefineSQL += "%lld"
		}
	case "uint64":
		{
			DefineSQL += "%llu"
		}
	case "double":
		{
			DefineSQL += "%g"
		}
	default:
		{
			DefineSQL += "\\\"%s\\\""
		}
	}
	return DefineSQL
}
