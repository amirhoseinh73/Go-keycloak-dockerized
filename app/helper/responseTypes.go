package helper

type AdminLoginResponse struct {
	Access_token       string `json:"access_token"`
	Expires_in         int    `json:"expires_in"`
	Refresh_expires_in int    `json:"refresh_expires_in"`
	Token_type         string `json:"token_type"`
	NotBeforePolicy    int    `json:"not-before-policy"`
	Scope              string `json:"scope"`
}

type ErrorResponseKC struct {
	Message string `json:"errorMessage"`
}

type UserInfoResponse struct {
	Sub               string `json:"sub"`
	EmailVerified     string `json:"email_verified"`
	Name              string `json:"name"`
	PreferredUsername string `json:"preferred_username"`
	GivenName         string `json:"given_name"`
	FamilyName        string `json:"family_name"`
	Email             string `json:"email"`
}
