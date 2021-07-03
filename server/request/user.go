package request


type UserLogin struct {
	UserName string		`json:"user_name" form:"user_name" binding:"required,min=3,max=30"`
	Password string 	`json:"password" form:"password" binding:"required,min=8,max=30"`
}

type UserRegister struct {
	UserLogin
	NickName string		`json:"nickname" form:"nickname" binding:"required,min=2,max=20"`
}

type ModifyUserData struct {
	Uid int64			`json:"u_id" form:"u_id" binding:"required"`
	NickName string		`json:"nickname" form:"nickname" binding:"min=2,max=20"`
	Portrait string		`json:"portrait,omitempty" form:"portrait" binding:"max=200"`
	Introduce string	`json:"introduce,omitempty" form:"introduce" binding:"max=300"`
}

type ModifyUserPassword struct {
	Uid int64			`json:"u_id" form:"u_id" binding:"required"`
	OldPassword string 	`json:"password" form:"password" binding:"required,min=8,max=30"`
	NewPassword string 	`json:"new_password" form:"new_password" binding:"required,min=8,max=30"`
}