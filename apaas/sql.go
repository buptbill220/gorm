package apaas

import (
	"bytes"
	"errors"
	"reflect"
	"strconv"
	"time"
)

type Rows interface {
	Columns() ([]string, error)
	ColumnTypes() ([]reflect.Type, error)
	Next() bool
	Scan(dest ...interface{}) error
	Err() error
	Close() error
}

type ApaasRows struct {
	rawResult []map[string]any
	columns   []string
	types     []reflect.Type
	iter      int
	err       error
}

func NewApaasRows(result []map[string]any, columns []string, err error) *ApaasRows {
	return &ApaasRows{
		rawResult: result,
		columns:   columns,
		iter:      0,
		err:       err,
	}
}

func (p *ApaasRows) Columns() ([]string, error) {
	return p.columns, nil
}

func (p *ApaasRows) ColumnTypes() ([]reflect.Type, error) {
	if len(p.types) == 0 && len(p.rawResult) > 0 {
		p.types = make([]reflect.Type, len(p.columns))
		dest := p.rawResult[0]
		for i, col := range p.columns {
			p.types[i] = reflect.TypeOf(dest[col])
		}
		return p.types, nil
	}
	return p.types, nil
}

func (p *ApaasRows) Next() bool {
	if p.iter >= len(p.rawResult) {
		return false
	}

	return true
}

var errNilPtr = errors.New("destination pointer is nil")

func initPtr(dest ...any) {
	for _, d := range dest {
		switch dt := d.(type) {
		case **int8:
			*dt = new(int8)
		case **int16:
			*dt = new(int16)
		case **int32:
			*dt = new(int32)
		case **int64:
			*dt = new(int64)
		case **uint:
			*dt = new(uint)
		case **uint8:
			*dt = new(uint8)
		case **uint16:
			*dt = new(uint16)
		case **uint32:
			*dt = new(uint32)
		case **uint64:
			*dt = new(uint64)
		case **uintptr:
			*dt = new(uintptr)
		case **float32:
			*dt = new(float32)
		case **float64:
			*dt = new(float64)
		case **string:
			*dt = new(string)
		case **[]byte:
			*dt = new([]byte)
		case **time.Time:
			*dt = new(time.Time)
		case **any:
			*dt = new(any)
		}
	}

}

type TimeMethod interface {
	String() string
	GoString() string
	UnixNano() int64
	Unix() int64
	Format(layout string) string
}

func (p *ApaasRows) Scan(dest ...interface{}) error {
	if !p.Next() {
		return nil
	}
	initPtr(dest...)
	row := p.rawResult[p.iter]
	for idx, col := range p.columns {
		switch rt := row[col].(type) {
		case int:
			switch dt := dest[idx].(type) {
			case *any:
				*dt = rt
			case *int:
				*dt = rt
			case *int8:
				*dt = int8(rt)
			case *int16:
				*dt = int16(rt)
			case *int32:
				*dt = int32(rt)
			case *int64:
				*dt = int64(rt)
			case *uint:
				*dt = uint(rt)
			case *uint8:
				*dt = uint8(rt)
			case *uint16:
				*dt = uint16(rt)
			case *uint32:
				*dt = uint32(rt)
			case *uint64:
				*dt = uint64(rt)
			case *uintptr:
				*dt = uintptr(rt)
			case *float32:
				*dt = float32(rt)
			case *float64:
				*dt = float64(rt)
			case **int:
				**dt = rt
			case **int8:
				**dt = int8(rt)
			case **int16:
				**dt = int16(rt)
			case **int32:
				**dt = int32(rt)
			case **int64:
				**dt = int64(rt)
			case **uint:
				**dt = uint(rt)
			case **uint8:
				**dt = uint8(rt)
			case **uint16:
				**dt = uint16(rt)
			case **uint32:
				**dt = uint32(rt)
			case **uint64:
				**dt = uint64(rt)
			case **uintptr:
				**dt = uintptr(rt)
			case **float32:
				**dt = float32(rt)
			case **float64:
				**dt = float64(rt)
			}
		case uint:
			switch dt := dest[idx].(type) {
			case *any:
				*dt = rt
			case *int:
				*dt = int(rt)
			case *int8:
				*dt = int8(rt)
			case *int16:
				*dt = int16(rt)
			case *int32:
				*dt = int32(rt)
			case *int64:
				*dt = int64(rt)
			case *uint:
				*dt = uint(rt)
			case *uint8:
				*dt = uint8(rt)
			case *uint16:
				*dt = uint16(rt)
			case *uint32:
				*dt = uint32(rt)
			case *uint64:
				*dt = uint64(rt)
			case *uintptr:
				*dt = uintptr(rt)
			case *float32:
				*dt = float32(rt)
			case *float64:
				*dt = float64(rt)
			case **int:
				**dt = int(rt)
			case **int8:
				**dt = int8(rt)
			case **int16:
				**dt = int16(rt)
			case **int32:
				**dt = int32(rt)
			case **int64:
				**dt = int64(rt)
			case **uint:
				**dt = uint(rt)
			case **uint8:
				**dt = uint8(rt)
			case **uint16:
				**dt = uint16(rt)
			case **uint32:
				**dt = uint32(rt)
			case **uint64:
				**dt = uint64(rt)
			case **uintptr:
				**dt = uintptr(rt)
			case **float32:
				**dt = float32(rt)
			case **float64:
				**dt = float64(rt)
			}
		case int32:
			switch dt := dest[idx].(type) {
			case *any:
				*dt = rt
			case *int:
				*dt = int(rt)
			case *int8:
				*dt = int8(rt)
			case *int16:
				*dt = int16(rt)
			case *int32:
				*dt = int32(rt)
			case *int64:
				*dt = int64(rt)
			case *uint:
				*dt = uint(rt)
			case *uint8:
				*dt = uint8(rt)
			case *uint16:
				*dt = uint16(rt)
			case *uint32:
				*dt = uint32(rt)
			case *uint64:
				*dt = uint64(rt)
			case *uintptr:
				*dt = uintptr(rt)
			case *float32:
				*dt = float32(rt)
			case *float64:
				*dt = float64(rt)
			case **int:
				**dt = int(rt)
			case **int8:
				**dt = int8(rt)
			case **int16:
				**dt = int16(rt)
			case **int32:
				**dt = int32(rt)
			case **int64:
				**dt = int64(rt)
			case **uint:
				**dt = uint(rt)
			case **uint8:
				**dt = uint8(rt)
			case **uint16:
				**dt = uint16(rt)
			case **uint32:
				**dt = uint32(rt)
			case **uint64:
				**dt = uint64(rt)
			case **uintptr:
				**dt = uintptr(rt)
			case **float32:
				**dt = float32(rt)
			case **float64:
				**dt = float64(rt)
			}
		case uint32:
			switch dt := dest[idx].(type) {
			case *any:
				*dt = rt
			case *int:
				*dt = int(rt)
			case *int8:
				*dt = int8(rt)
			case *int16:
				*dt = int16(rt)
			case *int32:
				*dt = int32(rt)
			case *int64:
				*dt = int64(rt)
			case *uint:
				*dt = uint(rt)
			case *uint8:
				*dt = uint8(rt)
			case *uint16:
				*dt = uint16(rt)
			case *uint32:
				*dt = uint32(rt)
			case *uint64:
				*dt = uint64(rt)
			case *uintptr:
				*dt = uintptr(rt)
			case *float32:
				*dt = float32(rt)
			case *float64:
				*dt = float64(rt)
			case **int:
				**dt = int(rt)
			case **int8:
				**dt = int8(rt)
			case **int16:
				**dt = int16(rt)
			case **int32:
				**dt = int32(rt)
			case **int64:
				**dt = int64(rt)
			case **uint:
				**dt = uint(rt)
			case **uint8:
				**dt = uint8(rt)
			case **uint16:
				**dt = uint16(rt)
			case **uint32:
				**dt = uint32(rt)
			case **uint64:
				**dt = uint64(rt)
			case **uintptr:
				**dt = uintptr(rt)
			case **float32:
				**dt = float32(rt)
			case **float64:
				**dt = float64(rt)
			}
		case int64:
			switch dt := dest[idx].(type) {
			case *any:
				*dt = rt
			case *int:
				*dt = int(rt)
			case *int8:
				*dt = int8(rt)
			case *int16:
				*dt = int16(rt)
			case *int32:
				*dt = int32(rt)
			case *int64:
				*dt = rt
			case *uint:
				*dt = uint(rt)
			case *uint8:
				*dt = uint8(rt)
			case *uint16:
				*dt = uint16(rt)
			case *uint32:
				*dt = uint32(rt)
			case *uint64:
				*dt = uint64(rt)
			case *uintptr:
				*dt = uintptr(rt)
			case *float32:
				*dt = float32(rt)
			case *float64:
				*dt = float64(rt)
			case **int:
				**dt = int(rt)
			case **int8:
				**dt = int8(rt)
			case **int16:
				**dt = int16(rt)
			case **int32:
				**dt = int32(rt)
			case **int64:
				**dt = rt
			case **uint:
				**dt = uint(rt)
			case **uint8:
				**dt = uint8(rt)
			case **uint16:
				**dt = uint16(rt)
			case **uint32:
				**dt = uint32(rt)
			case **uint64:
				**dt = uint64(rt)
			case **uintptr:
				**dt = uintptr(rt)
			case **float32:
				**dt = float32(rt)
			case **float64:
				**dt = float64(rt)
			}
		case uint64:
			switch dt := dest[idx].(type) {
			case *any:
				*dt = rt
			case *int:
				*dt = int(rt)
			case *int8:
				*dt = int8(rt)
			case *int16:
				*dt = int16(rt)
			case *int32:
				*dt = int32(rt)
			case *int64:
				*dt = int64(rt)
			case *uint:
				*dt = uint(rt)
			case *uint8:
				*dt = uint8(rt)
			case *uint16:
				*dt = uint16(rt)
			case *uint32:
				*dt = uint32(rt)
			case *uint64:
				*dt = uint64(rt)
			case *uintptr:
				*dt = uintptr(rt)
			case *float32:
				*dt = float32(rt)
			case *float64:
				*dt = float64(rt)
			case **int:
				**dt = int(rt)
			case **int8:
				**dt = int8(rt)
			case **int16:
				**dt = int16(rt)
			case **int32:
				**dt = int32(rt)
			case **int64:
				**dt = int64(rt)
			case **uint:
				**dt = uint(rt)
			case **uint8:
				**dt = uint8(rt)
			case **uint16:
				**dt = uint16(rt)
			case **uint32:
				**dt = uint32(rt)
			case **uint64:
				**dt = uint64(rt)
			case **uintptr:
				**dt = uintptr(rt)
			case **float32:
				**dt = float32(rt)
			case **float64:
				**dt = float64(rt)
			}
		case float32:
			switch dt := dest[idx].(type) {
			case *any:
				*dt = rt
			case *int:
				*dt = int(rt)
			case *int8:
				*dt = int8(rt)
			case *int16:
				*dt = int16(rt)
			case *int32:
				*dt = int32(rt)
			case *int64:
				*dt = int64(rt)
			case *uint:
				*dt = uint(rt)
			case *uint8:
				*dt = uint8(rt)
			case *uint16:
				*dt = uint16(rt)
			case *uint32:
				*dt = uint32(rt)
			case *uint64:
				*dt = uint64(rt)
			case *uintptr:
				*dt = uintptr(rt)
			case *float32:
				*dt = rt
			case *float64:
				*dt = float64(rt)
			case **int:
				**dt = int(rt)
			case **int8:
				**dt = int8(rt)
			case **int16:
				**dt = int16(rt)
			case **int32:
				**dt = int32(rt)
			case **int64:
				**dt = int64(rt)
			case **uint:
				**dt = uint(rt)
			case **uint8:
				**dt = uint8(rt)
			case **uint16:
				**dt = uint16(rt)
			case **uint32:
				**dt = uint32(rt)
			case **uint64:
				**dt = uint64(rt)
			case **uintptr:
				**dt = uintptr(rt)
			case **float32:
				**dt = rt
			case **float64:
				**dt = float64(rt)
			}
		case float64:
			switch dt := dest[idx].(type) {
			case *any:
				*dt = rt
			case *int:
				*dt = int(rt)
			case *int8:
				*dt = int8(rt)
			case *int16:
				*dt = int16(rt)
			case *int32:
				*dt = int32(rt)
			case *int64:
				*dt = int64(rt)
			case *uint:
				*dt = uint(rt)
			case *uint8:
				*dt = uint8(rt)
			case *uint16:
				*dt = uint16(rt)
			case *uint32:
				*dt = uint32(rt)
			case *uint64:
				*dt = uint64(rt)
			case *uintptr:
				*dt = uintptr(rt)
			case *float32:
				*dt = float32(rt)
			case *float64:
				*dt = rt
			case **int:
				**dt = int(rt)
			case **int8:
				**dt = int8(rt)
			case **int16:
				**dt = int16(rt)
			case **int32:
				**dt = int32(rt)
			case **int64:
				**dt = int64(rt)
			case **uint:
				**dt = uint(rt)
			case **uint8:
				**dt = uint8(rt)
			case **uint16:
				**dt = uint16(rt)
			case **uint32:
				**dt = uint32(rt)
			case **uint64:
				**dt = uint64(rt)
			case **uintptr:
				**dt = uintptr(rt)
			case **float32:
				**dt = float32(rt)
			case **float64:
				**dt = rt
			}
		case string:
			switch dt := dest[idx].(type) {
			case *any:
				*dt = rt
			case *string:
				*dt = rt
			case *[]byte:
				*dt = []byte(rt)
			case **string:
				**dt = rt
			case **[]byte:
				**dt = []byte(rt)
			}
		case any:
			switch dt := dest[idx].(type) {
			case *any:
				*dt = rt
			case **time.Time:
				rt, ok := rt.(time.Time)
				if ok {
					**dt = rt
				}
			case *time.Time:
				rt, ok := rt.(time.Time)
				if ok {
					*dt = rt
				}
			}
		case []byte:
			switch dt := dest[idx].(type) {
			case *any:
				*dt = rt
			case *string:
				*dt = string(rt)
			case *[]byte:
				*dt = bytes.Clone(rt)
			case **string:
				**dt = string(rt)
			case **[]byte:
				**dt = bytes.Clone(rt)
			}
		case bool:
			switch dt := dest[idx].(type) {
			case *any:
				*dt = rt
			case *bool:
				*dt = rt
			case *string:
				*dt = strconv.FormatBool(rt)
			case *[]byte:
				*dt = []byte(strconv.FormatBool(rt))
			case **bool:
				**dt = rt
			case **string:
				**dt = strconv.FormatBool(rt)
			case **[]byte:
				**dt = []byte(strconv.FormatBool(rt))
			}
		case time.Time:
			switch dt := dest[idx].(type) {
			case *any:
				*dt = rt
			case *time.Time:
				*dt = rt
			case *string:
				*dt = rt.Format(time.DateTime)
			case *[]byte:
				*dt = []byte(rt.Format(time.DateTime))
			case **time.Time:
				**dt = rt
			case **string:
				**dt = rt.Format(time.DateTime)
			case **[]byte:
				**dt = []byte(rt.Format(time.DateTime))
			}

		case TimeMethod:
			switch dt := dest[idx].(type) {
			case *any:
				*dt = rt
			case *time.Time:
				nano := rt.UnixNano()
				*dt = time.Unix(nano/1e9, nano%1e9)
			case *string:
				*dt = rt.Format(time.DateTime)
			case *[]byte:
				*dt = []byte(rt.Format(time.DateTime))
			case **time.Time:
				nano := rt.UnixNano()
				**dt = time.Unix(nano/1e9, nano%1e9)
			case **string:
				**dt = rt.Format(time.DateTime)
			case **[]byte:
				**dt = []byte(rt.Format(time.DateTime))
			}
		case nil:
			switch dt := dest[idx].(type) {
			case *any:
				*dt = nil
			case *[]byte:
				*dt = nil
			case **[]byte:
				**dt = nil
			}
		default:
			switch dt := dest[idx].(type) {
			case *any:
				*dt = rt
			}
		}
	}
	p.iter++
	return nil
}

func (p *ApaasRows) Err() error {
	return p.err
}

func (p *ApaasRows) Close() error {
	p.iter = 0
	return nil
}

type ApaasRow struct {
	rows *ApaasRows
	err  error
}

func NewApaasRow(result []map[string]any, columns []string, err error) *ApaasRow {
	return &ApaasRow{NewApaasRows(result, columns, err), err}
}

func NewApaasRowV2(rows *ApaasRows, err error) *ApaasRow {
	return &ApaasRow{rows, err}
}

func (p *ApaasRow) Scan(dest ...any) error {
	return p.rows.Scan(dest...)
}

func (p *ApaasRow) Err() error {
	if p.rows.err != nil {
		return p.rows.err
	}
	return p.err
}

type ApaasResult struct {
	lastInsertID int64
	rowsAffected int64
}

func (p ApaasResult) LastInsertId() (int64, error) {
	return p.lastInsertID, nil
}

func (p ApaasResult) RowsAffected() (int64, error) {
	return p.rowsAffected, nil
}
