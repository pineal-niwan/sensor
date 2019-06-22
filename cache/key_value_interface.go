package cache

import (
	"context"
	"time"
)

type ICacheClient interface {
	// 关闭
	Close() error
	// 拨号
	Dial(timeout time.Duration) error
	// 设置缓存
	Set(ctx context.Context, key string, value []byte) error
	// 设置缓存-带超时
	SetWithTimeout(ctx context.Context, key string, value []byte, timeout int64) error
	// 获取缓存
	Get(ctx context.Context, key string) ([]byte, error)
	// 获取缓存后刷新
	GetThenRefresh(ctx context.Context, key string, timeout int64) ([]byte, error)
	// 获取缓存长度
	GetLen(ctx context.Context) (hashLen int, listLen int, err error)
	// 清空缓存
	Clear(ctx context.Context) error
}
