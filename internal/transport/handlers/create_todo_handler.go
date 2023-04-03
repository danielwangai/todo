package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/danielwangai/todo-app/internal/svc"
	"github.com/sirupsen/logrus"
)

func CreateTodo(ctx context.Context, service svc.Svc, log *logrus.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Infoln("begin: create todo endpoint")
		var i svc.ItemServiceRequestType
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&i); err != nil {
			log.WithError(err).Error("invalid request body")
			respondWithError(w, http.StatusBadRequest, "Invalid request payload")
			return
		}
		// validate item
		errs := svc.ValidateCreateItemInput(&i)
		if len(errs) > 0 {
			log.Errorf("create item failed because of the following errors: %v", errs)
			respondWithJSON(w, http.StatusBadRequest, errs)
			return
		}

		newItem, err := service.CreateTodoItem(ctx, &i)
		if err != nil {
			log.WithError(err).Error("an error ocurred when creating item")
			respondWithError(w, http.StatusInternalServerError, "an error ocurred when creating item")
			return
		}

		log.Infof("item created successfully: %v", newItem)
		respondWithJSON(w, http.StatusCreated, newItem)
	}
}