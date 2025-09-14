package apaas

type ApaasQueryType int8

const (
	_skipQueryType ApaasQueryType = iota
	InsertType
	SelectType
	UpdateType
	DeleteType
	RawQueryType
	RawExecType
)

type JoinType string

const (
	JoinStr      JoinType = "JOIN"
	LeftJoinStr  JoinType = "LEFT JOIN"
	RightJoinStr JoinType = "RIGHT JOIN"
)

// Where.Type
type WhereType string

const (
	WhereStr  WhereType = "WHERE"
	HavingStr WhereType = "HAVING"
)

type OrderDirection string

const (
	DescDirection OrderDirection = "DESC"
	AscDirection  OrderDirection = "ASC"
)

type DDLType string

const (
	CreateDDL        DDLType = "create"
	AlterStr         DDLType = "alter"
	DropStr          DDLType = "drop"
	RenameStr        DDLType = "rename"
	TruncateStr      DDLType = "truncate"
	CreateVindexStr  DDLType = "create vindex"
	AddColVindexStr  DDLType = "add vindex"
	DropColVindexStr DDLType = "drop vindex"
)

type Expr string

type Column struct {
	Table string
	Name  string
	Alias string
	Raw   bool
}
type ViewColumn struct {
	Column
	LookupTag string
}
type SelectExpr struct {
	Distinct bool
	/*
		used when query all object fields and don't use Select Clause
		1) sql: select uid, roomid; select 2)*; 3) room.*
		2) case: db.Find(&obj); db.First(&obj);
	*/
	Column *Column
	/*
		used when query some fileds and use Select Clause
		1) select uid, count(1) as pv; 2) distinct(uid) as uid; select avg(score) as score
		2) case: db.Select("name", "age").Find(&users); db.Select("name", "count(1) as pv").Find(&users)
	*/
	SubExpr *Expr
}

type Table struct {
	Name  string
	Alias string
	Raw   bool
}

type TableExpr struct {
	/*
		used when query the main table
		1) sql: select *from User
		2) case: db.Find(&User{}); db.First(&User{});
	*/
	TableName *Table
	/*
		used when query the sub query
		1) sql: select *from (select *from User) as t
		2) sql: SELECT * FROM (SELECT `name` FROM `users`) as u, (SELECT `name` FROM `pets`) as p
	*/
	SubExpr  *Expr
	JoinExpr *JoinTableExpr
}

type JoinTableExpr struct {
	LeftExpr  TableExpr
	Join      JoinType
	RightExpr TableExpr
	Condition Expr
}

type JoinCondition struct {
	On    Expr
	Using Columns
}

type ColIdent struct {
	// This artifact prevents this struct from being compared
	// with itself. It consumes no space as long as it's not the
	// last field in the struct.
	_            [0]struct{ _ []byte }
	val, lowered string
}

type Columns []ColIdent

type Limit struct {
	Offset   *int
	Rowcount int
}

type Where struct {
	Type    WhereType
	SubExpr Expr
}

type Order struct {
	SubExpr   Expr
	Direction string
}

type Partitions Columns

type Comments [][]byte
type DDL string
type ApaasDDLArgs struct {
	Type   DDLType
	RawDDL *DDL
}

type ApaasDSLArgs struct {
	Type     ApaasQueryType
	DBName   string
	Table    string
	Query    *Query
	Update   *Update
	Delete   *Delete
	Insert   *Insert
	RawSQL   string
	SQL      string
	Vars     []any
	ViewCols []ViewColumn
}

type Query struct {
	Comments  Comments
	Select    []SelectExpr
	TableExpr []TableExpr
	Where     *Where
	GroupBy   []Expr
	OrderBy   []Order
	Having    *Where
	Limit     *Limit
	Lock      *string
}

type Update struct {
	Comments   Comments
	TableExpr  []TableExpr
	UpdateExpr []Expr
	Where      *Where
	OrderBy    *Order
	Limit      *Limit
}

type TableName struct {
	Name, Qualifier TableIdent
}
type TableIdent struct {
	v string
}
type TableNames []TableName

type Delete struct {
	Comments   Comments
	Targets    []TableName
	TableExpr  []TableExpr
	Partitions Partitions
	Where      *Where
	OrderBy    *Order
	Limit      *Limit
}

type Insert struct {
	Action     string
	Comments   Comments
	Ignore     string
	Table      TableName
	Partitions Partitions
	Columns    Columns
	Rows       [][]string
}
