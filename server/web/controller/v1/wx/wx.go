package wx

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/youngtrips/wxproxy/internal/config"
	"github.com/youngtrips/wxproxy/internal/log"
	"github.com/youngtrips/wxproxy/model"
)

func ValidateHandler(c echo.Context) error {
	p := &model.BaseParam{}
	if err := c.Bind(p); err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("%s", err))
	}

	cfg := config.Config()
	if !p.Validate(cfg.WX.Token) {
		log.Info("signature validate failed.")
		return c.String(http.StatusOK, "")
	}
	log.Info("signature validate success.")
	return c.String(http.StatusOK, p.EchoStr)
}

func MessageHandler(c echo.Context) error {

	p := &model.BaseParam{}
	p.Signature = c.QueryParam("signature")
	p.Timestamp = c.QueryParam("timestamp")
	p.Nonce = c.QueryParam("nonce")

	cfg := config.Config()
	if !p.Validate(cfg.WX.Token) {
		log.Info("signature validate failed.")
		return c.String(http.StatusOK, "success")
	}
	log.Info("signature validate success.")

	req := &model.BaseTextReq{}
	if err := c.Bind(req); err != nil {
		log.Error("parse req failed: ", err)
		return c.String(http.StatusOK, "success")
	}

	res := proc(req)
	return c.XML(http.StatusOK, res)
}

func buildTextRes(req *model.BaseTextReq, content string) *model.BaseTextRes {
	return &model.BaseTextRes{
		FromUserName: model.CDATAText{req.ToUserName},
		ToUserName:   model.CDATAText{req.FromUserName},
		CreateTime:   time.Duration(time.Now().Unix()),
		MsgType:      model.CDATAText{"text"},
		Content:      model.CDATAText{content},
	}
}

//
func proc(req *model.BaseTextReq) *model.BaseTextRes {
	if req.MsgType != "text" {
		return &model.BaseTextRes{
			FromUserName: model.CDATAText{req.ToUserName},
			ToUserName:   model.CDATAText{req.FromUserName},
			CreateTime:   time.Duration(time.Now().Unix()),
			MsgType:      model.CDATAText{"text"},
			Content:      model.CDATAText{fmt.Sprintf("message format error, expected text, but got %s", req.MsgType)},
		}
	}

	return buildTextRes(req, "hello, "+req.Content)
}
