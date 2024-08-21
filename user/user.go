package user

type User struct {
	ID       int64
	Email    string
	Password string
}

type NewUser struct {
	Email    string `binding:"required"`
	Password string `binding:"required"`
}