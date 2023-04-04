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

func DeleteTodoItemById(ctx context.Context, service svc.Svc, log *logrus.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Infoln("begin: delete todo item by id")
		_, err := svc.GetToken(r.Header.Get(literals.AuthorizationHeaderName))
		if err != nil {
			log.WithError(err).Error("unauthorized access. Failed to fetch todo item by id")
			respondWithError(w, http.StatusBadRequest, "failed to fetch todo item. Login to proceed.")
			return
		}

		// get id from request params
		params := mux.Vars(r)
		id, err := strconv.Atoi(params["id"])
		if err != nil {
			log.WithError(err).Error("invalid param id type. Must be an integer")
			respondWithError(w, http.StatusBadRequest, "invalid param[id] type. must be an integer")
			return
		}

		// delete todo item
		err = service.DeleteTodoItemById(ctx, id)
		if err != nil {
			log.WithError(err).Errorf("an error ocurred when deleting todo item of id: %d", id)
			respondWithError(w, http.StatusBadRequest, "an error ocurred when deleting todo item by id")
			return
		}

		log.Infof("successfully deleted todo item: %v by id: %d")
		respondWithJSON(w, http.StatusOK, "delete successsful")
		return
	}
}
