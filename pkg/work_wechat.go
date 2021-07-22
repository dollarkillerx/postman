package pkg

type WorkWechatV1Request struct {
	ToUser     string `json:"to_user"`
	ToGroup    string `json:"to_group"`
	Message    string `json:"message"`
	Encryption bool   `json:"encryption"`
}
