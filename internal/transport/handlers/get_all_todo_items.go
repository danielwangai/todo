package handlers

import (
	"context"
	"net/http"

	"github.com/danielwangai/todo-app/internal/literals"
	"github.com/danielwangai/todo-app/internal/svc"
	"github.com/sirupsen/logrus"
)

func GetTodoItems(ctx context.Context, service svc.Svc, log *logrus.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Infoln("begin: get user todo items endpoint")
		token, err := svc.GetToken(r.Header.Get(literals.AuthorizationHeaderName))
		if err != nil {
			log.WithError(err).Error("failed to fetch todo items by unauthorized user")
			respondWithError(w, http.StatusBadRequest, "failed to fetch todo items. Login to proceed.")
			return
		}
		// check user-password combination in database
		items, err := service.GetAllTodoItems(ctx, token.User.ID)
		if err != nil {
			log.WithError(err).Errorf("an error ocurred when getting todo items for user of id: %d", token.User.ID)
			respondWithError(w, http.StatusBadRequest, "an error ocurred when getting todo items")
			return
		}

		log.Infof("successfully fetched todo items for user of id: %d", token.User.ID)
		respondWithJSON(w, http.StatusOK, map[string]interface{}{"items": items})
		return
	}
}
