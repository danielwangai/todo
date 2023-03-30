package transport

import (
	"context"
	"fmt"
	"net/http"

	"github.com/danielwangai/todo-app/internal/config"
	"github.com/danielwangai/todo-app/internal/logging"
	"github.com/danielwangai/todo-app/internal/repository/psql"
	"github.com/danielwangai/todo-app/internal/svc"
	"github.com/sirupsen/logrus"
)

// Server ...
type Server struct {
	Router *Router
}

// NewServer ...
func NewServer() *Server {
	return &Server{
		Router: NewRouter(),
	}
}

// RunServer initializes services
func RunServer() error {
	log := logging.SetJSONFormatter(logrus.New())
	ctx := context.Background()

	cfg, err := config.FromEnv()
	if err != nil {
		fmt.Println("Error loading configs: ", err)
		return err
	}

	server := NewServer()

	dbClient, err := psql.NewDBClient(log, cfg.DB.Username, cfg.DB.Password, cfg.DB.Host, cfg.DB.DBName, cfg.DB.SSLMode)
	if err != nil {
		return err
	}

	// ping connection
	err = psql.PingDB(log, dbClient)
	if err != nil {
		return err
	}

	dao := psql.New(dbClient, log)
	service := svc.New(dao, log)
	server.Router.InitializeRoutes(ctx, service, log, dbClient)
	log.Infof("starting server on port %s", cfg.WebServer.Port)
	if err := http.ListenAndServe(":"+cfg.WebServer.Port, *server.Router); err != nil {
		log.WithError(err).Error("could not start the HTTP server")
		return err
	}

	return nil
}
