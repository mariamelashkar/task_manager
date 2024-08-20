package main

import (
	"net/http"
	"task/routers"
	"github.com/rs/cors"
	"log"
	"task/internal/redis"
	"task/config"
    "github.com/swaggo/http-swagger"
	_ "task/docs"

)
// @title Task Management System APIs
// @description This is a sample server for a Task Management System.
// @termsOfService http://swagger.io/terms/
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {

		config.InitConfig()	
		router := routers.InitRoutes()
		redis.InitRedis()
	// Configure CORS settings
	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"}, // Allow all origins, you can specify your frontend URL here
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Content-Type"},
		ExposedHeaders:   []string{"Content-Length", "Access-Control-Allow-Origin"},
		AllowCredentials: true,
	})

	handler := corsHandler.Handler(router)
    router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
