# postman  聚合通知
rocket.chat  企业微信  钉钉

## use

### rocket.chat
POST: `http://0.0.0.0:8675/v1/rocket_chat/send_message`

JSON BODY:
``` 
{
	"to": "#logs",
	"message": "hello world  1231"
}
```
HEADER:
``` 
PostmanToken: PostmanToken
```

### 企业微信
POST: `http://0.0.0.0:8675/v1/work_wechat/send_message`

JSON BODY 1:
``` 
{
	"to_user": "",
	"to_group": "01",  // 发送给组
	"message": "hello world",
	"encryption": false
}
```

JSON BODY 2:
``` 
{
	"to_user": "user", // 发送给某一个用户
	"to_group": "",  
	"message": "hello world",
	"encryption": false
}
```
HEADER:
``` 
PostmanToken: PostmanToken
```

## 部署
修改 `deploy/docker-compose.yaml` 中的配置

然后 `docker-compose up -d`
