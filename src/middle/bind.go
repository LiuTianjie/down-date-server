package middle

type User struct {
	User string `form:"user" json:"user"  binding:"required"`
	Pwd  string `form:"pwd" json:"pwd" binding:"required"`
}
