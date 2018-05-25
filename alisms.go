// Copyright 2018 by xruida, All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

// Package alisms 阿里云短信服务 SDK
package alisms

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// 签名信息
const (
	SignatureMethod  = "HMAC-SHA1"
	SignatureVersion = "1.0"
	HTTPSendURL      = "http://dysmsapi.aliyuncs.com"  //短信发送
	HTTPReqURL       = "http://dybaseapi.aliyuncs.com" //消息接收1
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
func (sms *AliSMS) Send(tplID, param, signname string, number ...string) (*Request, error) {
	req := &BodyRequest{}

	//处理批量号码
	req.PhoneNumbers = strings.Join(number, ",")
	req.Timestamp = time.Now().UTC().Format("2006-01-02T15:04:05Z")
	req.SignName = signname
	req.TemplateCode = tplID
	req.TemplateParam = param

	//提交参数列表

	params := map[string]string{}

	params["Timestamp"] = req.Timestamp
	params["AccessKeyId"] = sms.AccessKeyID
	params["SignatureMethod"] = SignatureMethod
	params["SignatureVersion"] = SignatureVersion
	rand.Seed(time.Now().Unix())
	rnd := rand.Int()
	params["SignatureNonce"] = strconv.Itoa(rnd)
	params["RegionId"] = RegionID
	params["Version"] = Version
	params["Action"] = Action
	params["Format"] = "JSON"
	params["PhoneNumbers"] = req.PhoneNumbers
	params["SignName"] = req.SignName
	params["TemplateCode"] = req.TemplateCode
	params["TemplateParam"] = req.TemplateParam

	signstr, err := Sign(sms.AccessSecret, params)
	if err != nil {
		return nil, err
	}

	url := HTTPSendURL + "?" + signstr

	reqd, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(reqd)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	relust := new(Request)

	if err := json.Unmarshal(body, relust); err != nil {
		return nil, err
	}

	return relust, nil
}
