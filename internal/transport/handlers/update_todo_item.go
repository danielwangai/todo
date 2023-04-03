package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/danielwangai/todo-app/internal/literals"
	"github.com/danielwangai/todo-app/internal/svc"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func UpdateTodoItem(ctx context.Context, service svc.Svc, log *logrus.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Infoln("begin: update todo item by id")
		_, err := svc.GetToken(r.Header.Get(literals.AuthorizationHeaderName))
		if err != nil {
			log.WithError(err).Error("unauthorized access. Failed to update todo item by id")
			respondWithError(w, http.StatusBadRequest, "failed to update todo item. Login to proceed.")
			return
		}

		params := mux.Vars(r)
		id, err := strconv.Atoi(params["id"])
		if err != nil {
			log.WithError(err).Error("invalid param id type. Must be an integer")
			respondWithError(w, http.StatusBadRequest, "invalid param[id] type. must be an integer")
			return
		}
		var i svc.ItemServiceRequestType
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&i); err != nil {
			log.WithError(err).Error("invalid request body")
			respondWithError(w, http.StatusBadRequest, "Invalid request payload")
			return
		}

		// find item by id
		item, err := service.FindTodoItemById(ctx, id)
		if err != nil {
			log.WithError(err).Errorf("an error ocurred when finding todo item of id: %d", id)
			respondWithError(w, http.StatusBadRequest, "an error ocurred when finding todo item by id")
			return
		}

		item, err = service.UpdateTodoItem(ctx, item, i.Name)
		if err != nil {
			log.WithError(err).Errorf("an error ocurred when updating todo item of id: %d", id)
			respondWithError(w, http.StatusBadRequest, "an error ocurred when updating todo item by id")
			return
		}

		log.Infof("successfully updated todo item: %v by id: %d")
		respondWithJSON(w, http.StatusOK, map[string]interface{}{"item": item})
		return
	}
}
