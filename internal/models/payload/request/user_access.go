package request

type UserAccess struct {
	Username string `header:"username"`
	Role     string `header:"role"`
}
