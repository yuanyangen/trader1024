package model

import "github.com/bytedance/sonic"

type HttpProxy struct {
	Addr       string
	ProxyType  string
	ExpireTime int64
	UserName   string
	Password   string
}

func (hp *HttpProxy) String() string {
	r, _ := sonic.MarshalString(hp)
	return r
}
