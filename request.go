// Copyright 2018 by xruida, All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package alisms

// Request 接受数据模型
type Request struct {
	RequestID string `json:"RequestId"` //请求ID
	Code      string `json:"Code"`      //状态码-返回OK代表请求成功,其他错误码详见错误码列表
	Message   string `json:"Message"`   //状态码的描述
	BizID     string `json:"BizId"`     //发送回执ID,可根据该ID查询具体的发送状态
}

//BodyRequest 入参参数
type BodyRequest struct {
	PhoneNumbers  string `json:"phonenumber"`   //电话号码
	SignName      string `json:"signname"`      //模板签名
	TemplateCode  string `json:"templatecode"`  //模板参数ID号
	TemplateParam string `json:"templateparam"` //模板变换json串
	OutID         string `json:"out_id"`        //调用发送短信接口时传的Outid
	Timestamp     string `json:"timestamp"`     //是请求的时间戳。日期格式按照ISO8601标准表示，并需要使用UTC时间。格式为YYYY-MM-DDThh:mm:ssZ 例如，2015-11-23T04:00:00Z（为北京时间2015年11月23日12点0分0秒）
}
