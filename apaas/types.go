package apaas

import (
	"fmt"
	"strings"
)

const (
	TagID      = "apass_engine_lookup_id"
	TagValue   = "apass_engine_lookup_value"
	MaxTagDeep = 4
)

type ApaasFieldType uint8
type FieldType uint8

func (f FieldType) IsInvalid() bool {
	return f == _skipFieldType
}

const (
	_skipFieldType FieldType = iota
	FieldBool
	FieldInt
	FieldUInt
	FieldInt64
	FieldUInt64
	FieldFloat
	FieldFloat64
	FieldString
	FieldTime
	FieldBit
	FieldBin
	FieldArray
	FieldObject
)

const (
	_skipApaasFieldType ApaasFieldType = iota
	ApaasLookupID
	ApaasLookupValue
	ApaasExtraType
	ApaasFormulaType
)

var FieldMapString = []string{
	"_skipApaasFieldTyle",
	"FieldBool",
	"FieldInt",
	"FieldUInt",
	"FieldInt64",
	"FieldUInt64",
	"FieldFloat",
	"FieldFloat64",
	"FieldString",
	"FieldTime",
	"FieldBit",
	"FieldBin",
	"FieldArray",
	"FieldObject",
}

func (p FieldType) String() string {
	return FieldMapString[p]
}

/*
description: store dynamic field info when is an instance like tag

	`apass_engine_lookup_value`

case:

	type RoomAnchorView struct {
			Room
			Gender string `gorm:"column:gender" json:"gender" apass_engine_lookup_value:"anchor_id.uid.gender"`
			Career string `gorm:"column:career" json:"career" apass_engine_lookup_value:"anchor_id.uid.career"`
			Name   string `gorm:"column:name" json:"name" apass_engine_lookup_value:"anchor_id.uid.name"`
		}

		type RoomAnchorFactionView struct {
			Room
			Gender    string `gorm:"column:gender" json:"gender" apass_engine_lookup_value:"anchor_id.uid.gender"`
			Career    string `gorm:"column:career" json:"career" apass_engine_lookup_value:"anchor_id.uid.career"`
			Name      string `gorm:"column:name" json:"name" apass_engine_lookup_value:"anchor_id.uid.name"`
			OrgName   string `gorm:"column:org_name" json:"org_name" apass_engine_lookup_value:"anchor_id.faction_id.faction_name"`
			UnionInfo string `gorm:"column:union_info" json:"union_info" apass_engine_lookup_value:"anchor_id.faction_id.org_id.union_info"`
		}
*/
type ApaasLookupMeta struct {
	// field gorm column tag name/table field's name. example: union_info
	CName string
	// lookup orgin meta. example: [anchor_id.faction_id.org_id.union_info]
	/*
		    1. anchor_id;       2. faction_id;        3: org_id;
			1. anchor.anchor_id;2. faction.faction_id;3: union.org_id;
	*/
	LookupMetas []*LookupMeta
	LastField   string // org_name/union_name

	/*
		OrgTag only set value when used in SDK mode
	*/
	// lookup org tag [anchor_id, faction_id, org_id, org_name]
	OrgTag []string
}

// lookup orgin meta. example: [anchor_id.faction_id.org_id.union_info]
type LookupMeta struct {
	FieldName   string
	ForeignMeta ForeignMeta
}

type ApaasTable struct {
	TableName     string
	DBName        string
	Fields        []*ApaasField
	FieldsByName  map[string]*ApaasField
	LookupIDField *ApaasField   // example: room.room_id, user.uid, faction.faction_id
	ForeignFields []*ApaasField // foreign key. example: room.anchor_id, named of relookupid
	FormulaFields []*ApaasField // example: union.title=update_time + org_name
	ExtraMeta     *ExtraMeta    // maybe both room_extra and contract extra
}

type TableExtraType uint8

const (
	RoomExtraType TableExtraType = iota
	ContractExtraType
)

type ExtraCond struct {
	FieldName string
	OP        string
	Value     string
}

type ContractExtra struct {
	ExtraCond *ExtraCond
	ExtraRule string
}

type ExtraMeta struct {
	ContrctExtra []*ContractExtra
	ExtraFields  []*ApaasField // example: room.extra is json object, or contract.data_key = 1 && contract.data_value = 1
}

func (p *ExtraMeta) Check(dest any, e map[string]any) error {
	var err error
	for key, checker := range extraChecker {
		switch d := dest.(type) {
		case []map[string]any:
			for _, ele := range d {
				err = checker(p, ele, e)
				if err != nil {
					return fmt.Errorf("ExtraMeta Check ExtraType=%s error=(%s)", key, err.Error())
				}
			}
		case map[string]any:
			err = checker(p, d, e)
			if err != nil {
				return fmt.Errorf("ExtraMeta Check ExtraType=%s error=(%s)", key, err.Error())
			}
		}
	}
	return nil
}

type ApaasField struct {
	Name      string
	Type      string
	FType     FieldType
	IsUniq    bool
	ApaasMeta *ApaasMeta
}

func (p *ApaasField) arseFieldType() {
	tp := strings.ToUpper(p.Type)
	ftp := _skipFieldType
	switch tp {
	case "BOOL":
		ftp = FieldBool
	case "INT", "TINYINT", "SMALLINT", "MEDIUMINT":
		ftp = FieldInt
	case "BIGINT":
		ftp = FieldInt64
	case "TINYINT UNSIGNED", "SMALLINT UNSIGNED", "MEDIUMINT UNSIGNED", "INT UNSIGNED":
		ftp = FieldUInt
	case "BIGINT UNSIGNED":
		ftp = FieldInt64
	case "FLOAT":
		ftp = FieldFloat
	case "DOUBLE", "DECIMAL", "REAL":
		ftp = FieldFloat64
	case "DATE", "DATETIME", "TIMESTAMP", "TIME", "YEAR":
		ftp = FieldTime
	case "VARCHAR", "CHAR", "ENUM", "TEXT", "TINYTEXT", "MEDIUMTEXT", "LONGTEXT", "JSON":
		ftp = FieldString
	case "BLOB", "TINYBLOB", "MEDIUMBLOB", "LONGBLOB", "BINARY", "VARBINARY":
		ftp = FieldBin
	case "BIT":
		ftp = FieldBit
	default:
		ftp = FieldString
	}
	p.FType = ftp
}

func (p *ApaasField) GetApaasMeta() *ApaasMeta {
	return p.ApaasMeta
}

type ForeignMeta struct {
	DBName string
	TName  string
	FName  string
	FTMeta *ApaasTable
}

type ApaasMeta struct {
	ForeignMeta *ForeignMeta // equal to LookupMeta
	ApaasFType  ApaasFieldType
	ExtraMeta   map[string]*ExtraFieldMeta
	FormulaMeta *FormulaMeta
}

func (p *ApaasMeta) IsApaasFieldType() bool {
	return p.ApaasFType == _skipApaasFieldType
}
func (p *ApaasMeta) IsExtraField() bool {
	return p.ApaasFType == ApaasExtraType
}
func (p *ApaasMeta) IsFormulaField() bool {
	return p.ApaasFType == ApaasFormulaType
}
func (p *ApaasMeta) IsLookupID() bool {
	return p.ApaasFType == ApaasLookupID
}
func (p *ApaasMeta) GetApaasFieldType() ApaasFieldType {
	return p.ApaasFType
}
func (p *ApaasMeta) GetExtraMeta() map[string]*ExtraFieldMeta {
	return p.ExtraMeta
}
func (p *ApaasMeta) GetFormulaMeta() *FormulaMeta {
	return p.FormulaMeta
}

type ExtraFieldMeta struct {
	Key        string
	Type       FieldType                  // Bool/Int/Float/String/Object/Array
	ObjectMeta map[string]*ExtraFieldMeta // if Type is Object, use ObjectMeta
	ArrayMeta  []*ExtraFieldMeta          // if Type is Array, need ArrayMeta
}

type FormulaMeta struct {
	FormulaRule string   // anchor_id.faction_name + anchor_id.org_id.org_name
	FormatRule  string   // record.anchor_id_faction_name + record.anchor_id_org_id_org_name
	LookupTags  []string // [anchor_id.faction_name, anchor_id.org_id.org_name]
}

func (p *FormulaMeta) ExtractAndReplace() {
	tags := ExtractNestedFields(p.FormulaRule)
	p.LookupTags = tags
	newRule := p.FormulaRule
	for _, tag := range tags {
		newRule = strings.Replace(newRule, tag, FormatByteScriptColumnTag(tag), 1)
	}
	p.FormatRule = newRule
	fmt.Printf("raw formular rule: %s\nrewrite formular rule: %s\n", p.FormulaRule, p.FormatRule)
}

type DBMeta struct {
	tableList    []*ApaasTable
	tableView    map[string]*ApaasTable
	lookupIDView map[string]*ApaasTable // each table has one one and only key to supprort lookup
}

func (p *DBMeta) GetTableByLookupID(lookID string) (*ApaasTable, bool) {
	v, ok := p.lookupIDView[lookID]
	return v, ok
}

func (p *DBMeta) GetTableByName(tableName string) (*ApaasTable, bool) {
	v, ok := p.tableView[tableName]
	return v, ok
}
