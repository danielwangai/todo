package literals

const (
	RoutePrefix             = "todo-app/v1"
	DatabaseName            = "todo_app"
	HealthcheckEndpoint     = "/v1/healthcheck"
	HealthcheckEndpointName = "healthcheck"

	// Users
	UsersBaseEndpoint      = "/v1/users"
	CreateUserEndpointName = "create-user"

	// Todo
	TodoBaseEndpoint       = "/v1/todo"
	CreateTodoEndpointName = "create-todo"
)
