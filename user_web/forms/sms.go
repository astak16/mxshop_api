package forms

type SendSmsForm struct {
	// tag 的 mobile 和 RegisterValidation 中注册的 mobile 一致
	Mobile string `form:"mobile" json:"mobile" binding:"required,mobile"`
	// 1 表示注册，2 表示登录
	Type uint `form:"type" json:"type" binding:"required,oneof=1 2"`
}
