package dto

type TokenDetail struct {
	AccessToken            string `json:"accessToken"`
	RefreshToken           string `json:"refreshToken"`
	AccessTokenExpireTime  int64  `json:"accessTokenExpireTime"`
	RefreshTokenExpireTime int64  `json:"refreshTokenExpireTime"`
}
type TokenResponse struct {
	AccessToken           string `json:"accessToken"`
	AccessTokenExpireTime int64  `json:"accessTokenExpireTime"`
}
type RegisterUserByUsernameRequest struct {
	Username string `json:"username" binding:"required,min=5"`
	Email    string `json:"email" binding:"min=6,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type LoginByUsernameRequest struct {
	Username string `json:"username" binding:"required,min=5"`
	Password string `json:"password" binding:"required,min=6"`
}
type TokenDto struct {
	UserId   int
	Username string
	Email    string
}
