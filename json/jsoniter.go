package json

import (
	"github.com/json-iterator/go"
)

var (
	jsonStd     = jsoniter.ConfigCompatibleWithStandardLibrary
	jsonDefault = jsoniter.ConfigDefault
	jsonFast    = jsoniter.ConfigFastest
)

//标准jsoniter json库 100%兼容
var (
	Marshal       = jsonStd.Marshal
	Unmarshal     = jsonStd.Unmarshal
	MarshalIndent = jsonStd.MarshalIndent
	NewDecoder    = jsonStd.NewDecoder
	NewEncoder    = jsonStd.NewEncoder
)

//缺省jsoniter json库
var (
	DefaultMarshal       = jsonDefault.Marshal
	DefaultUnmarshal     = jsonDefault.Unmarshal
	DefaultMarshalIndent = jsonDefault.MarshalIndent
	DefaultNewDecoder    = jsonDefault.NewDecoder
	DefaultNewEncoder    = jsonDefault.NewEncoder
)

//快速jsoniter json库 -- 浮点数会丢失精度，小数点最多后6位
var (
	FastMarshal       = jsonFast.Marshal
	FastUnmarshal     = jsonFast.Unmarshal
	FastMarshalIndent = jsonFast.MarshalIndent
	FastNewDecoder    = jsonFast.NewDecoder
	FastNewEncoder    = jsonFast.NewEncoder
)
