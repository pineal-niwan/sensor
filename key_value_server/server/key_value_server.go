package server

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/pineal-niwan/sensor/hash"
	"github.com/pineal-niwan/sensor/key_value_server/pb"
	"google.golang.org/grpc"
	"net"
)

type KeyValueService struct {
	kvHash *hash.LruHash
}

var (
	_emptyMsg = &empty.Empty{}
)

// 新建服务
func NewKeyValueService(size int) *KeyValueService {
	kService := &KeyValueService{}
	kService.kvHash = hash.NewLruHash(size)
	return kService
}

// 设置缓存
func (k *KeyValueService) Set(ctx context.Context, msg *pb.MsgKeyValue) (*empty.Empty, error) {
	k.kvHash.Set(msg.Key, msg.Value)
	return _emptyMsg, nil
}

// 设置缓存-带超时
func (k *KeyValueService) SetWithTimeout(ctx context.Context, msg *pb.MsgKeyValueTimeout) (*empty.Empty, error) {
	k.kvHash.SetWithTimeout(msg.Key, msg.Value, msg.Timeout)
	return _emptyMsg, nil
}

//设置value消息
func (k *KeyValueService) setupMsgValue(val interface{}, ok bool) *pb.MsgValue {
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
func (k *KeyValueService) Get(ctx context.Context, in *pb.MsgKey) (*pb.MsgValue, error) {
	val, ok := k.kvHash.Get(in.Key)
	return k.setupMsgValue(val, ok), nil
}

// 获取缓存后刷新
func (k *KeyValueService) GetThenRefresh(ctx context.Context, in *pb.MsgKeyTimeout) (*pb.MsgValue, error) {
	val, ok := k.kvHash.GetThenRefreshTimeout(in.Key, in.Timeout)
	return k.setupMsgValue(val, ok), nil
}

// 获取缓存长度
func (k *KeyValueService) GetLen(ctx context.Context, in *empty.Empty) (*pb.MsgLen, error) {
	hashLen, listLen := k.kvHash.Len()
	return &pb.MsgLen{
		HashLen: int32(hashLen),
		ListLen: int32(listLen),
	}, nil
}

// 清空缓存
func (k *KeyValueService) Clear(ctx context.Context, in *empty.Empty) (*empty.Empty, error) {
	k.kvHash.Clear()
	return _emptyMsg, nil
}

//启动Key-value服务
func StartKServer(port string, size int) error {
	ln, err := net.Listen("tcp", port)
	if err != nil {
		return err
	}

	kService := NewKeyValueService(size)

	rpcSrv := grpc.NewServer()
	pb.RegisterKeyValueServiceServer(rpcSrv, kService)
	err = rpcSrv.Serve(ln)
	return err
}
