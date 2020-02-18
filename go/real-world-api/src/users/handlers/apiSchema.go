package handlers

type UserSchema struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Bio      string `json:"bio"`
	Image    string `json:"image"`
}

// UserWithTokenSchema godoc
type UserWithTokenSchema struct {
	UserSchema
	Token string `json:"token"`
}

// ProfileSchema godoc
type ProfileSchema struct {
	Username  string `json:"username"`
	Bio       string `json:"bio"`
	Image     string `json:"image"`
	Following bool   `json:"following"`
}
