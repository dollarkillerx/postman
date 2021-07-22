package server

import (
	"encoding/json"
	"errors"
	"log"
	"time"

	"github.com/dollarkillerx/erguotou"
	"github.com/dollarkillerx/postman/internal/conf"
	"github.com/dollarkillerx/postman/pkg"
	"github.com/dollarkillerx/urllib"
)

type wechatTokenResp struct {
	Errcode     int    `json:"errcode"`
	Errmsg      string `json:"errmsg"`
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

func (s *server) updateWechatToken() {
	for {
		code, resp, err := urllib.Get("https://qyapi.weixin.qq.com/cgi-bin/gettoken").
			Queries("corpid", conf.Conf.WorkWechatCorpID).
			Queries("corpsecret", conf.Conf.WorkWechatCorpSecret).Byte()
		if err != nil {
			return
		}
		if code != 200 {
			log.Println("UpdateWechatToken ERROR: ", string(resp))
			time.Sleep(time.Second * 6)
			continue
		}

		token := wechatTokenResp{}
		err = json.Unmarshal(resp, &token)
		if err != nil {
			log.Println("UpdateWechatToken ERROR: ", err)
			time.Sleep(time.Second * 6)
			continue
		}

		if token.Errcode != 0 {
			log.Fatalln("UpdateWechatToken  Token Error: ", string(resp))
		}

		s.wxMu.Lock()
		s.wxToken = token.AccessToken
		s.wxMu.Unlock()
	}
}

// 发送到组
const weixinGroup = "https://qyapi.weixin.qq.com/cgi-bin/appchat/send"

// 单聊
const weixiUser = "https://qyapi.weixin.qq.com/cgi-bin/message/send"

func (s *server) workWechatV1Send(ctx *erguotou.Context) {
	var err error

	var request pkg.WorkWechatV1Request
	err = ctx.BindJson(&request)
	if err != nil {
		ctx.Json(400, pkg.Err400)
		return
	}

	s.wxMu.RLock()
	token := s.wxToken
	s.wxMu.RUnlock()

	switch {
	case request.ToGroup != "":
		err = wechatV1ToGroup(request, token)
	case request.ToUser != "":
		err = wechatV1ToUser(request, token)
	default:
		ctx.Json(400, pkg.Err400)
		return
	}

	if err != nil {
		ctx.Json(400, erguotou.H{"error": err.Error()})
		return
	}

	ctx.Json(200, erguotou.H{"message": "success"})
}

type wechatToGroupRequest struct {
	Chatid  string `json:"chatid"`
	Msgtype string `json:"msgtype"`
	Text    struct {
		Content string `json:"content"`
	} `json:"text"`
	Safe int `json:"safe"`
}

func wechatV1ToGroup(request pkg.WorkWechatV1Request, token string) error {
	safe := 0
	if request.Encryption {
		safe = 1
	}

	req := wechatToGroupRequest{
		Chatid:  request.ToGroup,
		Msgtype: "text",
		Text: struct {
			Content string `json:"content"`
		}{Content: request.Message},
		Safe: safe,
	}

	code, r, err := urllib.Post(weixinGroup).Queries("access_token", token).SetJsonObject(req).Byte()
	if err != nil {
		return err
	}

	if code != 200 {
		return errors.New(string(r))
	}

	return nil
}

type wechatToUserRequest struct {
	Touser  string `json:"touser"`
	Msgtype string `json:"msgtype"`
	Agentid int    `json:"agentid"`
	Text    struct {
		Content string `json:"content"`
	} `json:"text"`
	Safe                   int `json:"safe"`
	EnableIdTrans          int `json:"enable_id_trans"`
	EnableDuplicateCheck   int `json:"enable_duplicate_check"`
	DuplicateCheckInterval int `json:"duplicate_check_interval"`
}

func wechatV1ToUser(request pkg.WorkWechatV1Request, token string) error {
	safe := 0
	if request.Encryption {
		safe = 1
	}

	req := wechatToUserRequest{
		Touser:  request.ToUser,
		Msgtype: "text",
		Agentid: conf.Conf.WorkWechatAgentIDInt,
		Text: struct {
			Content string `json:"content"`
		}{Content: request.Message},
		Safe:                   safe,
		EnableIdTrans:          0,
		EnableDuplicateCheck:   0,
		DuplicateCheckInterval: 1800,
	}

	code, r, err := urllib.Post(weixiUser).Queries("access_token", token).SetJsonObject(req).Byte()
	if err != nil {
		return err
	}

	if code != 200 {
		return errors.New(string(r))
	}

	return nil
}
