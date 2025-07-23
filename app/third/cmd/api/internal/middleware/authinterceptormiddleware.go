package middleware

import (
	"fmt"
	"net/http"
)

type AuthInterceptorMiddleware struct {
}

func NewAuthInterceptorMiddleware() *AuthInterceptorMiddleware {
	return &AuthInterceptorMiddleware{}
}

func (m *AuthInterceptorMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("auth interceptor middleware before")
		next(w, r)
		fmt.Println("auth interceptor middleware after")
	}
}
