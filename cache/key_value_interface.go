package cache

import (
	"context"
	"time"
)

//cache接口
type ICacheClient interface {
	// 关闭
	Close() error
	// 拨号
	Dial(timeout time.Duration) error
	// 获取缓存长度
	GetLen(ctx context.Context) (hashLen int, listLen int, err error)
	// 清空缓存
	Clear(ctx context.Context) error
}

//以string为主键的cache
type IStringAsKeyCacheClient interface {
	ICacheClient
	// 设置缓存
	Set(ctx context.Context, key string, value []byte) error
	// 设置缓存-带超时
	SetWithTimeout(ctx context.Context, key string, value []byte, timeout int64) error
	// 获取缓存
	Get(ctx context.Context, key string) ([]byte, error)
	// 获取缓存后刷新
	GetThenRefresh(ctx context.Context, key string, timeout int64) ([]byte, error)
}

//以int64为主键的cache
type IInt64AsKeyCacheClient interface {
	ICacheClient
	// 设置缓存
	Set(ctx context.Context, key int64, value []byte) error
	// 设置缓存-带超时
	SetWithTimeout(ctx context.Context, key int64, value []byte, timeout int64) error
	// 获取缓存
	Get(ctx context.Context, key int64) ([]byte, error)
	// 获取缓存后刷新
	GetThenRefresh(ctx context.Context, key int64, timeout int64) ([]byte, error)
}
