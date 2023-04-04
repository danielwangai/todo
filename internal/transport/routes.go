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

	r.HandleFunc(literals.SignupEndpoint, handlers.CreateUser(ctx, service, log)).
		Methods(http.MethodPost).
		Name(literals.CreateUserEndpointName)

	r.HandleFunc(literals.LoginEndpoint, handlers.Login(ctx, service, log)).
		Methods(http.MethodPost).
		Name(literals.LoginEndpointName)

	r.HandleFunc(literals.TodoBaseEndpoint, handlers.CreateTodo(ctx, service, log)).
		Methods(http.MethodPost).
		Name(literals.CreateTodoItemEndpointName)

	r.HandleFunc(literals.TodoBaseEndpoint, handlers.GetTodoItems(ctx, service, log)).
		Methods(http.MethodGet).
		Name(literals.GetAllTodoItemsEndpointName)

	r.HandleFunc(literals.TodoByIdEndpoint, handlers.FindTodoItemById(ctx, service, log)).
		Methods(http.MethodGet).
		Name(literals.FindTodoByIdEndpointName)

	r.HandleFunc(literals.TodoByIdEndpoint, handlers.DeleteTodoItemById(ctx, service, log)).
		Methods(http.MethodDelete).
		Name(literals.DeleteTodoByIdEndpointName)

	r.HandleFunc(literals.TodoByIdEndpoint, handlers.UpdateTodoItem(ctx, service, log)).
		Methods(http.MethodPut).
		Name(literals.UpdateTodoByIdEndpointName)

	r.HandleFunc(literals.FindUserByIdEndpoint, handlers.FindUserById(ctx, service, log)).
		Methods(http.MethodGet).
		Name(literals.FindUserByIdEndpointName)

	r.Use(handlers.AuthMiddleware)
}
