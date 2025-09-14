package apaas

import (
	"context"
	"database/sql"
	"fmt"
)

type Rewriter interface {
	RewriteDSL(context.Context, *ApaasDSLArgs) (string, error)
	RewriteDDL(context.Context, *ApaasDDLArgs) (string, error)
}

type SqlExplain interface {
	Explain(sql string, vars ...interface{}) string
}

type Filter interface {
	Filt(*ApaasDSLArgs) error
}

type ApaasEngine struct {
	Rewriter Rewriter
	SqlPool  ConnPool
	Filter   Filter
}

type ConnPool interface {
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
}

func NewApaasEngine(rewriter Rewriter, pool ConnPool, filter Filter) *ApaasEngine {
	return &ApaasEngine{
		Rewriter: rewriter,
		SqlPool:  pool,
		Filter:   filter,
	}
}

func (p *ApaasEngine) RewriteDSL(ctx context.Context, dslArgs *ApaasDSLArgs) (string, error) {
	return p.Rewriter.RewriteDSL(ctx, dslArgs)
}

func (p *ApaasEngine) RewriteDDL(ctx context.Context, ddlArgs *ApaasDDLArgs) (string, error) {
	return p.Rewriter.RewriteDDL(ctx, ddlArgs)
}

func (p *ApaasEngine) PrepareContext(ctx context.Context, query string) (*sql.Stmt, error) {
	return nil, nil

}
func (p *ApaasEngine) ExecContext(ctx context.Context, dslArgs *ApaasDSLArgs) (sql.Result, error) {
	err := p.Filter.Filt(dslArgs)
	if err != nil {
		return nil, err
	}
	sqlStr, err := p.Rewriter.RewriteDSL(ctx, dslArgs)
	if err != nil {
		return nil, GenError(fmt.Sprintf("ApaasEngine ExecContext error=%s", err.Error()))
	}
	return p.SqlPool.ExecContext(ctx, sqlStr)
}
func (p *ApaasEngine) QueryContext(ctx context.Context, dslArgs *ApaasDSLArgs) (*sql.Rows, error) {
	err := p.Filter.Filt(dslArgs)
	if err != nil {
		return nil, err
	}
	sqlStr, err := p.Rewriter.RewriteDSL(ctx, dslArgs)
	if err != nil {
		return nil, GenError(fmt.Sprintf("ApaasEngine QueryContext error=%s", err.Error()))
	}
	return p.SqlPool.QueryContext(ctx, sqlStr)
}
func (p *ApaasEngine) QueryRowContext(ctx context.Context, dslArgs *ApaasDSLArgs) (*sql.Row, error) {
	err := p.Filter.Filt(dslArgs)
	if err != nil {
		return nil, err
	}
	sqlStr, err := p.Rewriter.RewriteDSL(ctx, dslArgs)
	if err != nil {
		return nil, GenError(fmt.Sprintf("ApaasEngine QueryRowContext error=%s", err.Error()))
	}
	return p.SqlPool.QueryRowContext(ctx, sqlStr), nil
}

type SimpleFilter struct{}

func (p SimpleFilter) Filt(dslArgs *ApaasDSLArgs) error {
	if dslArgs == nil {
		return GenError("ApaasEngine SimpleFilter error")
	}
	// TODO: filter db / table
	return nil
}
