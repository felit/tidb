// Copyright 2015 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package mysql

import "math"

// GetDefaultFieldLength is used for Interger Types, Flen is the display length.
// Call this when no Flen assigned in ddl.
// or column value is calculated from an expression.
// For example: "select count(*) from t;", the column type is int64 and Flen in ResultField will be 21.
// See https://dev.mysql.com/doc/refman/5.7/en/storage-requirements.html
func GetDefaultFieldLength(tp byte) int {
	switch tp {
	case TypeTiny:
		return 4
	case TypeShort:
		return 6
	case TypeInt24:
		return 9
	case TypeLong:
		return 11
	case TypeLonglong:
		return 21
	case TypeFloat:
		return 12
	case TypeDouble:
		return 22
	case TypeDate, TypeDuration:
		return 10
	case TypeDatetime, TypeTimestamp:
		return 19
	case TypeString, TypeBit:
		return 1
	case TypeNewDecimal:
		// See https://dev.mysql.com/doc/refman/5.7/en/fixed-point-types.html
		return 10
	case TypeBlob:
		return 65535
	case TypeJSON:
		return math.MaxUint32
	default:
		//TODO: Add more types.
		return -1
	}
}

// GetDefaultDecimal returns the default decimal length for column.
func GetDefaultDecimal(tp byte) int {
	switch tp {
	case TypeNewDecimal:
		// See https://dev.mysql.com/doc/refman/5.7/en/fixed-point-types.html
		return 0
	case TypeDatetime, TypeDuration:
		return 0
	default:
		//TODO: Add more types.
		return -1
	}
}
