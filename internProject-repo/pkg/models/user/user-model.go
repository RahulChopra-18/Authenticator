package user

type User struct {
	Password []byte `json:"password" form:"password"`
	Username string `json:"username" form:"username"`
	Email    string `json:"email"`
}

type Check struct {
	id       string ` bson:"_id"`
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
	Username string `json:"username" form:"username"`
}
