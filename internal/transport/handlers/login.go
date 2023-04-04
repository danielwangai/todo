package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/danielwangai/todo-app/internal/svc"
	"github.com/sirupsen/logrus"
)

func Login(ctx context.Context, service svc.Svc, log *logrus.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Infoln("begin: login endpoint")
		var u svc.UserLoginAPIRequestType
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&u); err != nil {
			log.WithError(err).Error("invalid request body")
			respondWithError(w, http.StatusBadRequest, "Invalid request payload")
			return
		}
		// validation
		errs := svc.ValidateUserLoginDetails(&u)
		if len(errs) > 0 {
			respondWithJSON(w, http.StatusBadRequest, map[string]interface{}{"errors": errs})
			return
		}

		// find user by email
		user, err := service.FindUserByEmail(ctx, u.Email)
		if err != nil {
			log.WithError(err).Errorf("login failed due to invalid credentials from email: %s", u.Email)
			respondWithError(w, http.StatusBadRequest, "invalid credentials")
			return
		}

		// verify password
		if !svc.CheckPasswordHash(u.Password, user.Password) {
			log.WithError(err).Errorf("login failed due to invalid password: %s", u.Email)
			respondWithError(w, http.StatusBadRequest, "invalid credentials")
			return
		}

		// generate jwt token
		tokenString, err := svc.GenerateJWT(user)
		if err != nil {
			log.WithError(err).Errorf("login failed due to invalid JWT claims for email: %s", u.Email)
			respondWithError(w, http.StatusBadRequest, "could not authenticate at this time try again later.")
			return
		}

		res := map[string]string{
			"message": "login successful",
			"token":   tokenString,
		}

		log.Infof("user logged in successfully with email: %s", u.Email)
		respondWithJSON(w, http.StatusOK, res)
		return
	}
}
