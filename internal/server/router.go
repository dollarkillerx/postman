package server

import (
	"github.com/dollarkillerx/erguotou"
	"github.com/dollarkillerx/postman/internal/conf"
)

func (s *server) router() {
	v1 := s.app.Group("/v1", func(ctx *erguotou.Context) {
		if conf.Conf.PostmanToken != "" {
			header := ctx.Header("PostmanToken")
			if header != conf.Conf.PostmanToken {
				ctx.Json(401, erguotou.H{"error": "PostmanToken 401"})
				return
			}
		}
		ctx.Next()
	})

	rocketChatV1 := v1.Group("/rocket_chat")
	workWechatV1 := v1.Group("/work_wechat")

	rocketChatV1.Post("/send_message", s.rocketChatV1Send)
	workWechatV1.Post("/send_message", s.workWechatV1Send)
}
