package server

import (
	"github.com/dollarkillerx/erguotou"
	"github.com/dollarkillerx/postman/internal/conf"
	"github.com/dollarkillerx/postman/pkg"
	"github.com/dollarkillerx/urllib"
)

func (s *server) rocketChatV1Send(ctx *erguotou.Context) {
	var request pkg.RocketChatV1Request
	err := ctx.BindJson(&request)
	if err != nil {
		ctx.Json(400, pkg.Err400)
		return
	}

	rocketMessage := rocketMessage{Text: request.Message, Channel: request.To}
	code, resp, err := urllib.Post(conf.Conf.RocketChatAddr).SetHeaderMap(map[string]string{
		"X-Auth-Token": conf.Conf.RocketChatToken,
		"X-User-Id":    conf.Conf.RocketChatUserID,
	}).SetJsonObject(rocketMessage).Byte()
	if err != nil {
		ctx.Json(500, erguotou.H{"error": err.Error()})
		return
	}

	if code != 200 {
		ctx.Json(400, erguotou.H{"error": string(resp)})
		return
	}

	ctx.Json(200, erguotou.H{"message": "success"})
}

type rocketMessage struct {
	Text    string `json:"text"`
	Channel string `json:"channel"`
}
