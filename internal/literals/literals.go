package literals

const (
	RoutePrefix             = "todo-app/v1"
	DatabaseName            = "todo_app"
	HealthcheckEndpoint     = "/v1/healthcheck"
	HealthcheckEndpointName = "healthcheck"
	AuthorizationHeaderName = "Authorization"

	// Users
	UsersBaseEndpoint      = "/v1/users"
	SignupEndpoint         = UsersBaseEndpoint + "/signup"
	CreateUserEndpointName = "create-user"
	LoginEndpoint          = UsersBaseEndpoint + "/login"
	LoginEndpointName      = "login"

	// Todo
	TodoBaseEndpoint            = "/v1/todo"
	CreateTodoItemEndpointName  = "create-todo"
	GetAllTodoItemsEndpointName = "get-all-todo"
)
