package transport

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/danielwangai/todo-app/internal/literals"
	"github.com/danielwangai/todo-app/internal/svc"
	"github.com/danielwangai/todo-app/internal/transport/handlers"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// Router ...
type Router struct {
	*mux.Router
}

// NewRouter ...
func NewRouter() *Router {
	return &Router{mux.NewRouter()}
}

func (r *Router) InitializeRoutes(ctx context.Context, service svc.Svc, log *logrus.Logger, db *sql.DB) {
	r.HandleFunc(literals.HealthcheckEndpoint, handlers.Healthcheck(ctx, log, db)).
		Methods(http.MethodGet).
		Name(literals.HealthcheckEndpointName)

	r.HandleFunc(literals.UsersBaseEndpoint, handlers.CreateUser(ctx, service, log)).
		Methods(http.MethodPost).
		Name(literals.CreateUserEndpointName)

	r.HandleFunc(literals.TodoBaseEndpoint, handlers.CreateTodo(ctx, service, log)).
		Methods(http.MethodPost).
		Name(literals.CreateTodoEndpointName)
}
