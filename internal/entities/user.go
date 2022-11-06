package entities

import "github.com/golang-jwt/jwt"

type User struct {
	Id       int    `json:"-" db:"id"`
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type SignInInput struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type TokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}
