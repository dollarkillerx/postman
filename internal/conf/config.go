package conf

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"strconv"

	"github.com/dollarkillerx/env"
)

type conf struct {
	RocketChatAddr   string `json:"RocketChatAddr"` // exp: http://127.0.0.1:8000/api/v1/chat.postMessage
	RocketChatToken  string `json:"RocketChatToken"`
	RocketChatUserID string `json:"RocketChatUserID"`

	WorkWechatAgentIDInt int
	WorkWechatAgentID    string `json:"WorkWechatAgentID"`
	WorkWechatCorpID     string `json:"WorkWechatCorpID"`
	WorkWechatCorpSecret string `json:"WorkWechatCorpSecret"`

	PostmanAddr  string `json:"PostmanAddr"`
	PostmanToken string `json:"PostmanToken"`
}

var Conf *conf

func init() {
	var cnf conf

	file, err := ioutil.ReadFile("./configs/config.json")
	if err != nil {
		// get env
		err := env.Fill(&cnf)
		if err != nil {
			log.Fatalln("Config Parse Error")
		}

		Conf = &cnf
		return
	}

	err = json.Unmarshal(file, &cnf)
	if err != nil {
		panic(err)
	}

	if cnf.WorkWechatAgentID != "" {
		workWechatAgentIDInt, err := strconv.Atoi(cnf.WorkWechatAgentID)
		if err != nil {
			log.Fatalln("WorkWechatAgentID Error")
		}
		cnf.WorkWechatAgentIDInt = workWechatAgentIDInt
	}

	Conf = &cnf
}
