package input

type Register struct {
	FirstName string `form:"first_name" json:"first_name" binding:"required"`
	LastName  string `form:"last_name" json:"last_name" binding:"required"`
	Email     string `form:"email" json:"email" binding:"required"`
	Password  string `form:"password" json:"password" binding:"required"`
}
