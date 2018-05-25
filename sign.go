package alisms

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"net/url"
	"sort"
	"strings"
)

var replacer = strings.NewReplacer("+", "%20", "*", "%2A", "%7E", "~")

// 签名方法
func sign(apikey string, params map[string]string) (string, error) {
	/* 排序 */
	keys := make([]string, 0, len(params))
	for k := range params {
		if k == "sign" {
			continue
		}

		keys = append(keys, k)
	}
	sort.Strings(keys)

	/* 拼接字符串 */
	buf := new(bytes.Buffer)
	bufSign := new(bytes.Buffer)

	bufSign.WriteString("GET&")
	bufSign.WriteString(specialURLEncode("/"))
	bufSign.WriteString("&")

	for _, k := range keys {
		v := params[k]
		if len(v) == 0 {
			continue
		}

		buf.WriteString(specialURLEncode(k))
		buf.WriteByte('=')
		buf.WriteString(specialURLEncode(v))
		buf.WriteByte('&')
	}

	buf.Truncate(buf.Len() - 1)

	bufSign.WriteString(specialURLEncode(buf.String()))

	h := hmac.New(sha1.New, []byte(apikey+"&"))

	h.Write([]byte(bufSign.String()))

	bufstr := new(bytes.Buffer)
	bufstr.WriteString("Signature=")
	bufstr.WriteString(specialURLEncode(base64.StdEncoding.EncodeToString(h.Sum(nil))))
	bufstr.WriteString("&")
	bufstr.Write([]byte(buf.String()))

	return bufstr.String(), nil
}

func specialURLEncode(str string) string {
	str = url.QueryEscape(str)
	return replacer.Replace(str)
}
