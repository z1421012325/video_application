package request

type UserLogin struct {
	UserName string		`json:"user_name" form:"user_name" binding:"required"`
	Password string 	`json:"password" form:"password" binding:"required"`
}
