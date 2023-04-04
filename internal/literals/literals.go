package literals

const (
	RoutePrefix             = "todo-app/v1"
	DatabaseName            = "todo_app"
	HealthcheckEndpoint     = "/v1/healthcheck"
	HealthcheckEndpointName = "healthcheck"
	AuthorizationHeaderName = "Authorization"

	// Users
	UsersBaseEndpoint        = "/v1/users"
	FindUserByIdEndpoint     = UsersBaseEndpoint + "/{id}"
	SignupEndpoint           = UsersBaseEndpoint + "/signup"
	CreateUserEndpointName   = "create-user"
	FindUserByIdEndpointName = "find-user-by-id"
	LoginEndpoint            = UsersBaseEndpoint + "/login"
	LoginEndpointName        = "login"

	// Todo
	TodoBaseEndpoint            = "/v1/todo"
	CreateTodoItemEndpointName  = "create-todo"
	GetAllTodoItemsEndpointName = "get-all-todo"
	TodoByIdEndpoint            = TodoBaseEndpoint + "/{id}"
	FindTodoByIdEndpointName    = "find-todo-by-id"
	DeleteTodoByIdEndpointName  = "delete-todo-by-id"
	UpdateTodoByIdEndpointName  = "update-todo-by-id"
)
