package user

type UpdateUserData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type DeleteUserData struct {
	Password string `json:"password"`
}
