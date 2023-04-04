package handlers

import (
	"context"
	"net/http"
	"strconv"

	"github.com/danielwangai/todo-app/internal/literals"
	"github.com/danielwangai/todo-app/internal/svc"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func FindUserById(ctx context.Context, service svc.Svc, log *logrus.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Infoln("begin: find user by id endpoint")
		_, err := svc.GetToken(r.Header.Get(literals.AuthorizationHeaderName))
		if err != nil {
			log.WithError(err).Error("unauthorized access. Failed to fetch todo item by id")
			respondWithError(w, http.StatusBadRequest, "failed to fetch todo item. Login to proceed.")
			return
		}

		params := mux.Vars(r)
		id, err := strconv.Atoi(params["id"])
		if err != nil {
			log.WithError(err).Error("invalid param id type. Must be an integer")
			respondWithError(w, http.StatusBadRequest, "invalid param[id] type. must be an integer")
			return
		}

		// find user by id
		user, err := service.FindUserById(ctx, id)
		if err != nil {
			log.WithError(err).Errorf("an error ocurred when finding user by id: %d", id)
			respondWithError(w, http.StatusBadRequest, "an error ocurred when finding user by id")
			return
		}

		log.Infof("successfully found user : %v by id: %d", user, id)
		respondWithJSON(w, http.StatusOK, map[string]interface{}{"userr": user})
		return
	}
}
