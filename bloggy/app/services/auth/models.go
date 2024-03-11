package auth

type TokenData struct {
	Id    string `json:"id"`
	Email string `json:"email"`
}

type JwtToken struct {
	Type  string `json:"type"`
	Token string `json:"token"`
}
