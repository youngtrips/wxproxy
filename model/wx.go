package model

import (
	"crypto/sha1"
	"encoding/xml"
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/youngtrips/wxproxy/internal/log"
)

type BaseParam struct {
	Signature string `json:"signature" form:"signature" query:"signature"`
	Timestamp string `json:"timestamp" form:"timestamp" query:"timestamp"`
	Nonce     string `json:"nonce" form:"nonce" query:"nonce"`
	EchoStr   string `json:"echostr" form:"echostr" query:"echostr"`
}

func (p *BaseParam) Validate(token string) bool {
	items := []string{token, p.Timestamp, p.Nonce}

	sort.Strings(items)
	ctx := strings.Join(items, "")

	log.Debug("signature: ", p.Signature)
	log.Debug("timestamp: ", p.Timestamp)
	log.Debug("Nonce: ", p.Nonce)
	log.Debug("echostr: ", p.EchoStr)

	expected := fmt.Sprintf("%x", sha1.Sum([]byte(ctx)))
	log.Debug("expected: ", expected)
	return expected == p.Signature
}

type BaseTextReq struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string
	FromUserName string
	CreateTime   time.Duration
	MsgType      string
	Content      string
	MsgId        int
}

type CDATAText struct {
	Value string `xml:",cdata"`
}

type BaseTextRes struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   CDATAText
	FromUserName CDATAText
	CreateTime   time.Duration
	MsgType      CDATAText
	Content      CDATAText
}
