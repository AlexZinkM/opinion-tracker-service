package main

import (
	"log"
	"net/http"

	"opinion-tracker-service/adapter/controller/rest"
	_ "opinion-tracker-service/docs"
	usecaseimpl "opinion-tracker-service/domain/usecaseimpl"
	"opinion-tracker-service/infra/restserver/nethttp"
)

// @title Opinion Tracker Service API
// @version 1.0
// @description This is a sample server for Opinion Tracker Service.
func main() {
	createUserUseCase := usecaseimpl.NewCreateUserUseCaseImpl()
	userController := rest.NewUserController(createUserUseCase)

	server := nethttp.NewNetHttpServer()

	server.RegisterPublicRoute(http.MethodPost, "/users", userController.CreateUserHandler)
	server.RegisterSwaggerRoutes()

	log.Println("Server starting on localhost:8080")
	log.Println("Swagger UI available at: http://localhost:8080/swagger/")

	if err := server.Start(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}

}
