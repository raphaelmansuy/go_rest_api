package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"

	"net/http"

	_ "github.com/raphaelmansuy/go_rest_api/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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

// @title			A simple REST API
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

	var port = os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	r.GET("/users", getUsers)

	endpointUrl := fmt.Sprintf("http://localhost:%s", port)

	// docs route

	// add swagger middleware
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	log.Println("🚀 Simple service is starting on port " + port)
	// inter
	log.Println(fmt.Sprintf("✅ Simple service is ready to handle requests at %s", endpointUrl))
	// print how to access the Swagger UI
	log.Println(fmt.Sprintf("👉 You can access the Swagger UI at %s/swagger/", endpointUrl))

	// start the server
	r.Run(":" + port)

}
