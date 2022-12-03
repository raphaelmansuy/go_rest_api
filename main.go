package main

import (
	"log"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	_ "github.com/raphaelmansuy/go_rest_api/docs"
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
//
//	@Summary		Get a list of users
//	@Description	Returns a list of all users in the system.
//	@Tags			users
//	@Success		200	{array}	User
//	@Router			/users [get]
func getUsers(c *gin.Context) {
	c.JSON(http.StatusOK, users)
}

// @title			My Simple REST API
// @version		1.0
// @description	This is a simple API for managing users.
// @termsOfService	http://swagger.io/terms/
// @contact.name	Jane Doe
// @contact.email	jane.doe@example.com
// @license.name	Apache 2.0
// @license.url	http://www.apache.org/licenses/LICENSE-2.0.html
// @host			localhost:8080
// @BasePath		/
func main() {
	r := gin.Default()


	r.GET("/users", getUsers)

	// docs route
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))



	log.Println("🚀 Simple service is starting on http://localhost:8080/ ...")
	log.Println("🚀 Simple service is ready to handle requests at http://localhost:8080/")
	// print how to access the Swagger UI
	log.Println("👉 You can access the Swagger UI at http://localhost:8080/docs/")

	// start the server
	r.Run(":8080")

}
