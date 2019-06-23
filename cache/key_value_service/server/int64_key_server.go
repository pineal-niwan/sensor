package server

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/pineal-niwan/sensor/cache/key_value_service/pb"
	"github.com/pineal-niwan/sensor/cache/lru"
	"google.golang.org/grpc"
	"net"
)

type Int64KeyValueService struct {
	kvHash *lru.LruInt64AsKeyHash
}

// 新建服务
func NewInt64KeyValueService(size int) *Int64KeyValueService {
	kService := &Int64KeyValueService{}
	kService.kvHash = lru.NewLruInt64AsKeyHash(size)
	return kService
}

// 设置缓存
func (k *Int64KeyValueService) Set(ctx context.Context, msg *pb.MsgInt64KeyValue) (*empty.Empty, error) {
	k.kvHash.Set(msg.Key, msg.Value)
	return _emptyMsg, nil
}

// 设置缓存-带超时
func (k *Int64KeyValueService) SetWithTimeout(ctx context.Context, msg *pb.MsgInt64KeyValueTimeout) (*empty.Empty, error) {
	k.kvHash.SetWithTimeout(msg.Key, msg.Value, msg.Timeout)
	return _emptyMsg, nil
}

//设置value消息
func (k *Int64KeyValueService) setupMsgValue(val interface{}, ok bool) *pb.MsgValue {
	msgValue := &pb.MsgValue{}
	if !ok {
		msgValue.Ok = false
		return msgValue
	} else {
		buffer, ok := val.([]byte)
		if !ok {
			msgValue.Ok = false
			return msgValue
		} else {
			msgValue.Ok = true
			msgValue.Value = buffer
			return msgValue
		}
	}
}

// 获取缓存
func (k *Int64KeyValueService) Get(ctx context.Context, in *pb.MsgInt64Key) (*pb.MsgValue, error) {
	val, ok := k.kvHash.Get(in.Key)
	return k.setupMsgValue(val, ok), nil
}

// 获取缓存后刷新
func (k *Int64KeyValueService) GetThenRefresh(ctx context.Context, in *pb.MsgInt64KeyTimeout) (*pb.MsgValue, error) {
	val, ok := k.kvHash.GetThenRefreshTimeout(in.Key, in.Timeout)
	return k.setupMsgValue(val, ok), nil
}

// 获取缓存长度
func (k *Int64KeyValueService) GetLen(ctx context.Context, in *empty.Empty) (*pb.MsgLen, error) {
	hashLen, listLen := k.kvHash.Len()
	return &pb.MsgLen{
		HashLen: int32(hashLen),
		ListLen: int32(listLen),
	}, nil
}

// 清空缓存
func (k *Int64KeyValueService) Clear(ctx context.Context, in *empty.Empty) (*empty.Empty, error) {
	k.kvHash.Clear()
	return _emptyMsg, nil
}

//启动int64 Key-value服务
func StartInt64KeyServer(port string, size int) error {
	ln, err := net.Listen("tcp", port)
	if err != nil {
		return err
	}

	kService := NewInt64KeyValueService(size)

	rpcSrv := grpc.NewServer()
	pb.RegisterInt64KeyValueServiceServer(rpcSrv, kService)
	err = rpcSrv.Serve(ln)
	return err
}
