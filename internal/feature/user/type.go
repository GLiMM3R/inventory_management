package user

type UserCreateDto struct {
	BranchID string `json:"branch_id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}
