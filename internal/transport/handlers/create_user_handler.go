package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/danielwangai/todo-app/internal/svc"
	"github.com/sirupsen/logrus"
)

func CreateUser(ctx context.Context, service svc.Svc, log *logrus.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Infoln("begin: create user endpoint")
		var u svc.UserAPIRequestType
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&u); err != nil {
			log.WithError(err).Error("invalid request body")
			respondWithError(w, http.StatusBadRequest, "Invalid request payload")
			return
		}

		// validate user input
		errs := svc.ValidateUserInput(&u)
		if len(errs) > 0 {
			respondWithJSON(w, http.StatusBadRequest, map[string]interface{}{"errors": errs})
			return
		}

		// convert api object to service layer object
		svcRequest := ConvertUserRequestToServiceObject(&u)

		s, err := service.CreateUser(ctx, svcRequest)
		if err != nil {
			log.WithError(err).Error("error creating user")
			respondWithError(w, http.StatusBadRequest, "an error ocurred when creating user")
			return
		}

		log.Infof("user created successfully: %v", s)
		respondWithJSON(w, http.StatusCreated, s)
	}
}
