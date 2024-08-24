package user

type User struct {
	ID       int64
	Email    string
	Password string
}

type UserData struct {
	ID       int64  `json:"-"`
	Email    string `binding:"required,email"`
	Password string `binding:"required"`
}