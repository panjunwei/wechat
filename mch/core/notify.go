package core

import (
	"bytes"
	"encoding/xml"
	"strings"
	"strconv"
)

type WeChatNotifyResponse struct {
	ReturnCode string `xml:"return_code"`
	ReturnMsg  string `xml:"return_msg"`
}

func (this *WeChatNotifyResponse) ToXmlString() (xmlStr string) {
	buffer := new(bytes.Buffer)
	buffer.WriteString("<xml><return_code><![CDATA[")
	buffer.WriteString(this.ReturnCode)
	buffer.WriteString("]]></return_code>")

	buffer.WriteString("<return_msg><![CDATA[")
	buffer.WriteString(this.ReturnMsg)
	buffer.WriteString("]]></return_msg></xml>")
	xmlStr = buffer.String()
	return
}

type Params map[string]string

func XmlToMap(xmlStr string) Params {

	params := make(Params)
	decoder := xml.NewDecoder(strings.NewReader(xmlStr))

	var (
		key   string
		value string
	)

	for t, err := decoder.Token(); err == nil; t, err = decoder.Token() {
		switch token := t.(type) {
		case xml.StartElement: // 开始标签
			key = token.Name.Local
		case xml.CharData: // 标签内容
			content := string([]byte(token))
			value = content
		}
		if key != "xml" {
			if value != "\n" {
				params.SetString(key, value)
			}
		}
	}

	return params
}

// 微信异步接口处理函数 map本来已经是引用类型了，所以不需要 *Params
func (p Params) SetString(k, s string) Params {
	p[k] = s
	return p
}

func (p Params) GetString(k string) string {
	s, _ := p[k]
	return s
}

func (p Params) SetInt64(k string, i int64) Params {
	p[k] = strconv.FormatInt(i, 10)
	return p
}

func (p Params) GetInt64(k string) int64 {
	i, _ := strconv.ParseInt(p.GetString(k), 10, 64)
	return i
}

// 判断key是否存在
func (p Params) ContainsKey(key string) bool {
	_, ok := p[key]
	return ok
}
