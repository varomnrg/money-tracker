package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	uh "github.com/varomnrg/money-tracker/handler/user"
	ur "github.com/varomnrg/money-tracker/repository/user"
	us "github.com/varomnrg/money-tracker/service/user"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	DB_URL := os.Getenv("PSQL_DB_URL")

	// User Service
	usersPostgresRepo := ur.NewPostgresqlUserRepository(DB_URL)
	userService := us.NewUserService(usersPostgresRepo)
	userHandler := uh.NewUserHandler(userService)

	router.GET("/users", loggerHandler(userHandler.GetUsers))
	router.GET("/users/:id", loggerHandler(userHandler.GetUser))
	router.POST("/users", loggerHandler(userHandler.CreateUser))
	router.PUT("/users/:id", loggerHandler(userHandler.UpdateUser))
	router.DELETE("/users/:id", loggerHandler(userHandler.DeleteUser))

	log.Println("Server is running at 8000 port.")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func loggerHandler(h httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		startTime := time.Now()

		// Call the original handler
		h(w, r, ps)

		// Log information about the request
		duration := time.Since(startTime)
		log.Printf("[%s] %s %s %s", r.Method, r.RequestURI, r.RemoteAddr, duration)
	}
}
