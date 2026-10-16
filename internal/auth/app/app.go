package app

import (
	"github.com/mikhail-kold/authGo/internal/auth/adapters/jwt"
	"github.com/mikhail-kold/authGo/internal/auth/app/http"
)

type Application struct {
	Http         Http
	TokenService jwt.Token
}

type Http struct {
	RegisterUser http.RegisterUserHandler
	LoginUser    http.LoginUserHandler
}
