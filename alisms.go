// Copyright 2018 by xruida, All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

// Package alisms 阿里云短信服务 SDK
package alisms

// 签名信息
const (
	SignatureMethod  = "HMAC_SHA1"
	SignatureVersion = "1.0"
)

// 一些固定的业务参数
const (
	Action   = "SendSms"
	Version  = "2017-05-25"
	RegionID = "cn-hangzhou"
)

// AliSMS 短信的配置项
type AliSMS struct {
	AccessKeyID  string
	AccessSecret string
}

// New 声明一个 AliSMS 实例
func New(keyid, secret string) *AliSMS {
	return &AliSMS{
		AccessKeyID:  keyid,
		AccessSecret: secret,
	}
}

// Send 发送短信
func (sms *AliSMS) Send(tplID string, number ...string) {
	req := &Request{}

	// TODO
}
