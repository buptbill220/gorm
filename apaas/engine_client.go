package apaas

import (
	"context"
	"fmt"
	"os"
	"sync"
)

type ClientPool struct {
	rwLock sync.RWMutex
	m      map[string]ClientConnPool
}

type ClientConnPool interface {
	ExecContext(ctx context.Context, dslArgs *ApaasDSLArgs) (ApaasResult, error)
	QueryContext(ctx context.Context, dslArgs *ApaasDSLArgs) (*ApaasRows, error)
	QueryRowContext(ctx context.Context, dslArgs *ApaasDSLArgs) *ApaasRow
}

var globalConnPool ClientPool = ClientPool{
	m: make(map[string]ClientConnPool, 4),
}

/*
example：

	register global uniq pool: RegisterClientPool("",  pool)
	some special db has a pool: RegisterClientPool("webcast", pool)
*/
func RegisterClientPool(name string, pool ClientConnPool) bool {
	if pool == nil {
		fmt.Fprintf(os.Stderr, "%s RegisterClientPool name=%s pool nil", MSG_PREFIX, name)
		return false
	}
	globalConnPool.rwLock.RLock()
	_, ok := globalConnPool.m[name]
	if ok {
		fmt.Fprintf(os.Stderr, "%s RegisterClientPool name=%s have been registered", MSG_PREFIX, name)
		globalConnPool.rwLock.RUnlock()
		return false
	}
	globalConnPool.rwLock.RUnlock()
	globalConnPool.rwLock.Lock()
	globalConnPool.m[name] = pool
	globalConnPool.rwLock.Unlock()
	return true
}

func GetClientPool(name string) ClientConnPool {
	globalConnPool.rwLock.RLock()
	defer globalConnPool.rwLock.RUnlock()
	pool, ok := globalConnPool.m[name]
	if ok {
		return pool
	}
	pool, ok = globalConnPool.m[""]
	if ok {
		return pool
	}
	return nil
}

func NewApaasEngineClient() ApaasEngineClient {
	return ApaasEngineClient{}
}

type ApaasEngineClient struct{}

func (p ApaasEngineClient) ExecContext(ctx context.Context, dslArgs *ApaasDSLArgs) (ApaasResult, error) {
	pool := GetClientPool(dslArgs.DBName)
	if pool == nil {
		return ApaasResult{}, GenError(fmt.Sprintf("ApaasEngineClient ExecContext dbname=%s get pool nil", dslArgs.DBName))
	}
	return pool.ExecContext(ctx, dslArgs)
}

func (p ApaasEngineClient) QueryContext(ctx context.Context, dslArgs *ApaasDSLArgs) (*ApaasRows, error) {
	pool := GetClientPool(dslArgs.DBName)
	if pool == nil {
		return nil, GenError(fmt.Sprintf("ApaasEngineClient QueryContext dbname=%s get pool nil", dslArgs.DBName))
	}
	return pool.QueryContext(ctx, dslArgs)
}
func (p ApaasEngineClient) QueryRowContext(ctx context.Context, dslArgs *ApaasDSLArgs) (*ApaasRow, error) {
	pool := GetClientPool(dslArgs.DBName)
	if pool == nil {
		return nil, GenError(fmt.Sprintf("ApaasEngineClient QueryRowContext dbname=%s get pool nil", dslArgs.DBName))
	}
	return pool.QueryRowContext(ctx, dslArgs), nil
}
