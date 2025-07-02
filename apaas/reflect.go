package apaas

import (
	"reflect"
	"strings"
)

var _unknown_field = reflect.StructField{}

func GetFieldTypeByColumnNameV2(value reflect.Value, fieldColumnName string) (reflect.StructField, bool) {
	// check obj is pointer or not
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}
	typ := value.Type()
	for i := 0; i < value.NumField(); i++ {
		field := typ.Field(i)
		fieldValue := value.Field(i)
		if field.Anonymous {
			return GetFieldTypeByColumnNameV2(fieldValue, fieldColumnName)
		}
		tag := field.Tag.Get("gorm")
		if tag == "" {
			continue
		}
		if cname, ok := getColumnNameByColumnTag(tag); ok && cname == fieldColumnName {
			return field, ok
		}
	}
	return _unknown_field, false
}

/*
case：

	type Faction struct {
			ID          int64  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
			LiveID      int64  `gorm:"column:live_id" json:"live_id"`
			BizID       int64  `gorm:"column:biz_id" json:"biz_id"`
			FactionID   int64  `gorm:"column:faction_id;not null" json:"faction_id"`
			FactionName string `gorm:"column:faction_name" json:"faction_name"`
			OrgID       int64  `gorm:"column:org_id;comment:union外键" json:"org_id" apass_engine_lookup_id:"webcast.union.org_id"` // union外键
		}

fieldColumnName: faction_id
*/
func GetFieldTypeByColumnName(obj any, fieldColumnName string) (reflect.StructField, bool) {
	return GetFieldTypeByColumnNameV2(reflect.ValueOf(obj), fieldColumnName)
}

/*
case: `gorm:"column:id;primaryKey;autoIncrement:true"
*/
func getColumnNameByColumnTag(tag string) (string, bool) {
	fs := strings.Split(tag, ";")
	if len(fs) >= 1 {
		sfs := strings.Split(fs[0], ":")
		if sfs[0] == "column" {
			return sfs[1], true
		}
	}
	return "", false
}

func GetFieldTypeByColumnNameV3(typ reflect.Type, fieldColumnName string) (reflect.StructField, bool) {
	// check obj is pointer or not
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		if field.Anonymous {
			return GetFieldTypeByColumnNameV3(field.Type, fieldColumnName)
		}
		tag := field.Tag.Get("gorm")
		if tag == "" {
			continue
		}
		if cname, ok := getColumnNameByColumnTag(tag); ok && cname == fieldColumnName {
			return field, ok
		}
	}
	return _unknown_field, false
}
