package server

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/pineal-niwan/sensor/cache/key_value_service/pb"
	"github.com/pineal-niwan/sensor/cache/lru"
	"google.golang.org/grpc"
	"net"
)

type StringKeyValueService struct {
	kvHash *lru.StringAsKeyLruHash
}

var (
	_emptyMsg = &empty.Empty{}
)

// 新建服务
func NewStringKeyValueService(size int) *StringKeyValueService {
	kService := &StringKeyValueService{}
	kService.kvHash = lru.NewLruHash(size)
	return kService
}

// 设置缓存
func (k *StringKeyValueService) Set(ctx context.Context, msg *pb.MsgStringKeyValue) (*empty.Empty, error) {
	k.kvHash.Set(msg.Key, msg.Value)
	return _emptyMsg, nil
}

// 设置缓存-带超时
func (k *StringKeyValueService) SetWithTimeout(ctx context.Context, msg *pb.MsgStringKeyValueTimeout) (*empty.Empty, error) {
	k.kvHash.SetWithTimeout(msg.Key, msg.Value, msg.Timeout)
	return _emptyMsg, nil
}

//设置value消息
func (k *StringKeyValueService) setupMsgValue(val interface{}, ok bool) *pb.MsgValue {
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

// 设置缓存 -- if not exit
func (k *StringKeyValueService) SetIfKeyNotExist(ctx context.Context, msg *pb.MsgStringKeyValue) (*pb.MsgValue, error) {

	val, ok := k.kvHash.SetIfKeyNotExist(msg.Key, msg.Value)
	return k.setupMsgValue(val, ok), nil
}

// 设置缓存-带超时 -- if not exist
func (k *StringKeyValueService) SetWithTimeoutIfKeyNotExist(ctx context.Context, msg *pb.MsgStringKeyValueTimeout) (
	*pb.MsgValue, error) {

	val, ok := k.kvHash.SetWithTimeoutIfKeyNotExist(msg.Key, msg.Value, msg.Timeout)
	return k.setupMsgValue(val, ok), nil
}

// 设置缓存 -- if exit
func (k *StringKeyValueService) SetIfKeyExist(ctx context.Context, msg *pb.MsgStringKeyValue) (*pb.MsgValue, error) {

	val, ok := k.kvHash.SetIfKeyExist(msg.Key, msg.Value)
	return k.setupMsgValue(val, ok), nil
}

// 设置缓存-带超时 -- if exist
func (k *StringKeyValueService) SetWithTimeoutIfKeyExist(ctx context.Context, msg *pb.MsgStringKeyValueTimeout) (
	*pb.MsgValue, error) {

	val, ok := k.kvHash.SetWithTimeoutIfKeyExist(msg.Key, msg.Value, msg.Timeout)
	return k.setupMsgValue(val, ok), nil
}

// 获取缓存
func (k *StringKeyValueService) Get(ctx context.Context, in *pb.MsgStringKey) (*pb.MsgValue, error) {
	val, ok := k.kvHash.Get(in.Key)
	return k.setupMsgValue(val, ok), nil
}

// 获取缓存后刷新
func (k *StringKeyValueService) GetThenRefresh(ctx context.Context, in *pb.MsgStringKeyTimeout) (*pb.MsgValue, error) {
	val, ok := k.kvHash.GetThenRefreshTimeout(in.Key, in.Timeout)
	return k.setupMsgValue(val, ok), nil
}

// 获取缓存长度
func (k *StringKeyValueService) GetLen(ctx context.Context, in *empty.Empty) (*pb.MsgLen, error) {
	hashLen, listLen := k.kvHash.Len()
	return &pb.MsgLen{
		HashLen: int32(hashLen),
		ListLen: int32(listLen),
	}, nil
}

// 清空缓存
func (k *StringKeyValueService) Clear(ctx context.Context, in *empty.Empty) (*empty.Empty, error) {
	k.kvHash.Clear()
	return _emptyMsg, nil
}

//启动string Key-value服务
func StartStringKeyValueServer(port string, size int) error {
	ln, err := net.Listen("tcp", port)
	if err != nil {
		return err
	}

	kService := NewStringKeyValueService(size)

	rpcSrv := grpc.NewServer()
	pb.RegisterStringKeyValueServiceServer(rpcSrv, kService)
	err = rpcSrv.Serve(ln)
	return err
}
