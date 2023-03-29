package transport

import (
	"context"
	"net/http"

	"github.com/danielwangai/todo-app/internal/literals"
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

func (r *Router) InitializeRoutes(ctx context.Context, log *logrus.Logger) {
	r.HandleFunc(literals.HealthcheckEndpoint, handlers.Healthcheck(ctx, log)).
		Methods(http.MethodGet).
		Name(literals.HealthcheckEndpointName)
}
