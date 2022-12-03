package main 

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"

	_ "github.com/swaggo/http-swagger/example/docs"

	"github.com/swaggo/http-swagger"
)

// User represents a user in the system.
type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

// users is a mock in-memory store of users.
var users = []User{
	{ID: 1, FirstName: "John", LastName: "Doe", Email: "john.doe@example.com"},
	{ID: 2, FirstName: "Jane", LastName: "Doe", Email: "jane.doe@example.com"},
}

// GetUsers returns a list of all users.
//	@Summary		Get a list of users
//	@Description	Returns a list of all users in the system.
//	@Tags			users
//	@Success		200	{array}	User
//	@Router			/users [get]
func getUsers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(users)
}

//	@title			My SimpleAPI
//	@version		1.0
//	@description	This is a simple API for managing users.
//	@termsOfService	http://swagger.io/terms/
//	@contact.name	Jane Doe
//	@contact.email	jane.doe@example.com
//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html
//	@host			localhost:8080
//	@BasePath		/
func main() {
	r := mux.NewRouter()

	
	r.HandleFunc("/users", getUsers).Methods("GET")

	// Serve the API documentation at the "/docs" endpoint using the Swagger 2.0 format.
	// The Swagger 2.0 specification is available at https://swagger.io/specification/v2/.
	// The local directory "docs" contains the Swagger UI files.
	// docs contains swagger.json and swagger.yaml and docs.go


	r.PathPrefix("/docs/").Handler(httpSwagger.WrapHandler)




	log.Println("🚀 Simple service is starting on http://localhost:8080/ ...")
	log.Println("🚀 Simple service is ready to handle requests at http://localhost:8080/")
	// print how to access the Swagger UI
	log.Println("👉 You can access the Swagger UI at http://localhost:8080/documentation/")

	// start the server
	log.Fatal(http.ListenAndServe(":8080", r))

}
