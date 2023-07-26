package forms

type PassWordLoginForm struct {
	// tag 的 mobile 和 RegisterValidation 中注册的 mobile 一致
	Mobile    string `form:"mobile" json:"mobile" binding:"required,mobile"`
	PassWord  string `form:"password" json:"password" binding:"required,min=3,max=20"`
	Captcha   string `form:"captcha" json:"captcha" binding:"required,len=5"`
	CaptchaId string `form:"captchaId" json:"captchaId" binding:"required"`
}

type RegisterForm struct {
	Mobile   string `form:"mobile" json:"mobile" binding:"required,mobile"`
	Password string `form:"password" json:"password" binding:"required,min=3,max=20"`
	Code     string `form:"code" json:"code" binding:"required,len=6"`
}
