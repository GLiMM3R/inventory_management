package user

type UserCreateDto struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
