package handlers

import (
	"net/http"

	"github.com/danielwangai/todo-app/internal/literals"
	"github.com/danielwangai/todo-app/internal/svc"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// endpoints that don't need authentications to be used
		noAuthEndpoints := []string{"/v1/users/signup", "/v1/users/login", "/v1/healthcheck"}
		for _, url := range noAuthEndpoints {
			if url == r.URL.Path {
				next.ServeHTTP(w, r)
				return
			}
		}

		tokenString := r.Header.Get(literals.AuthorizationHeaderName)
		if tokenString == "" {
			// err := errors.New("request does not contain access token")
			// log.WithError(err)
			respondWithError(w, http.StatusUnauthorized, "please login to gain access")
			return
		}

		if _, err := svc.GetToken(tokenString); err != nil {
			// log.WithError(err)
			respondWithError(w, http.StatusUnauthorized, "auth token has expired. Please log in again")
			return
		}

		next.ServeHTTP(w, r)
	})
}
