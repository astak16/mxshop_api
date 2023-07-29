package global

import (
	"mxshop_api/user_web/config"
	"mxshop_api/user_web/proto"

	ut "github.com/go-playground/universal-translator"
)

var (
	Trans        ut.Translator
	ServerConfig *config.ServerConfig = &config.ServerConfig{}

	UserSrvClient proto.UserClient
)
