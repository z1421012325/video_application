package response

const (

	// 正常
	SUCCESS_CODE = 0


	// 5位数2000x 为用户账号异常
	// 登录异常
	ValueCode = 20001
	// 参数异常
	BindErr = 20002
	// 密码验证错误
	UserPassword = 20003
	// token加密错误
	TokenErr = 20004
	// 注册异常
	RegisterErr = 20005
	// 修改异常
	ModifyErr = 20006
	// 身份异常
	IdentityErr = 20009

	// 数据库异常
	DatabaseErr = 40004

)
