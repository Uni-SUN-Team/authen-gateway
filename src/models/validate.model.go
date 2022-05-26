package models

type Validate struct {
	Token string `json:"token"`
}

type Signin struct {
	Token  string `json:"refresh_token"`
	UserId int    `json:"user_id"`
}

type ReponseRefreshToken struct {
	Status  bool       `json:"status"`
	Message string     `json:"message"`
	Claims  RefreshJWT `json:"claims"`
}
