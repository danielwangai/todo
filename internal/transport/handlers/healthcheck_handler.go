package handlers

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/sirupsen/logrus"
)

func Healthcheck(ctx context.Context, log *logrus.Logger, db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res := make(map[string]string)
		err := db.Ping()
		if err != nil {
			res["db"] = "db connection failed"
			res["overallSystemStatus"] = "not okay"
		} else {
			res["overallSystemStatus"] = "ok"
			res["dbStatus"] = "ok"
		}

		respondWithJSON(w, http.StatusOK, res)
	}
}
