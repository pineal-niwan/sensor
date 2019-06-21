package client

import (
	"context"
	"errors"
	"github.com/pineal-niwan/sensor/key_value_server/pb"
	"google.golang.org/grpc"
	"sync"
	"time"
)

var (
	ErrClientNotReady = errors.New(`client not ready`)

	RetEmptyBytes = make([]byte, 0)
)

type KClient struct {
	conn   *grpc.ClientConn
	client pb.KeyValueServiceClient
	url    string
	sync.RWMutex
}

//新建Client
func NewKClient(url string) *KClient {
	return &KClient{
		url: url,
	}
}

//关闭
func (c *KClient) Close() error {
	c.Lock()
	defer c.Unlock()

	if c.conn != nil {
		err := c.conn.Close()
		if err != nil {
			return err
		}
		//已经关闭了
		c.conn = nil
		c.client = nil
	}
	return nil
}

//连接
func (c *KClient) Dial(url string, timeout time.Duration, maxBackOff time.Duration) error {
	c.Lock()
	defer c.Unlock()

	//拨号连接
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	conn, err := grpc.DialContext(ctx, c.url, grpc.WithInsecure(), grpc.WithBackoffMaxDelay(maxBackOff))
	defer cancel()

	if err != nil {
		return err
	}

	//设置client和conn
	c.client = pb.NewKeyValueServiceClient(conn)
	c.conn = conn
	return nil
}

// 设置缓存
func (c *KClient) Set(ctx context.Context, key string, value []byte) error {
	c.RLock()
	defer c.RUnlock()

	if c.client == nil {
		return ErrClientNotReady
	}
	_, err := c.client.Set(ctx, &pb.MsgKeyValue{
		Key:   key,
		Value: value,
	})
	return err
}

// 设置缓存-带超时
func (c *KClient) SetWithTimeout(ctx context.Context, key string, value []byte, timeout int64) error {
	c.RLock()
	defer c.RUnlock()

	if c.client == nil {
		return ErrClientNotReady
	}
	_, err := c.client.SetWithTimeout(ctx, &pb.MsgKeyValueTimeout{
		Timeout: timeout,
		Key:     key,
		Value:   value,
	})
	return err
}

// 获取缓存
func (c *KClient) Get(ctx context.Context, key string) ([]byte, error) {
	var buff []byte

	c.RLock()
	defer c.RUnlock()

	if c.client == nil {
		return RetEmptyBytes, ErrClientNotReady
	}
	msgValue, err := c.client.Get(ctx, &pb.MsgKey{
		Key: key,
	})

	if err != nil {
		return RetEmptyBytes, err
	}

	if msgValue.Ok {
		buff = msgValue.Value
	} else {
		buff = RetEmptyBytes
	}

	return buff, err
}
