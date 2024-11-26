package entities

import "github.com/golang-jwt/jwt/v5"

type User struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
}

type Password struct {
	Value string `json:"password"`
}

type HashedPassword struct {
	Value string `json:"hashed_password"`
}

type UserClaim struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
	jwt.RegisteredClaims
}
