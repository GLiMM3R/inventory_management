package auth

type AuthResponse struct {
	User struct {
		Username string `json:"username"`
		Email    string `json:"email"`
	}
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int    `json:"expires_in"`
}
