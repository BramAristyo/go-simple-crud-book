package middleware

import (
	"net/http"
	"simple-restful-api/helper"
	"simple-restful-api/response"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{Handler: handler}
}

func (middleware *AuthMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if "MyAPIKey" == r.Header.Get("X-API-Key") {
		middleware.Handler.ServeHTTP(w, r)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)

		webResponse := response.WebResponse{
			Code: http.StatusUnauthorized,
			Status: "UNAUTHORIZED",
		}

		helper.WriteToResponse(w, webResponse)
	}
}