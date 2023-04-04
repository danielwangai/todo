# TODO API

### Stack used
- Golang
- Postgres

### Architecture used
**Clean Architecture**
The trade off here having an adaptable and easy to understand codebase making it easy to make changes but the downside is writing more code.
The API would easily be implemented using a simple 3-layered architecture pattern which is faster to build and prototytpe, but harder to make changes to the codebase. This is because of coupled dependencies e.g. database methods/queries in the API controller/handler. In this case, if we want to change the database engine from say Postgres to Mongo we'll have to also make changes in the controller which get's harder when the codebase is large.

In the architecture chosen for this codebase, if a change in database is required, the service and API layers will not be changed hence the choice.

### App components and layers
##### Logging
Logging is used to capture activity of the API at every stage of the request from initiation to it's success/failure. This eventually make it easier to investigate on-call incidences by narrowing down to where the issue ocurred in the "lifecycle" of the request.

#### Service layer(svc)
The service layer performs business logic operations. It also calls database functions.

#### API layer
This is where API requests and data validations are handled. The API sends data to the service layer which performs operations and sends result(success/error) back to the API layer to be presented to the end user.

### Running the API
#### Dependencies to be installed
- Go
- Postgres for your operating system

1. After cloning cloning the repo, at the root of the codebase, run the command below to fetch necessary packages:-
```
$ go mod tidy
```
2. Create your postgres database. You can use psql on terminal or your favorite GUI client
3. Use the schema at the root of the codebase to create the necessary tables.
4. Open `./etc/config/config.local.env` on your editor and replace the values with appropriate ones
5. At the root of the codebase, export environment variables by running to make them available to be used when running the API:-
```
$ source ./etc/config/config.local.env
```
6. To run the API, change directory to:-
```
$ cd ./cmd/server
```
7. Run the server:-
```
$ go run main.go
```

### Endpoints

All endpoints are protected except healtcheck, signup and login.
Protected Endpoints can't be accessed without logging in i.e. the Authrization Header must contain a valid JWT token string.
The JWT token has a lifespan of 1 hour after which the user has to login again to access the protected endpoints.


| Endpoint  | REST Method | Request Body | Protected Endpoint?|  Notes |
| ------------- | ------------- | ------------- | ------------- | ------------- |
| /v1/healthcheck  | GET  | N/A | False |   Checks the overall health of the API |
| /v1/users/signup  | POST | {"firstName": "first-name", "lastName": "last-name", 	"email": "email@host.com", "password": "password"} | False | User sign up |
| /v1/users/login  | POST | {"email": "email@host.com", "password": "password"} | False | User login |
| /v1/users/{id}  | GET | N/A | True | Get user by id |
| /v1/todo  | POST | {"name": "what to do"} | True | Create new todo item |
| /v1/todo  | GET |  N/A |True | Get all todo items for currently logged in user |
| /v1/todo/{id} | GET | N/A | True  | Get todo item by id |
| /v1/todo/{id}  | PUT | {"name": "what to do"} | True  | Update todo item by id |
| /v1/todo/{id}  | DELETE | N/A | True  | Delete user by id |
