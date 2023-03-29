package handlers

import (
	"context"
	"net/http"

	"github.com/sirupsen/logrus"
)

func Healthcheck(ctx context.Context, log *logrus.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res := map[string]string{
			"status": "ok",
		}
		respondWithJSON(w, http.StatusOK, res)
	}
}
